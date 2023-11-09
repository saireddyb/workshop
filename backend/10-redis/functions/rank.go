package functions

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var leaderboardKey = "leaderboard"

type RedisMovie struct {
	Id string `bson:"id" binding:"required"`
	Rating   int `json:"points" binding:"required"`
	Rank     float64    `json:"rank"`
}

func AddMovieInRedis(movie RedisMovie){
	fmt.Print("reached here")
	member := &redis.Z{
		Score: float64(movie.Rating),
		Member: movie.Id,
	}

	pipe := redisClient.TxPipeline()
	pipe.ZAdd(context.TODO(), leaderboardKey, *member)
	pipe.ZRank(context.TODO(), leaderboardKey, movie.Id)
	pipe.Exec(context.TODO())

}
