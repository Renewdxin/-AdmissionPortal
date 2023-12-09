package mail

import (
	"fmt"
	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
	"log"
	"os"
	"strconv"
)

type Mail struct{}

func NewMail() *Mail {
	return &Mail{}
}

func (m *Mail) Send(to string, subject string, body string) error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("无法加载 .env 文件")
	}

	// 获取配置信息
	from := os.Getenv("EMAIL_FROM")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))
	smtpUsername := os.Getenv("SMTP_USERNAME")
	smtpPassword := os.Getenv("SMTP_PASSWORD")

	// 创建邮件消息
	message := gomail.NewMessage()
	message.SetHeader("From", from)
	message.SetHeader("To", to)
	message.SetHeader("Subject", subject)
	message.SetBody("text/plain", body)

	// 创建 SMTP 发送器
	dialer := gomail.NewDialer(smtpHost, smtpPort, smtpUsername, smtpPassword)

	// 发送邮件
	err = dialer.DialAndSend(message)
	if err != nil {
		fmt.Println("发送邮件失败:", err)
		return err
	}

	fmt.Println("邮件发送成功！")
	return nil
}

func (m *Mail) CodeSend(to string, subject string, code string) error {
	return m.Send(to, subject, "your code is: "+code)
}

func (m *Mail) WelcomeMail(to string) error {
	return m.Send(to, "Welcome to Wonderland", "Welcome")
}
