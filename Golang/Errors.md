### 1. docker
docker run returns
```
exec /opt/app/run.bin: no such file or directory
```
manually execute with `docker run -it --entrypoint=/bin/sh  916980name/demo-httpserver:v1` returns
```
/bin/sh: ./run.bin: not found
```

reason: [a typical symptom of dynamic link failure](https://stackoverflow.com/questions/66963068/docker-alpine-executable-binary-not-found-even-if-in-path?noredirect=1&lq=1)

check executable binary file: `file exefile` returns
```
exefile: ELF 64-bit LSB executable, x86-64, version 1 (SYSV), statically linked, Go BuildID=M-Ae_vnjHZcR58Dks3dB/GZAZBhvHyCiK9ErpApZl/9GtrcjYgXtLbec4Ugxyp/cwTHOGZSriDAJh1Gf6fP, with debug_info, not stripped
```

check un-executable binary file: `file unexefile` returns
```
unexefile: ELF 64-bit LSB executable, x86-64, version 1 (SYSV), dynamically linked, interpreter /lib64/ld-linux-x86-64.so.2, Go BuildID=JCBkK1O2Eqw0Ml-E5JcQ/WO_XGIynh7vepDPkIuds/DW_Z7R1XJn05jJSKFzf7/FyfwhIdvlk3Iw8bzQ0n4, with debug_info, not stripped
```

solution: `@CGO_ENABLED=0 go build ...`