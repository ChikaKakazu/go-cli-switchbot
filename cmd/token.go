/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/ChikaKakazu/go-cli-switchbot/config"
	"github.com/spf13/cobra"
)

// tokenCmd represents the token command
var tokenCmd = &cobra.Command{
	Use:   "token [set|get]",
	Short: "Manage SwitchBot API token and secret",
	Long:  `Set or get the SwitchBot API token and secret used for authenticating requests`,
	Run: func(cmd *cobra.Command, args []string) {
		config := config.NewConfig()
		switch args[0] {
		case "set":
			set(config)
		case "get":
			get(config)
		default:
			fmt.Println("Invalid argument. Please specify 'set' or 'get'.")
		}
	},
}

func init() {
	rootCmd.AddCommand(tokenCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tokenCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tokenCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// トークンを設定する
func set(config *config.Config) {
	var token string
	var secret string
	// トークンを入力してもらう
	fmt.Print("Enter your token: ")
	fmt.Scanln(&token)
	fmt.Print("Enter your secret: ")
	fmt.Scanln(&secret)
	config.Token = token
	config.Secret = secret
	err := config.SetConfig()
	if err != nil {
		fmt.Println("Failed to set token or secret: ", err)
	} else {
		fmt.Println("Token and Secret set successfully.")
	}
}

// トークンを取得する
func get(config *config.Config) {
	config, err := config.GetConfig()
	if err != nil {
		fmt.Println("Failed to get token and secret: ", err)
	} else {
		fmt.Println("Token: ", config.Token)
		fmt.Println("Secret: ", config.Secret)
	}
}
