package main

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
	"time"
)

func initRedis() (conn redis.Conn, err error) {
	conn, err = redis.Dial("tcp", "192.168.247.133:6379")
	if err != nil {
		fmt.Printf("connect redis failed, err: %s\n", err)
		return
	}
	return
}

func testSetGet(conn redis.Conn) {
	key := "aaa"
	_, err := conn.Do("set", key, "hahaha")
	if err != nil {
		fmt.Printf("set redis failed, err: %s\n", err)
		return
	}

	data, err := redis.String(conn.Do("get", key))
	if err != nil {
		fmt.Printf("get redis failed, err: %s\n", err)
		return
	}
	fmt.Printf("key: %s value: %s\n", key, data)
}

func testHSetGet(conn redis.Conn) {
	key := "aaa"
	_, err := conn.Do("hset", "books", key, "hahaha")
	if err != nil {
		fmt.Printf("set redis failed, err: %s\n", err)
		return
	}

	data, err := redis.String(conn.Do("hget", "books", key))
	if err != nil {
		fmt.Printf("get redis failed, err: %s\n", err)
		return
	}
	fmt.Printf("key: %s value: %s\n", key, data)
}

func testMSetGet(conn redis.Conn) {
	key := "aaa"
	key1 := "bbb"
	_, err := conn.Do("mset", key, "hahaha", key1, "hehehe")
	if err != nil {
		fmt.Printf("set redis failed, err: %s\n", err)
		return
	}

	data, err := redis.Strings(conn.Do("mget", key, key1))
	if err != nil {
		fmt.Printf("get redis failed, err: %s\n", err)
		return
	}
	for _, val := range data {
		fmt.Printf("value: %s\n", val)
	}

}

func testExpireKey(conn redis.Conn) {
	key := "aaa"
	_, err := conn.Do("expire", key, "30")
	if err != nil {
		fmt.Printf("set redis failed, err: %s\n", err)
		return
	}

	data, err := redis.String(conn.Do("get", key))
	if err != nil {
		fmt.Printf("get redis failed, err: %s\n", err)
		return
	}
	fmt.Printf("key: %s value: %s\n", key, data)

}

func newPool(server, password string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     64,
		MaxActive:   1000,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}

			if len(password) > 0 {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			}

			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}

func testRedisPool()  {
	pool := newPool("192.168.247.133:6379", "")

	// 获取一个连接
	conn := pool.Get()
	conn.Do("set", "qwe", "1234567890")
	val, err := redis.String(conn.Do("get", "qwe"))
	fmt.Printf("val: %s , err: %v\n", val, err)

	// 归还连接至连接池
	conn.Close()
}

func main() {
	//conn, err := initRedis()
	//if err != nil {
	//	return
	//}
	//testSetGet(conn)
	//testHSetGet(conn)
	//testMSetGet(conn)
	//testExpireKey(conn)
	testRedisPool()
}
