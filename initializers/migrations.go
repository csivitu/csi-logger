package initializers

import (
	"fmt"

	"github.com/csivitu/csi-logger/models"
)

func AutoMigrate() {
	fmt.Println("\nStarting Migrations...")
	DB.AutoMigrate(
		&models.Log{},

	)
	fmt.Println("Migrations Finished!")
}
