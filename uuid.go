// 基于redis
package helper

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/itmisx/redisx"
)

// 唯一键
type UUID struct {
	cli redisx.Client
}

// Init redis初始化连接池
func (uuid *UUID) Init(conf redisx.Config) {
	uuid.cli = redisx.New(conf)
}

func (uuid *UUID) Int64() (ID int64, err error) {
	return uuid.getUniqueKey()
}

func (uuid *UUID) String() (ID string, err error) {
	idInt, err := uuid.getUniqueKey()
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(idInt, 10), nil
}

// getUniqueKey 获取唯一主键
func (uuid *UUID) getUniqueKey() (ID int64, err error) {
	times := 0
	for {
		// 最大尝试次数
		times++
		if times > 100 {
			return 0, errors.New("获取唯一键失败")
		}

		// 获取纳秒
		now := time.Now()
		ID = now.UnixNano() / 1000

		// 通过redis setnx 保证id是唯一的
		key := "primary_key:" + strconv.FormatInt(ID, 10)

		// 判断id是否唯一
		success, err := uuid.cli.SetNX(context.Background(), key, 1, time.Second*1).Result()
		if !success || err != nil {
			// 休眠1毫秒，避免最大尝试次数内获得的是同一个微妙
			time.Sleep(time.Microsecond)
			continue
		}
		break
	}
	return
}
