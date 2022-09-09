module github.com/k8snetworkplumbingwg/kubectl-multinet

go 1.14

require (
	github.com/emicklei/go-restful v2.15.0+incompatible // indirect
	github.com/go-openapi/spec v0.20.3 // indirect
	github.com/google/go-cmp v0.5.2 // indirect
	github.com/k8snetworkplumbingwg/network-attachment-definition-client v1.1.1-0.20201119153432-9d213757d22d
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/spf13/cobra v1.1.1
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.6.1 // indirect
	k8s.io/apimachinery v0.19.7
	k8s.io/cli-runtime v0.19.7
	k8s.io/client-go v0.19.7
	k8s.io/klog/v2 v2.5.0 // indirect
	k8s.io/kube-openapi v0.0.0-20210211043216-66d8d84e87dd // indirect
)

replace (
	github.com/Azure/go-autorest/autorest => github.com/Azure/go-autorest/autorest v0.11.28
	github.com/containernetworking/cni => github.com/containernetworking/cni v0.8.1
	github.com/go-openapi/spec => github.com/go-openapi/spec v0.19.5
	github.com/gogo/protobuf => github.com/gogo/protobuf v1.3.2
	sigs.k8s.io/kustomize/api => sigs.k8s.io/kustomize/api v0.8.1
)
