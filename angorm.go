package angorm

import (
	"gorm.io/gorm"
)

type Angorm struct {
	*gorm.DB
}

func New() *Angorm {
	return &Angorm{}
}

func (angorm *Angorm) Name() (name string) {
	return ""
}

func (angorm *Angorm) Initialize(db *gorm.DB) error {
	angorm.DB = db
	return registerCallback(db)
}
