/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/ChikaKakazu/go-cli-switchbot/domain"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// humidifierCmd represents the humidifier command
var humidifierCmd = &cobra.Command{
	Use:   "humidifier",
	Short: "Interact with Humidifier devices",
	Long:  `Interact with Humidifier devices to turn them on or off.`,
	Run: func(cmd *cobra.Command, args []string) {
		humidifier()
	},
}

func init() {
	rootCmd.AddCommand(humidifierCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// humidifierCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// humidifierCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func humidifier() {
	// デバイスの一覧から加湿器のデバイスを取得する
	devices, sighReq := domain.GetDevices()
	if len(devices) == 0 {
		return
	}

	humidifierDevices := make(map[string]domain.Humidifier)
	for _, device := range devices {
		if device.DeviceType == "Humidifier" {
			humidifierDevices[device.DeviceId] = domain.Humidifier{
				DeviceId:         device.DeviceId,
				DeviceName:       device.DeviceName,
				DeviceSelectName: fmt.Sprintf("%s: %s", device.DeviceId, device.DeviceName),
			}
		}
	}

	if (len(humidifierDevices)) == 0 {
		fmt.Println("No humidifier devices found")
		return
	}

	var items []string
	for _, h := range humidifierDevices {
		items = append(items, h.DeviceSelectName)
	}

	// 加湿器のデバイスの一覧を表示する
	prompt := promptui.Select{
		Label: "Select Humidifier",
		Items: items,
	}

	_, result, err := prompt.Run()
	if err != nil {
		fmt.Println("Prompt failed: ", err)
		return
	}

	// 選択されたデバイスに対して処理を行う
	var humidifier *domain.Humidifier
	for _, h := range humidifierDevices {
		if h.DeviceSelectName == result {
			humidifier = &h
			break
		}
	}

	if humidifier == nil {
		fmt.Println("Failed to get humidifier")
		return
	}

	// 加湿器のデバイスに対して処理を行う
	prompt = promptui.Select{
		Label: "Select Action",
		Items: []string{"Turn on", "Turn off"},
	}

	_, action, err := prompt.Run()
	if err != nil {
		fmt.Println("Prompt failed: ", err)
		return
	}

	var resp []byte
	if action == "Turn off" {
		resp, _ = humidifier.TurnOff(sighReq)
	} else {
		resp, _ = humidifier.TurnOn(sighReq)
	}

	// 処理結果を表示する
	fmt.Println(string(resp))
}
