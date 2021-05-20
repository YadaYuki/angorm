package angorm

import (
	"fmt"

	"gorm.io/gorm"
)

type Angorm struct {
	Db *gorm.DB
}

func New() *Angorm {
	return &Angorm{}
}

func (angorm *Angorm) Name() (name string) {
	return ""
}

func (angorm *Angorm) Initialize(db *gorm.DB) error {
	angorm.Db = db
	angorm.Db.Callback().Query().Register("printValue", printValue)
	return nil
}

func printValue(db *gorm.DB) {
	fmt.Println(db.Statement.ReflectValue)
}
