package Util

import (
	extra "DT/Extra"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	// "github.com/spacemonkeygo/openssl"
)

// Trim ...
func Trim(stringn string) map[string]interface{} {
	// stringn := "�,$��K��oh���� {\"Data\":\"mmmI am string\",\"IV\":\"kkkkkkkk\"}"
	delimiter := "{"

	leftOfDelimiter := strings.Split(stringn, delimiter)[0]
	rightOfDelimiter := strings.Join(strings.Split(stringn, delimiter)[1:], delimiter)
	fmt.Println("Full: ", stringn)
	fmt.Println("Left of Delimiter: ", leftOfDelimiter)
	fmt.Println("Right of Delimiter: ", "{"+rightOfDelimiter)
	left := "{" + rightOfDelimiter
	var f interface{}
	json.Unmarshal([]byte(left), &f)
	data := f.(map[string]interface{})
	// data := f.(map[string]interface{})
	return data
}

//DecryptData1 ....
func DecryptData1(SecretKey string) string {
	sha256 := sha256.Sum256([]byte(SecretKey))
	str := fmt.Sprintf("%x", sha256)
	return string(str)

	// Vmr-uU5mA2_Zr_13
}

// Decryptnew ...
func Decryptnew(ciphertext string, iv string) (map[string]interface{}, error) {
	SecretKey := os.Getenv("SECRET_KEY")
	key := GetHashSha256(SecretKey)
	btext, err := base64.StdEncoding.DecodeString(ciphertext)
	iv1, err1 := base64.StdEncoding.DecodeString(iv)
	// btext := []byte(ciphertext)
	var newiv []byte
	if err1 != nil {
		newiv = []byte(iv)
		fmt.Println("error")
		fmt.Println("Aes:", aes.BlockSize, "Iv:", len(newiv))

	} else {

		if aes.BlockSize != len(iv1) {
			newiv = []byte(iv)
		} else {
			newiv = iv1
		}

		fmt.Println("erro1")
	}

	aesCipher, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}
	fmt.Println("Aes:", aes.BlockSize, "Iv:", len(newiv))
	cipher.NewCBCDecrypter(aesCipher, newiv).
		CryptBlocks(btext, btext)
	fmt.Println("eoorr1")
	// ciphertext1, _ := pkcs7Unpad(btext, aes.BlockSize)
	ciphertext1, _ := extra.Unpad(btext)

	response := Trim(string(ciphertext1))
	return response, nil
}

// Encryptnew ...
func Encryptnew(aeskey string, filename string) (string, string) {
	key := []byte(aeskey)
	// Create the AES cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	// plaintext, _ := pkcs7Pad([]byte(filename), block.BlockSize())
	plaintext := extra.Pad([]byte(filename))
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]

	fmt.Println("block:", aes.BlockSize, "iv", len(iv))

	// fmt.Println("iv:", iv)
	// os.Exit(2)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	bm := cipher.NewCBCEncrypter(block, iv)
	bm.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	return base64.StdEncoding.EncodeToString(ciphertext), base64.StdEncoding.EncodeToString(iv)
}

var (
	// ErrInvalidBlockSize indicates hash blocksize <= 0.
	ErrInvalidBlockSize = errors.New("invalid blocksize")

	// ErrInvalidPKCS7Data indicates bad input to PKCS7 pad or unpad.
	ErrInvalidPKCS7Data = errors.New("invalid PKCS7 data (empty or not padded)")

	// ErrInvalidPKCS7Padding indicates PKCS7 unpad fails to bad input.
	ErrInvalidPKCS7Padding = errors.New("invalid padding on input")
)

// PKCS7Paddingnew ...
func PKCS7Paddingnew(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)

}
func pkcs7Pad(b []byte, blocksize int) ([]byte, error) {
	if blocksize <= 0 {
		return nil, ErrInvalidBlockSize
	}
	if b == nil || len(b) == 0 {
		return nil, ErrInvalidPKCS7Data
	}
	n := blocksize - (len(b) % blocksize)
	pb := make([]byte, len(b)+n)
	copy(pb, b)
	copy(pb[len(b):], bytes.Repeat([]byte{byte(n)}, n))
	return pb, nil
}

func pkcs7Unpad(b []byte, blocksize int) ([]byte, error) {
	if blocksize <= 0 {
		return nil, ErrInvalidBlockSize
	}
	if b == nil || len(b) == 0 {
		return nil, ErrInvalidPKCS7Data
	}
	if len(b)%blocksize != 0 {
		return nil, ErrInvalidPKCS7Padding
	}
	c := b[len(b)-1]
	n := int(c)
	if n == 0 || n > len(b) {
		return nil, ErrInvalidPKCS7Padding
	}
	for i := 0; i < n; i++ {
		if b[len(b)-n+i] != c {
			return nil, ErrInvalidPKCS7Padding
		}
	}
	return b[:len(b)-n], nil
}
