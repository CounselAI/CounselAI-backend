package seeds

import (
	"github.com/BearTS/go-gin-monolith/database/tables"
	"gorm.io/gorm"
)

func Reports(db *gorm.DB) error {
	err := db.Create(&tables.Reports{
		PID:        "rep_89b2f151110644cb84b0912bb8ce7885",
		Url:        "https://cdn.discordapp.com/attachments/1113692548825894913/1114661078408908860/test6.pdf",
		UserID:     "usr_1efec19b64bb42d68adb0852a628f809",
		IsSandbox:  false,
		IsDeleted:  false,
		IsArchived: false,
	}).Error

	if err != nil {
		return err
	}

	err = db.Create(&tables.Reports{
		PID:        "rep_89b2f151110644cb84b0912bb8ce7886",
		Url:        "https://cdn.discordapp.com/attachments/1113692548825894913/1114661078798975016/test5.pdf",
		UserID:     "usr_1efec19b64bb42d68adb0852a628f809",
		IsSandbox:  false,
		IsDeleted:  false,
		IsArchived: false,
	}).Error

	if err != nil {
		return err
	}

	err = db.Create(&tables.Reports{
		PID:        "rep_89b2f151110644cb84b0912bb8ce7887",
		Url:        "https://cdn.discordapp.com/attachments/1113692548825894913/1114661079109337190/test4.pdf",
		UserID:     "usr_1efec19b64bb42d68adb0852a628f809",
		IsSandbox:  false,
		IsDeleted:  false,
		IsArchived: true,
	}).Error

	if err != nil {
		return err
	}

	return err
}
