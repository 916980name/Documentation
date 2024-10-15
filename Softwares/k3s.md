### Errors
1. When compiling `make`

    1.  go: golang.org/x/tools/cmd/goimports@gopls/v0.11.0: unrecognized import path "golang.org/x/tools/cmd/goimports": https fetch: Get "https://golang.org/x/tools/cmd/goimports?go-get=1": dial tcp: lookup golang.org on [fd33:5417:76de::1]:53: dial udp [fd33:5417:76de::1]:53: connect: network is unreachable
        > disable IPv6 on the host

    1. Git not found
        > docker cached layer error, need force rebuild, add ARG in dockerfile
        ```
        ARG GOLANG=golang:1.22.6-alpine3.20
        FROM ${GOLANG}
        ...

        ##### Add this:
        ARG CACHEBUST=1
        ##### You add a CACHEBUST argument to your Dockerfile at the location you want to enforce a rebuild.

        # Install necessary packages
        RUN apk -U --no-cache add \
        ...
        ```

1. When running

    1. FATA[0114] failed to retrieve agent configuration: failed to find host-local: exec: "host-local": executable file not found in $PATH