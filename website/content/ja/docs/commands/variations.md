---
title: "variations"
weight: 7
---

# variations

DevCycle フィーチャーバリエーションを管理するコマンド。

バリエーションは、フィーチャーがユーザーに提供できる異なる値を定義します。各バリエーションには、一緒に配信される変数値のセットが含まれます。

## list

フィーチャーの全バリエーションを一覧表示します。

### 使用方法

```bash
dvcx variations list [flags]
```

### フラグ

| フラグ | 短縮形 | 説明 | 必須 |
|------|-------|------|------|
| `--project` | `-p` | プロジェクトキー | はい（または設定ファイルで指定） |
| `--feature` | `-f` | フィーチャーキー | はい |
| `--output` | `-o` | 出力形式 (table, json, yaml) | いいえ |

### 例

```bash
# フィーチャーのバリエーションを一覧表示
$ dvcx variations list -p my-app -f dark-mode
KEY           NAME           VARIABLES
variation-on  Dark Mode On   {"enabled": true}
variation-off Dark Mode Off  {"enabled": false}

# JSON 形式でバリエーションを一覧表示
$ dvcx variations list -p my-app -f dark-mode -o json
[
  {
    "key": "variation-on",
    "name": "Dark Mode On",
    "variables": {"enabled": true}
  },
  {
    "key": "variation-off",
    "name": "Dark Mode Off",
    "variables": {"enabled": false}
  }
]
```

---

## get

特定のバリエーションの詳細を取得します。

### 使用方法

```bash
dvcx variations get <variation-key> [flags]
```

### 引数

| 引数 | 説明 |
|------|------|
| `variation-key` | バリエーションの一意キー |

### フラグ

| フラグ | 短縮形 | 説明 | 必須 |
|------|-------|------|------|
| `--project` | `-p` | プロジェクトキー | はい（または設定ファイルで指定） |
| `--feature` | `-f` | フィーチャーキー | はい |
| `--output` | `-o` | 出力形式 (table, json, yaml) | いいえ |

### 例

```bash
# バリエーションの詳細を取得
$ dvcx variations get variation-on -p my-app -f dark-mode
KEY:       variation-on
NAME:      Dark Mode On
VARIABLES: {"enabled": true}

# JSON 形式でバリエーションの詳細を取得
$ dvcx variations get variation-on -p my-app -f dark-mode -o json
{
  "key": "variation-on",
  "name": "Dark Mode On",
  "variables": {
    "enabled": true
  }
}
```

---

## create

フィーチャーに新しいバリエーションを作成します。

### 使用方法

```bash
dvcx variations create [flags]
```

### フラグ

| フラグ | 短縮形 | 説明 | 必須 |
|------|-------|------|------|
| `--project` | `-p` | プロジェクトキー | はい（または設定ファイルで指定） |
| `--feature` | `-f` | フィーチャーキー | はい |
| `--name` | `-n` | バリエーション名 | はい |
| `--key` | `-k` | バリエーションキー | はい |
| `--variables` | `-v` | JSON 形式の変数値 | いいえ |
| `--output` | `-o` | 出力形式 (table, json, yaml) | いいえ |

### 例

```bash
# シンプルなバリエーションを作成
$ dvcx variations create -p my-app -f dark-mode \
  --key variation-dim \
  --name "Dim Mode"
KEY:       variation-dim
NAME:      Dim Mode
VARIABLES: -

# 変数付きのバリエーションを作成
$ dvcx variations create -p my-app -f dark-mode \
  --key variation-custom \
  --name "Custom Theme" \
  --variables '{"enabled": true, "brightness": 80}'
KEY:       variation-custom
NAME:      Custom Theme
VARIABLES: {"enabled": true, "brightness": 80}

# バリエーションを作成して JSON で出力
$ dvcx variations create -p my-app -f new-checkout \
  --key v2 \
  --name "Checkout V2" \
  -v '{"version": 2}' \
  -o json
{
  "key": "v2",
  "name": "Checkout V2",
  "variables": {"version": 2}
}
```

### 注意事項

- バリエーションキーはフィーチャー内で一意である必要があります
- バリエーションキーには小文字、数字、ハイフンを使用できます
- 変数はフィーチャーの変数スキーマで定義された型と一致する必要があります

---

## update

既存のバリエーションを更新します。

### 使用方法

```bash
dvcx variations update <variation-key> [flags]
```

### 引数

| 引数 | 説明 |
|------|------|
| `variation-key` | 更新するバリエーションの一意キー |

### フラグ

| フラグ | 短縮形 | 説明 | 必須 |
|------|-------|------|------|
| `--project` | `-p` | プロジェクトキー | はい（または設定ファイルで指定） |
| `--feature` | `-f` | フィーチャーキー | はい |
| `--name` | `-n` | 新しいバリエーション名 | いいえ |
| `--variables` | `-v` | 新しい変数値（JSON 形式） | いいえ |
| `--output` | `-o` | 出力形式 (table, json, yaml) | いいえ |

### 例

```bash
# バリエーション名を更新
$ dvcx variations update variation-on -p my-app -f dark-mode \
  --name "Dark Theme Enabled"

# バリエーションの変数を更新
$ dvcx variations update variation-on -p my-app -f dark-mode \
  --variables '{"enabled": true, "theme": "midnight"}'

# 名前と変数の両方を更新
$ dvcx variations update variation-custom -p my-app -f dark-mode \
  --name "Custom Dark Theme" \
  --variables '{"enabled": true, "brightness": 70}'
```

### 注意事項

- 指定されたフィールドのみが更新されます
- バリエーションキーは作成後に変更できません

---

## delete

フィーチャーからバリエーションを削除します。

### 使用方法

```bash
dvcx variations delete <variation-key> [flags]
```

### 引数

| 引数 | 説明 |
|------|------|
| `variation-key` | 削除するバリエーションの一意キー |

### フラグ

| フラグ | 短縮形 | 説明 | 必須 |
|------|-------|------|------|
| `--project` | `-p` | プロジェクトキー | はい（または設定ファイルで指定） |
| `--feature` | `-f` | フィーチャーキー | はい |
| `--force` | | 確認プロンプトをスキップ | いいえ |

### 例

```bash
# バリエーションを削除（確認プロンプトあり）
$ dvcx variations delete variation-old -p my-app -f dark-mode
Are you sure you want to delete variation 'variation-old'? [y/N]: y
Variation 'variation-old' deleted successfully

# 確認なしでバリエーションを削除
$ dvcx variations delete variation-test -p my-app -f dark-mode --force
Variation 'variation-test' deleted successfully
```

### 警告

- この操作は取り消せません
- ターゲティングルールでアクティブに使用されているバリエーションは削除できません
- 自動化スクリプトでは `--force` フラグを使用して確認をスキップできます
