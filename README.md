### Practice Go

Go言語を学ぶ為のリポジトリです。

日本のことわざに「郷に入ってはGoに従え」というのがあるように、ここではしっかりGoの文法やお作法を学びます。

### What is Go?

- 2009年1月にGoogleで開発された。
- 強力でシンプルな言語設計と文法
  - 繰り返し処理は`for`しかない
  - 文法が意図的に少なくなっている。
  - Go2.xは[ジェネリクス](https://www.sejuku.net/blog/22699)やエラー処理について検討している。
- 並行処理
- 豊富な標準ライブラリ群
  - `fmt`:書式に関するライブラリ
  - `nrt/http`:HTTPサーバ、HTTPクライアント
  - `crypto`:暗号化
  - `encoding`:JSON、CSV、XMLなどのエンコーディング
  - `html/template`:HTMLテンプレート
  - `os、path/filepath`:OS固有の処理やファイル関連
- 周辺ツールの充実
  - `go build`:ソースコードのビルド
  - `go install`:ライブラリやコマンドラインツールのインストール
  - `go mod`:ライブラリのバージョン管理
  - `go test`:単体テストを行う
  - `go doc`:ドキュメンテーション
  - `go fmt`:コードフォーマットチェック
  - `go vet`:コードチェッカー(静的解析ツール)
  - Goは割と静的解析がしやすい。
- シングルバイナリ・クロスコンパイル
  - 開発の環境が違くても動作する
- [peco](https://github.com/peco/peco)もGoが使われている
  - 標準入力されたデータをインクリメンタルサーチをして選択したら出力する
- サクッとCLIツールを作るのに適している

### SetUp

Macだとbrewコマンドで簡単に入る。

```sh
 brew install go
```

よさそう

確認

```sh
go version
go version go1.19.4 darwin/amd64
```

### Run

Goで書かれたファイルを実行します。

```go
go run helloworld.go
```

### Build

- `o`でbuildしてできたファイルに名前をつける

```sh
go build -o hellowold hello.go
```

### Codeing

[Effective Go — プログラミング言語 Go ドキュメント v0.1 documentation](http://go.shibu.jp/effective_go.html)

`go fmt`でコーディングスタイルを整える事ができる。

```sh
go fmt
```
※VSCodeで書く場合は自動で`go fmt`をしてくれているので、便利！

#### エラーチェック

[kisielk/errcheck: errcheck checks that you checked errors.](https://github.com/kisielk/errcheck)

#### メモリを考えた順序に並び替える

アライメント(メモリに載せるデータの並びを意識する)

[mdempsky/maligned: Tool to detect Go structs that would take less memory if their fields were sorted.](https://github.com/mdempsky/maligned)


```go
type A struct {
    a [4]byte
    b [3]byte
    s string
}

type B struct{
    a [4]byte
    s string
    b [3]byte
}
````

GoはC言語と同じようにアライメント(メモリに載せる順番)を意識する原義で、この例ではAとBが同じメモリサイズと思いきや、OSのワードサイズや書き込みの順番によって扱うメモるサイズに差が生じる。
例えばBに関して、ワードサイズは`8byte`とすると、`string`に必要な数は2ワード(`16byte`)になるので、Aに関しては

`4(a)+16(s)` =20

となる。が、bのアクセス位置を埋める為に更に`4byte`を追加する。`32byte`が次のワードサイズなので

`4(a)+16(s)+4(パディング)+3(b)+8(次のワードサイズを埋める為)` =32

一方でAの場合

`4(a)+3(b)+1(パディング)+8` = 24

となり`8byte`の差が生まれる。これは小さいけど大きめの配列を作る際は無視できない大きさになってしまう。

[Go Meta Linter がサポートするツールまとめ - SideCI TechBlog](http://tech.sideci.com/entry/2017/01/10/110000)

### Document

- [ソフトウェアデザイン 2019年5月号](https://booklog.jp/item/1/B07PZN49RW)
- [The Go Programming Language](https://golang.org/#)
- [A Tour of Go](https://go-tour-jp.appspot.com/welcome/1)

