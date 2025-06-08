``` mermaid
erDiagram
    USER {
        int id PK "ユーザーID"
        string email "メールアドレス"
        string password_hash "パスワードハッシュ"
        string display_name "表示名"
        datetime created_at "作成日時"
        datetime updated_at "更新日時"
    }
    PASSKEY {
        int id PK "パスキーID"
        int user_id FK "ユーザーID"
        string credential_id "クレデンシャルID"
        string public_key "公開鍵"
        string device_name "デバイス名"
        datetime created_at "作成日時"
        datetime last_used_at "最終使用日時"
    }
    AUTH_HISTORY {
        int id PK "履歴ID"
        int user_id FK "ユーザーID"
        string method "認証方式"
        datetime authenticated_at "認証日時"
        string ip_address "IPアドレス"
        string user_agent "ユーザーエージェント"
    }
    SESSION {
        string id PK "セッションID"
        int user_id FK "ユーザーID"
        datetime created_at "作成日時"
        datetime expires_at "有効期限"
        string refresh_token "リフレッシュトークン"
        string ip_address "IPアドレス"
        string user_agent "ユーザーエージェント"
    }

    USER ||--o{ PASSKEY : "1人のユーザーは複数のパスキーを持つ"
    USER ||--o{ AUTH_HISTORY : "1人のユーザーは複数の認証履歴を持つ"
    USER ||--o{ SESSION : "1人のユーザーは複数のセッションを持つ"
```