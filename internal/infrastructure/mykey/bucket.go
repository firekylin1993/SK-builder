package mykey

import (
	"SK-builder/internal/conf"
	"SK-builder/internal/data"
	"context"
	"crypto/rsa"
	"errors"
	"os"
	"path"
	"strconv"
	"sync"
)

var SnowIDMap *sync.Map

type RsaBucket struct {
	Path     string
	Limit    int32
	RsaKey   *RsaKey
	BucketDb *data.RsaBucketRepo
	SnowNode *SnowNode
}

func NewRsaBucket(c *conf.Server, r *RsaKey, sn *SnowNode, db *data.RsaBucketRepo) *RsaBucket {
	return &RsaBucket{
		Path:     c.RsaBucket.Path,
		Limit:    c.RsaBucket.Limit,
		RsaKey:   r,
		BucketDb: db,
		SnowNode: sn,
	}
}

func (b *RsaBucket) Generate(ctx context.Context) (*rsa.PrivateKey, error) {
	return b.RsaKey.GenerateKey() //生成私钥对象
}

func (b *RsaBucket) Fill(ctx context.Context, pk *rsa.PrivateKey) (string, error) {
	snowIDBtye, snowIDInt64 := b.SnowNode.GetID() //获取雪花算法的ID
	//组装桶密钥路径
	path := path.Join(b.Path, strconv.Itoa(int(snowIDInt64)))
	if os.IsNotExist(os.MkdirAll(path, 0755)) {
		return "", errors.New("mkdir error")
	}

	err := b.RsaKey.GetKey(pk, path) // 生成私钥文件
	if err != nil {
		return "", err
	}

	err = b.RsaKey.GetPublicKey(pk, snowIDBtye, path) // 生成公钥文件
	if err != nil {
		return "", err
	}

	return path, nil
}

func (b *RsaBucket) Remove(ctx context.Context, path string) error {
	return os.RemoveAll(path) //删除桶密钥
}
