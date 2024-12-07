# Context について

## Context とは

- コンテキストは、Go の並行処理において、キャンセル、タイムアウト、その他のリクエストスコープの情報を伝達するためのもの
- コンテキストは、並行処理の中で、キャンセルやタイムアウトを管理するために使用される

```go
package main

import (
	"context"
	"fmt"
	"time"
)

func doSomething(ctx context.Context) {
	for {
		select {
		case <-ctx.Done(): // 親からキャンセル通知を受け取る
			fmt.Println("キャンセルされました:", ctx.Err())
			return
		default:
			fmt.Println("処理中...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	// キャンセル可能なコンテキストを作成
	ctx, cancel := context.WithCancel(context.Background())

	// 子ゴルーチンを実行
	go doSomething(ctx)

	// 3秒後にキャンセル
	time.Sleep(3 * time.Second)
	cancel()

	// キャンセル後も少し待機
	time.Sleep(1 * time.Second)
	fmt.Println("終了")
}
```
