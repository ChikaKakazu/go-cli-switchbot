package helper

import "fmt"

const BaseURL string = "https://api.switch-bot.com/v1.1"

func CommandUrl(deviceId string) string {
	// /devices/{deviceId}/commands
	return fmt.Sprintf("%s/devices/%s/commands", BaseURL, deviceId)
}

func SceneExecuteUrl(sceneId string) string {
	// /scenes/{SceneId}/execute
	return fmt.Sprintf("%s/scenes/%s/execute", BaseURL, sceneId)
}
