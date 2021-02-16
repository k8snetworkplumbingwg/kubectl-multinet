# kubectl-podnet
[![Build](https://github.com/k8snetworkplumbingwg/kubectl-podnet/actions/workflows/build.yml/badge.svg)](https://github.com/k8snetworkplumbingwg/kubectl-podnet/actions/workflows/build.yml)[![Test](https://github.com/k8snetworkplumbingwg/kubectl-podnet/actions/workflows/test.yml/badge.svg)](https://github.com/k8snetworkplumbingwg/kubectl-podnet/actions/workflows/test.yml)

This is a kubectl plugin which outputs Pods' network-status of [multi-net-spec](https://github.com/k8snetworkplumbingwg/multi-net-spec).

## Description

kubectl-podnet shows Pods' network-status annotation of [multi-net-spec](https://github.com/k8snetworkplumbingwg/multi-net-spec).

## Usage

Show pod network status, in namespace 'default'

```
$ kubectl podnet
POD      NET                    IF     ADDRESS         MAC
centos                          eth0   [10.244.1.16]   0a:f7:47:2c:a6:dd
centos   default/macvlan-net1   net1   [10.1.1.11]     ae:e8:85:88:bb:d0
nginx                           eth0   [10.244.2.16]   c2:d5:70:24:b8:68
nginx    default/macvlan-net1   net1   [10.1.1.13]     fe:f2:5c:4a:ce:62
```

Show pod network status, in namespace 'test1'

```
$ kubectl podnet -n test1
POD            NET                          IF     ADDRESS         MAC
test1-centos                                eth0   [10.244.1.17]   02:47:1b:dd:9d:f4
test1-centos   test1/test1-macvlan-conf-1   net1   [20.1.1.101]    9e:4b:ed:a9:07:07
```

Show pod network status, in namespace 'default', JSON format (note: JSON format is experimental, so format might be changed)

```
$ kubectl podnet -o json
[{"pod":"centos","net":"","interface":"eth0","address":["10.244.1.16"],"mac":"0a:f7:47:2c:a6:dd"},{"pod":"centos","net":"default/macvlan-net1","interface":"net1","address":["10.1.1.11"],"mac":"ae:e8:85:88:bb:d0"},{"pod":"nginx","net":"","interface":"eth0","address":["10.244.2.16"],"mac":"c2:d5:70:24:b8:68"},{"pod":"nginx","net":"default/macvlan-net1","interface":"net1","address":["10.1.1.13"],"mac":"fe:f2:5c:4a:ce:62"}]
```

## Installation as kubectl plugin

See [Installing kubectl plugins](https://kubernetes.io/docs/tasks/extend-kubectl/kubectl-plugins/#installing-kubectl-plugins).
Krew support is TBD for now.

## License

This software is released under the Apache License 2.0.
