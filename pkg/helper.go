// Package pkg public utils.
package pkg

import (
	Rand "crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jordan-wright/email"
	"google.golang.org/protobuf/proto"
	"io"
	"log/slog"
	rand "math/rand/v2"
	"net/smtp"
	"regexp"
	"strings"
	"xrChat_backend/internal/proto/pb"
)

// GenValidateCode generate random code for email verify.
func GenValidateCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	var sb strings.Builder
	for i := 0; i < width; i++ {
		_, _ = fmt.Fprintf(&sb, "%d", numeric[rand.IntN(r)])
	}
	return sb.String()
}

// VerifyEmailAddress verify email address.
func VerifyEmailAddress(address string) (err error) {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if emailRegex.MatchString(address) {
		return
	}
	err = errors.New("邮箱地址不合法")
	return
}

// GenerateSalt generate salt for encrypt password.
func GenerateSalt() (salt string, err error) {
	// 定义盐值的长度，通常16或32字节就足够
	const saltLength = 32

	// 创建一个足够大的缓冲区以存储随机字节
	buf := make([]byte, saltLength)

	// 从crypto/rand读取随机字节
	if _, err = io.ReadFull(Rand.Reader, buf); err != nil {
		return "", err
	}

	// 将字节转换为十六进制字符串作为盐值
	salt = fmt.Sprintf("%x", buf)
	return
}

// EncryptPassword encrypt password.
func EncryptPassword(salt, password string) (pass string) {
	pass = salt + password
	h := sha256.New()
	h.Write([]byte(pass))
	pass = base64.StdEncoding.EncodeToString(h.Sum(nil))
	return
}

// SendEmail send verify code to email.
func SendEmail(fem, tem, body string, emCode string) error {
	e := email.NewEmail()
	e.From = fem
	e.To = []string{tem}
	e.Subject = "xrChat邮箱验证"
	e.Text = []byte(body)
	err := e.Send("smtp.qq.com:25", smtp.PlainAuth("", fem, emCode, "smtp.qq.com"))
	if err != nil {
		return errors.New("error sending email: " + err.Error())
	}
	return nil

}

func BindProto(c *gin.Context, data proto.Message) (err error) {
	all, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return err
	}
	err = proto.Unmarshal(all, data)
	if err != nil {
		return err
	}
	return nil
}

func WriteProto(c *gin.Context, data proto.Message) {
	bytes, _ := proto.Marshal(data)
	_, err := c.Writer.Write(bytes)
	if err != nil {
		slog.Error("response err" + err.Error())
	}
}

func HandleError(c *gin.Context, er error) {
	resp := &pb.Response{}
	resp.Code = 500
	resp.Message = er.Error()
	res, _ := proto.Marshal(resp)
	_, err := c.Writer.Write(res)
	if err != nil {
		slog.Error("success response err" + err.Error())
	}
}

func HandleSuccess(c *gin.Context, msg string) {
	resp := &pb.Response{}
	resp.Code = 200
	resp.Message = msg
	res, _ := proto.Marshal(resp)
	_, err := c.Writer.Write(res)
	if err != nil {
		slog.Error("response 200 error" + err.Error())
	}
}
