## Packages, variables, and functions.

Go プログラムの基本的なコンポーネントを学びます。

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

その後に必要となる処理をmain関数内で記述します。

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

### Named return values

Goでは戻り値となる変数に名前をつける事ができます。
名前をつけた戻り値の変数を使うと、 `return` ステートメントに何も書かずに戻すことができます。これを `naked" return` と呼びます。ただし、`naked return`ステートメントは、短い関数でのみ利用すべきです。

```go
func split(sum int) (x, y int) {
	x = sum % 17
	y = sum * 2
	return
}

func main() {
	fmt.Println(split(17))
}
```

### Variables

Go言語における`var`ステートメントは変数を宣言します。関数の引数のリストと同様に。複数の変数の最後に型を書く事で、変数のリストを宣言します。この場合、`c`, `python` , `java` はboolean で `i` だけint型になる。

```go
package main

import "fmt"

var c, python , java bool

func main() {
	var i int
	fmt.Println(i, c, python, java)
}
```

### Short variable declarations

関数の中では、`var`宣言の代わりに、短い `:=` の代入文を使い、暗黙的な型宣言ができます。

```go
package main

import "fmt"

func main() {
	var i, j int = 1, 2
	k := 3  // var k = 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
```

### Type conversions

Goで型変換をする場合は明示的に型を指定して変換する。

```go
func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}
```

またGoは明示的な型を指定せずに変換した場合には右側の新しい変数は同じ型になります。

```go
var i int
j := i // j is an int
```

### Constants

Goで定数を宣言する場合には `const` キーワードを設定する。
定数は、文字(character)、文字列(string)、boolean、数値(numeric)のみで使えます。また定数は省略形である `:=` は使えない。

```go
package main
import "fmt"

const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")
}
```

### Numeric Constants

数値の定数は、高精度な 値 ( values )です。
型のない定数は、その状況によって必要な型を取ることになります。

```go
package main

import "fmt"

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small)) // 21
 	fmt.Println(needFloat(Small)) // 0.2
	fmt.Println(needFloat(Big)) // 1.267506
}
```

### 参考

- https://go-tour-jp.appspot.com/list
- https://go-tour-jp.appspot.com/welcome/1
- https://www.tohoho-web.com/ex/golang.html
