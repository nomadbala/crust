package service

import (
	"encoding/hex"
	"github.com/google/uuid"
	"github.com/nomadbala/crust/server/db/postgres/sqlc"
	"github.com/nomadbala/crust/server/internal/domain/post"
	"github.com/nomadbala/crust/server/pkg/resend"
	"log"
)

type PostsService struct {
	repository post.Repository
	cache      post.Cache
}

func NewPostsService(repository post.Repository, cache post.Cache) *PostsService {
	return &PostsService{repository, cache}
}

func (p PostsService) List() ([]*post.Response, error) {
	posts, err := p.repository.List()
	if err != nil {
		return nil, err
	}

	err = resend.SendResendMessage("nurkenkenes7@gmail.com", "nigga")
	if err != nil {
		return nil, err
	}

	err = resend.SendResendMessage("mounjoyer@gmail.com", "nigga")
	if err != nil {
		return nil, err
	}

	responses := post.ConvertEntitiesToResponses(posts)

	return responses, nil
}

func (p PostsService) Get(id uuid.UUID) (*post.Response, error) {
	cachedPost, err := p.cache.Get(id)
	if err == nil {
		return &post.Response{
			Id:        cachedPost.ID,
			UserId:    cachedPost.UserID,
			Content:   cachedPost.Content + " Redis",
			CreatedAt: cachedPost.CreatedAt,
		}, nil
	}

	savedPost, err := p.repository.Get(id)
	if err != nil {
		return nil, err
	}

	go func() {
		if cacheErr := p.cache.Set(id, savedPost); cacheErr != nil {
			log.Printf("Error caching post: %v", cacheErr)
		}
	}()

	return post.ConvertEntityToResponse(savedPost), nil
}

func (p PostsService) GetPopular(limit, offset int) ([]*post.Response, error) {
	params := sqlc.GetPopularPostsParams{Limit: int32(limit), Offset: int32(offset)}

	posts, err := p.repository.GetPopular(params)
	if err != nil {
		return nil, err
	}

	responses := post.ConvertEntitiesToResponses(posts)

	return responses, nil
}

func (p PostsService) Create(params sqlc.CreatePostParams) (*post.Response, error) {
	savedPost, err := p.repository.Create(params)
	if err != nil {
		return nil, err
	}

	response := post.ConvertEntityToResponse(savedPost)

	return response, nil
}

func encodeUUID(src [16]byte) string {
	var buf [36]byte

	hex.Encode(buf[0:8], src[:4])
	buf[8] = '-'
	hex.Encode(buf[9:13], src[4:6])
	buf[13] = '-'
	hex.Encode(buf[14:18], src[6:8])
	buf[18] = '-'
	hex.Encode(buf[19:23], src[8:10])
	buf[23] = '-'
	hex.Encode(buf[24:], src[10:])

	return string(buf[:])
}
