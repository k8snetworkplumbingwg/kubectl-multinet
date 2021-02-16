module github.com/k8snetworkplumbingwg/kubectl-podnet

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
	golang.org/x/sys v0.0.0-20210216163648-f7da38b97c65 // indirect
	golang.org/x/text v0.3.4 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	k8s.io/apimachinery v0.19.7
	k8s.io/cli-runtime v0.19.7
	k8s.io/client-go v0.19.7
	k8s.io/klog/v2 v2.5.0 // indirect
	k8s.io/kube-openapi v0.0.0-20210211043216-66d8d84e87dd // indirect
)

replace (
	github.com/go-openapi/spec => github.com/go-openapi/spec v0.19.5
	sigs.k8s.io/kustomize/api => sigs.k8s.io/kustomize/api v0.8.1
)
