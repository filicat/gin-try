package global

import (
	"context"
	"filicat/jassue-gin-try/utils"
	"github.com/go-redis/redis/v8"
	"time"
)

type Interface interface {
	Get() bool
	Block(seconds int64) bool
	Release() bool
	ForceRelease()
}

type lock struct {
	context context.Context
	name    string
	owner   string
	seconds int64 // 有效期
}

// 释放锁 Lua 脚本，防止任何客户端都能解锁
const releaseLockLuaScript = `
if redis.call("get",KEYS[1]) == ARGV[1] then
    return redis.call("del",KEYS[1])
else
    return 0
end
`

// Lock 生成锁实例
func Lock(name string, seconds int64) Interface {
	return &lock{
		context: context.Background(),
		name:    name,
		owner:   utils.RandString(16),
		seconds: seconds,
	}
}

// Get 获取锁
func (l *lock) Get() bool {
	return App.Redis.SetNX(l.context, l.name, l.owner, time.Duration(l.seconds)*time.Second).Val()
}

// Block 阻塞一段时间，尝试获取锁
func (l *lock) Block(seconds int64) bool {
	starting := time.Now().Unix()
	for {
		if !l.Get() {
			time.Sleep(time.Duration(1) * time.Second)
			if time.Now().Unix()-seconds > starting {
				return false
			}
		} else {
			return true
		}
	}
}

// Release 释放锁
func (l *lock) Release() bool {
	luaScript := redis.NewScript(releaseLockLuaScript)
	result := luaScript.Run(l.context, App.Redis, []string{l.name}, l.owner).Val().(int64)
	return result != 0
}

func (l *lock) ForceRelease() {
	App.Redis.Del(l.context, l.name).Val()
}
