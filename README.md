### Goを学ぶ

Go言語を学ぶ為のリポジトリです。
日本のことわざに「郷に入ってはGoに従え」というのがあるように、ここではしっかりGoの文法やお作法を学びます。

### Goはどんな言語なのか?

- 2009年1月にGoogleで開発された。
- 協力でシンプルな言語設計と文法
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
  - `go vet`"コードチェッカー(静的解析ツール)
  - Goは割と静的解析がしやすい。
- シングルバイナリ・クロスコンパイル
  - 開発の環境が違くても動作する
- [peco](https://github.com/peco/peco)もGoが使われている
  - 標準入力されたデータをインクリメンタルサーチをして選択したら出力する
- サクッとCLIツールを作ルるのに適している
-

### SetUp


### 教材

- [ソフトウェアデザイン 2019年5月号](https://booklog.jp/item/1/B07PZN49RW)
- [The Go Programming Language](https://golang.org/#)
- [A Tour of Go](https://go-tour-jp.appspot.com/welcome/1)

