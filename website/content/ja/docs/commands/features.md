---
title: "features"
weight: 3
---

# features

DevCycleフィーチャー（フィーチャーフラグ）を管理するためのコマンド。

## list

プロジェクト内のすべてのフィーチャーを一覧表示します。

### 使用方法

```bash
dvcx features list [flags]
```

### フラグ

| フラグ | 短縮形 | 説明 | 必須 |
|------|-------|------|------|
| `--project` | `-p` | プロジェクトキー | はい（または設定で指定） |
| `--output` | `-o` | 出力形式 (table, json, yaml) | いいえ |

### 例

```bash
# 特定のプロジェクトのフィーチャーを一覧表示
$ dvcx features list -p my-app
KEY                 NAME                TYPE      STATUS
dark-mode           Dark Mode           release   active
new-checkout        New Checkout Flow   release   active
beta-features       Beta Features       permission inactive

# JSON形式でフィーチャーを一覧表示
$ dvcx features list -p my-app -o json
[
  {
    "key": "dark-mode",
    "name": "Dark Mode",
    "type": "release",
    "status": "active"
  },
  ...
]

# 設定のデフォルトプロジェクトを使用
$ dvcx features list
```

### 注意事項

- `--project`が指定されていない場合、設定のデフォルトプロジェクトが使用されます
- フィーチャータイプ: `release`, `experiment`, `permission`, `ops`
- フィーチャーステータス: `active`, `inactive`, `archived`

---

## get

特定のフィーチャーの詳細を取得します。

### 使用方法

```bash
dvcx features get <feature-key> [flags]
```

### 引数

| 引数 | 説明 |
|------|------|
| `feature-key` | フィーチャーの一意のキー |

### フラグ

| フラグ | 短縮形 | 説明 | 必須 |
|------|-------|------|------|
| `--project` | `-p` | プロジェクトキー | はい（または設定で指定） |
| `--output` | `-o` | 出力形式 (table, json, yaml) | いいえ |

### 例

```bash
# フィーチャーの詳細を取得
$ dvcx features get dark-mode -p my-app
KEY:         dark-mode
NAME:        Dark Mode
TYPE:        release
STATUS:      active
DESCRIPTION: Enable dark mode theme for the application
CREATED:     2024-01-20T10:00:00Z
UPDATED:     2024-06-15T14:30:00Z

# JSON形式でフィーチャーの詳細を取得
$ dvcx features get dark-mode -p my-app -o json
{
  "key": "dark-mode",
  "name": "Dark Mode",
  "type": "release",
  "status": "active",
  "description": "Enable dark mode theme for the application",
  "createdAt": "2024-01-20T10:00:00Z",
  "updatedAt": "2024-06-15T14:30:00Z",
  "variables": [
    {
      "key": "dark-mode-enabled",
      "type": "Boolean"
    }
  ]
}
```

### フィーチャータイプ

| タイプ | 説明 |
|------|------|
| `release` | 機能リリース用の標準フィーチャーフラグ |
| `experiment` | A/Bテストと実験 |
| `permission` | ユーザー権限ベースの機能 |
| `ops` | システム設定用の運用フラグ |

### 注意事項

- レスポンスには関連する変数とバリエーションが含まれます
- デプロイ前にフィーチャー設定を確認するためにこのコマンドを使用してください
