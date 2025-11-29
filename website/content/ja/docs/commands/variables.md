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

---

## create

プロジェクトに新しい変数を作成します。

### 使用方法

```bash
dvcx variables create [flags]
```

### フラグ

| フラグ | 短縮形 | 説明 | 必須 |
|------|-------|------|------|
| `--project` | `-p` | プロジェクトキー | はい（または設定で指定） |
| `--name` | `-n` | 変数名 | はい |
| `--key` | `-k` | 変数キー | はい |
| `--description` | `-d` | 変数の説明 | いいえ |
| `--type` | `-t` | 変数タイプ (String, Boolean, Number, JSON) | いいえ（デフォルト: Boolean） |
| `--feature` | | 関連付けるフィーチャーキー | いいえ |
| `--output` | `-o` | 出力形式 (table, json, yaml) | いいえ |

### 例

```bash
# Boolean変数を作成
$ dvcx variables create -p my-app -n "Dark Mode Enabled" -k dark-mode-enabled
KEY:          dark-mode-enabled
TYPE:         Boolean
STATUS:       active
CREATED:      2024-06-20T10:00:00Z

# 説明付きのString変数を作成
$ dvcx variables create -p my-app -n "Theme Color" -k theme-color -t String -d "Primary theme color"

# フィーチャーに関連付けた変数を作成
$ dvcx variables create -p my-app -n "Checkout Version" -k checkout-version -t Number --feature new-checkout

# 変数を作成してJSON形式で出力
$ dvcx variables create -p my-app -n "Config" -k config-json -t JSON -o json
```

### 注意事項

- 変数キーはプロジェクト内で一意である必要があります
- 変数キーには小文字、数字、ハイフン、アンダースコアを含めることができます
- 指定しない場合、デフォルトの変数タイプは`Boolean`です
- `--feature`を使用して既存のフィーチャーに変数を関連付けることができます

---

## update

既存の変数を更新します。

### 使用方法

```bash
dvcx variables update <variable-key> [flags]
```

### 引数

| 引数 | 説明 |
|------|------|
| `variable-key` | 更新する変数の一意のキー |

### フラグ

| フラグ | 短縮形 | 説明 | 必須 |
|------|-------|------|------|
| `--project` | `-p` | プロジェクトキー | はい（または設定で指定） |
| `--name` | `-n` | 新しい変数名 | いいえ |
| `--description` | `-d` | 新しい変数の説明 | いいえ |
| `--output` | `-o` | 出力形式 (table, json, yaml) | いいえ |

### 例

```bash
# 変数名を更新
$ dvcx variables update dark-mode-enabled -p my-app -n "Dark Theme Enabled"

# 変数の説明を更新
$ dvcx variables update dark-mode-enabled -p my-app -d "Controls whether dark theme is enabled"

# 名前と説明の両方を更新
$ dvcx variables update dark-mode-enabled -p my-app -n "Dark Theme" -d "Dark theme toggle"
```

### 注意事項

- 指定されたフィールドのみが更新されます
- 変数キーとタイプは作成後に変更できません

---

## delete

プロジェクトから変数を削除します。

### 使用方法

```bash
dvcx variables delete <variable-key> [flags]
```

### 引数

| 引数 | 説明 |
|------|------|
| `variable-key` | 削除する変数の一意のキー |

### フラグ

| フラグ | 短縮形 | 説明 | 必須 |
|------|-------|------|------|
| `--project` | `-p` | プロジェクトキー | はい（または設定で指定） |
| `--force` | `-f` | 確認プロンプトをスキップ | いいえ |

### 例

```bash
# 変数を削除（確認プロンプトあり）
$ dvcx variables delete dark-mode-enabled -p my-app
Are you sure you want to delete variable 'dark-mode-enabled'? [y/N]: y
Variable 'dark-mode-enabled' deleted successfully

# 確認なしで変数を削除
$ dvcx variables delete dark-mode-enabled -p my-app --force
Variable 'dark-mode-enabled' deleted successfully
```

### 注意事項

- この操作は元に戻せません
- 自動化スクリプトでは`--force`フラグを使用して確認をスキップできます
- 変数がフィーチャーに関連付けられている場合、先に関連付けを解除する必要がある場合があります
