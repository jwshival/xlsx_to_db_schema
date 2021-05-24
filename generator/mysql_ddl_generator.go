package generator

import (
	"fmt"
	"strings"
	c "xlsx_to_db_schema/constants"
)

/*
Generate to create DLL String
*/

func createTable(tableName string, createColumns []string) string {
	tmpCreateTable := fmt.Sprintf("CREATE TABLE %s (\n", tableName)
	tmpCreateColumns := strings.Join(createColumns, ",\n")
	tmpDestCreateTable := fmt.Sprintf("\n)")
	var tmpOptions = fmt.Sprintf(" ENGINE=%s AUTO_INCREMENT=%s DEFAULT "+
		"CHARSET=%s COLLATE=%s ROW_FORMAT=%s;", c.EngineDefault, c.AutoIncrementDefault,
		c.CharsetDefault, c.CollateDefault, c.RowFormatDefault)
	table := []string{tmpCreateTable, tmpCreateColumns, tmpDestCreateTable, tmpOptions}
	return strings.Join(table, "")
}

func createColumn(physicalName string, cType string, notNullChecked string, primaryKeyChecked string) string {
	var notNull, primaryKey, options string
	if notNullChecked == "x" {
		notNull = "NOT NULL"
	}
	if primaryKeyChecked == "x" {
		primaryKey = "PRIMARY KEY"
	}
	if physicalName == "id" {
		options = "AUTO_INCREMENT"
	}
	tmpCreateColumn := fmt.Sprintf("`%s` %s %s %s %s", physicalName, cType,
		notNull, primaryKey, options)
	return strings.TrimSpace(tmpCreateColumn)
}

func createDBTemplate() string{
	return ""
}
