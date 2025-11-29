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

---

## create

DevCycle組織に新しいプロジェクトを作成します。

### 使用方法

```bash
dvcx projects create [flags]
```

### フラグ

| フラグ | 短縮形 | 説明 | 必須 |
|------|-------|------|------|
| `--name` | `-n` | プロジェクト名 | はい |
| `--key` | `-k` | プロジェクトキー | はい |
| `--description` | `-d` | プロジェクトの説明 | いいえ |
| `--output` | `-o` | 出力形式 (table, json, yaml) | いいえ |

### 例

```bash
# 新しいプロジェクトを作成
$ dvcx projects create -n "My New App" -k my-new-app
KEY:         my-new-app
NAME:        My New App
CREATED:     2024-06-20T10:00:00Z

# 説明付きのプロジェクトを作成
$ dvcx projects create -n "Production App" -k production-app -d "Main production application"

# プロジェクトを作成してJSON形式で出力
$ dvcx projects create -n "Staging App" -k staging-app -o json
{
  "key": "staging-app",
  "name": "Staging App",
  "createdAt": "2024-06-20T10:00:00Z"
}
```

### 注意事項

- プロジェクトキーは組織内で一意である必要があります
- プロジェクトキーには小文字、数字、ハイフンを含めることができます
- 作成後、プロジェクトキーは変更できません

---

## update

既存のプロジェクトを更新します。

### 使用方法

```bash
dvcx projects update <project-key> [flags]
```

### 引数

| 引数 | 説明 |
|------|------|
| `project-key` | 更新するプロジェクトの一意のキー |

### フラグ

| フラグ | 短縮形 | 説明 | 必須 |
|------|-------|------|------|
| `--name` | `-n` | 新しいプロジェクト名 | いいえ |
| `--description` | `-d` | 新しいプロジェクトの説明 | いいえ |
| `--output` | `-o` | 出力形式 (table, json, yaml) | いいえ |

### 例

```bash
# プロジェクト名を更新
$ dvcx projects update my-app -n "My Updated Application"

# プロジェクトの説明を更新
$ dvcx projects update my-app -d "Updated description for my application"

# 名前と説明の両方を更新
$ dvcx projects update my-app -n "New Name" -d "New description"
```

### 注意事項

- 指定されたフィールドのみが更新されます
- プロジェクトキーは作成後に変更できません
