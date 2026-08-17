package bookmark

import (
	"context"
	"fmt"
	"testing"

	"github.com/HemlockPham7/common-libs/pkg/utils/mocks"
	"github.com/HemlockPham7/worker-service/internal/app/model"
	mock_bookmark "github.com/HemlockPham7/worker-service/internal/app/repository/bookmark/mocks"
	mock_cache "github.com/HemlockPham7/worker-service/internal/app/repository/cache/mocks"
	"github.com/HemlockPham7/worker-service/internal/app/service/queue"
	"github.com/stretchr/testify/assert"
)

func TestBookmarkService_CreateBatchBookmarks(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string

		setupMockCacheRepo     func(ctx context.Context) *mock_cache.DB
		setupMockBookmarkRepo  func(ctx context.Context) *mock_bookmark.Repository
		setupMockCodeGenerator func(ctx context.Context) *mocks.GenCode

		expectedError error
	}{
		{
			name: "success",

			setupMockCacheRepo: func(ctx context.Context) *mock_cache.DB {
				mockCache := mock_cache.NewDB(t)
				mockCache.On("DeleteCache", ctx, fmt.Sprintf(getBookmarksCacheGroupKeyFormat, setupInputUserID())).Return(nil)
				return mockCache
			},

			setupMockCodeGenerator: func(ctx context.Context) *mocks.GenCode {
				mockCodeGenerator := mocks.NewGenCode(t)
				mockCodeGenerator.On("GenerateCode", codeLength).Return("example1", nil)
				return mockCodeGenerator
			},

			setupMockBookmarkRepo: func(ctx context.Context) *mock_bookmark.Repository {
				mockBookmark := mock_bookmark.NewRepository(t)
				mockBookmark.On("CreateBatchBookmarks", ctx, setupMockModelBookmark()).Return(nil)
				return mockBookmark
			},

			expectedError: nil,
		},
		{
			name: "delete cache fail",

			setupMockCacheRepo: func(ctx context.Context) *mock_cache.DB {
				mockCache := mock_cache.NewDB(t)
				mockCache.On("DeleteCache", ctx, fmt.Sprintf(getBookmarksCacheGroupKeyFormat, setupInputUserID())).Return(assert.AnError)
				return mockCache
			},

			setupMockCodeGenerator: func(ctx context.Context) *mocks.GenCode {
				return mocks.NewGenCode(t)
			},

			setupMockBookmarkRepo: func(ctx context.Context) *mock_bookmark.Repository {
				return mock_bookmark.NewRepository(t)
			},

			expectedError: assert.AnError,
		},
		{
			name: "failed to generate code",

			setupMockCacheRepo: func(ctx context.Context) *mock_cache.DB {
				mockCache := mock_cache.NewDB(t)
				mockCache.On("DeleteCache", ctx, fmt.Sprintf(getBookmarksCacheGroupKeyFormat, setupInputUserID())).Return(nil)
				return mockCache
			},

			setupMockCodeGenerator: func(ctx context.Context) *mocks.GenCode {
				mockCodeGenerator := mocks.NewGenCode(t)
				mockCodeGenerator.On("GenerateCode", codeLength).Return("", assert.AnError)
				return mockCodeGenerator
			},

			setupMockBookmarkRepo: func(ctx context.Context) *mock_bookmark.Repository {
				return mock_bookmark.NewRepository(t)
			},

			expectedError: assert.AnError,
		},
		{
			name: "failed to create batch bookmarks",

			setupMockCacheRepo: func(ctx context.Context) *mock_cache.DB {
				mockCache := mock_cache.NewDB(t)
				mockCache.On("DeleteCache", ctx, fmt.Sprintf(getBookmarksCacheGroupKeyFormat, setupInputUserID())).Return(nil)
				return mockCache
			},

			setupMockCodeGenerator: func(ctx context.Context) *mocks.GenCode {
				mockCodeGenerator := mocks.NewGenCode(t)
				mockCodeGenerator.On("GenerateCode", codeLength).Return("example1", nil)
				return mockCodeGenerator
			},

			setupMockBookmarkRepo: func(ctx context.Context) *mock_bookmark.Repository {
				mockBookmark := mock_bookmark.NewRepository(t)
				mockBookmark.On("CreateBatchBookmarks", ctx, setupMockModelBookmark()).Return(assert.AnError)
				return mockBookmark
			},

			expectedError: assert.AnError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			ctx := t.Context()

			mockCacheRepo := tc.setupMockCacheRepo(ctx)
			mockCodeGenerator := tc.setupMockCodeGenerator(ctx)
			mockBookmarkRepo := tc.setupMockBookmarkRepo(ctx)

			mockService := NewService(mockBookmarkRepo, mockCacheRepo, mockCodeGenerator)

			err := mockService.CreateBatchBookmarks(ctx, setupInputUserID(), setupInputBookmarksList())

			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func setupInputUserID() string {
	return "d7c13097-67a7-4eae-a60e-0b9b533b3636"
}

func setupInputBookmarksList() []*queue.ImportBookmarkInput {
	return []*queue.ImportBookmarkInput{
		{
			Description: "Bookmark Example",
			URL:         "https://www.example.com",
		},
	}
}

func setupMockModelBookmark() []*model.Bookmark {
	return []*model.Bookmark{
		{
			Description: "Bookmark Example",
			URL:         "https://www.example.com",
			Code:        "example1",
			UserID:      "d7c13097-67a7-4eae-a60e-0b9b533b3636",
		},
	}
}
