1. Get the binary tls lib
    ```
    ldd `which wget` | grep -E 'tls|ssl|nss|nspr'
    ```
    ```
    pldd [pid] | grep -E 'tls|ssl|nss|nspr'
    ```
1. View symbols of ELF
    ``` 
    nm [binary]

    readelf：
    -h：文件头
    -S：段表
    -s：符号表
    -d: 查看依赖库
    -p：查看某个段内容，非常重要。如：readelf -p .comment libc.so （通过-p对只读段的查看就可以替代strings命令）

    objdump：
    -d：反汇编（objdump我基本只用这一个）
    -h：段表，同readelf -S，所以可以不用记
    –s：代码段、数据段、只读数据段，各个段二进制
    -a：看一个.a静态库文件中包含了哪些目标文件

    od：
    如：十六进制输出数据并且地址以十进制打印：od -A d -t xCc 文件
    命令中各选项的含义：
    - A 指定地址基数，包括：
    d 十进制
    o 八进制（系统默认值）
    x 十六进制
    n 不打印位移值
    - t 指定数据的显示格式，主要的参数有：
    c ASCII字符或反斜杠序列
    d 有符号十进制数
    f 浮点数
    o 八进制（系统默认值为02）
    u 无符号十进制数
    x 十六进制数
    ```