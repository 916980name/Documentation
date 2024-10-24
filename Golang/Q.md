### params copy
1. One int param
    ```go
    // a = 10
    func ptest2(a int) int
    ```
    The param stored in register `$rax`
1. Three int param

    The three args already in `$rax, $rbx, $rcx`
    ```go
    func sum3(a, b, c int) int {    0x5ad1e0 push %rbp              main.sum3+0
                    [build stack]   0x5ad1e1 mov  %rsp,%rbp         main.sum3+1
                    -------------   0x5ad1e4 sub  $0x8,%rsp         main.sum3+4
                                    0x5ad1e8 mov  %rax,0x18(%rsp)   main.sum3+8
     [copy args to new stack mem]   0x5ad1ed mov  %rbx,0x20(%rsp)   main.sum3+13
                    -------------   0x5ad1f2 mov  %rcx,0x28(%rsp)   main.sum3+18
                                    0x5ad1f7 movq $0x0,(%rsp)       main.sum3+23

    return a + b + c                0x5ad1ff lea (%rax,%rbx,1),%rdx main.sum3+31
               [return v -> $rax]   0x5ad203 lea (%rdx,%rcx,1),%rax main.sum3+35
                                    0x5ad207 mov %rax,(%rsp)        main.sum3+39
                    -------------   0x5ad20b add $0x8,%rsp          main.sum3+43
                  [recover stack]   0x5ad20f pop %rbp               main.sum3+47
                                    0x5ad210 ret                    main.sum3+48
    }
    [return v -> caller stack mem]  0x5ad146 mov %rax,0x18(%rsp)    main.ptest3+70
    ```

1. One string param

    String memory address in `$rax`, string length in `$rbx`
    ```go
    func string4(s string) string {
    [cp str addr]   0x5ad1f5 mov %rax,0x8(%rsp)     main.string4+85
    [cp str len]    0x5ad1fa mov %rbx,0x10(%rsp)    main.string4+90
                    0x5ad1ff nop                    main.string4+95
                    0x5ad200 call 0x478680 <runtime.morestack_noctxt>main.string4+96
                    0x5ad205 mov 0x8(%rsp),%rax     main.string4+101
                    0x5ad20a mov 0x10(%rsp),%rbx    main.string4+106
                    0x5ad20f jmp 0x5ad1a0 <main.string4>main.string4+111
        s += " pesicola"
                    0x5ad1be mov 0x48(%rsp),%rbx    main.string4+30
                    0x5ad1c3 mov 0x50(%rsp),%rcx    main.string4+35
                    0x5ad1c8 xor %eax,%eax          main.string4+40
                    0x5ad1ca lea 0x45f7d(%rip),%rdi # 0x5f314e  main.string4+42
                    0x5ad1d1 mov $0x9,%esi          main.string4+49
    [str cp here]   0x5ad1d6 call 0x4593e0 <runtime.concatstring2>main.string4+54
                    0x5ad1db mov %rax,0x48(%rsp)    main.string4+59
                    0x5ad1e0 mov %rbx,0x50(%rsp)    main.string4+64
        return s
                    0x5ad1e5 mov %rax,0x28(%rsp)    main.string4+69
                    0x5ad1ea mov %rbx,0x30(%rsp)    main.string4+74
                    0x5ad1ef add $0x38,%rsp         main.string4+79
                    0x5ad1f3 pop %rbp               main.string4+83
                    0x5ad1f4 ret                    main.string4+84
    }
    ```

1. One pointer param

    Struct
    ```golang
    type paramstestStructInner struct {
        a bool
    }
    type paramstestStruct struct {
        a int
        b float32
        c string
        d bool
        e paramstestStructInner
    }
    ```

    Pointer address in `$rax`, address copied from reg to mem

1. One struct param

    Struct fields already been set into each(six in this example) registers `$rax ~ $rdi`
    ```go
    func structmod6(v paramstestStruct) paramstestStruct {
                            0x5ad2e0 push %rbp
                            0x5ad2e1 mov %rsp,%rbp
                            0x5ad2e4 sub $0x28,%rsp
                            0x5ad2e8 mov %rax,0x38(%rsp)
                            0x5ad2ed movss %xmm0,0x40(%rsp)
                            0x5ad2f3 mov %rbx,0x48(%rsp)
                            0x5ad2f8 mov %rcx,0x50(%rsp)
                            0x5ad2fd mov %dil,0x58(%rsp)
                            0x5ad302 mov %sil,0x59(%rsp)
                            0x5ad307 movups %xmm15,(%rsp)
     [arg copy reg -> mem]  0x5ad30c movups %xmm15,0x8(%rsp)
                            0x5ad312 movups %xmm15,0x18(%rsp)
        v.d = false         0x5ad318 movb $0x0,0x58(%rsp)
        v.e.a = false       0x5ad31d movb $0x0,0x59(%rsp)
        v.c = "no no no"    0x5ad322 lea 0x45db5(%rip),%rdx # 0x5f30de
                            0x5ad329 mov %rdx,0x48(%rsp)
                            0x5ad32e movq $0x8,0x50(%rsp)
        v.a = 8             0x5ad337 movq $0x8,0x38(%rsp)
        v.b = 8.8           0x5ad340 movss 0x7880c(%rip),%xmm1 # 0x625b54 <$f32.410ccccd>
                            0x5ad348 movss %xmm1,0x40(%rsp)
        return v            0x5ad34e mov 0x38(%rsp),%rdx
                            0x5ad353 mov %rdx,(%rsp)
                            0x5ad357 movups 0x40(%rsp),%xmm1
                            0x5ad35c movups %xmm1,0x8(%rsp)
                            0x5ad361 movups 0x50(%rsp),%xmm1
                            0x5ad366 movups %xmm1,0x18(%rsp)
                            0x5ad36b mov (%rsp),%rax
                            0x5ad36f movss 0x8(%rsp),%xmm0
                            0x5ad375 mov 0x10(%rsp),%rbx
                            0x5ad37a mov 0x18(%rsp),%rcx
                            0x5ad37f movzbl 0x20(%rsp),%edi
    [result copy mem -> reg]0x5ad384 movzbl 0x21(%rsp),%esi
                            0x5ad389 add $0x28,%rsp
                            0x5ad38d pop %rbp
                            0x5ad38e ret
    }
    ```