package test

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"io"
	"miniapp/utils"
	"testing"
)

// AES-GCM加密 (返回base64编码结果)
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

func TestA(t *testing.T) {
	key := "your-secret-key-32bytes" // 任意长度，代码自动处理为32字节

	// 加密示例
	plaintext := []byte("sensitive data")
	encrypted, err := EncryptGCM(plaintext, key)
	if err != nil {
		panic(err)
	}
	fmt.Println("加密结果:", encrypted)

	// 解密示例
	decrypted, err := DecryptGCM(encrypted, key)
	if err != nil {
		panic(err)
	}
	fmt.Println("解密结果:", string(decrypted))
}

func TestB(t *testing.T) {

	// 加密示例
	plaintext := []byte("333333")
	encodeStr, err := utils.EncryptGCM(plaintext, "sss")

	fmt.Println("加密结果:", encodeStr, "--", err)

	// 解密示例
	utils.DecryptGCM(encodeStr, "sss")

}

func TestC(t *testing.T) {
	hash := sha256.New()
	hash.Write([]byte("333"))
	bytes := hash.Sum(nil)
	fmt.Println(string(bytes))

	key := sha256.Sum256([]byte("ssss"))
	fmt.Println(string(key[:]))
}

func TestD(t *testing.T) {
	var b = make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, b)
	if err != nil {
		panic(err)
	}
	encode := hexutil.Encode(b)
	fmt.Println("encode: ", encode)
}
