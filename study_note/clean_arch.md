# Clean Architecture

# Directory Structure

project/
│
├── domain/ # エンティティやドメインロジック（アプリケーションの中心部分）
│ ├── model/ # ドメインモデル（ビジネスデータの定義）
│ ├── repository/ # インターフェース（データ操作の契約）
│ └── service/ # ドメインルールやロジック
│
├── usecase/ # ユースケース（アプリケーションの具体的な操作）
│ ├── transition_usecase.go # 各ユースケースごとのロジック
│ └── ...
│
├── handler/ # ハンドラー（リクエストとレスポンスの処理）
│ ├── transition_handler.go # HTTP エンドポイントの処理
│ └── ...
│
├── infrastructure/ # 外部リソースとの接続（データベース、API など）
│ ├── database/ # データベースの実装
│ ├── external/ # 外部 API のクライアント
│ └── ...
│
├── middleware/ # 認証やロギングなどの共通機能
├── config/ # 設定ファイルや環境変数の管理
└── main.go # アプリケーションのエントリーポイント
