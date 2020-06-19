module shippy-cli-consignment

go 1.14

replace github.com/ryanyogan/shippy/shippy-service-consignment => ../shippy-service-consignment

require (
	github.com/ryanyogan/shippy/shippy-service-consignment v0.0.0-00010101000000-000000000000
	google.golang.org/genproto v0.0.0-20200619004808-3e7fca5c55db // indirect
	google.golang.org/grpc v1.29.1
)
