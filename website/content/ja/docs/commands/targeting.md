---
title: "targeting"
weight: 6
---

# targeting

DevCycle フィーチャーターゲティング設定を管理するコマンド。

ターゲティングを使用すると、ルール、パーセンテージ、ユーザー属性に基づいて、特定のユーザーにフィーチャーフラグの特定のバリエーションを表示するように制御できます。

## get

フィーチャーのターゲティング設定を取得します。

### 使用方法

```bash
dvcx targeting get [flags]
```

### フラグ

| フラグ | 短縮形 | 説明 | 必須 |
|------|-------|------|------|
| `--project` | `-p` | プロジェクトキー | はい（または設定ファイルで指定） |
| `--feature` | `-f` | フィーチャーキー | はい |
| `--output` | `-o` | 出力形式 (table, json, yaml) | いいえ |

### 例

```bash
# フィーチャーのターゲティング設定を取得
$ dvcx targeting get -p my-app -f dark-mode
ENVIRONMENT   STATUS    TARGETS
development   active    2 rule(s)
staging       active    1 rule(s)
production    inactive  0 rule(s)

# JSON 形式でターゲティング設定を取得
$ dvcx targeting get -p my-app -f dark-mode -o json
{
  "development": {
    "status": "active",
    "targets": [...]
  },
  "staging": {
    "status": "active",
    "targets": [...]
  },
  "production": {
    "status": "inactive",
    "targets": []
  }
}
```

---

## update

JSON ファイルを使用してフィーチャーのターゲティング設定を更新します。

### 使用方法

```bash
dvcx targeting update [flags]
```

### フラグ

| フラグ | 短縮形 | 説明 | 必須 |
|------|-------|------|------|
| `--project` | `-p` | プロジェクトキー | はい（または設定ファイルで指定） |
| `--feature` | `-f` | フィーチャーキー | はい |
| `--from-file` | `-F` | 設定用 JSON 入力ファイル、標準入力には '-' を使用 | はい |
| `--output` | `-o` | 出力形式 (table, json, yaml) | いいえ |

### 例

```bash
# JSON ファイルからターゲティングを更新
$ dvcx targeting update -p my-app -f dark-mode -F targeting-config.json

# 標準入力からターゲティングを更新
$ cat targeting-config.json | dvcx targeting update -p my-app -f dark-mode -F -

# targeting-config.json の構造例:
{
  "development": {
    "status": "active",
    "targets": [
      {
        "name": "All Users",
        "distribution": [
          {"_variation": "variation-on", "percentage": 1.0}
        ],
        "audience": {
          "name": "All Users",
          "filters": {
            "filters": [{"type": "all"}],
            "operator": "and"
          }
        }
      }
    ]
  }
}
```

### 注意事項

- JSON ファイルには環境キーから設定オブジェクトへのマップを含める必要があります
- 最大ファイルサイズは 10MB です
- 他のコマンドからの設定をパイプするには標準入力（`-F -`）を使用します

---

## enable

特定の環境でフィーチャーを有効にします。

### 使用方法

```bash
dvcx targeting enable [flags]
```

### フラグ

| フラグ | 短縮形 | 説明 | 必須 |
|------|-------|------|------|
| `--project` | `-p` | プロジェクトキー | はい（または設定ファイルで指定） |
| `--feature` | `-f` | フィーチャーキー | はい |
| `--environment` | `-e` | 環境キー | はい |

### 例

```bash
# 開発環境でフィーチャーを有効化
$ dvcx targeting enable -p my-app -f dark-mode -e development
Feature 'dark-mode' enabled for environment 'development'

# 本番環境でフィーチャーを有効化
$ dvcx targeting enable -p my-app -f new-checkout -e production
Feature 'new-checkout' enabled for environment 'production'
```

### 注意事項

- フィーチャーを有効にすると、指定した環境でターゲティングルールがアクティブになります
- ユーザーは設定されたターゲティングルールに基づいてバリエーションを受け取り始めます

---

## disable

特定の環境でフィーチャーを無効にします。

### 使用方法

```bash
dvcx targeting disable [flags]
```

### フラグ

| フラグ | 短縮形 | 説明 | 必須 |
|------|-------|------|------|
| `--project` | `-p` | プロジェクトキー | はい（または設定ファイルで指定） |
| `--feature` | `-f` | フィーチャーキー | はい |
| `--environment` | `-e` | 環境キー | はい |

### 例

```bash
# 本番環境でフィーチャーを無効化
$ dvcx targeting disable -p my-app -f dark-mode -e production
Feature 'dark-mode' disabled for environment 'production'

# ステージング環境でフィーチャーを無効化
$ dvcx targeting disable -p my-app -f experimental-feature -e staging
Feature 'experimental-feature' disabled for environment 'staging'
```

### 注意事項

- フィーチャーを無効にすると、指定した環境でユーザーへの配信が停止します
- フィーチャーが無効の場合、ユーザーはデフォルト/オフのバリエーションを受け取ります
- ターゲティングルールは保持され、フィーチャーを再度有効にすると再び適用されます
