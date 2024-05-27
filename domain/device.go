package domain

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ChikaKakazu/go-cli-switchbot/config"
)

type Device struct {
	DeviceId           string `json:"deviceId"`
	DeviceName         string `json:"deviceName"`
	DeviceType         string `json:"deviceType"`
	EnableCloudService bool   `json:"enableCloudService"`
	HubDeviceId        string `json:"hubDeviceId"`
}

type DeviceList struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Body       struct {
		DeviceList []Device `json:"deviceList"`
	} `json:"body"`
}

func GetDeviceList() (*DeviceList, *config.SignRequest) {
	config := config.NewConfig()
	config, err := config.GetConfig()
	if err != nil {
		fmt.Println("Failed to get token or secret: ", err)
		return nil, nil
	}

	if config.Token == "" || config.Secret == "" {
		fmt.Println("Token or Secret is not set. Please set them first.")
		return nil, nil
	}

	// 署名を生成する
	sighReq, err := config.GenerateSignature()
	if err != nil {
		fmt.Println("Failed to generate signature: ", err)
		return nil, nil
	}

	// デバイス一覧を取得する
	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("GET", "https://api.switch-bot.com/v1.1/devices", nil)
	if err != nil {
		fmt.Println("Failed to create request: ", err)
		return nil, nil
	}

	req.Header.Set("Authorization", sighReq.Token)
	req.Header.Set("sign", sighReq.Signature)
	req.Header.Set("t", sighReq.Time)
	req.Header.Set("nonce", sighReq.Nonce)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to send request: ", err)
		return nil, nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Failed to get devices: ", resp.Status)
		return nil, nil
	}

	var result DeviceList
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		fmt.Println("Failed to decode response: ", err)
		return nil, nil
	}

	return &result, sighReq
}

func GetDevices() ([]Device, *config.SignRequest) {
	deviceList, req := GetDeviceList()
	return deviceList.Body.DeviceList, req
}
