package boot

import (
	"flag"
	"xlsx_to_db_schema/constants"
)

func InitArguments(adapter *string, tableName *string, sourcePath *string, destPath *string) {
	flag.StringVar(adapter, constants.AdapterName, constants.MySQLAdapter, constants.AdapterName)
	flag.StringVar(tableName, constants.TableName, constants.ExampleTableName, constants.TableName)
	flag.StringVar(sourcePath, constants.SourcePathName, constants.SourcePathDefault, constants.SourcePathName)
	flag.StringVar(destPath, constants.DestPathName, constants.DestPathDefault, constants.DestPathName)
	flag.Parse()
}
