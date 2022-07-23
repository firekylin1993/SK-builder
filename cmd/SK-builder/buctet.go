package main

import (
	"SK-builder/internal/infrastructure/mykey"
	"context"
	"sync"

	"github.com/go-kratos/kratos/v2/log"
)

func newBucket(ctx context.Context, b *mykey.RsaBucket, logger log.Logger) error {
	log.NewHelper(logger).Info("初始化密钥桶...")
	defer log.NewHelper(logger).Info("初始化密钥桶完成")

	pk, err := b.Generate(ctx)
	if err != nil {
		return err
	}

	c := make(chan struct{}, b.Limit)
	for i := 0; i < int(b.Limit); i++ {
		c <- struct{}{}
	}
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
		Loop:
			for {
				select {
				case <-c:
					path, err := b.Fill(ctx, pk) // 生成密钥文件,且放入密钥桶
					if err != nil {
						log.NewHelper(logger).Errorf("密钥文件入库失败：%s，路径：%s", err.Error(), path)
						b.Remove(ctx, path) // 删除密钥桶
					}

					err = b.BucketDb.Add(ctx, path) // 将密钥桶路径添加到数据库
					if err != nil {
						log.NewHelper(logger).Errorf("密钥文件入库失败：%s，路径：%s", err.Error(), path)
						b.Remove(ctx, path) // 删除密钥桶
					}
				default:
					wg.Done()
					break Loop
				}
			}
		}()
	}
	wg.Wait()
	close(c)
	return nil
}
