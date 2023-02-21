## Flow control statements: for, if, else, switch and defer

条件文とループ、switch、defer を使ってコードの流れをコントロールする方法を学びます。

### For

基本的に、 for ループはセミコロン `;` で3つの部分に分かれています。

- `初期化ステートメント`: 最初のイテレーション(繰り返し)の前に初期化が実行されます
- `条件式`: イテレーション毎に評価されます
- `後処理ステートメント`: イテレーション毎の最後に実行されます

```go
package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum) //45
}
```

また、`初期化ステートメント`と`後処理ステートメント`の記述は省略しても処理の内容は同じになる

```go
package main

import "fmt"

func main() {
	sum := 1
	for sum < 10 {
		sum += sum
	}
	fmt.Println(sum)
}
```

### Forever Go

ループ条件を省略すれば無限ループを実現する事ができます。


```go
package main

func main() {
	for {
	}
}
```

### If

Go言語の if ステートメントは、先ほどの for ループと同様に、`括弧 ( )` は不要で、`中括弧 { }` は必要です。

```go
package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		// 0より下回って場合は平方根にする
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}
```

また、`if` ステートメントは、 `for` のように、条件の前に、評価するための簡単なステートメントを書くことができます。

```go
func pow(x,n,lim float64) float64 {
	// べき乗の値が(xのn乗)がlimよりも小さい場合はべき乗の値を返す
	if v := math.Pow(x,n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3,2,10),
		pow(3,3,20),
	)
}
```

### If and else

`if` ステートメントで宣言された変数は、 `else` ブロック内でも使うことができます。

```go
package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
```

### Switch

`switch`ステートメントは `if-else`ステートメントの
シーケンスを短く記述する事ができます。
他の言語と似ていますが、GoのSwitchは選択されたcaseだけを実行して、**続く全てのcaseは実行されません** 。つまり他の言語で必要な `brake`ステートメントがGoでは自動的に提供されます。

```go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OX X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}
}
```

### 参考

- https://go-tour-jp.appspot.com/flowcontrol/1
