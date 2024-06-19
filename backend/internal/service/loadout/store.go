package loadout

import (
	"backend/gen/postgres/public/model"
	"backend/gen/postgres/public/table"
	"database/sql"
	"fmt"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

// stmt := table.Loadouts.SELECT(
// 	table.Loadouts.AllColumns,
// 	table.Attachments.AllColumns,
// ).FROM(
// 	table.Loadouts.
// 		LEFT_JOIN(table.LoadAttachment, table.Loadouts.ID.EQ(table.LoadAttachment.LoadoutID)).
// 		LEFT_JOIN(table.Attachments, table.Attachments.ID.EQ(table.LoadAttachment.AttachmentID))).
// 	ORDER_BY(table.Loadouts.Title.ASC())

func (s *Store) ListLoadouts() ([]model.Loadouts, error) {
	var dest []model.Loadouts

	stmt := table.Loadouts.SELECT(
		table.Loadouts.AllColumns.Except(table.Loadouts.Attachments),
	).FROM(table.Loadouts).
		ORDER_BY(table.Loadouts.Title.ASC())

	debugSql := stmt.DebugSql()
	fmt.Println(debugSql)

	err := stmt.Query(s.db, &dest)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return dest, nil
}
