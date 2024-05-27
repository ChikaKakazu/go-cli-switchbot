package domain

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ChikaKakazu/go-cli-switchbot/config"
	"github.com/ChikaKakazu/go-cli-switchbot/helper"
)

type Scene struct {
	SceneId         string `json:"sceneId"`
	SceneName       string `json:"sceneName"`
	SceneSelectName string
}

type SceneList struct {
	StatusCode int     `json:"statusCode"`
	Message    string  `json:"message"`
	Body       []Scene `json:"body"`
}

func GetSceneList() (*SceneList, *config.SignRequest) {
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

	// シーン一覧を取得する
	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/scenes", helper.BaseURL), nil)
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

	var result SceneList
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		fmt.Println("Failed to decode response: ", err)
		return nil, nil
	}

	return &result, sighReq
}

func GetScenes() ([]Scene, *config.SignRequest) {
	sceneList, sighReq := GetSceneList()
	if sceneList == nil {
		return nil, nil
	}

	return sceneList.Body, sighReq
}

func (s *Scene) ExecScene(signReq *config.SignRequest) ([]byte, error) {
	url := helper.SceneExecuteUrl(s.SceneId)

	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

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
