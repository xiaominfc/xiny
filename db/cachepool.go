package db

import (
    "github.com/go-redis/redis"
)

type CacheConn struct{
    Addr string
    Password string
    DB    int
    RedisClient *redis.Client
}



func NewCacheConn(addr,password string,dbIndex int) *CacheConn {    
    return &CacheConn{Addr:addr, Password:password,DB:dbIndex}
}

func (this *CacheConn) Init() bool {
    if this.RedisClient != nil {
        return true
    } else {
        this.RedisClient = redis.NewClient(&redis.Options{
            Addr:     this.Addr,
            Password: this.Password, // no password set
            DB:       this.DB,  // use default DB
        })
        pong, err := this.RedisClient.Ping().Result()
        if err != nil {
            return false
        }
        println(pong)
        return true
    }
    return false
}


func (this *CacheConn) HGet(key,field string) string {
    if this.Init() {
        result,err := this.RedisClient.HGet(key,field).Result()
        if err != nil {
            return ""
        }

        return result
    }
    return ""
}

func (this *CacheConn) HSet(key,field,value string) bool{
    if this.Init() {
        result,err := this.RedisClient.HSet(key,field,value).Result()
        if err != nil {
            return false
        }
        return result
    }
    return false
}

func (this *CacheConn) HDel(key,field string) int64{
    if this.Init() {
        result,err := this.RedisClient.HDel(key,field).Result()
        if err != nil {
            return 0
        }

        return result
    }
    return 0
}

