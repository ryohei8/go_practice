# ポインタとは？

ポインタは、「変数がどこに格納されているか（メモリ上の住所）」を指す特別な変数です。

普通の変数: 値そのものを持つ
ポインタ: 値がどこにあるかを指す

# ポインタの記法

Go でポインタを使うときの主な記法は次の 2 つ：

### 「 & 」アドレス演算子

- 変数の「住所」を取得します。
- 例: &x は変数 x の住所。

### 「 \* 」デリファレンス演算子

- ポインタの指している住所から「中身」を取得します。
- 例: \*p はポインタ p が指している値。

# Example

```go
package main

import "fmt"

func main() {
    x := 10         // 普通の変数
    p := &x         // xのアドレスをポインタpに保存

    fmt.Println(x)  // 出力: 10
    fmt.Println(p)  // 出力: xのメモリアドレス (例: 0xc000014080)
    fmt.Println(*p) // 出力: 10 (ポインタpが指している値)

    *p = 20         // pが指す値を変更 (xの中身が変わる)
    fmt.Println(x)  // 出力: 20
}
```

# よくある使い方

構造体のフィールドを変更
構造体を関数で変更したい場合、ポインタが必要です。

```go
type User struct {
    Name string
    Age  int
}

func updateUser(u *User) {
    u.Name = "Alice"
    u.Age = 25
}

func main() {
    user := User{Name: "Bob", Age: 20}
    updateUser(&user)
    fmt.Println(user) // 出力: {Alice 25}
}
```

&user を渡すことで、updateUser 関数内で元の user を変更できます。
