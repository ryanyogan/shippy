module github.com/ryanyogan/shippy/shippy-cli-consignment

go 1.14

replace github.com/ryanyogan/shippy/shippy-service-consignment => ../shippy-service-consignment

require (
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/ryanyogan/shippy/shippy-service-consignment v0.0.0-20200619044641-2ef30574e6c0
	google.golang.org/genproto v0.0.0-20200619004808-3e7fca5c55db // indirect
	google.golang.org/grpc v1.29.1
)
