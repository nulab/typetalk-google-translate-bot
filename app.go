package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/context"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
	"google.golang.org/api/option"
)

type Post struct {
	Id      int64
	Message string
}

type WebhookResponse struct {
	Post Post `json:"post"`
}

type WebhookReturnResponse struct {
	Message string `json:"message"`
	ReplyTo int64  `json:"replyTo"`
}

func isDefineApiKey() bool {
	return getApiKey() != ""
}

func getApiKey() string {
	return os.Getenv("GOOGLE_API_KEY")
}

func translateInEnglish(text string) (string, error) {
	ctx := context.Background()

	client, err := translate.NewClient(ctx, option.WithAPIKey(getApiKey()))
	if err != nil {
		return "", err
	}
	defer client.Close()

	resp, err := client.Translate(ctx, []string{text}, language.English, nil)
	if err != nil {
		return "", err
	}
	return resp[0].Text, nil
}

func detectLanguage(text string) (*translate.Detection, error) {
	ctx := context.Background()
	client, err := translate.NewClient(ctx, option.WithAPIKey(getApiKey()))
	if err != nil {
		return nil, err
	}
	defer client.Close()

	lang, err := client.DetectLanguage(ctx, []string{text})
	if err != nil {
		return nil, err
	}
	return &lang[0][0], nil
}

func parse(w http.ResponseWriter, req *http.Request) (Post, error) {
	decoder := json.NewDecoder(req.Body)
	var webhookRes WebhookResponse
	err := decoder.Decode(&webhookRes)

	defer req.Body.Close()
	return webhookRes.Post, err
}

func handle(w http.ResponseWriter, req *http.Request) {
	post, err := parse(w, req)
	if err != nil {
		log.Fatal("Failed parsed Webhook response: ", err)
		return
	}

	detection, err := detectLanguage(post.Message)
	if err != nil {
		log.Fatal("Failed detect language: ", err)
		return
	}

	if detection.Language == language.English {
		return
	}

	translatedText, err := translateInEnglish(post.Message)
	if err != nil {
		log.Fatal("Failed translate text: ", err)
		return
	}

	returnRes := WebhookReturnResponse{translatedText, post.Id}
	json, _ := json.Marshal(returnRes)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func main() {
	if !isDefineApiKey() {
		panic("Set GOOGLE_API_KEY on your OS environment")
	}

	http.HandleFunc("/", handle)
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
