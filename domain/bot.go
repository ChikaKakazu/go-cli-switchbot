package domain

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ChikaKakazu/go-cli-switchbot/config"
	"github.com/ChikaKakazu/go-cli-switchbot/helper"
)

// Bot SwitchBotデバイスの情報を保持する構造体
type Bot struct {
	DeviceId         string `json:"deviceId"`   // デバイスの一意の識別子
	DeviceName       string `json:"deviceName"` // デバイスの表示名
	DeviceSelectName string // CLIでのデバイス選択時に使用する名前
}

// TurnOff 選択したSwitchBotデバイスをオフにする
// 戻り値: APIレスポンスのJSONバイトとエラー情報
func (b *Bot) TurnOff(signReq *config.SignRequest) ([]byte, error) {
	url := helper.CommandUrl(b.DeviceId)

	// コマンドのペイロードを準備
	body := map[string]string{
		"command":     "turnOff",
		"parameter":   "default",
		"commandType": "command",
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %w", err)
	}

	// タイムアウト付きのHTTPクライアントを作成
	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// 認証に必要なヘッダーを設定
	req.Header.Set("Authorization", signReq.Token)
	req.Header.Set("sign", signReq.Signature)
	req.Header.Set("t", signReq.Time)
	req.Header.Set("nonce", signReq.Nonce)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to turn off device: %s", resp.Status)
	}

	// レスポンスを解析して整形
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

// TurnOn 選択したSwitchBotデバイスをオンにする
// 戻り値: APIレスポンスのJSONバイトとエラー情報
func (b *Bot) TurnOn(signReq *config.SignRequest) ([]byte, error) {
	url := helper.CommandUrl(b.DeviceId)

	// コマンドのペイロードを準備
	body := map[string]string{
		"command":     "turnOn",
		"parameter":   "default",
		"commandType": "command",
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %w", err)
	}

	// タイムアウト付きのHTTPクライアントを作成
	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// 認証に必要なヘッダーを設定
	req.Header.Set("Authorization", signReq.Token)
	req.Header.Set("sign", signReq.Signature)
	req.Header.Set("t", signReq.Time)
	req.Header.Set("nonce", signReq.Nonce)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to turn off device: %s", resp.Status)
	}

	// レスポンスを解析して整形
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
