---
title: "environments"
weight: 5
---

# environments

DevCycle環境を管理するためのコマンド。

環境は、フィーチャーが異なる設定を持つことができるデプロイメントステージ（開発、ステージング、本番）を表します。

**エイリアス**: `envs`, `env`

## list

プロジェクト内のすべての環境を一覧表示します。

### 使用方法

```bash
dvcx environments list [flags]
# または
dvcx envs list [flags]
dvcx env list [flags]
```

### フラグ

| フラグ | 短縮形 | 説明 | 必須 |
|------|-------|------|------|
| `--project` | `-p` | プロジェクトキー | はい（または設定で指定） |
| `--output` | `-o` | 出力形式 (table, json, yaml) | いいえ |

### 例

```bash
# プロジェクト内の環境を一覧表示
$ dvcx environments list -p my-app
KEY           NAME          TYPE
development   Development   development
staging       Staging       staging
production    Production    production

# JSON形式で環境を一覧表示
$ dvcx envs list -p my-app -o json
[
  {
    "key": "development",
    "name": "Development",
    "type": "development"
  },
  {
    "key": "staging",
    "name": "Staging",
    "type": "staging"
  },
  {
    "key": "production",
    "name": "Production",
    "type": "production"
  }
]
```

### 環境タイプ

| タイプ | 説明 |
|------|------|
| `development` | テスト用の開発環境 |
| `staging` | 本番前テスト用のステージング環境 |
| `production` | ライブユーザー用の本番環境 |

---

## get

特定の環境の詳細を取得します。

### 使用方法

```bash
dvcx environments get <environment-key> [flags]
# または
dvcx envs get <environment-key> [flags]
dvcx env get <environment-key> [flags]
```

### 引数

| 引数 | 説明 |
|------|------|
| `environment-key` | 環境の一意のキー |

### フラグ

| フラグ | 短縮形 | 説明 | 必須 |
|------|-------|------|------|
| `--project` | `-p` | プロジェクトキー | はい（または設定で指定） |
| `--output` | `-o` | 出力形式 (table, json, yaml) | いいえ |

### 例

```bash
# 環境の詳細を取得
$ dvcx environments get production -p my-app
KEY:         production
NAME:        Production
TYPE:        production
DESCRIPTION: Live production environment
CREATED:     2024-01-15T10:00:00Z

# JSON形式で環境の詳細を取得
$ dvcx env get production -p my-app -o json
{
  "key": "production",
  "name": "Production",
  "type": "production",
  "description": "Live production environment",
  "createdAt": "2024-01-15T10:00:00Z",
  "updatedAt": "2024-01-15T10:00:00Z"
}
```

### 注意事項

- 各プロジェクトは複数の環境を持つことができます
- 環境設定は独立しています - フィーチャーは開発では有効で本番では無効にできます
- SDKキーは環境ごとに一意です

---

## create

プロジェクトに新しい環境を作成します。

### 使用方法

```bash
dvcx environments create [flags]
# または
dvcx envs create [flags]
dvcx env create [flags]
```

### フラグ

| フラグ | 短縮形 | 説明 | 必須 |
|------|-------|------|------|
| `--project` | `-p` | プロジェクトキー | はい（または設定で指定） |
| `--name` | `-n` | 環境名 | はい |
| `--key` | `-k` | 環境キー | はい |
| `--type` | `-t` | 環境タイプ (development, staging, production) | いいえ（デフォルト: development） |
| `--color` | | 環境の色（16進数形式） | いいえ |
| `--description` | `-d` | 環境の説明 | いいえ |

### 例

```bash
# ステージング環境を作成
$ dvcx environments create -p my-app \
  --key staging \
  --name "Staging" \
  --type staging \
  --color "#ffff00" \
  --description "本番前テスト環境"
KEY:         staging
NAME:        Staging
TYPE:        staging
DESCRIPTION: 本番前テスト環境
CREATED:     2024-01-15T10:00:00Z

# 最小限の環境を作成
$ dvcx envs create -p my-app --key qa --name "QA"
KEY:         qa
NAME:        QA
TYPE:        development
CREATED:     2024-01-15T10:00:00Z
```

---

## update

既存の環境を更新します。

### 使用方法

```bash
dvcx environments update <environment-key> [flags]
# または
dvcx envs update <environment-key> [flags]
dvcx env update <environment-key> [flags]
```

### 引数

| 引数 | 説明 |
|------|------|
| `environment-key` | 更新する環境の一意のキー |

### フラグ

| フラグ | 短縮形 | 説明 | 必須 |
|------|-------|------|------|
| `--project` | `-p` | プロジェクトキー | はい（または設定で指定） |
| `--name` | `-n` | 新しい環境名 | いいえ |
| `--color` | | 新しい環境の色（16進数形式） | いいえ |
| `--description` | `-d` | 新しい環境の説明 | いいえ |

### 例

```bash
# 環境名と説明を更新
$ dvcx environments update staging -p my-app \
  --name "ステージング環境" \
  --description "更新されたステージング環境"
KEY:         staging
NAME:        ステージング環境
TYPE:        staging
DESCRIPTION: 更新されたステージング環境
UPDATED:     2024-01-16T10:00:00Z

# 色のみを更新
$ dvcx envs update staging -p my-app --color "#00ffff"
KEY:         staging
NAME:        ステージング環境
TYPE:        staging
COLOR:       #00ffff
UPDATED:     2024-01-16T10:00:00Z
```

---

## delete

プロジェクトから環境を削除します。

### 使用方法

```bash
dvcx environments delete <environment-key> [flags]
# または
dvcx envs delete <environment-key> [flags]
dvcx env delete <environment-key> [flags]
```

### 引数

| 引数 | 説明 |
|------|------|
| `environment-key` | 削除する環境の一意のキー |

### フラグ

| フラグ | 短縮形 | 説明 | 必須 |
|------|-------|------|------|
| `--project` | `-p` | プロジェクトキー | はい（または設定で指定） |
| `--force` | `-f` | 確認プロンプトをスキップ | いいえ |

### 例

```bash
# 確認付きで環境を削除
$ dvcx environments delete staging -p my-app
環境 'staging' を削除しますか？ [y/N]: y
環境が正常に削除されました

# 確認なしで環境を削除
$ dvcx envs delete qa -p my-app --force
環境が正常に削除されました
```

### 警告

環境の削除は元に戻すことができず、その環境のすべてのフィーチャー設定が削除されます。
