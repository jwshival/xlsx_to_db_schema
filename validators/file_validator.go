package validators

import (
	"errors"
	"os"
	"path/filepath"
	"reflect"
	"xlsx_to_db_schema/constants"
)

var defaultHeader = []string{"Physical Name", "Type", "Not Null",
	"Primary Key"}

/*
Check console from user.

1. File extensions (csv or xlsx)
2. PATH (exist for 2 arguments sourcePath - destPath)
3. File formats (file header, cell)
- file header (value of a cell).
*/

func checkFilePath(sourcePath string) error {
	/*
	os.IsNotExist(err)
	- Catch this error then will check.
	-> if not_exist => true
	-> else if exist -> false
	 */
	if _, err := os.Stat(sourcePath); os.IsNotExist(err) {
		return errors.New(constants.FilePathNotExist)
	}
	return nil
}

func checkFileExtension(sourcePath string) error {
	extensionName := filepath.Ext(sourcePath)
	if extensionName != constants.CsvExtName && extensionName != constants.XlsxExtName {
		return errors.New(constants.FileExtensionWrong)
	}
	return nil
}

func checkFileHeader(header []string) error {
	if reflect.DeepEqual(header, defaultHeader) {
		return nil
	}
	return errors.New(constants.FileHeaderWrong)
}
