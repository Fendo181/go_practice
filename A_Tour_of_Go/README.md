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

関数の2つ以上の引数が同じ型である場合には最後の型を残して省略して記述ができます。

```go
func add(x, y int) int {
	return x + y
}
```

また関数は複数の戻り値の型を指定する事ができます。

```go
func swap(x, y string) (string, string) {
	return x, y
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
```

### 参考

- https://go-tour-jp.appspot.com/welcome/1
- https://www.tohoho-web.com/ex/golang.html
