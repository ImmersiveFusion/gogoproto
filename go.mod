module github.com/cosmos/gogoproto

go 1.19

require (
	github.com/google/go-cmp v0.6.0
	golang.org/x/exp v0.0.0-20230811145659-89c5cff77bcb
	google.golang.org/grpc v1.61.1
	google.golang.org/protobuf v1.31.0
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.18.0 // indirect
	golang.org/x/sys v0.14.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231106174013-bbf56f31fb17 // indirect
)

// API changed in an incompatible way
retract v1.4.8
