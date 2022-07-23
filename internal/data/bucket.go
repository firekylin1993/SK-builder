package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type RsaBucketRepo struct {
	data *Data
	log  *log.Helper
}

func NewBucketRepo(data *Data, logger log.Logger) *RsaBucketRepo {
	return &RsaBucketRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (b *RsaBucketRepo) Add(ctx context.Context, path string) error {
	// _, err := b.data.mysql.Exec("INSERT INTO bucket(id, bucket_path) VALUES(?, ?)", 1, path)
	// if err != nil {
	// 	b.log.Error("添加密钥桶路径失败", err)
	// 	return err
	// }
	return nil
}
