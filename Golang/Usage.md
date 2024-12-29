#### Option
see `go.uber.org/zap/zapcore.Option`

#### discoverEndpoints()
- startGateway
- discoverEndpoints
- startEtcd(cfg *embed.Config)

#### 一种request timeout的方式
`https://github.com/kubernetes/kubernetes/blob/3a4c35cc89c0ce132f8f5962ce4b9a48fae77873/staging/src/k8s.io/apiserver/pkg/endpoints/handlers/finisher/finisher.go#L87-L122`

#### 处理http response
`net/http/response.go#ReadResponse`

#### 特性开关
staging/src/k8s.io/component-base/featuregate/feature_gate.go#L109

### Debug Build
`-ldflags=-compressdwarf=false -gcflags "all=-N -l"`
- -N, turning off optimization
- -l, disable inlining
- -ldflags ... false`, generate uncompressed debug information

### pprof
1. trace
    ```
    go tool trace trace.out
    ```
1. cpu & mem
    ```
    go tool pprof [xxx].prof
    ```
1. webui
    ```
    go tool pprof -http=:8080  [binary file] ./[xxx].prof
    ```
1. webui with running
    ```golang
    // 先启动 pprof http 服务器
    go func() {
        http.ListenAndServe(":6060", nil)
    }()
    ```
    ```
    http://localhost:6060/debug/pprof/
    ```