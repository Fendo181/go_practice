###  A Tour of Go

https://go-tour-jp.appspot.com/welcome/1

### Packagesとは

Goのプログラムはパッケージで管理されています。
プログラムは`main`というパッケージから始まります。
その後に`fmt`、`math/rand`パッケージを指定しています。

```go
package main

import (
    "fmt"
    "math/rand"
)
```

その後に

###
