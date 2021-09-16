package Util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	b64 "encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

//GetHashSha256 ....
func GetHashSha256(SecretKey string) string {
	hash := sha256.Sum256([]byte(SecretKey))
	key := hex.EncodeToString(hash[:])[:32]
	return key
}

// Base64decode ...
func Base64decode(s string) (map[string]interface{}, error) {
	sDec, _ := b64.StdEncoding.DecodeString(s)
	var dat map[string]interface{}
	if err := json.Unmarshal(sDec, &dat); err != nil {
		return nil, err
	}

	return dat, nil
}

// Base64encode ...
func Base64encode(s string, iv string) string {
	encode := fmt.Sprintf(`{"Data":"%v","IV":"%v"}`, s, iv)
	sDec := b64.StdEncoding.EncodeToString([]byte(encode))
	return sDec
}

// Decrypt ...
// func Decrypt(ciphertext string, iv string) (map[string]interface{}, error) {
// 	SecretKey := os.Getenv("SECRET_KEY")
// 	key := GetHashSha256(SecretKey)
// 	btext, err := base64.StdEncoding.DecodeString(ciphertext)
// 	// btext := []byte(ciphertext)
// 	if err != nil {
// 		return nil, err
// 	}

// 	aesCipher, err := aes.NewCipher([]byte(key))
// 	if err != nil {
// 		return nil, err
// 	}
// 	if len(aesCipher) != len(iv) {

// 	}

// 	cipher.NewCBCDecrypter(aesCipher, []byte(iv)).
// 		CryptBlocks(btext, btext)
// 	delimiter := "}"
// 	leftOfDelimiter := strings.SplitAfter(string(btext), delimiter)[0]
// 	var f interface{}
// 	json.Unmarshal([]byte(leftOfDelimiter), &f)
// 	data := f.(map[string]interface{})
// 	return data, nil
// }

// Enrypt ...
func Enrypt(ciphertext string) string {
	// SecretKey := os.Getenv("SECRET_KEY")
	// key := GetHashSha256(SecretKey)
	key, _ := hex.DecodeString("6368616e676520746869732070617373")
	btext := []byte("exampleplaintext")
	// btext := []byte("llllllllll")
	// if err != nil {
	// 	return "nil"
	// }

	// if len(ciphertext)%aes.BlockSize != 0 {
	// 	panic("plaintext is not a multiple of the block size")
	// }

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}
	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	newtext := make([]byte, aes.BlockSize+len(btext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, []byte(iv)); err != nil {
		panic(err)
	}

	mode := cipher.NewCBCEncrypter(block, []byte(iv))
	mode.CryptBlocks(newtext[aes.BlockSize:], []byte(ciphertext))

	// It's important to remember that ciphertexts must be authenticated
	// (i.e. by using crypto/hmac) as well as being encrypted in order to
	// be secure.

	fmt.Printf("%x\n", ciphertext)
	// data := f.(map[string]interface{})
	return "data"
}

// PKCS7Padding ...
func PKCS7Padding(ciphertext []byte) []byte {
	padding := aes.BlockSize - len(ciphertext)%aes.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// PKCS7UnPadding ...
func PKCS7UnPadding(plantText []byte) []byte {
	length := len(plantText)
	unpadding := int(plantText[length-1])
	return plantText[:(length - unpadding)]
}

// OpensslEncrypt ...
func OpensslEncrypt(ivStr string, text string) string {

	plaintext := []byte(text)
	SecretKey := os.Getenv("SECRET_KEY")
	key := GetHashSha256(SecretKey)
	// iv, _ := hex.DecodeString(ivStr)

	// plaintext = PKCS7Padding(plaintext)
	// ciphertext := make([]byte, len(plaintext))
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}
	// ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	// iv := ciphertext[:aes.BlockSize]
	// fmt.Println("IV:", iv)
	mode := cipher.NewCBCEncrypter(block, []byte("fnuhqyanUXKVxey7"))
	mode.CryptBlocks(plaintext, plaintext)

	return fmt.Sprintf("%x\n", plaintext)
}
