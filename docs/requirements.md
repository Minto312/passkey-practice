# 要件定義書（requirements.md）

## 1. プロジェクト概要
パスキー認証の学習用Webアプリ

## 2. 背景・目的
- パスキー認証の仕組みを理解するため
- 実際にパスキー認証を実装・体験するため
- セキュリティのベストプラクティスを学ぶため

## 3. 機能要件

### MUST（必須機能）
- ユーザー登録（パスキーによる新規登録）
- ログイン（パスキー認証）
- ログアウト
- パスキーの管理（登録済みパスキーの一覧表示・削除）
- メールアドレスとパスワードによる認証（登録・ログイン）
- パスキーの追加登録（複数端末対応）

### WANT（あれば嬉しい機能）
- ユーザー一覧表示ページ（テスト・デモ用）
- 認証履歴の表示

## 4. 制約事項
- フロントエンド：Next.js, Tailwind CSS, shadcn/ui を使用する
- バックエンド：Go を使用する
