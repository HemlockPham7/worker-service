package worker

import (
	"context"
	"encoding/json"
	"testing"

	mock_bookmark "github.com/HemlockPham7/worker-service/internal/app/service/bookmark/mocks"
	"github.com/HemlockPham7/worker-service/internal/app/service/queue"
	"github.com/stretchr/testify/assert"
)

func TestHandler_Handle(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string

		setupService func(ctx context.Context) *mock_bookmark.Service

		inputMessage []byte

		expected error
	}{
		{
			name: "successfully to handle message",

			setupService: func(ctx context.Context) *mock_bookmark.Service {
				mockService := mock_bookmark.NewService(t)
				mockService.On("CreateBatchBookmarks", ctx, "user-123", setupImportBookmarkInput()).Return(nil)
				return mockService
			},

			inputMessage: setupInputMessage(),

			expected: nil,
		},
		{
			name: "There is an error",

			setupService: func(ctx context.Context) *mock_bookmark.Service {
				mockService := mock_bookmark.NewService(t)
				mockService.On("CreateBatchBookmarks", ctx, "user-123", setupImportBookmarkInput()).Return(assert.AnError)
				return mockService
			},

			inputMessage: setupInputMessage(),

			expected: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			ctx := t.Context()

			mockService := tc.setupService(ctx)
			mockHandler := NewHandler(mockService)

			err := mockHandler.Handle(ctx, tc.inputMessage)

			assert.Equal(t, tc.expected, err)
		})
	}
}

func setupInputMessage() []byte {
	bookmarkList := setupImportBookmarkInput()
	input := &queue.ImportMessage{
		UID:       "user-123",
		Bookmarks: bookmarkList,
	}
	message, err := json.Marshal(input)
	if err != nil {
		return []byte("{}")
	}
	return message
}

func setupImportBookmarkInput() []*queue.ImportBookmarkInput {
	return []*queue.ImportBookmarkInput{
		{
			Description: "Bookmark Example",
			URL:         "https://www.example.com",
		},
	}
}
