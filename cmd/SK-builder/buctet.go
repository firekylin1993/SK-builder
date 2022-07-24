package main

import (
	"SK-builder/internal/infrastructure/mykey"
	"context"
	"errors"
	"sync"
	"sync/atomic"

	"github.com/go-kratos/kratos/v2/log"
)

func newBucket(ctx context.Context, b *mykey.RsaBucket, logger log.Logger) error {
	log.NewHelper(logger).Info("初始化密钥桶...")
	defer log.NewHelper(logger).Info("初始化密钥桶完成")

	pk, err := b.Generate(ctx)
	if err != nil {
		return err
	}

	i, err := b.BucketDb.GetAll(ctx)
	if err != nil {
		return err
	}

	var diff int32 = 0
	diff = b.Limit - i
	if diff <= 0 {
		return nil
	}

	c := make(chan struct{}, b.Limit)
	for i := 0; i < int(diff); i++ {
		c <- struct{}{}
	}
	var keys int32
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
		Loop:
			for {
				select {
				case <-c:
					path, snowIDInt64, err := b.Fill(ctx, pk) // 生成密钥文件,且放入密钥桶
					if err != nil {
						log.NewHelper(logger).Errorf("密钥文件入库失败：%s，路径：%s", err.Error(), path)
						b.Remove(ctx, path) // 删除密钥桶
						break
					}

					err = b.BucketDb.Add(ctx, snowIDInt64) // 将密钥桶路径添加到数据库
					if err != nil {
						log.NewHelper(logger).Errorf("密钥文件入库失败：%s，路径：%s", err.Error(), path)
						b.Remove(ctx, path) // 删除密钥桶
					}
					atomic.AddInt32(&keys, 1)
				default:
					break Loop
				}
			}
		}()
	}
	wg.Wait()
	close(c)
	if (keys+diff) != b.Limit {
		log.NewHelper(logger).Errorf("当前密钥桶密钥对数%d\n", keys)
		return errors.New("密钥桶数量不一致")
	}
	return nil
}
