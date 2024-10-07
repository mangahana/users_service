package sms

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"users_service/internal/configuration"
)

type SmsService struct {
	accessToken string
	baseUrl     string
}

func New(config *configuration.SMSConfig) *SmsService {
	return &SmsService{
		accessToken: config.ApiKey,
		baseUrl:     config.ApiDomain,
	}
}

func (s *SmsService) Send(recipient, text string) error {
	path := fmt.Sprintf("%s%s?apiKey=%s", s.baseUrl, "/service/Message/SendSmsMessage", s.accessToken)
	data := url.Values{
		"recipient": []string{recipient},
		"text":      []string{text},
	}
	resp, err := http.PostForm(path, data)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var res struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
	if err := json.Unmarshal(body, &res); err != nil {
		return err
	}

	if res.Code != 0 {
		return errors.New("SMS Error")
	}

	return nil
}
