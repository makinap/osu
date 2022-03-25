package migration

import (
	"github.com/makinap/osu/config"
	"github.com/makinap/osu/graph/model"
)

func MigrateTable() {
	db := config.GetDB()

	db.AutoMigrate(&model.User{})
}