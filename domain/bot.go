package domain

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ChikaKakazu/go-cli-switchbot/config"
)

type Bot struct {
	DeviceId         string `json:"deviceId"`
	DeviceName       string `json:"deviceName"`
	DeviceSelectName string
}

func url(deviceId string) string {
	return fmt.Sprintf("%s/devices/%s/commands", BaseURL, deviceId)
}

// 選択したbotのデバイスをオフにする
func (b *Bot) TurnOff(sighReq *config.SignRequest) ([]byte, error) {
	// /v1.1/devices/{deviceId}/commands
	url := url(b.DeviceId)
	body := map[string]string{
		"command":     "turnOff",
		"parameter":   "default",
		"commandType": "command",
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %w", err)
	}

	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", sighReq.Token)
	req.Header.Set("sign", sighReq.Signature)
	req.Header.Set("t", sighReq.Time)
	req.Header.Set("nonce", sighReq.Nonce)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to turn off device: %s", resp.Status)
	}

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	resJson, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to encode response: %w", err)
	}

	return resJson, nil
}

// 選択したbotのデバイスをオンにする
func (b *Bot) TurnOn(sighReq *config.SignRequest) ([]byte, error) {
	// /v1.1/devices/{deviceId}/commands
	url := url(b.DeviceId)
	body := map[string]string{
		"command":     "turnOn",
		"parameter":   "default",
		"commandType": "command",
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %w", err)
	}

	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", sighReq.Token)
	req.Header.Set("sign", sighReq.Signature)
	req.Header.Set("t", sighReq.Time)
	req.Header.Set("nonce", sighReq.Nonce)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to turn off device: %s", resp.Status)
	}

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	resJson, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to encode response: %w", err)
	}

	return resJson, nil
}
