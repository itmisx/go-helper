package helper

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/itmisx/redisx"
)

// 唯一键
type UUID struct{}

var cli redisx.Client

// Init redis初始化连接池
func (UUID) Init(conf redisx.Config) {
	cli = redisx.New(conf)
}

// GetUniqueKey 获取唯一主键
func (UUID) GetUniqueKey() (ID string, err error) {
	ID = ""
	times := 0
	// 获取开一个redis客户端连接
	ctx := context.Background()
	for {
		// 最大尝试次数
		times++
		if times > 100 {
			return "", errors.New("获取唯一键失败")
		}

		// 获取纳秒
		now := time.Now()
		ID = strconv.FormatInt(now.UnixNano()/1000, 10) // 取微妙,并转为字符串

		// 通过redis setnx 保证id是唯一的
		key := "primary_key:" + ID

		// 判断id是否唯一
		success, err := cli.SetNX(ctx, key, 1, time.Second*1).Result()
		if !success || err != nil {
			time.Sleep(5 * time.Microsecond) // 休眠2微妙，避免最大尝试次数内获得的是同一个微妙
			continue
		}
		break
	}
	return
}
