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
