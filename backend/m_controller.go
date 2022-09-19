package main

import (
 "log"
 "time"
 "bytes"
 "strconv"
 "strings"
 "net/http"
 "crypto/aes"
 "crypto/md5"
 "crypto/hmac"
 "crypto/sha1"
 "encoding/hex"
 "crypto/cipher"
 "encoding/base32"
 "encoding/binary"
 "encoding/base64"
 "github.com/gin-gonic/gin"
 _ "github.com/mattn/go-sqlite3"
 "github.com/sethvargo/go-password/password"
)

var sbytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

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

func Encrypt(text, SKey string) (string, error) {
  hasher := md5.New()
  hasher.Write([]byte(SKey))
  block, err := aes.NewCipher([]byte(hex.EncodeToString(hasher.Sum(nil))))
  if err != nil {
    return "", err
  }
  plainText := []byte(text)
  cfb := cipher.NewCFBEncrypter(block, sbytes)
  cipherText := make([]byte, len(plainText))
  cfb.XORKeyStream(cipherText, plainText)
  return Encode(cipherText), nil
}

func Decrypt(text, SKey string) (string, error) {
  hasher := md5.New()
  hasher.Write([]byte(SKey))
  block, err := aes.NewCipher([]byte(hex.EncodeToString(hasher.Sum(nil))))
  if err != nil {
    return "", err
  }
  cipherText := Decode(text)
  cfb := cipher.NewCFBDecrypter(block, sbytes)
  plainText := make([]byte, len(cipherText))
  cfb.XORKeyStream(plainText, cipherText)
  return string(plainText), nil
}

func GenRand(l int, d int, sym int, ulC bool, drC bool) (string) {
  res, err := password.Generate(l, d, sym, ulC, drC)
  if err != nil {
    log.Fatal(err)
  }
  return res
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func prefix0(otp string) string {
	if len(otp) == 6 {
		return otp
	}
	for i := (6 - len(otp)); i > 0; i-- {
		otp = "0" + otp
	}
	return otp
}

func getHOTPToken(secret string, interval int64) string {
	key, err := base32.StdEncoding.DecodeString(strings.ToUpper(secret))
	check(err)
	bs := make([]byte, 8)
	binary.BigEndian.PutUint64(bs, uint64(interval))
	hash := hmac.New(sha1.New, key)
	hash.Write(bs)
	h := hash.Sum(nil)
	o := (h[19] & 15)
	var header uint32
	r := bytes.NewReader(h[o : o+4])
	err = binary.Read(r, binary.BigEndian, &header)
	check(err)
	h12 := (int(header) & 0x7fffffff) % 1000000
	otp := strconv.Itoa(int(h12))
	return prefix0(otp)
}

func getTOTPToken(secret string) string {
	interval := time.Now().Unix() / 30
	return getHOTPToken(secret, interval)
}

func getUsers(c *gin.Context) {
  auth := c.Param("auth")
  tOki, _ := Encrypt(auth, auth)
	db, _ := fDB()
  rows, err := db.Query("SELECT User, Pass, Email, Otp, Tag, Oki FROM users WHERE Oki = ?", tOki)
  if (err != nil){
    log.Fatal(err)
  }
  if rows.Next() {
    for rows.Next() {
      var gUserData User
      rows.Scan(&gUserData.User, &gUserData.Pass, &gUserData.Email, &gUserData.Otp, &gUserData.Tag, &gUserData.Oki)
      tUser, _ := Decrypt(gUserData.User, auth)
      tPass, _ := Decrypt(gUserData.Pass, auth)
      tEmail, _ := Decrypt(gUserData.Email, auth)
      tOtp, _ := Decrypt(gUserData.Otp, auth)
      cOtp := getTOTPToken(tOtp)
      c.JSON(http.StatusOK, gin.H{
        "sMsg": "ok",
        "user": tUser,
        "pass": tPass,
        "email": tEmail,
        "otp": cOtp,
        "tag": gUserData.Tag,
      }) 
    }
  }else{
    c.JSON(http.StatusOK, gin.H{
      "sMsg": "error",
    }) 
  }
  rows.Close()
}

func addUser(c *gin.Context) {
  var addUserData User
  auth := c.Param("auth")
  if err := c.BindJSON(&addUserData); err != nil {
    log.Fatal(err);
  }
	db, _ := fDB()
  st, err := db.Prepare("INSERT INTO users (User, Pass, Email, Otp, Tag, Oki) VALUES (?, ?, ?, ?, ?, ?)")
  if (err != nil){
    log.Fatal(err)
  }
  tUser, _ := Encrypt(addUserData.User, auth)
  tPass, _ := Encrypt(addUserData.Pass, auth)
  tEmail, _ := Encrypt(addUserData.Email, auth)
  tOtp, _ := Encrypt(strings.ReplaceAll(addUserData.Otp, " ", ""), auth)
  tTag, _ := Encrypt(addUserData.Tag, auth)
  tOki, _ := Encrypt(auth, auth)
  st.Exec(tUser, tPass, tEmail, tOtp, tTag, tOki)
  c.JSON(http.StatusOK, gin.H{
      "sMsg": "ok",
    })
}