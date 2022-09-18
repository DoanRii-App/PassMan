package main

import (
 "log"
 "fmt"
 "os"
 "crypto/aes"
 "crypto/cipher"
 "encoding/base64"
 "github.com/joho/godotenv"

 "./passgen"
)

func goDotEnvVariable(key string) string {
    err := godotenv.Load(".env")
  
    if err != nil {
      log.Fatalf("Error loading .env file")
    }
  
    return os.Getenv(key)
}

var bytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

// .env SKEY="secretekey"
var MySecret string = goDotEnvVariable("SKEY")

func Encode(b []byte) string {
 return base64.StdEncoding.EncodeToString(b)
}

func Decode(s string) []byte {
 data, err := base64.StdEncoding.DecodeString(s)
 if err != nil {
  panic(err)
 }
 return data
}

// Encrypt method is to encrypt or hide any classified text
func Encrypt(text, MySecret string) (string, error) {
 block, err := aes.NewCipher([]byte(MySecret))
 if err != nil {
  return "", err
 }
 plainText := []byte(text)
 cfb := cipher.NewCFBEncrypter(block, bytes)
 cipherText := make([]byte, len(plainText))
 cfb.XORKeyStream(cipherText, plainText)
 return Encode(cipherText), nil
}

// Decrypt method is to extract back the encrypted text
func Decrypt(text, MySecret string) (string, error) {
 block, err := aes.NewCipher([]byte(MySecret))
 if err != nil {
  return "", err
 }
 cipherText := Decode(text)
 cfb := cipher.NewCFBDecrypter(block, bytes)
 plainText := make([]byte, len(cipherText))
 cfb.XORKeyStream(plainText, cipherText)
 return string(plainText), nil
}

func main() {
 fmt.Println(GenPass(24,12,0,true,false))
 StringToEncrypt := "Encrypting this string"

 // To encrypt the StringToEncrypt
 encText, err := Encrypt(StringToEncrypt, MySecret)
 if err != nil {
  fmt.Println("error encrypting your classified text: ", err)
 }
 fmt.Println(encText)

 // To decrypt the original StringToEncrypt
 decText, err := Decrypt("Li5E8RFcV/EPZY/neyCXQYjrfa/atA==", MySecret)
 if err != nil {
  fmt.Println("error decrypting your encrypted text: ", err)
 }
 fmt.Println(decText)
}