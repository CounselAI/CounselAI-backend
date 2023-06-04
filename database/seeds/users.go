package seeds

import (
	"time"

	"github.com/BearTS/go-gin-monolith/database/tables"
	"gorm.io/gorm"
)

func Users(db *gorm.DB) error {
	// Seed 1
	err := db.Create(&tables.Users{
		PID:       "usr_gh67fhgCU123jkhfdsjkdfhkds",
		IsSandbox: false,
		IsDeleted: false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}).Error

	if err != nil {
		return err
	}

	// Seed 2
	err = db.Create(&tables.Users{
		PID:          "usr_1efec19b64bb42d68adb0852a628f808",
		Email:        "anuj.parihar2021@vitstudent.ac.in",
		Name:         "Anuj Parihar",
		MobileNumber: "7828310234",
		IsDeleted:    false,
		IsSandbox:    false,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}).Error

	if err != nil {
		return err
	}

	// Seed 3
	err = db.Create(&tables.Users{
		PID:          "usr_1efec19b64bb42d68adb0852a628f809",
		Email:        "harshaljanjani@gmail.com",
		Name:         "Harshal Janjani",
		MobileNumber: "7828310235",
		IsDeleted:    false,
		IsSandbox:    false,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}).Error

	if err != nil {
		return err
	}

	return err
}
