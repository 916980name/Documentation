```
x86 word = 2 bytes
x86 dword = 4 bytes (double word)
x86 qword = 8 bytes (quad word)
x86 double-quad or xmmword = 16 bytes, e.g. movdqa xmm0, [rdi].
Also in the cqo mnemonic, oct-word. (Sign-extend RAX into RDX:RAX, e.g. before idiv)
```

### TUI
https://sourceware.org/gdb/current/onlinedocs/gdb.html/TUI-Keys.html#TUI-Keys
- `C-x a`, Enter or leave the TUI mode
- `C-x 1`, Use a TUI layout with only one window.
- `C-x 2`, Use a TUI layout with at least two windows
- `C-x o`, Change the active window

#### print register
- p $register
- info r register

#### in memory data
- string: `10\n`
- byte: `31 30 0a 00 00 00 00 00`