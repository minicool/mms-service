package db

import (
	_"github.com/jinzhu/gorm"
)

var models []interface{}
var migrations []string

func RegisterModel(model interface{}) {
	models = append(models, model)
}

func RegisterMigration(migration string) {
	migrations = append(migrations, migration)
}

func AutoMigrate() {
	for _, migration := range migrations {
		Conn.Exec(migration)
	}
	Conn.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8 AUTO_INCREMENT=1").AutoMigrate(models...)
}