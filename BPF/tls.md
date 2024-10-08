1. Get the binary tls lib
    ```
    ldd `which wget` | grep -E 'tls|ssl|nss|nspr'
    ```
    ```
    pldd [pid] | grep -E 'tls|ssl|nss|nspr'
    ```