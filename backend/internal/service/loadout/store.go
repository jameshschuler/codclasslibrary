package loadout

import (
	"backend/gen/postgres/public/model"
	"backend/gen/postgres/public/table"
	"database/sql"
	"fmt"

	"github.com/go-jet/jet/v2/postgres"
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
	ListLoadoutsByUser(userId string) (*[]model.Loadouts, error)
	GetLoadoutByUser(userId, loadoutId string) (*LoadoutDetail, error)
}

func (s *Store) ListCommunityLoadouts(page, pageSize int64) (*[]model.Loadouts, error) {
	var dest []model.Loadouts

	limit := pageSize
	offset := (page - 1) * pageSize

	stmt := table.Loadouts.SELECT(
		table.Loadouts.AllColumns.Except(table.Loadouts.Attachments),
	).FROM(table.Loadouts).
		LIMIT(limit).
		OFFSET(offset).
		ORDER_BY(table.Loadouts.Title.ASC())

	debugSql := stmt.DebugSql()
	fmt.Println(debugSql)

	err := stmt.Query(s.db, &dest)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &dest, nil
}

func (s *Store) ListLoadoutsByUser(userId string) (*[]model.Loadouts, error) {
	if userId == "" {
		return nil, fmt.Errorf("userId is required")
	}

	var userIdString = Uuid(userId)

	var dest []model.Loadouts

	stmt := table.Loadouts.SELECT(
		table.Loadouts.AllColumns.Except(table.Loadouts.Attachments),
	).FROM(table.Loadouts).
		WHERE(
			table.Loadouts.CreatedBy.EQ(postgres.UUID(userIdString))).
		ORDER_BY(table.Loadouts.Title.ASC())

	err := stmt.Query(s.db, &dest)

	if err != nil {
		fmt.Println(err)
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
		LEFT_JOIN(table.LoadAttachment, table.Loadouts.ID.EQ(table.LoadAttachment.LoadoutID)).
		LEFT_JOIN(table.Attachments, table.Attachments.ID.EQ(table.LoadAttachment.AttachmentID))).
		WHERE(
			table.Loadouts.CreatedBy.EQ(postgres.UUID(userIdString)).
				AND(table.Loadouts.ID.EQ(postgres.UUID(loadoutIdString))))

	err := stmt.Query(s.db, &foundLoadout)

	if err != nil {
		fmt.Println(err)
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
		LEFT_JOIN(table.LoadAttachment, table.Loadouts.ID.EQ(table.LoadAttachment.LoadoutID)).
		LEFT_JOIN(table.Attachments, table.Attachments.ID.EQ(table.LoadAttachment.AttachmentID))).
		WHERE(
			table.Loadouts.ID.EQ(postgres.UUID(loadoutIdString)))

	err := stmt.Query(s.db, &foundLoadout)

	if err != nil {
		fmt.Println("Error: ", err)
		return nil, err
	}

	return &foundLoadout, nil
}
