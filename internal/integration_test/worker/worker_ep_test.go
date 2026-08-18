package worker

//
//import (
//	"context"
//	"os"
//	"testing"
//	"time"
//
//	redisPkg "github.com/HemlockPham7/common-libs/pkg/redis"
//	"github.com/HemlockPham7/common-libs/pkg/sqldb"
//	"github.com/HemlockPham7/common-libs/pkg/utils"
//	workerHdl "github.com/HemlockPham7/worker-service/internal/app/handler/worker"
//	"github.com/HemlockPham7/worker-service/internal/app/model"
//	bookmarkRepo "github.com/HemlockPham7/worker-service/internal/app/repository/bookmark"
//	cacheRepo "github.com/HemlockPham7/worker-service/internal/app/repository/cache"
//	queueRepo "github.com/HemlockPham7/worker-service/internal/app/repository/queue"
//	bookmarkSvc "github.com/HemlockPham7/worker-service/internal/app/service/bookmark"
//	"github.com/HemlockPham7/worker-service/internal/worker"
//	"github.com/stretchr/testify/assert"
//	"github.com/stretchr/testify/require"
//	"gorm.io/gorm"
//)
//
//func TestWorkerEngine_Start(t *testing.T) {
//	ctx := t.Context()
//
//	messages := []string{
//		`{"user_id":"4d9326d6-980c-4c62-9709-dbc70a82cbfe","bookmarks":[{"url":"https://example.com/newbookmark1","description":"New bookmark 1 for Test User 1"},{"url":"https://example.com/newbookmark2","description":"New bookmark 2 for Test User 1"}]}`,
//		`{"user_id":"4d9326d6-980c-4c62-9709-dbc70a82cbfe","bookmarks":[{"url":"https://example.com/newbookmark3","description":"New bookmark 3 for Test User 1"},{"url":"https://example.com/newbookmark4","description":"New bookmark 4 for Test User 1"}]}`,
//		`{"user_id":"4d9326d6-980c-4c62-9709-dbc70a82cbfe","bookmarks":[{"url":"https://example.com/newbookmark5","description":"New bookmark 5 for Test User 1"},{"url":"https://example.com/newbookmark6","description":"New bookmark 6 for Test User 1"}]}`,
//	}
//
//	testWorkerEngine, db := makeTestWorkerEngine(t, ctx, messages)
//
//	go testWorkerEngine.Start(ctx)
//	time.Sleep(1 * time.Second)
//
//	process, err := os.FindProcess(os.Getpid())
//	require.NoError(t, err)
//
//	err = process.Signal(os.Interrupt)
//	assert.NoError(t, err)
//
//	time.Sleep(10 * time.Second)
//
//	var count int64
//	err = db.WithContext(ctx).Model(&model.Bookmark{}).Count(&count).Error
//	assert.NoError(t, err)
//	assert.Equal(t, int64(6), count)
//}
//
//func makeTestWorkerEngine(t *testing.T, ctx context.Context, messages []string) (worker.Engine, *gorm.DB) {
//	mockRedisClient := redisPkg.InitMockRedis(t)
//	for _, message := range messages {
//		err := mockRedisClient.LPush(ctx, "test_queue", message).Err()
//		assert.NoError(t, err)
//	}
//
//	db := sqldb.InitMockDB(t)
//
//	sqlDB, err := db.DB()
//	assert.NoError(t, err)
//	sqlDB.SetMaxOpenConns(1)
//
//	err = db.AutoMigrate(&model.Bookmark{})
//	assert.NoError(t, err)
//
//	mockCodeGenerator := utils.NewGenCode()
//
//	mockQueueRepo := queueRepo.NewRedisQueue(mockRedisClient, "test_queue")
//	mockCacheRepo := cacheRepo.NewRedisDB(mockRedisClient)
//
//	bookmarkRepository := bookmarkRepo.NewRepository(db)
//	bookmarkService := bookmarkSvc.NewService(bookmarkRepository, mockCacheRepo, mockCodeGenerator)
//
//	workerHandler := workerHdl.NewHandler(bookmarkService)
//	workerEngine := worker.NewEngine(mockQueueRepo, workerHandler)
//
//	return workerEngine, db
//}
