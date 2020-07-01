package main

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/go-redis/redis/v7"
	"log"
	"os"
	"time"
)

func exit(err *error) {
	if *err != nil {
		log.Printf("exited with error: %s", (*err).Error())
		os.Exit(1)
	} else {
		log.Println("exited")
	}
}

func main() {
	var err error
	defer exit(&err)

	var opts *redis.Options
	if opts, err = redis.ParseURL(os.Getenv("REDIS_URL")); err != nil {
		return
	}
	opts.PoolSize = 1

	r := redis.NewClient(opts)

	buf := make([]byte, 128, 128)

	log.Println("started")

	count := int64(0)

	for {
		count++
		for i := 0; i < 100; i++ {
			if _, err = rand.Read(buf); err != nil {
				return
			}
			key := "r-s-t-" + hex.EncodeToString(buf[0:16])
			val := hex.EncodeToString(buf)
			if err = r.Set(key, val, time.Second*5).Err(); err != nil {
				return
			}
			if err = r.Del(key).Err(); err != nil {
				return
			}
		}
		log.Printf("round: %012d", count)
		time.Sleep(time.Second)
	}
}
