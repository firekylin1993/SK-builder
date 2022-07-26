package myrsa

import (
	"SK-builder-demo/internal/conf"
	"SK-builder-demo/internal/data/mysnowflake"
	"SK-builder-demo/internal/db"
	"context"
	"crypto/rsa"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"os"
	"path"
	"strconv"
)

type rsaBucketRepo struct {
	data *db.Data
	log  *log.Helper
}

func NewBucketRepo(data *db.Data, logger log.Logger) *rsaBucketRepo {
	return &rsaBucketRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

type RsaBucket struct {
	Path     string
	Limit    int32
	RsaKey   *RsaKey
	Repo     *rsaBucketRepo
	SnowNode *mysnowflake.SnowNode
}

func NewRsaBucket(c *conf.Server, r *RsaKey, sn *mysnowflake.SnowNode, repo *rsaBucketRepo) *RsaBucket {
	return &RsaBucket{
		Path:     c.RsaBucket.Path,
		Limit:    c.RsaBucket.Limit,
		RsaKey:   r,
		Repo:     repo,
		SnowNode: sn,
	}
}

func (b *RsaBucket) Generate(ctx context.Context) (*rsa.PrivateKey, error) {
	return b.RsaKey.GenerateKey() //生成私钥对象
}

func (b *RsaBucket) Fill(ctx context.Context, pk *rsa.PrivateKey) (string, int64, error) {
	snowIDBtye, snowIDInt64 := b.SnowNode.GetID() //获取雪花算法的ID
	//组装桶密钥路径
	path := path.Join(b.Path, strconv.Itoa(int(snowIDInt64)))
	if os.IsNotExist(os.MkdirAll(path, 0755)) {
		return "", 0, errors.New("mkdir error")
	}

	err := b.RsaKey.GetKey(pk, path) // 生成私钥文件
	if err != nil {
		return "", 0, err
	}

	err = b.RsaKey.GetPublicKey(pk, snowIDBtye, path) // 生成公钥文件
	if err != nil {
		return "", 0, err
	}

	return path, snowIDInt64, nil
}

func (b *RsaBucket) Remove(ctx context.Context, path string) error {
	return os.RemoveAll(path) //删除桶密钥
}

func (b *rsaBucketRepo) Add(ctx context.Context, snowId int64) error {
	// _, err := b.data.mysql.Exec("INSERT INTO bucket(id, bucket_path) VALUES(?, ?)", 1, path)
	// if err != nil {
	// 	b.log.Error("添加密钥桶路径失败", err)
	// 	return err
	// }
	return nil
}

func (b *rsaBucketRepo) GetAll(ctx context.Context) (int32, error) {
	// TODO 获取密钥桶所有数据
	return 0, nil
}

func (b *rsaBucketRepo) Delete(ctx context.Context, snowId int64) error {
	return nil
}
