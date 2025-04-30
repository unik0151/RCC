package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"github.com/sirupsen/logrus"
	"io"
)

var PassKey = "dwuqjjsndsb"

func EncodeMd5(buffer []byte) string {
	md5Hash := md5.New()
	md5Hash.Write(buffer)
	md5HashSum := md5Hash.Sum(nil)
	md5Str := hex.EncodeToString(md5HashSum)
	return md5Str
}

func EncryptGCM(plaintext []byte, keyStr string) (string, error) {
	// 生成固定长度密钥
	key := sha256.Sum256([]byte(keyStr))

	block, err := aes.NewCipher(key[:])
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// AES-GCM解密
func DecryptGCM(ciphertext string, keyStr string) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return nil, err
	}

	key := sha256.Sum256([]byte(keyStr))
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}

	nonce, ciphertextBytes := data[:nonceSize], data[nonceSize:]
	return gcm.Open(nil, nonce, ciphertextBytes, nil)
}

func getGcm(key []byte) cipher.AEAD {
	block, err := aes.NewCipher(key)
	if err != nil {
		logrus.Error(err)
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		logrus.Error(err)
	}
	return gcm
}

func GetRandSizeKeyBytes(size int) ([]byte, error) {
	bytes := make([]byte, size)
	_, err := io.ReadFull(rand.Reader, bytes)
	if err != nil {
		logrus.Error(err)
	}
	return bytes, nil
}
