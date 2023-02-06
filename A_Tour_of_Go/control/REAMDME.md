## Flow control statements: for, if, else, switch and defer

条件文とループ、switch、defer を使ってコードの流れをコントロールする方法を学びます。

### For

基本的に、 for ループはセミコロン ; で3つの部分に分かれています:

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

### 参考

- https://go-tour-jp.appspot.com/flowcontrol/1
