package helper

import "fmt"

const BaseURL string = "https://api.switch-bot.com/v1.1"

func CommandUrl(deviceId string) string {
	// /v1.1/devices/{deviceId}/commands
	return fmt.Sprintf("%s/devices/%s/commands", BaseURL, deviceId)
}
