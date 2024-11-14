module github.com/coinbase-samples/exchange-sdk-go

go 1.22.2

//require github.com/coinbase-samples/core-go v0.1.0

replace github.com/coinbase-samples/core-go => github.com/rcbgr/core-go v0.0.0-20241030193543-031fda14e3ff

//require github.com/coinbase-samples/core-go v0.0.0-00010101000000-000000000000

require github.com/coinbase-samples/core-go v0.0.0-00010101000000-000000000000

require (
	github.com/gorilla/websocket v1.5.3 // indirect
	github.com/shopspring/decimal v1.4.0 // indirect
	golang.org/x/net v0.30.0 // indirect
	golang.org/x/text v0.19.0 // indirect
)
