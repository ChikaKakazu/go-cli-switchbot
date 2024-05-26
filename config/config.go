package config

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
)

const configFilePath = "config.json"

type Config struct {
	Token  string `json:"token"`
	Secret string `json:"secret"`
}

type SignRequest struct {
	Token     string
	Signature string
	Time      string
	Nonce     string
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) SetConfig() error {
	// ファイルを開く
	// ファイルが存在しない場合は作成する
	// 既存のファイルがある場合は上書きする
	file, err := os.OpenFile(configFilePath, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// トークンを書き込む
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(c); err != nil {
		return err
	}
	return nil
}

func (c *Config) GetConfig() (*Config, error) {
	file, err := os.Open(configFilePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return &Config{}, nil
		}
		return nil, err
	}
	defer file.Close()

	config := &Config{}
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}
	return config, nil
}

// 署名を生成する
func (c *Config) GenerateSignature() (*SignRequest, error) {
	token := c.Token
	secret := c.Secret
	// 現在のUnixタイムをミリ秒単位で取得
	time := time.Now().UnixNano() / int64(time.Millisecond)
	// ランダムなuuidを生成
	nonce := uuid.New().String()

	// 署名を生成するためのデータを作成
	data := fmt.Sprintf("%s%d%s", token, time, nonce)

	// 署名を生成
	hmac := hmac.New(sha256.New, []byte(secret))
	_, err := hmac.Write([]byte(data))
	if err != nil {
		return &SignRequest{}, err
	}

	// 署名をBase64エンコードして返す
	signature := base64.StdEncoding.EncodeToString(hmac.Sum(nil))

	return &SignRequest{
		Token:     token,
		Signature: signature,
		Time:      fmt.Sprintf("%d", time),
		Nonce:     nonce,
	}, nil
}
