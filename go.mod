module github.com/k8snetworkplumbingwg/kubectl-multinet

go 1.14

require (
	github.com/emicklei/go-restful v2.16.0+incompatible // indirect
	github.com/k8snetworkplumbingwg/network-attachment-definition-client v1.1.1-0.20201119153432-9d213757d22d
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/spf13/cobra v1.4.0
	github.com/spf13/pflag v1.0.5
	k8s.io/apimachinery v0.24.10
	k8s.io/cli-runtime v0.24.10
	k8s.io/client-go v0.24.10
)

replace (
	github.com/Azure/go-autorest/autorest => github.com/Azure/go-autorest/autorest v0.11.28
	github.com/containernetworking/cni => github.com/containernetworking/cni v0.8.1
	github.com/go-openapi/spec => github.com/go-openapi/spec v0.19.5
	github.com/gogo/protobuf => github.com/gogo/protobuf v1.3.2
	golang.org/x/text => golang.org/x/text v0.3.8
	sigs.k8s.io/kustomize/api => sigs.k8s.io/kustomize/api v0.12.1
)
