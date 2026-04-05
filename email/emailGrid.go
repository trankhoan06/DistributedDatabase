package emailSend

//
//import (
//	"bytes"
//	"encoding/base64"
//	"encoding/json"
//	"fmt"
//	"io"
//	"net/http"
//	"time"
//)
//
//func (sender *GmailSender) SendMail443(
//	//apiKey string, // SG.xxxxx
//	//fromName string, // Tên hiển thị
//	//fromEmail string, // Email đã verify trên SendGrid
//	subject string,
//	html string,
//	to []string,
//	fileName string, // "" nếu không có
//	fileBytes []byte, // nil nếu không có
//) error {
//
//	type Email struct {
//		Email string `json:"email"`
//		Name  string `json:"name,omitempty"`
//	}
//	type Personalization struct {
//		To []Email `json:"to"`
//	}
//	type Content struct {
//		Type  string `json:"type"`
//		Value string `json:"value"`
//	}
//	type Attachment struct {
//		Content     string `json:"content"`
//		Filename    string `json:"filename"`
//		Type        string `json:"type"`
//		Disposition string `json:"disposition"`
//		ContentID   string `json:"content_id"`
//	}
//
//	toList := make([]Email, 0, len(to))
//	for _, x := range to {
//		toList = append(toList, Email{Email: x})
//	}
//
//	payload := map[string]any{
//		"from": Email{Email: sender.fromEmailAddr, Name: sender.name},
//		"personalizations": []Personalization{
//			{To: toList},
//		},
//		"subject": subject,
//		"content": []Content{
//			{Type: "text/html", Value: html},
//		},
//	}
//
//	// Attachment inline (ảnh)
//	if fileName != "" && len(fileBytes) > 0 {
//		payload["attachments"] = []Attachment{
//			{
//				Content:     base64.StdEncoding.EncodeToString(fileBytes),
//				Filename:    fileName,
//				Type:        "image/png",
//				Disposition: "inline",
//				ContentID:   fileName,
//			},
//		}
//	}
//
//	body, _ := json.Marshal(payload)
//
//	req, _ := http.NewRequest(
//		"POST",
//		"https://api.sendgrid.com/v3/mail/send",
//		bytes.NewReader(body),
//	)
//	req.Header.Set("Authorization", "Bearer "+sender.fromEmailPassword)
//	req.Header.Set("Content-Type", "application/json")
//
//	client := &http.Client{Timeout: 20 * time.Second}
//	resp, err := client.Do(req)
//	if err != nil {
//		return err
//	}
//	defer resp.Body.Close()
//
//	if resp.StatusCode >= 300 {
//		b, _ := io.ReadAll(resp.Body)
//		return fmt.Errorf("sendgrid error: %s | %s", resp.Status, string(b))
//	}
//	return nil
//}
