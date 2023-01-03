###  A Tour of Go

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

その後に必要となるパッケージを指定します。

###　型について

Goの場合は変数名の後ろに型名を書くようにしています。

```go
func add(x int, y int) int {
	return x + y
}

```


### 参考

- https://go-tour-jp.appspot.com/welcome/1
- https://www.tohoho-web.com/ex/golang.html
