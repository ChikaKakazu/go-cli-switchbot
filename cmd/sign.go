/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/ChikaKakazu/go-cli-switchbot/config"
	"github.com/spf13/cobra"
)

// signCmd represents the sign command
var signCmd = &cobra.Command{
	Use:   "sign",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		sign()
	},
}

func init() {
	rootCmd.AddCommand(signCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// signCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// signCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func sign() {
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
	// signature, time, nonce, err := config.GenerateSignature()
	req, err := config.GenerateSignature()
	if err != nil {
		fmt.Println("Failed to generate signature: ", err)
		return
	}

	fmt.Println("Token: ", req.Token)
	fmt.Println("Signature: ", req.Signature)
	fmt.Println("Time: ", req.Time)
	fmt.Println("Nonce: ", req.Nonce)
}
