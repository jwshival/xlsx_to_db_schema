package validators

import (
	"xlsx_to_db_schema/utils"
)

type Validator struct {
	adapter    string
	sourcePath string
	destPath   string
	Err        error
}

func NewValidator(adapter string, sourcePath string, destPath string) *Validator {
	return &Validator{adapter: adapter, sourcePath: sourcePath, destPath: destPath}
}

func getValidatorErr(err error, validatorErr *error)  {
	*validatorErr = err
}

func (validator *Validator) Valid() *[][]string {
	var err error
	err = checkSupportedAdapter(validator.adapter)
	if err != nil {
		getValidatorErr(err, &validator.Err)
		return nil
	}

	err = checkSupportedAdapter(validator.adapter)
	if err != nil {
		getValidatorErr(err, &validator.Err)
		return nil
	}

	err = checkFilePath(validator.sourcePath)
	if err != nil {
		getValidatorErr(err, &validator.Err)
		return nil
	}

	err = checkFileExtension(validator.sourcePath)
	if err != nil {
		getValidatorErr(err, &validator.Err)
		return nil
	}

	records, err := utils.ReadData(validator.sourcePath)
	if err != nil {
		getValidatorErr(err, &validator.Err)
		return nil
	}

	err = checkFileHeader(records[0])
	if err != nil {
		getValidatorErr(err, &validator.Err)
		return nil
	}

	return &records

}
