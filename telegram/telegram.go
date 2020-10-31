package telegram

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/getsentry/sentry-go"
	"github.com/hashicorp/go-retryablehttp"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type Client struct {
	Request *http.Client
	Token   string
}

type Message struct {
	ChatId string `json:"chat_id"`
	Text   string `json:"text"`
}

type TelegramResponse struct {
	Ok          bool   `json:"ok"`
	ErrorCode   int    `json:"error_code"`
	Description string `json:"description"`
	Result      struct {
		MessageID int `json:"message_id"`
		Chat      struct {
			ID       int64  `json:"id"`
			Title    string `json:"title"`
			Username string `json:"username"`
			Type     string `json:"type"`
		} `json:"chat"`
		Date int    `json:"date"`
		Text string `json:"text"`
	} `json:"result"`
}

func New(token string) Client {
	httpRetryClient := retryablehttp.NewClient()
	httpRetryClient.RetryMax = 2
	httpRetryClient.HTTPClient.Timeout = time.Second * 30
	httpClient := httpRetryClient.StandardClient()
	return Client{
		Request: httpClient,
		Token:   token,
	}
}

func (c Client) SendMessage(chatId string, message string) error {
	data, err := json.Marshal(
		Message{
			ChatId: chatId,
			Text:   message,
		},
	)
	if err != nil {
		sentry.CaptureException(err)
	}
	endpoint := "https://api.telegram.org/bot" + c.Token + "/sendMessage"
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(data))
	if err != nil {
		sentry.CaptureException(err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.Request.Do(req)
	if err != nil {
		sentry.CaptureException(err)
	}
	defer resp.Body.Close()
	var response TelegramResponse
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		return err
	}
	if response.Ok {
		return nil
	}
	strError := strconv.Itoa(response.ErrorCode) + ": " + response.Description
	return errors.New(strError)
}
