# switchbot CLI

## about
CLI tool for SwitchBot API(currently available only two commands: list and status)

## prepare your keys
1. get open token and secret key from SwitchBot app (see [SwitchBot API Docs](https://github.com/OpenWonderLabs/SwitchBotAPI#getting-started))
2. create config.toml in current directory:
```toml
openToken="your open token"
secretKey="your secret key"
```
## build
```shell
go build
```

## usage
### `list`: list devices
```shell
./switchbot list
```
will shown like below(JSON array):
<details>

```json
[
  {
    "deviceId": "DEVICE_ID",
    "deviceName": "Device Name",
    "deviceType": "Humidifier",
    "enableCloudService": true,
    "hubDeviceId": "000000000000"
  },
  {
    "deviceId": "DEVICE_ID",
    "deviceName": "Device Name",
    "deviceType": "Humidifier",
    "hubDeviceId": "000000000000"
  },
  {
    "deviceId": "DEVICE_ID",
    "deviceName": "Device Name",
    "deviceType": "Plug Mini (JP)",
    "enableCloudService": true
  },
  {
    "deviceId": "DEVICE_ID",
    "deviceName": "Device Name",
    "deviceType": "Meter",
    "enableCloudService": true,
    "hubDeviceId": "HUB_DEVICE_ID"
  },
  {
    "deviceId": "DEVICE_ID",
    "deviceName": "Device Name",
    "deviceType": "Blind Tilt",
    "hubDeviceId": "000000000000",
    "master": true
  },
  {
    "deviceId": "DEVICE_ID",
    "deviceName": "Device Name",
    "deviceType": "Meter",
    "enableCloudService": true,
    "hubDeviceId": "HUB_DEVICE_ID"
  },
  {
    "deviceId": "DEVICE_ID",
    "deviceName": "Device Name",
    "deviceType": "Hub Mini",
    "hubDeviceId": "000000000000"
  },
  {
    "deviceId": "DEVICE_ID",
    "deviceName": "Device Name",
    "deviceType": "Plug",
    "enableCloudService": true,
    "hubDeviceId": "000000000000"
  }
]
```
</details>


### `status`: get device status
```shell
./switchbot status device_id [device_id ...]
```
will shown like below(JSON array):
<details>

```json
[{"deviceId":"DEVICE_ID","deviceType":"Meter","hubDeviceId":"HUB_DEVICE_ID","humidity":43,"temperature":24.3,"brightness":{},"battery":100}]
```
</details>
