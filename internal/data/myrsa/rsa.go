package myrsa

import (
	"SK-Builder/internal/conf"
	"SK-Builder/kits"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"
	"path/filepath"
)

type PublicKey struct {
}

type PrivateKey struct {
}

type RsaKey struct {
	Key        *PrivateKey
	PubKey     *PublicKey
	Path       string
	confServer *conf.Server
}

func NewPublicKey() *PublicKey {
	return &PublicKey{}
}

func NewProviderKey() *PrivateKey {
	return &PrivateKey{}
}

func NewRsaKey(c *conf.Server, k *PrivateKey, pk *PublicKey) *RsaKey {
	return &RsaKey{
		Key:        k,
		PubKey:     pk,
		confServer: c,
	}
}

func (r *RsaKey) GenerateKey() (*rsa.PrivateKey, error) {
	// 生成私钥
	return rsa.GenerateKey(rand.Reader, int(r.confServer.RsaBucket.KeySize))
}

func (r *RsaKey) GetKey(k *rsa.PrivateKey, path string) error {
	// 生成私钥文件
	derStream := x509.MarshalPKCS1PrivateKey(k)
	block := &pem.Block{
		Bytes: derStream,
	}

	keyPath := filepath.Join(path, "private.pem")
	file, err := os.Create(keyPath)
	if err != nil {
		return err
	}
	defer file.Close()
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	return nil
}

func (r *RsaKey) GetPublicKey(k *rsa.PrivateKey, sid [8]byte, path string) error {
	if len(sid) != 8 {
		return errors.New("snowId length error")
	}

	// 生成公钥文件
	publicKey := &k.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}
	keyPath := filepath.Join(path, "public.pem")
	file, err := os.Create(keyPath)
	if err != nil {
		return err
	}
	defer file.Close()
	combine := kits.BytesCombine([]byte("A"), sid[:], derPkix)
	crc32Byte := kits.IntToBytes(kits.GetCRC32Key(combine))

	bytesCombine := kits.BytesCombine(combine, crc32Byte)
	n, err := file.Write(bytesCombine)
	if err != nil {
		return err
	}

	if n != 307 {
		return errors.New("publicKey length error")
	}
	return nil
}
