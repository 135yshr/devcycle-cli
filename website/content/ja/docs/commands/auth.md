---
title: "auth"
weight: 1
---

# auth

DevCycle API用の認証コマンド。

## login

OAuth2認証情報を使用してDevCycleで認証します。

### 使用方法

```bash
dvcx auth login
```

### 説明

このコマンドは、DevCycle API認証情報の入力を求めます：

- **Client ID**: DevCycle APIクライアントID
- **Client Secret**: DevCycle APIクライアントシークレット

認証成功後、アクセストークンはカレントディレクトリの`.devcycle/token.json`に保存されます。

### 例

```bash
$ dvcx auth login
Enter Client ID: your-client-id
Enter Client Secret: ********
Successfully authenticated!
```

### 注意事項

- 認証情報はDevCycleダッシュボードの**Settings** → **API Credentials**から取得します
- トークンはローカルに保存され、以降のAPI呼び出しに使用されます
- トークンの有効期限は自動的に処理されます。定期的に再認証が必要な場合があります

---

## logout

保存された認証情報を削除します。

### 使用方法

```bash
dvcx auth logout
```

### 説明

このコマンドは`.devcycle/token.json`から保存されたアクセストークンを削除します。

### 例

```bash
$ dvcx auth logout
Successfully logged out!
```

### 注意事項

- ログアウト後、他のコマンドを使用する前に`dvcx auth login`を再度実行する必要があります
- これはローカルトークンのみを削除します。サーバー上でトークンを無効化するわけではありません
