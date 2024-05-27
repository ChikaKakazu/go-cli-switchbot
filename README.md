# go-cli-switchbot
SwitchBotを操作するCLIです。
- [Authentication](#authentication)
- [Authentication Settings](#authentication-settings)
    - [token set](#token-set)
    - [token get](#token-get)
- [Getting Sigh](#getting-sigh)
    - [sign](#sign)
- [Getting Devices](#getting-devices)
    - [devices](#devices)
- [Control Bot](#control-bot)
    - [bot](#bot)
- [Control Humidifier](#control-humidifier)
    - [humidifier](#humidifier)
- [Execute Scene](#execute-scene)
    - [scene list](#scene-list)
    - [scene exec](#scene-exec)

## Authentication
SwitchBotApiを操作するには、TokenとSecretKeyが必要です。
[SwitchBotAPI](https://github.com/OpenWonderLabs/SwitchBotAPI?tab=readme-ov-file#getting-started)をのReadMeを参考にTokenとSecretKeyを取得してください。

## Authentication Settings

### token set
- `token set`コマンドでTokenとSecretを保存します
- TokenとSecretを入力します
- これは一度入力すると保存されます。また、`token set`コマンドで何度でも設定し直すことができます
```sh
# 入力されている文字列は疑似値です
./go-cli-switchbot token set
Enter your token: 6b67d2fe1a592db1181102576e5e97be0e136c7c07b7973a5efeaffc5e2d010ca4f4954de9c7d2f62f01d61
Enter your secret: 72e728e58f7064950bd81c6c0
Token and Secret set successfully.
```

### token get
- `token get`コマンドで保存したTokenとSecretを表示します
```sh
./go-cli-switchbot token get
Token:  6b67d2fe1a592db1181102576e5e97be0e136c7c07b7973a5efeaffc5e2d010ca4f4954de9c7d2f62f01d61
Secret:  72e728e58f7064950bd81c6c0
```

## Getting Sigh
SwitchBotApiへApiをリクエストするにはヘッダー情報に以下の情報を含める必要があります。
| Parameter | Type |
| --------- | ---- |
| Authorization | String |
| sign | String |
| t    | Long   |
| nonce| Long   |

### sign
- `sign`コマンドで、リクエストに必要なヘッダー情報を取得します
- このCLIで対応していないApiを呼び出す際に利用してください
```sh
./go-cli-switchbot sign
Token:  6b67d2fe1a592db1181102576e5e97be0e136c7c07b7973a5efeaffc5e2d010ca4f4954de9c7d2f62f01d61
Signature: QfzSS7PRBukzjh8bHjs44WTxcPGbgEQWQsqXfT=
Time:  1716736135779
Nonce:  79d5be9-659f-43b-915-0c881eda815
```

## Getting Devices
### devices
- `devices`コマンドで自身が利用しているSwitchBotデバイスの一覧を返します
- 各デバイスに対して操作を行う場合、deviceIdが必要になります。このCLIで対応していないApiを呼び出す際でも利用してください
```sh
./go-cli-switchbot devices
{
  "body": {
    "deviceList": [
      {
        "deviceId": "943CC68C85E",
        "deviceName": "加湿器 ",
        "deviceType": "Humidifier",
        "enableCloudService": true,
        "hubDeviceId": "000000000000"
      },
      {
        "deviceId": "C49B7555C1F",
        "deviceName": "ハブミニ",
        "deviceType": "Hub Mini",
        "enableCloudService": false,
        "hubDeviceId": "000000000000"
      },
      {
        "deviceId": "EB657390730",
        "deviceName": "温湿度計",
        "deviceType": "Meter",
        "enableCloudService": true,
        "hubDeviceId": "C49B75559C1F"
      },
      {
        "deviceId": "F62E81F2571",
        "deviceName": "部屋の電気ボット",
        "deviceType": "Bot",
        "enableCloudService": true,
        "hubDeviceId": "C49B75559C1F"
      }
    ]
  },
  "message": "success",
  "statusCode": 100
}
```

## Control Bot
SwitchBot ボットの操作を行います。具体的にはスイッチのオン・オフを行います。

### bot
- `bot`コマンドで自身が利用しているボットの一覧を表示します。その中から操作したいボットを選んでください
    ```sh
    ./go-cli-switchbot bot
    Use the arrow keys to navigate: ↓ ↑ → ← 
    ? Select a bot device: 
    ▸ F62E81F2571: 部屋の電気ボット
    　C32D8HF2550: 寝室の電気ボット
    ```
- 選んだボットのスイッチに対してオン・オフを実行するコマンドが表示されるので選んでください
    ```sh
    ✔ F62E81F2571: 部屋の電気ボット
    Use the arrow keys to navigate: ↓ ↑ → ← 
    ? Select Action: 
    ▸ Turn off
      Turn on
    ```

## Control Humidifier
SwitchBot 加湿器の操作を行います。具体的にはスイッチのオン・オフを行います。

### humidifier
- `humidifier`コマンドで自身が利用している加湿器の一覧を表示します。その中から操作したい加湿器を選んでください
  ```sh
  ./go-cli-switchbot humidifier
  Use the arrow keys to navigate: ↓ ↑ → ← 
  ? Select Humidifier: 
    ▸ 943C6885FE: 加湿器 
  ```
- 選んだ加湿器に対してオン・オフを実行するコマンドが表示されるので選んでください
  ```sh
  ✔ 943C6885FE: 加湿器 
  Use the arrow keys to navigate: ↓ ↑ → ← 
  ? Select Action: 
    ▸ Turn on
      Turn off
  ```

## Execute Scene
シーンの一覧表示と実行を行います。登録してあるシーンの中から選択したシーンを実行します。

### scene list
- `scene list`コマンドで自身が登録してあるシーンの一覧を表示します。
```sh
./go-cli-switchbot scene list
{
  "body": [
    {
      "SceneSelectName": "",
      "sceneId": "ce8b607-aa60-4631-a30-408cfcf1985",
      "sceneName": "電気とエアコンをつける"
    },
    {
      "SceneSelectName": "",
      "sceneId": "283e9a6-1d68-450b-bb3-e7cdedd2a02",
      "sceneName": "電気とエアコンを消す"
    }
  ],
  "message": "success",
  "statusCode": 100
}
```

### scene exec
- `scene exec`コマンドで登録してあるシーンの中から実行したいシーンを選択し、実行します。
```sh
./go-cli-switchbot scene exec
Use the arrow keys to navigate: ↓ ↑ → ← 
? Select Scene: 
  ▸ "ce8b607-aa60-4631-a30-408cfcf1985: 電気とエアコンをつける
    283e9a6-1d68-450b-bb3-e7cdedd2a02: 電気とエアコンを消す
```