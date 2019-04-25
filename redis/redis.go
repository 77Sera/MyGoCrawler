package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

func standalon_redis_test() {

	client := redis.NewClient(&redis.Options{
		Addr:		"localhost:6379",
		Password: 	"",
		DB: 		0,
	})

	defer client.Close()

	client.Set("test_key","test_value123", 0)

	time.Sleep( 10 * time.Second )

	value, _ := client.Get("test_key").Result()

	fmt.Print(value)




	//pong, err := client.Ping().Result()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Println(pong)

}