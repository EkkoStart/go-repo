module repo

go 1.23.5

require github.com/EkkoStart/go-repo/wlog v1.0.1

require github.com/EkkoStart/go-repo/werrors v1.0.0

require (
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/modern-go/concurrent v0.0.0-20180228061459-e0a39a4cb421 // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
)

replace github.com/EkkoStart/go-repo/wlog => ./wlog

replace github.com/EkkoStart/go-repo/werrors => ./werrors
