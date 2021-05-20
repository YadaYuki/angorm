package angorm

import (
	"fmt"

	"gorm.io/gorm"
)

func registerCallback(db *gorm.DB) error {
	db.Callback().Query().Register("printValue", printValue)
	return nil
}

func printValue(db *gorm.DB) {
	fmt.Println(db.Statement.ReflectValue)
}
