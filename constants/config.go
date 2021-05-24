package constants

type StringFunc func(notNull string) string

const (
	CsvExtName        string = ".csv"
	XlsxExtName       string = ".xlsx"
	AdapterName       string = "adapter"
	TableName		  string = "table_name"
	ExampleTableName  string = "example"
	MySQLAdapter      string = "mysql"
	PSQLAdapter       string = "postgre"
	ExecuteDone       string = "Done!"
	SourcePathName    string = "source"
	SourcePathDefault string = "/"
	DestPathName      string = "dest"
	DestPathDefault   string = "/"
	// Error messages
	NotSupportAdapter  string = "no support this adapter"
	FilePathNotExist   string = "file path is not exist"
	CanNotOpenFile     string = "can not open this file"
	FileExtensionWrong string = "file extension is wrong"
	CanNotReadCsvFile  string = "can not read file csv"
	FileHeaderWrong    string = "file header is wrong"
	CanNotCreateFile   string = "can not create this file"
	CanNotWriteFile	   string = "can not write this file"
)
