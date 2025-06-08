# backend ディレクトリ構成

```
backend/
  cmd/
    server/         # メインエントリポイント（main.goなど）
  internal/
    domain/         # ドメイン層（エンティティ、値オブジェクト、ドメインサービス）
      user/
        user.go
    usecase/        # ユースケース層（アプリケーションサービス）
      user/
        user_usecase.go
    infra/          # インフラ層（DBアクセス、外部APIクライアントなど）
      repository/
        user_repository.go
    web/            # プレゼンテーション層（HTTPハンドラなど）
      handler/
        user_handler.go
      middleware/
        auth.go
  go.mod
  README.md
```

この構成はドメイン駆動設計（DDD）やクリーンアーキテクチャを意識しつつ、小規模プロジェクト向けにシンプルにまとめているよ。 