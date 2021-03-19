package lib

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

type Locker struct {
	key string
	expire time.Duration
	unlock bool
	incrScript *redis.Script
}

const  incrLua = `
if redis.call('get',KEYS[1]) == ARGV[1] then 
	return redis.call('expire',KEYS[1],ARGV[2]) 
else 
	return '0' 
end`

func NewLocker(key string) *Locker {
	return &Locker{
		key: key,
		expire: 30 * time.Second,
	}
}

func NewLockerWithTTL(key string,expire time.Duration) *Locker {
	if expire.Seconds() <= 0 {
		panic("error expire")
	}

	return &Locker{
		key:    key,
		expire: expire,
		incrScript: redis.NewScript(incrLua),
	}
}

func (l *Locker) Lock() *Locker {
	boolCmd := redisClient.SetNX(context.Background(), l.key, "1", l.expire)
	if ok, err := boolCmd.Result(); err != nil || !ok {
		panic(fmt.Sprintf("lock error with key:%s", l.key))
	}
	l.expandLockTime()

	return l
}

func (l *Locker) expandLockTime() {
	sleepTime := l.expire.Seconds() * 2 / 3
	go func() {
		for {
			time.Sleep(time.Second * time.Duration(sleepTime))
			if l.unlock {
				break
			}
			l.resetExpire()
		}
	}()
}

// 重新设置过期时间
func (l *Locker) resetExpire() {
	cmd := l.incrScript.Run(context.Background(),redisClient,[]string{l.key},1,l.expire.Seconds())
	v,err := cmd.Result()
	log.Printf("key=%s,续期结果：%v,%v\n",l.key,err,v)
}

func (l *Locker) UnLock() {
	l.unlock = true
	redisClient.Del(context.Background(),l.key)
}
