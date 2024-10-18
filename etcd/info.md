## Release-3.5
### Dir
```
.
├── api                     # module go.etcd.io/etcd/api/v3
│   ├── authpb
│   ├── etcdserverpb
│   │   └── gw
│   ├── membershippb
│   ├── mvccpb
│   ├── v3rpc
│   │   └── rpctypes
│   └── version
├── client
│   ├── pkg                 # module go.etcd.io/etcd/client/pkg/v3
│   │   ├── fileutil
│   │   ├── logutil
│   │   ├── pathutil
│   │   ├── srv
│   │   ├── systemd
│   │   ├── testutil
│   │   ├── tlsutil
│   │   ├── transport
│   │   └── types
│   ├── v2                  # module go.etcd.io/etcd/client/v2
│   └── v3                  # module go.etcd.io/etcd/client/v3
│       ├── clientv3util
│       ├── concurrency
│       ├── credentials
│       ├── experimental
│       │   └── recipes
│       │       └── grpc_gateway
│       ├── internal
│       │   ├── endpoint
│       │   └── resolver
│       ├── leasing
│       ├── mirror
│       ├── mock
│       │   └── mockserver
│       ├── namespace
│       ├── naming
│       │   ├── endpoints
│       │   │   └── internal
│       │   └── resolver
│       ├── ordering
│       ├── snapshot
│       └── yaml
├── contrib
│   ├── lock
│   │   ├── client
│   │   └── storage
│   ├── mixin
│   ├── raftexample
│   └── systemd
│       ├── etcd3-multinode
│       └── sysusers.d
├── Documentation
│   └── dev-guide
│       └── apispec
│           └── swagger
│               ├── api
│               │   └── etcdserverpb
│               └── server
│                   └── etcdserver
│                       └── api
│                           ├── v3election
│                           │   └── v3electionpb
│                           └── v3lock
│                               └── v3lockpb
├── etcdctl                 # module go.etcd.io/etcd/etcdctl/v3
│   ├── ctlv2
│   │   └── command
│   ├── ctlv3
│   │   └── command
│   └── doc
├── etcdutl                 # module go.etcd.io/etcd/etcdutl/v3
│   ├── etcdutl
│   └── snapshot
├── hack
│   ├── benchmark
│   ├── insta-discovery
│   ├── kubernetes-deploy
│   ├── patch
│   └── tls-setup
│       └── config
├── logos
├── pkg                     # module go.etcd.io/etcd/pkg/v3
│   ├── adt
│   │   └── img
│   ├── cobrautl
│   ├── contention
│   ├── cpuutil
│   ├── crc
│   ├── debugutil
│   ├── expect
│   ├── flags
│   ├── grpc_testing
│   ├── httputil
│   ├── idutil
│   ├── ioutil
│   ├── netutil
│   ├── osutil
│   ├── pbutil
│   ├── proxy
│   │   └── fixtures
│   ├── report
│   ├── runtime
│   ├── schedule
│   ├── stringutil
│   ├── traceutil
│   └── wait
├── raft                    # module go.etcd.io/etcd/raft/v3
│   ├── confchange
│   │   └── testdata
│   ├── quorum
│   │   └── testdata
│   ├── raftpb
│   ├── rafttest
│   ├── testdata
│   └── tracker
├── scripts
├── security
├── server                  # module go.etcd.io/etcd/server/v3
│   ├── auth
│   ├── config
│   ├── datadir
│   ├── default.etcd
│   │   └── member
│   │       ├── snap
│   │       └── wal
│   ├── embed
│   ├── etcdmain
│   ├── etcdserver
│   │   ├── api
│   │   │   ├── etcdhttp
│   │   │   ├── membership
│   │   │   ├── rafthttp
│   │   │   ├── snap
│   │   │   │   └── snappb
│   │   │   ├── v2auth
│   │   │   ├── v2discovery
│   │   │   ├── v2error
│   │   │   ├── v2http
│   │   │   │   ├── httptypes
│   │   │   │   └── testdata
│   │   │   ├── v2stats
│   │   │   ├── v2store
│   │   │   ├── v2v3
│   │   │   ├── v3alarm
│   │   │   ├── v3client
│   │   │   ├── v3compactor
│   │   │   ├── v3election
│   │   │   │   └── v3electionpb
│   │   │   │       └── gw
│   │   │   ├── v3lock
│   │   │   │   └── v3lockpb
│   │   │   │       └── gw
│   │   │   └── v3rpc
│   │   └── cindex
│   ├── lease
│   │   ├── leasehttp
│   │   └── leasepb
│   ├── mock
│   │   ├── mockstorage
│   │   ├── mockstore
│   │   └── mockwait
│   ├── mvcc
│   │   ├── backend
│   │   │   └── testing
│   │   └── buckets
│   ├── proxy
│   │   ├── grpcproxy
│   │   │   ├── adapter
│   │   │   └── cache
│   │   ├── httpproxy
│   │   └── tcpproxy
│   ├── storage
│   │   └── mvcc
│   │       └── testutil
│   ├── verify
│   └── wal
│       └── walpb
├── tests                   # module go.etcd.io/etcd/tests/v3
│   ├── docker-dns
│   │   ├── certs
│   │   ├── certs-common-name-auth
│   │   ├── certs-common-name-multi
│   │   ├── certs-gateway
│   │   ├── certs-san-dns
│   │   ├── certs-wildcard
│   │   └── insecure
│   ├── docker-dns-srv
│   │   ├── certs
│   │   ├── certs-gateway
│   │   └── certs-wildcard
│   ├── docker-static-ip
│   │   ├── certs
│   │   └── certs-metrics-proxy
│   ├── e2e
│   ├── fixtures
│   ├── framework
│   │   ├── e2e
│   │   └── testutils
│   ├── functional
│   │   ├── agent
│   │   ├── cmd
│   │   │   ├── etcd-agent
│   │   │   ├── etcd-proxy
│   │   │   ├── etcd-runner
│   │   │   └── etcd-tester
│   │   ├── rpcpb
│   │   ├── runner
│   │   ├── scripts
│   │   └── tester
│   ├── integration
│   │   ├── client
│   │   │   └── examples
│   │   ├── clientv3
│   │   │   ├── concurrency
│   │   │   ├── connectivity
│   │   │   ├── examples
│   │   │   ├── experimental
│   │   │   │   └── recipes
│   │   │   ├── lease
│   │   │   ├── naming
│   │   │   └── snapshot
│   │   ├── embed
│   │   ├── fixtures-expired
│   │   ├── proxy
│   │   │   └── grpcproxy
│   │   ├── snapshot
│   │   │   └── testdata
│   │   └── v2store
│   └── testutils
└── tools
    ├── benchmark
    │   └── cmd
    ├── etcd-dump-db
    ├── etcd-dump-logs
    │   ├── expectedoutput
    │   └── testdecoder
    ├── etcd-dump-metrics
    ├── local-tester
    │   └── bridge
    └── mod
```