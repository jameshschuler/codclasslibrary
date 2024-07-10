package loadout

import (
	"backend/gen/postgres/public/model"
	"backend/gen/postgres/public/table"
	"context"
	"database/sql"
	"time"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

type LoadoutStore interface {
	ListCommunityLoadouts(page, pageSize int64) (*[]model.Loadouts, error)
	GetCommunityLoadout(loadoutId string) (*LoadoutDetail, error)
	ListLoadoutsByUser(userId string, page, pageSize int64) (*[]model.Loadouts, error)
	GetLoadoutByUser(userId, loadoutId string) (*LoadoutDetail, error)
	CreateLoadout(loadout *model.Loadouts, attachments []uuid.UUID) (uuid.UUID, error)
}

func (s *Store) ListCommunityLoadouts(page, pageSize int64) (*[]model.Loadouts, error) {
	var dest []model.Loadouts

	limit := pageSize
	offset := (page - 1) * pageSize

	stmt := table.Loadouts.SELECT(
		table.Loadouts.AllColumns,
	).FROM(table.Loadouts).
		LIMIT(limit).
		OFFSET(offset).
		ORDER_BY(table.Loadouts.Title.ASC())

	err := stmt.Query(s.db, &dest)

	if err != nil {
		return nil, err
	}

	return &dest, nil
}

func (s *Store) ListLoadoutsByUser(userId string, page, pageSize int64) (*[]model.Loadouts, error) {
	var userIdString = Uuid(userId)

	var dest []model.Loadouts

	limit := pageSize
	offset := (page - 1) * pageSize

	stmt := table.Loadouts.SELECT(
		table.Loadouts.AllColumns,
	).FROM(table.Loadouts).
		WHERE(
			table.Loadouts.CreatedBy.EQ(postgres.UUID(userIdString))).
		LIMIT(limit).
		OFFSET(offset).
		ORDER_BY(table.Loadouts.Title.ASC())

	err := stmt.Query(s.db, &dest)

	if err != nil {
		return nil, err
	}

	return &dest, nil
}

func (s *Store) GetLoadoutByUser(userId, loadoutId string) (*LoadoutDetail, error) {
	var userIdString = Uuid(userId)
	var loadoutIdString = Uuid(loadoutId)

	var foundLoadout LoadoutDetail

	stmt := table.Loadouts.SELECT(
		table.Loadouts.AllColumns,
		table.Attachments.AllColumns,
	).FROM(table.Loadouts.
		LEFT_JOIN(table.LoadoutAttachment, table.Loadouts.ID.EQ(table.LoadoutAttachment.LoadoutID)).
		LEFT_JOIN(table.Attachments, table.Attachments.ID.EQ(table.LoadoutAttachment.AttachmentID))).
		WHERE(
			table.Loadouts.CreatedBy.EQ(postgres.UUID(userIdString)).
				AND(table.Loadouts.ID.EQ(postgres.UUID(loadoutIdString))))

	err := stmt.Query(s.db, &foundLoadout)

	if err != nil {
		return nil, err
	}

	return &foundLoadout, nil
}

func (s *Store) GetCommunityLoadout(loadoutId string) (*LoadoutDetail, error) {
	var loadoutIdString = Uuid(loadoutId)
	var foundLoadout LoadoutDetail

	stmt := table.Loadouts.SELECT(
		table.Loadouts.AllColumns,
		table.Attachments.AllColumns,
	).FROM(table.Loadouts.
		LEFT_JOIN(table.LoadoutAttachment, table.Loadouts.ID.EQ(table.LoadoutAttachment.LoadoutID)).
		LEFT_JOIN(table.Attachments, table.Attachments.ID.EQ(table.LoadoutAttachment.AttachmentID))).
		WHERE(
			table.Loadouts.ID.EQ(postgres.UUID(loadoutIdString)))

	err := stmt.Query(s.db, &foundLoadout)

	if err != nil {
		return nil, err
	}

	return &foundLoadout, nil
}

func (s *Store) CreateLoadout(loadout *model.Loadouts, attachments []uuid.UUID) (uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var dest model.Loadouts

	tx, err := s.db.Begin()

	if err != nil {
		return uuid.Nil, err
	}

	defer tx.Rollback()

	stmt := table.Loadouts.INSERT(
		table.Loadouts.Title,
		table.Loadouts.Source,
		table.Loadouts.SourceURL,
		table.Loadouts.WeaponName,
		table.Loadouts.WeaponCategory,
		table.Loadouts.CreatedBy,
		table.Loadouts.GameID,
	).MODEL(loadout).RETURNING(
		table.Loadouts.ID,
	)

	err = stmt.QueryContext(ctx, tx, &dest)

	if err != nil {
		return uuid.Nil, err
	}

	err = insertAttachments(tx, dest.ID, attachments)

	if err != nil {
		return uuid.Nil, err
	}

	err = tx.Commit()

	if err != nil {
		return uuid.Nil, err
	}

	return dest.ID, nil
}

func insertAttachments(tx *sql.Tx, loadoutId uuid.UUID, attachments []uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	att := []model.LoadoutAttachment{}
	for _, a := range attachments {
		att = append(att, model.LoadoutAttachment{
			LoadoutID:    loadoutId,
			AttachmentID: a,
		})
	}

	insertAttachmentsStmt := table.LoadoutAttachment.INSERT(table.LoadoutAttachment.LoadoutID, table.LoadoutAttachment.AttachmentID).MODELS(att).RETURNING(
		table.LoadoutAttachment.AllColumns,
	)

	_, err := insertAttachmentsStmt.ExecContext(ctx, tx)

	return err
}
