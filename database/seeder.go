package database

import (
	"fmt"

	"github.com/koficypher/simple_server/models"
	"gorm.io/gorm"
)

var articles = []models.Article{
	{
		Title:   "A stitch in time",
		Author:  "Robert Orwell",
		Content: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Curabitur ac lectus vel libero lacinia eleifend. Vestibulum viverra sed nibh eu porta. Proin nec mauris nec metus ultricies feugiat in vitae risus. Vestibulum pulvinar nisi et est sagittis pulvinar.",
	},
	{
		Title:   "Bambara",
		Author:  "Anita Smith",
		Content: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Curabitur ac lectus vel libero lacinia eleifend. Vestibulum viverra sed nibh eu porta. Proin nec mauris nec metus ultricies feugiat in vitae risus. Vestibulum pulvinar nisi et est sagittis pulvinar.",
	},
	{
		Title:   "Knife Tool",
		Author:  "Kaugh Lemar",
		Content: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Curabitur ac lectus vel libero lacinia eleifend. Vestibulum viverra sed nibh eu porta. Proin nec mauris nec metus ultricies feugiat in vitae risus. Vestibulum pulvinar nisi et est sagittis pulvinar.",
	},
}

func RunSeeder(db *gorm.DB) error {
	fmt.Println("Seeding Tables ...")
	if err := db.Create(&articles); err.Error != nil {
		return err.Error
	}

	fmt.Println("Done Seeding Tables ...")

	return nil
}
