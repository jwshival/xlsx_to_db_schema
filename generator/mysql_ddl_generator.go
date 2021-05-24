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
	var tmpOptions = fmt.Sprintf("ENGINE=%s AUTO_INCREMENT=%s DEFAULT"+
		"CHARSET=%s COLLATE=%s ROW_FORMAT=%s;", c.EngineDefault, c.AutoIncrementDefault,
		c.CharsetDefault, c.CollateDefault, c.RowFormatDefault)
	table := []string{tmpCreateTable, tmpCreateColumns, tmpDestCreateTable, tmpOptions}
	return strings.Join(table, "")
}

func createColumn(columnName string, columnType string, options map[string]string) string {
	tmpCreateColumn := fmt.Sprintf("%s %s ", columnName, columnType)
	for _, value := range options {
		tmpCreateColumn += " " + value
	}
	return tmpCreateColumn
}
