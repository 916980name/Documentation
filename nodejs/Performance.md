1. in node shell
    ```
    > require('child_process').execSync('ps -Af | grep -q -E -c "\\-\\-user-data-dir=\\.+App"').stdout.toString()
    ```
    tracing in bash shell
    ```
    $ pgrep -l node
    6521 node
    $ strace -t -f -p 6521 -o strace_err.txt
    ```
    The `exec` will search everything in `$PATH` until find the binary.
    ```
    7500  09:47:09 newfstatat(AT_FDCWD, "/home/[x]/.local/share/pnpm/ps", 0x7ffd65dea9e0, 0) = -1 ENOENT (No such file or directory)
    7500  09:47:09 newfstatat(AT_FDCWD, "/home/[x]/Softwares/CNI-plugins/ps", 0x7ffd65dea9e0, 0) = -1 ENOENT (No such file or directory)
    7500  09:47:09 newfstatat(AT_FDCWD, "/home/[x]/go/bin/ps", 0x7ffd65dea9e0, 0) = -1 ENOENT (No such file or directory)
    7500  09:47:09 newfstatat(AT_FDCWD, "/home/[x]/Softwares/go/bin/ps", 0x7ffd65dea9e0, 0) = -1 ENOENT (No such file or directory)
    7500  09:47:09 newfstatat(AT_FDCWD, "/home/[x]/Softwares/node-v20.14.0/bin/ps", 0x7ffd65dea9e0, 0) = -1 ENOENT (No such file or directory)
    7500  09:47:09 newfstatat(AT_FDCWD, "/home/[x]/Softwares/jdk-11.0.2/bin/ps", 0x7ffd65dea9e0, 0) = -1 ENOENT (No such file or directory)
    7500  09:47:09 newfstatat(AT_FDCWD, "/home/[x]/Softwares/gradle-8.0.2/bin/ps", 0x7ffd65dea9e0, 0) = -1 ENOENT (No such file or directory)
    7500  09:47:09 newfstatat(AT_FDCWD, "/home/[x]/.local/bin/ps", 0x7ffd65dea9e0, 0) = -1 ENOENT (No such file or directory)
    7500  09:47:09 newfstatat(AT_FDCWD, "/usr/local/sbin/ps", 0x7ffd65dea9e0, 0) = -1 ENOENT (No such file or directory)
    7500  09:47:09 newfstatat(AT_FDCWD, "/usr/local/bin/ps", 0x7ffd65dea9e0, 0) = -1 ENOENT (No such file or directory)
    7500  09:47:09 newfstatat(AT_FDCWD, "/usr/sbin/ps", 0x7ffd65dea9e0, 0) = -1 ENOENT (No such file or directory)
    7500  09:47:09 newfstatat(AT_FDCWD, "/usr/bin/ps", {st_mode=S_IFREG|0755, st_size=141776, ...}, 0) = 0
    ```
