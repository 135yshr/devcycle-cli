---
title: "webhooks"
weight: 22
---

# webhooks

DevCycle Webhook を管理するコマンドです。Webhook を使用すると、プロジェクト内でイベントが発生したときに通知を受け取ることができます。

## list

プロジェクト内の全 Webhook を一覧表示します。

### 使用方法

```bash
dvcx webhooks list [flags]
```

### フラグ

| フラグ | 短縮形 | 説明 | 必須 |
|-------|--------|------|------|
| `--project` | `-p` | プロジェクトキー | はい |
| `--output` | `-o` | 出力形式 | いいえ |

### 例

```bash
$ dvcx webhooks list -p my-app
```

---

## get

特定の Webhook の詳細を取得します。

### 使用方法

```bash
dvcx webhooks get <webhook-id> [flags]
```

### 例

```bash
$ dvcx webhooks get wh-123 -p my-app -o json
```

---

## create

新しい Webhook を作成します。

### 使用方法

```bash
dvcx webhooks create [flags]
```

### フラグ

| フラグ | 短縮形 | 説明 | 必須 |
|-------|--------|------|------|
| `--project` | `-p` | プロジェクトキー | はい |
| `--url` | | Webhook URL | はい |
| `--description` | | 説明 | いいえ |
| `--enabled` | | Webhook を有効化（デフォルト: true） | いいえ |

### 例

```bash
# 有効な Webhook を作成
$ dvcx webhooks create -p my-app \
  --url "https://example.com/webhook" \
  --description "本番通知" \
  --enabled

# 無効な Webhook を作成（テスト用）
$ dvcx webhooks create -p my-app \
  --url "https://staging.example.com/webhook" \
  --description "テスト Webhook"
```

---

## update

既存の Webhook を更新します。

### 使用方法

```bash
dvcx webhooks update <webhook-id> [flags]
```

### フラグ

| フラグ | 短縮形 | 説明 | 必須 |
|-------|--------|------|------|
| `--project` | `-p` | プロジェクトキー | はい |
| `--url` | | 新しい URL | いいえ |
| `--description` | | 新しい説明 | いいえ |
| `--enabled` | | Webhook を有効化 | いいえ |
| `--disabled` | | Webhook を無効化 | いいえ |

### 例

```bash
# Webhook を有効化
$ dvcx webhooks update wh-123 -p my-app --enabled

# Webhook を無効化
$ dvcx webhooks update wh-123 -p my-app --disabled
```

---

## delete

Webhook を削除します。

### 使用方法

```bash
dvcx webhooks delete <webhook-id> [flags]
```

### フラグ

| フラグ | 短縮形 | 説明 | 必須 |
|-------|--------|------|------|
| `--project` | `-p` | プロジェクトキー | はい |
| `--force` | | 確認プロンプトをスキップ | いいえ |

### 例

```bash
$ dvcx webhooks delete wh-123 -p my-app --force
```
