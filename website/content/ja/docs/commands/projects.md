---
title: "projects"
weight: 2
---

# projects

DevCycleプロジェクトを管理するためのコマンド。

## list

組織内のすべてのプロジェクトを一覧表示します。

### 使用方法

```bash
dvcx projects list [flags]
```

### フラグ

| フラグ | 短縮形 | 説明 |
|------|-------|------|
| `--output` | `-o` | 出力形式 (table, json, yaml) |

### 例

```bash
# テーブル形式でプロジェクトを一覧表示（デフォルト）
$ dvcx projects list
KEY             NAME                    CREATED
my-app          My Application          2024-01-15
staging-app     Staging Application     2024-02-20
production      Production App          2024-03-10

# JSON形式でプロジェクトを一覧表示
$ dvcx projects list -o json
[
  {
    "key": "my-app",
    "name": "My Application",
    "createdAt": "2024-01-15T10:00:00Z"
  },
  ...
]

# YAML形式でプロジェクトを一覧表示
$ dvcx projects list -o yaml
- key: my-app
  name: My Application
  createdAt: 2024-01-15T10:00:00Z
...
```

---

## get

特定のプロジェクトの詳細を取得します。

### 使用方法

```bash
dvcx projects get <project-key> [flags]
```

### 引数

| 引数 | 説明 |
|------|------|
| `project-key` | プロジェクトの一意のキー |

### フラグ

| フラグ | 短縮形 | 説明 |
|------|-------|------|
| `--output` | `-o` | 出力形式 (table, json, yaml) |

### 例

```bash
# テーブル形式でプロジェクトの詳細を取得
$ dvcx projects get my-app
KEY:         my-app
NAME:        My Application
DESCRIPTION: Main application project
CREATED:     2024-01-15T10:00:00Z

# JSON形式でプロジェクトの詳細を取得
$ dvcx projects get my-app -o json
{
  "key": "my-app",
  "name": "My Application",
  "description": "Main application project",
  "createdAt": "2024-01-15T10:00:00Z",
  "updatedAt": "2024-06-01T15:30:00Z"
}
```

### 注意事項

- プロジェクトキーは大文字小文字を区別します
- 利用可能なプロジェクトキーを確認するには`projects list`を使用してください
