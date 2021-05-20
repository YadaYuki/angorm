package angorm

import (
	"fmt"

	"gorm.io/gorm"
)

type Angorm struct {
}

func New() *Angorm {
	return &Angorm{}
}

func (p *Angorm) Name() (name string) {
	return ""
}

func (p *Angorm) Initialize(db *gorm.DB) error {
	db.Callback().Query().Register("printValue", printValue)
	return nil
}

func printValue(db *gorm.DB) {
	fmt.Println(db.Statement.ReflectValue)
}
