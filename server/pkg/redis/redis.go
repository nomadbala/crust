package redis

import (
	"context"
	"fmt"
	"github.com/nomadbala/crust/server/db/postgres/sqlc"
	"github.com/nomadbala/crust/server/pkg/log"
	"github.com/redis/go-redis/v9"
	"github.com/vmihailenco/msgpack/v5" // Импортируем MessagePack
	"go.uber.org/zap"
	"os"
	"time"
)

var ctx context.Context
var rdb *redis.Client

func ConfigureRedisClient(ctx_2 context.Context) {
	redisUrl := os.Getenv("TEST_REDIS_URL")
	if redisUrl == "" {
		log.Logger.Error(fmt.Sprintf("error occurred while getting .env redis_url variable: %s", redisUrl))
		return // Добавьте return, чтобы избежать дальнейших ошибок
	}

	redisOptions, err := redis.ParseURL(redisUrl)
	if err != nil {
		log.Logger.Error("error occurred while parsing .env redis_url variable", zap.Error(err))
		return // Добавьте return, чтобы избежать дальнейших ошибок
	}

	rdb = redis.NewClient(redisOptions)
	ctx = ctx_2
}

func GetPostFromCache(postID string) (sqlc.Post, error) {
	data, err := rdb.Get(ctx, postID).Result()
	if err != nil {
		return sqlc.Post{}, err // Возвращаем пустой sqlc.Post и ошибку
	}

	var postData sqlc.Post
	if err := msgpack.Unmarshal([]byte(data), &postData); err != nil {
		return sqlc.Post{}, fmt.Errorf("error unmarshalling post data: %v", err)
	}

	return postData, nil // Возвращаем десериализованный пост
}

func CachePost(postID string, postData sqlc.Post) error {
	packedData, err := msgpack.Marshal(postData) // Сериализация в MessagePack
	if err != nil {
		return err
	}

	return rdb.Set(ctx, postID, packedData, 15*time.Minute).Err() // Сохранение в Redis
}

func GetPopularPostsIDsFromCache(page, pageSize int) ([]string, error) {
	start := (page - 1) * pageSize
	end := start + pageSize - 1
	return rdb.LRange(ctx, "popular_posts", int64(start), int64(end)).Result()
}

func CachePopularPostsIDs(postIDs []string) error {
	return rdb.RPush(ctx, "popular_posts", postIDs).Err()
}

func ClearPopularPostsCache() error {
	return rdb.Del(ctx, "popular_posts").Err()
}
