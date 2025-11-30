---
title: "keys"
weight: 6
---

# keys

DevCycle環境のSDKキーを管理するためのコマンド。

SDKキーはDevCycleとアプリケーションの認証に使用されます。各環境には、異なるプラットフォーム（クライアント、サーバー、モバイル）用のSDKキーセットがあります。

## list

環境のすべてのSDKキーを一覧表示します。

### 使用方法

```bash
dvcx keys list [flags]
```

### フラグ

| フラグ | 短縮形 | 説明 | 必須 |
|------|-------|------|------|
| `--project` | `-p` | プロジェクトキー | はい（または設定で指定） |
| `--environment` | `-e` | 環境キー | はい |
| `--output` | `-o` | 出力形式 (table, json, yaml) | いいえ |

### 例

```bash
# development環境のSDKキーを一覧表示
$ dvcx keys list -p my-app -e development
ENVIRONMENT   TYPE     KEY
development   client   dvc_client_abc123...
development   server   dvc_server_def456...
development   mobile   dvc_mobile_ghi789...

# JSON形式でSDKキーを一覧表示
$ dvcx keys list -p my-app -e development -o json
{
  "environment": "development",
  "keys": {
    "client": "dvc_client_abc123...",
    "server": "dvc_server_def456...",
    "mobile": "dvc_mobile_ghi789..."
  }
}
```

### キータイプ

| タイプ | 説明 | 用途 |
|------|------|------|
| `client` | クライアント側SDKキー | Webブラウザ、フロントエンドアプリケーション |
| `server` | サーバー側SDKキー | バックエンドサービス、API |
| `mobile` | モバイルSDKキー | iOS、Androidアプリケーション |

### 注意事項

- SDKキーは環境ごとに一意です
- サーバー側キーをクライアントアプリケーションに公開しないでください
- クライアントキーとモバイルキーはクライアント側コードに含めても安全です

---

## rotate

環境のSDKキーをローテーションします。

### 使用方法

```bash
dvcx keys rotate [flags]
```

### フラグ

| フラグ | 短縮形 | 説明 | 必須 |
|------|-------|------|------|
| `--project` | `-p` | プロジェクトキー | はい（または設定で指定） |
| `--environment` | `-e` | 環境キー | はい |
| `--type` | `-t` | ローテーションするキータイプ (client, server, mobile) | はい |
| `--force` | `-f` | 確認プロンプトをスキップ | いいえ |

### 例

```bash
# 確認付きでクライアントSDKキーをローテーション
$ dvcx keys rotate -p my-app -e development --type client
環境 'development' のclient SDKキーをローテーションしますか？
既存のキーは無効になります。 [y/N]: y
Previous Key: dvc_client_abc123...
New Key:      dvc_client_xyz987...

# 確認なしでサーバーSDKキーをローテーション
$ dvcx keys rotate -p my-app -e development --type server --force
Previous Key: dvc_server_def456...
New Key:      dvc_server_uvw654...

# 本番環境のモバイルSDKキーをローテーション
$ dvcx keys rotate -p my-app -e production --type mobile --force
Previous Key: dvc_mobile_ghi789...
New Key:      dvc_mobile_rst321...
```

### 警告

**キーのローテーションは元に戻せません！**

- 古いキーは即座に無効化されます
- 古いキーを使用しているすべてのアプリケーションはDevCycleへのアクセスを失います
- 古いキーの有効期限が切れる前に、アプリケーションを新しいキーで更新してください
- トラフィックが少ない時間帯にキーローテーションを実行することを検討してください

### ベストプラクティス

1. **ローテーションの計画**: メンテナンスウィンドウ中にキーローテーションを計画する
2. **アプリケーションを先に更新**: ローテーション前に新しいキーをデプロイ準備する
3. **ローテーション後の監視**: ローテーション後に認証エラーを監視する
4. **定期的なローテーション**: 定期的なキーローテーションはセキュリティを向上させます
