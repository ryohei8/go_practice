# JWT 認証の流れを理解する

https://qiita.com/asagohan2301/items/cef8bcb969fef9064a5c

# 具体例

トークンを作るときの流れ

トークンのヘッダーとペイロードを次のように作成します：

```
Header: {
  "alg": "HS256",
  "typ": "JWT"
}
Payload: {
  "user_id": "12345",
  "exp": 1712345678
}
```

サーバーはこれを次のように署名します：

```
Base64URL(Header) + "." + Base64URL(Payload) + " を鍵（jwtKey）で署名"
```

結果として次のようなトークンが作成されます：

```
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMTIzNDUiLCJleHAiOjE3MTIzNDU2Nzh9.7FJgXOKl73jJ9TzFQH0RgjxLDlRqs6YI_c2bQvEzSi8
```

トークンを受け取るときの流れ

サーバーは以下を確認します：

ヘッダーとペイロード部分を取り出し、秘密鍵（jwtKey）を使って署名を再計算。
クライアントが送ったトークンの署名と一致するかを比較。
署名が一致する場合: 「このトークンはサーバーで作られたもので、改ざんされていない」と判断。

署名が一致しない場合: トークンが改ざんされた可能性があるため、「不正なトークン」として拒否します。
