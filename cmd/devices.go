/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ChikaKakazu/go-cli-switchbot/config"
	"github.com/spf13/cobra"
)

// devicesCmd represents the devices command
var devicesCmd = &cobra.Command{
	Use:   "devices",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		devices()
	},
}

func init() {
	rootCmd.AddCommand(devicesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// devicesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// devicesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func devices() {
	config := config.NewConfig()
	config, err := config.GetConfig()
	if err != nil {
		fmt.Println("Failed to get token or secret: ", err)
		return
	}

	if config.Token == "" || config.Secret == "" {
		fmt.Println("Token or Secret is not set. Please set them first.")
		return
	}

	// 署名を生成する
	sighReq, err := config.GenerateSignature()
	if err != nil {
		fmt.Println("Failed to generate signature: ", err)
		return
	}

	// デバイス一覧を取得する
	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("GET", "https://api.switch-bot.com/v1.1/devices", nil)
	if err != nil {
		fmt.Println("Failed to create request: ", err)
		return
	}

	req.Header.Set("Authorization", sighReq.Token)
	req.Header.Set("sign", sighReq.Signature)
	req.Header.Set("t", sighReq.Time)
	req.Header.Set("nonce", sighReq.Nonce)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to send request: ", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Failed to get devices: ", resp.Status)
		return
	}

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		fmt.Println("Failed to decode response: ", err)
		return
	}

	resJson, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		fmt.Println("Failed to encode response: ", err)
		return
	}

	fmt.Println(string(resJson))
}
