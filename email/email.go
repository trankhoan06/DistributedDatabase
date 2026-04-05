package emailSend

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/jordan-wright/email"
	"io"
	"net/http"
	"net/smtp"
	"time"
)

const (
	SmtpAuthorAddress = "smtp.sendgrid.net"
	SmtpSeverService  = "smtp.sendgrid.net:587"
)

type Sender interface {
	SendEmail(
		Title string,
		Content string,
		To []string,
		Cc []string,
		Bcc []string,
		FileName string, // ví dụ: "qrcode.png"
		Buf *bytes.Buffer,

	) error
	SendMail443(
		subject string,
		html string,
		to []string,
		Cc []string,
		Bcc []string,
		fileName string, // "" nếu không có
		Buf *bytes.Buffer, // nil nếu không có

	) error
	SendMailResend(
		subject string,
		html string,
		to []string,
		Cc []string,
		Bcc []string,
		fileName string, // "" nếu không có
		Buf *bytes.Buffer, // nil nếu không có

	) error
}
type GmailSender struct {
	name              string
	fromEmailAddr     string
	fromEmailPassword string
}

func NewGmailSender(name string, fromEmailAddr string, fromEmailPassword string) Sender {
	return &GmailSender{
		name:              name,
		fromEmailAddr:     fromEmailAddr,
		fromEmailPassword: fromEmailPassword,
	}
}
func (sender *GmailSender) SendEmail(
	Title string,
	Content string,
	To []string,
	Cc []string,
	Bcc []string,
	FileName string, // ví dụ: "qrcode.png"
	Buf *bytes.Buffer,
) error {

	e := email.NewEmail()
	e.From = fmt.Sprintf("<%s>", sender.fromEmailAddr)
	e.To = To
	e.Subject = Title
	e.HTML = []byte(Content)
	e.Bcc = Bcc
	e.Cc = Cc
	if FileName != "" {
		att, err := e.Attach(
			Buf,
			FileName,
			"image/png",
		)
		if err != nil {
			return fmt.Errorf("failed to attach inline file: %w", err)
		}
		att.Header.Set("Content-ID", "<"+FileName+">")
		att.Header.Set("Content-Disposition", "inline")

	}

	auPath := smtp.PlainAuth("", sender.fromEmailAddr, sender.fromEmailPassword, SmtpAuthorAddress)
	return e.Send(SmtpSeverService, auPath)
}
func (sender *GmailSender) SendMail443(
	subject string,
	html string,
	to []string,
	cc []string,
	bcc []string,
	fileName string, // "" nếu không có
	Buf *bytes.Buffer, // nil nếu không có
) error {

	type Email struct {
		Email string `json:"email"`
		Name  string `json:"name,omitempty"`
	}
	type Personalization struct {
		To  []Email `json:"to,omitempty"`
		Cc  []Email `json:"cc,omitempty"`
		Bcc []Email `json:"bcc,omitempty"`
	}
	type Content struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	}
	type Attachment struct {
		Content     string `json:"content"`
		Filename    string `json:"filename"`
		Type        string `json:"type,omitempty"`
		Disposition string `json:"disposition,omitempty"`
		ContentID   string `json:"content_id,omitempty"`
	}

	mapList := func(xs []string) []Email {
		out := make([]Email, 0, len(xs))
		for _, x := range xs {
			out = append(out, Email{Email: x})
		}
		return out
	}

	payload := map[string]any{
		"from": Email{Email: sender.fromEmailAddr, Name: sender.name},
		"personalizations": []Personalization{
			{
				To:  mapList(to),
				Cc:  mapList(cc),
				Bcc: mapList(bcc),
			},
		},
		"subject": subject,
		"content": []Content{
			{Type: "text/html", Value: html},
		},
	}

	// Attachment inline (ảnh). HTML dùng: <img src="cid:qrcode.png" />
	if fileName != "" && Buf != nil && Buf.Len() > 0 {
		payload["attachments"] = []Attachment{
			{
				Content:     base64.StdEncoding.EncodeToString(Buf.Bytes()),
				Filename:    fileName,
				Type:        "image/png",
				Disposition: "inline",
				ContentID:   fileName,
			},
		}
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("marshal payload: %w", err)
	}

	req, err := http.NewRequest(
		"POST",
		"https://api.sendgrid.com/v3/mail/send",
		bytes.NewReader(body),
	)
	if err != nil {
		return fmt.Errorf("new request: %w", err)
	}

	// sender.fromEmailPassword = SendGrid API KEY (SG....)
	req.Header.Set("Authorization", "Bearer "+sender.fromEmailPassword)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 20 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("sendgrid request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		b, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("sendgrid error: %s | %s", resp.Status, string(b))
	}
	return nil
}
func (sender *GmailSender) SendMailResend(
	subject string,
	html string,
	to []string,
	cc []string,
	bcc []string,
	fileName string, // "" nếu không có
	Buf *bytes.Buffer, // nil nếu không có
) error {

	url := "https://api.resend.com/emails"

	payload := map[string]interface{}{
		"from":    sender.fromEmailAddr,
		"to":      to,
		"subject": subject,
		"html":    html,
	}

	// Nếu có file, thêm phần attachments
	if fileName != "" && Buf != nil && Buf.Len() > 0 {
		attachment := map[string]string{
			"content":    base64.StdEncoding.EncodeToString(Buf.Bytes()),
			"filename":   fileName,
			"content_id": fileName,
		}
		payload["attachments"] = []map[string]string{attachment}
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("marshal payload: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+sender.fromEmailPassword)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("resend error: %s | %s", resp.Status, string(body))
	}

	return nil
}
