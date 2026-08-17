package fixtures

import (
	"github.com/HemlockPham7/worker-service/internal/app/model"
	"gorm.io/gorm"
)

type BookmarkCommonTestDB struct {
	base
}

func (b *BookmarkCommonTestDB) Migrate() error {
	return b.db.AutoMigrate(&model.Bookmark{})
}

func (b *BookmarkCommonTestDB) GenerateData() error {
	db := b.db.Session(&gorm.Session{SkipHooks: true})

	bookmarks := []*model.Bookmark{
		{
			Base:        GetTestBase("d7c13097-67a7-4eae-a60e-0b9b533b7bd6"),
			Description: "Bookmark 1",
			URL:         "https://www.google.com",
			Code:        "bookmark1",
			UserID:      "d7c13097-67a7-4eae-a60e-0b9b533b7bd4",
		},
		{
			Base:        GetTestBase("d7c13097-67a7-4eae-a60e-0b9b533b7bd7"),
			Description: "Bookmark 2",
			URL:         "https://www.google.com",
			Code:        "bookmark2",
			UserID:      "d7c13097-67a7-4eae-a60e-0b9b533b7bd5",
		},
	}

	return db.CreateInBatches(bookmarks, 10).Error
}
