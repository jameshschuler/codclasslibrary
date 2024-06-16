//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var LoadAttachment = newLoadAttachmentTable("public", "load_attachment", "")

type loadAttachmentTable struct {
	postgres.Table

	// Columns
	LoadoutID    postgres.ColumnString
	AttachmentID postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type LoadAttachmentTable struct {
	loadAttachmentTable

	EXCLUDED loadAttachmentTable
}

// AS creates new LoadAttachmentTable with assigned alias
func (a LoadAttachmentTable) AS(alias string) *LoadAttachmentTable {
	return newLoadAttachmentTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new LoadAttachmentTable with assigned schema name
func (a LoadAttachmentTable) FromSchema(schemaName string) *LoadAttachmentTable {
	return newLoadAttachmentTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new LoadAttachmentTable with assigned table prefix
func (a LoadAttachmentTable) WithPrefix(prefix string) *LoadAttachmentTable {
	return newLoadAttachmentTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new LoadAttachmentTable with assigned table suffix
func (a LoadAttachmentTable) WithSuffix(suffix string) *LoadAttachmentTable {
	return newLoadAttachmentTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newLoadAttachmentTable(schemaName, tableName, alias string) *LoadAttachmentTable {
	return &LoadAttachmentTable{
		loadAttachmentTable: newLoadAttachmentTableImpl(schemaName, tableName, alias),
		EXCLUDED:            newLoadAttachmentTableImpl("", "excluded", ""),
	}
}

func newLoadAttachmentTableImpl(schemaName, tableName, alias string) loadAttachmentTable {
	var (
		LoadoutIDColumn    = postgres.StringColumn("loadout_id")
		AttachmentIDColumn = postgres.StringColumn("attachment_id")
		allColumns         = postgres.ColumnList{LoadoutIDColumn, AttachmentIDColumn}
		mutableColumns     = postgres.ColumnList{}
	)

	return loadAttachmentTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		LoadoutID:    LoadoutIDColumn,
		AttachmentID: AttachmentIDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}