package angorm

import (
	"fmt"

	"gorm.io/gorm"
)

type AngormPlugin struct {
}

func NewPlugin() *AngormPlugin {
	return &AngormPlugin{}
}

func (p *AngormPlugin) Name() (name string) {
	return ""
}

func (p *AngormPlugin) Initialize(db *gorm.DB) error {
	db.Callback().Query().Register("printValue", printValue)
	return nil
}

func printValue(db *gorm.DB) {
	fmt.Println(db.Statement.ReflectValue)
}
