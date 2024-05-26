/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/ChikaKakazu/go-cli-switchbot/domain"
	"github.com/spf13/cobra"
)

// devicesCmd represents the devices command
var devicesCmd = &cobra.Command{
	Use:   "devices",
	Short: "Fetch SwitchBot devices information",
	Long:  `Fetch information about all SwitchBot devices associated with the user's account.`,
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
	deviceList, _ := domain.GetDeviceList()

	var result map[string]interface{}
	byteDeviceList, _ := json.Marshal(deviceList)
	err := json.Unmarshal(byteDeviceList, &result)
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
