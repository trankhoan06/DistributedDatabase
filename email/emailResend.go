package emailSend

//func main() {
//	apiKey := "re_75Zomb3A_HHV2HsE93XYBcYJ6jFUoepwN"
//
//	url := "https://api.resend.com/emails"
//
//	payload := map[string]interface{}{
//		"from":    "onboarding@resend.dev",
//		"to":      []string{"your@email.com"},
//		"subject": "Hello from Go",
//		"html":    "<strong>It works!</strong>",
//	}
//
//	jsonData, _ := json.Marshal(payload)
//
//	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
//	req.Header.Set("Authorization", "Bearer "+apiKey)
//	req.Header.Set("Content-Type", "application/json")
//
//	client := &http.Client{}
//	resp, err := client.Do(req)
//
//	if err != nil {
//		panic(err)
//	}
//	defer resp.Body.Close()
//
//	fmt.Println("Status:", resp.Status)
//}
