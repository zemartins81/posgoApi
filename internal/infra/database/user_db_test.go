package database

import (
	"github.com/stretchr/testify/assert"
	"github.com/zemartins81/posgoApi/internal/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func TestCreateUSer(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	if err := db.AutoMigrate(&User{}); err != nil {
		t.Error(err)
	}

	user, _ := entity.NewUser{"Zé", "ze@Ze.com", "123456"}
	userDB := NewUser(db)

	err = userDB.Create(&user)
	assert.Nil(t, err)

}
