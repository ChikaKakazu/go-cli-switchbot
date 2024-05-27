/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/ChikaKakazu/go-cli-switchbot/domain"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// sceneCmd represents the scene command
var sceneCmd = &cobra.Command{
	Use:   "scene [list|exec]",
	Short: "Interact with Scene devices. You can `list` scenes or `exec` a scene.",
	Long:  `Interact with Scene devices. You can list scenes or execute a scene.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please specify 'scene list' or 'scene exec'")
			return
		}
		switch args[0] {
		case "list":
			sceneList()
		case "exec":
			scene()
		default:
			fmt.Println("Invalid argument. Please specify 'scene list' or 'scene exec'")
		}
	},
}

func init() {
	rootCmd.AddCommand(sceneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sceneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sceneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func scene() {
	scenes, signReq := domain.GetScenes()
	if len(scenes) == 0 {
		return
	}

	// scene.SceneSelectNameを書き換える
	var selectScenes []domain.Scene
	for _, scene := range scenes {
		scene.SceneSelectName = fmt.Sprintf("%s: %s", scene.SceneId, scene.SceneName)
		selectScenes = append(selectScenes, scene)
	}

	if len(selectScenes) == 0 {
		fmt.Println("No scenes found")
		return
	}

	var items []string
	for _, scene := range selectScenes {
		items = append(items, scene.SceneSelectName)
	}

	// sceneの一覧を表示する
	prompt := promptui.Select{
		Label: "Select Scene",
		Items: items,
	}

	_, result, err := prompt.Run()
	if err != nil {
		fmt.Println("Prompt failed: ", err)
		return
	}

	var scene *domain.Scene
	for _, s := range selectScenes {
		if s.SceneSelectName == result {
			scene = &s
			break
		}
	}

	if scene == nil {
		fmt.Println("Failed to get scene")
		return
	}

	// 選択されたsceneを実行する
	resp, _ := scene.ExecScene(signReq)

	fmt.Println(string(resp))
}

func sceneList() {
	sceneList, _ := domain.GetSceneList()

	var result map[string]interface{}
	byteSceneList, _ := json.Marshal(sceneList)
	err := json.Unmarshal(byteSceneList, &result)
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
