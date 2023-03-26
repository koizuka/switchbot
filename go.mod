module github.com/koizuka/switchbot/v2

go 1.18

require (
	github.com/BurntSushi/toml v1.2.1
	github.com/nasa9084/go-switchbot/v2 v2.1.0
)

require github.com/google/uuid v1.3.0 // indirect

replace github.com/nasa9084/go-switchbot/v2 v2.1.0 => github.com/koizuka/go-switchbot/v2 v2.1.1-0.20230326100016-3c0abdfaea9c
