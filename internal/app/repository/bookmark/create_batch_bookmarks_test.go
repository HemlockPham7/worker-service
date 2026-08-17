package bookmark

import (
	"testing"

	"github.com/HemlockPham7/common-libs/pkg/dbutils"
	"github.com/HemlockPham7/worker-service/internal/app/model"
	"github.com/HemlockPham7/worker-service/internal/integration_test/data/fixtures"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestBookmarkRepository_CreateBatchBookmarks(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string

		setupMock func(t *testing.T) *gorm.DB

		inputBookmarks []*model.Bookmark

		expectedError error
	}{
		{
			name: "success",
			setupMock: func(t *testing.T) *gorm.DB {
				return fixtures.NewFixture(t, &fixtures.BookmarkCommonTestDB{})
			},
			inputBookmarks: []*model.Bookmark{
				{
					Base:        fixtures.GetTestBase("a1b2c3d4-e5f6-7890-abcd-ef0000000077"),
					UserID:      "4d9326d6-980c-4c62-9709-dbc70a82cbfe",
					URL:         "https://example.com/newbookmark1",
					Description: "New bookmark 1 for Test User 1",
					Code:        "New-Bookmark-1",
				},
				{
					Base:        fixtures.GetTestBase("b1c2d3e4-f5g6-7890-abcd-ef0000000088"),
					UserID:      "4d9326d6-980c-4c62-9709-dbc70a82cbfe",
					URL:         "https://example.com/newbookmark2",
					Description: "New bookmark 2 for Test User 1",
					Code:        "New-Bookmark-2",
				},
			},
			expectedError: nil,
		},
		{
			name: "failed due to duplicate code",
			setupMock: func(t *testing.T) *gorm.DB {
				return fixtures.NewFixture(t, &fixtures.BookmarkCommonTestDB{})
			},
			inputBookmarks: []*model.Bookmark{
				{
					Base:        fixtures.GetTestBase("a1b2c3d4-e5f6-7890-abcd-ef0000000077"),
					UserID:      "4d9326d6-980c-4c62-9709-dbc70a82cbfe",
					URL:         "https://example.com/newbookmark1",
					Description: "New bookmark 1 for Test User 1",
					Code:        "duplicate-code",
				},
				{
					Base:        fixtures.GetTestBase("b1c2d3e4-f5g6-7890-abcd-ef0000000088"),
					UserID:      "4d9326d6-980c-4c62-9709-dbc70a82cbfe",
					URL:         "https://example.com/newbookmark2",
					Description: "New bookmark 2 for Test User 1",
					Code:        "duplicate-code",
				},
			},
			expectedError: dbutils.ErrDuplicationType,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			ctx := t.Context()

			dbMock := tc.setupMock(t)
			repo := NewRepository(dbMock)

			err := repo.CreateBatchBookmarks(ctx, tc.inputBookmarks)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}
