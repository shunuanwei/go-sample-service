package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"sync"
	"time"
)

var Client = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "root", // no password set
	DB:       12,     // use default DB
})

//加锁
func MyLock() bool {
	nx := Client.SetNX("distributed_lock", 1, time.Second*10)
	result, err := nx.Result()
	return result && err == nil
}

//解锁
func UnMyLock() {
	Client.Del("distributed_lock")
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for {
				if MyLock() {
					fmt.Println(i)
					UnMyLock()
					return
				}
			}
		}(i)
	}
	wg.Wait()
}

// Redis 分布式锁
func incr() {
	//client := redis.NewClient(&redis.Options{
	//	Addr:     "localhost:6379",
	//	Password: "root", // no password set
	//	DB:       12,     // use default DB
	//})

	var lockKey = "counter_lock"
	var counterKey = "counter"

	// lock
	resp := Client.SetNX(lockKey, 1, time.Second*5)
	lockSuccess, err := resp.Result()

	if err != nil || !lockSuccess {
		fmt.Println(err, "lock result: ", lockSuccess)
		return
	}

	// counter ++
	getResp := Client.Get(counterKey)
	cntValue, err := getResp.Int64()
	if err == nil || err == redis.Nil {
		cntValue++
		resp := Client.Set(counterKey, cntValue, 0)
		_, err := resp.Result()
		if err != nil {
			// log err
			println("set value error!")
		}
	}
	println("current counter is ", cntValue)

	delResp := Client.Del(lockKey)
	unlockSuccess, err := delResp.Result()
	if err == nil && unlockSuccess > 0 {
		println("unlock success!")
	} else {
		println("unlock failed", err)
	}
}

func mai1n() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			incr()
		}()
	}
	wg.Wait()
}
