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

# Flow

+-------------------+
| main.go | (アプリのエントリーポイント)
| - 各層を初期化して依存関係を注入
| - サーバーを起動しリクエストを受け付ける
+-------------------+
↓
+-------------------+
| Handler | (エンドポイント)
| - リクエストを受け取り解析
| - Usecase を呼び出し、レスポンスを返す
+-------------------+
↓
+-------------------+
| Usecase | (ユースケース)
| - ビジネスロジックの流れを管理
| - Service, Repository を利用
+-------------------+
↓
+-------------------+
| Service | (ビジネスロジック)
| - ドメインルールを実現
| - 複雑な処理をカプセル化
| - Repository を利用
+-------------------+
↓
+-------------------+
| Repository | (抽象インターフェース)
| - データ操作の契約を定義
| - Domain に依存
+-------------------+
↓
+-------------------+
| Infrastructure | (具体的な実装)
| - Repository を実装
| - データベースや外部 API に接続
+-------------------+
↓
+-------------------+
| Domain | (モデルとビジネスルール)
| - アプリケーションの核
| - モデルとドメインロジックを定義
| - 他の層に依存しない
+-------------------+

# 依存関係

main.go (依存関係を構築)
↓
Handler (Usecase に依存)
↓
Usecase (Service, Repository に依存)
↓
Service (Domain, Repository に依存)
↓
Repository (Domain に依存)
↓
Infrastructure (Repository に依存)
↓
Domain (依存なし)

## 依存関係の説明

1. Handler → Usecase
   依存関係の意味
   Handler は Usecase に依存しています。
   Usecase がなければ、Handler はビジネスロジックを実行できないため、動作しません。
   具体例
   クライアントからリクエストが来た場合：

Handler はリクエストを受け取るだけで、具体的なビジネスロジック（例えば、データの保存）は Usecase に任せます。
Usecase が提供する機能を使って、Handler はレスポンスを生成します。

2. Usecase → Service/Repository
   依存関係の意味
   Usecase は Service や Repository に依存しています。
   ビジネスロジックを処理するために Service を呼び出し、データ操作のために Repository を利用します。
   Service や Repository がないと、Usecase は何もできません。

3. Service → Repository
   依存関係の意味
   Service は Repository に依存しています。
   Repository がなければ、データベースや外部リソースとの通信ができません。
