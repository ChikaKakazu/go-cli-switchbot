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

// botCmd represents the bot command
var botCmd = &cobra.Command{
	Use:   "bot",
	Short: "Interact with Bot devices",
	Long:  `Interact with Bot devices to turn them on or off.`,
	Run: func(cmd *cobra.Command, args []string) {
		bot()
	},
}

func init() {
	rootCmd.AddCommand(botCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// botCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// botCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// bot(物理ボタン)の処理を行う
func bot() {
	// デバイスの一覧からbotのデバイスを取得する
	devices, sighReq := domain.GetDevices()
	if len(devices) == 0 {
		return
	}

	botDevices := make(map[string]domain.Bot)
	for _, device := range devices {
		if device.DeviceType == "Bot" {
			botDevices[device.DeviceId] = domain.Bot{
				DeviceId:         device.DeviceId,
				DeviceName:       device.DeviceName,
				DeviceSelectName: fmt.Sprintf("%s: %s", device.DeviceId, device.DeviceName),
			}
		}

	}

	if (len(botDevices)) == 0 {
		fmt.Println("No bot devices found")
		return
	}

	var items []string
	for _, b := range botDevices {
		items = append(items, b.DeviceSelectName)
	}

	// botのデバイスの一覧を表示する
	prompt := promptui.Select{
		Label: "Select a bot device",
		Items: items,
	}

	_, result, err := prompt.Run()
	if err != nil {
		fmt.Println("Prompt failed: ", err)
		return
	}

	// 選択したbotのデバイスに対して処理を行う
	var bot *domain.Bot
	for _, b := range botDevices {
		if b.DeviceSelectName == result {
			bot = &b
			break
		}
	}

	if bot == nil {
		fmt.Println("Failed to get selected device")
		return
	}

	// ボタンに対するアクションを選択する
	prompt = promptui.Select{
		Label: "Select Action",
		Items: []string{"Turn off", "Turn on"},
	}
	_, action, err := prompt.Run()
	if err != nil {
		fmt.Println("Prompt failed: ", err)
		return
	}

	var resp []byte
	if action == "Turn off" {
		resp, _ = bot.TurnOff(sighReq)
	} else {
		resp, _ = bot.TurnOn(sighReq)
	}

	// 処理結果を表示する
	fmt.Println(string(resp))
}
