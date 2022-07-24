package data

import (
	"SK-builder/internal/infrastructure/db"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type RsaBucketRepo struct {
	data *db.Data
	log  *log.Helper
}

func NewBucketRepo(data *db.Data, logger log.Logger) *RsaBucketRepo {
	return &RsaBucketRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (b *RsaBucketRepo) Add(ctx context.Context, snowId int64) error {
	// _, err := b.data.mysql.Exec("INSERT INTO bucket(id, bucket_path) VALUES(?, ?)", 1, path)
	// if err != nil {
	// 	b.log.Error("添加密钥桶路径失败", err)
	// 	return err
	// }
	return nil
}

func (b *RsaBucketRepo) GetAll(ctx context.Context) (int32, error) {
	// TODO 获取密钥桶所有数据
	return 0, nil
}

func (b *RsaBucketRepo) Delete(ctx context.Context, snowId int64) error {
	return nil
}
