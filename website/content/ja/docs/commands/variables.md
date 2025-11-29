---
title: "variables"
weight: 4
---

# variables

DevCycle変数を管理するためのコマンド。

変数は、フィーチャーフラグが制御する値です。各フィーチャーは異なるタイプの複数の変数を持つことができます。

## list

プロジェクト内のすべての変数を一覧表示します。

### 使用方法

```bash
dvcx variables list [flags]
```

### フラグ

| フラグ | 短縮形 | 説明 | 必須 |
|------|-------|------|------|
| `--project` | `-p` | プロジェクトキー | はい（または設定で指定） |
| `--output` | `-o` | 出力形式 (table, json, yaml) | いいえ |

### 例

```bash
# プロジェクト内の変数を一覧表示
$ dvcx variables list -p my-app
KEY                   TYPE      FEATURE
dark-mode-enabled     Boolean   dark-mode
theme-color           String    dark-mode
checkout-version      Number    new-checkout
config-json           JSON      beta-features

# JSON形式で変数を一覧表示
$ dvcx variables list -p my-app -o json
[
  {
    "key": "dark-mode-enabled",
    "type": "Boolean",
    "feature": "dark-mode"
  },
  ...
]
```

### 変数タイプ

| タイプ | 説明 | 値の例 |
|------|------|--------|
| `Boolean` | 真偽値 | `true`, `false` |
| `String` | テキスト値 | `"dark"`, `"light"` |
| `Number` | 数値 | `1`, `2.5`, `100` |
| `JSON` | 複雑なJSONオブジェクト | `{"key": "value"}` |

---

## get

特定の変数の詳細を取得します。

### 使用方法

```bash
dvcx variables get <variable-key> [flags]
```

### 引数

| 引数 | 説明 |
|------|------|
| `variable-key` | 変数の一意のキー |

### フラグ

| フラグ | 短縮形 | 説明 | 必須 |
|------|-------|------|------|
| `--project` | `-p` | プロジェクトキー | はい（または設定で指定） |
| `--output` | `-o` | 出力形式 (table, json, yaml) | いいえ |

### 例

```bash
# 変数の詳細を取得
$ dvcx variables get dark-mode-enabled -p my-app
KEY:          dark-mode-enabled
TYPE:         Boolean
FEATURE:      dark-mode
DESCRIPTION:  Controls whether dark mode is enabled
CREATED:      2024-01-20T10:00:00Z

# JSON形式で変数の詳細を取得
$ dvcx variables get dark-mode-enabled -p my-app -o json
{
  "key": "dark-mode-enabled",
  "type": "Boolean",
  "feature": "dark-mode",
  "description": "Controls whether dark mode is enabled",
  "createdAt": "2024-01-20T10:00:00Z",
  "updatedAt": "2024-06-15T14:30:00Z"
}
```

### 注意事項

- 変数は常にフィーチャーに関連付けられています
- 変数キーはプロジェクト内で一意である必要があります
- 変数タイプは作成後に変更できません
