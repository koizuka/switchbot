module github.com/koizuka/switchbot/v2

go 1.24

require (
	github.com/BurntSushi/toml v1.5.0
	github.com/nasa9084/go-switchbot/v3 v3.2.0
)

require github.com/google/uuid v1.6.0 // indirect

replace github.com/nasa9084/go-switchbot/v3 => github.com/koizuka/go-switchbot/v3 v3.0.2-0.20230805050340-e721d4da7f5d
