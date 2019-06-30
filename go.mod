module sample_model

go 1.12

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.21.1

replace google.golang.org/appengine => github.com/golang/appengine v1.6.1

replace google.golang.org/genproto => github.com/ilisin/genproto v0.0.0-20181026194446-8b5d7a19e2d9

replace cloud.google.com/go => github.com/googleapis/google-cloud-go v0.40.1-0.20190612163021-8a8d2c2fb096

replace golang.org/x/lint => github.com/golang/lint v0.0.0-20190313153728-d0100b6bd8b3

replace golang.org/x/sys => github.com/golang/sys v0.0.0-20190616124812-15dcb6c0061f

replace golang.org/x/text => github.com/golang/text v0.3.2

replace golang.org/x/sync => github.com/golang/sync v0.0.0-20190423024810-112230192c58

replace golang.org/x/oauth2 => github.com/golang/oauth2 v0.0.0-20190604053449-0f29369cfe45

replace golang.org/x/tools => github.com/golang/tools v0.0.0-20190614205625-5aca471b1d59

replace golang.org/x/net => github.com/golang/net v0.0.0-20190613194153-d28f0bde5980

replace golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190611184440-5c40567a22f8

replace golang.org/x/time => github.com/golang/time v0.0.0-20190308202827-9d24e82272b4

require (
	github.com/golang/protobuf v1.3.1
	google.golang.org/grpc v1.20.1
)
