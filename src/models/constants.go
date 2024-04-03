package models

import "fmt"

const CREATE string = "create"
const EDIT string = "edit"

func FieldisEmptyMessage(field string) string {
	return fmt.Sprintf("The field %s is needed and cannot be empty", field)
}
