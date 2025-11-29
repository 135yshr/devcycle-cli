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

---

## create

プロジェクトに新しいフィーチャーを作成します。

### 使用方法

```bash
# シンプルな作成（v1 API）
dvcx features create --name <name> --key <key> [flags]

# JSONファイルから作成（v2 API）
dvcx features create --from-file <file.json> [flags]

# 標準入力から作成（v2 API）
dvcx features create --from-file - [flags]
```

### フラグ

| フラグ | 短縮形 | 説明 | 必須 |
|------|-------|------|------|
| `--project` | `-p` | プロジェクトキー | はい（または設定で指定） |
| `--name` | `-n` | フィーチャー名 | はい（シンプル作成の場合） |
| `--key` | `-k` | フィーチャーキー | はい（シンプル作成の場合） |
| `--description` | `-d` | フィーチャーの説明 | いいえ |
| `--type` | `-t` | フィーチャータイプ (release, experiment, permission, ops) | いいえ（デフォルト: release） |
| `--from-file` | `-F` | フィーチャー作成用のJSON入力ファイル（v2 APIを使用）、標準入力の場合は`-`を指定 | いいえ |
| `--dry-run` | | 作成せずに設定を検証 | いいえ |
| `--output` | `-o` | 出力形式 (table, json, yaml) | いいえ |

### シンプル作成の例

```bash
# リリースフィーチャーを作成
$ dvcx features create -p my-app -n "Dark Mode" -k dark-mode
KEY:         dark-mode
NAME:        Dark Mode
TYPE:        release
STATUS:      active
CREATED:     2024-06-20T10:00:00Z

# 説明付きの実験フィーチャーを作成
$ dvcx features create -p my-app -n "New Checkout" -k new-checkout -t experiment -d "A/B test for new checkout flow"

# フィーチャーを作成してJSON形式で出力
$ dvcx features create -p my-app -n "Beta Feature" -k beta-feature -o json
```

### JSONファイルから作成（v2 API）

`--from-file`フラグを使用すると、v2 APIを使用して変数、バリエーション、ターゲティングルールを含む完全なフィーチャー設定が可能になります。

```bash
# JSONファイルから作成
$ dvcx features create -p my-app --from-file feature.json

# 作成せずに検証（dry-run）
$ dvcx features create -p my-app --from-file feature.json --dry-run
```

#### JSONファイル形式

```json
{
  "name": "Dark Mode",
  "key": "dark-mode",
  "description": "ダークモードテーマを有効化",
  "type": "release",
  "tags": ["ui", "theme"],
  "variables": [
    {
      "key": "enabled",
      "name": "Enabled",
      "type": "Boolean"
    }
  ],
  "variations": [
    {
      "key": "off",
      "name": "Off",
      "variables": { "enabled": false }
    },
    {
      "key": "on",
      "name": "On",
      "variables": { "enabled": true }
    }
  ],
  "controlVariation": "off",
  "configurations": {
    "development": {
      "status": "active",
      "targets": [
        {
          "name": "All Users",
          "audience": {
            "filters": {
              "operator": "and",
              "filters": [{ "type": "all" }]
            }
          },
          "distribution": [
            { "_variation": "on", "percentage": 1.0 }
          ]
        }
      ]
    }
  }
}
```

### 標準入力から作成

ファイルパスとして`-`を指定することで、JSONコンテンツを直接コマンドにパイプできます。

```bash
# ファイルからパイプ
$ cat feature.json | dvcx features create -p my-app --from-file -

# ヒアドキュメントを使用
$ dvcx features create -p my-app --from-file - <<EOF
{
  "name": "Quick Feature",
  "key": "quick-feature",
  "type": "release"
}
EOF

# 他のコマンドからパイプ（例: jq）
$ jq '.features[0]' features.json | dvcx features create -p my-app --from-file -
```

### v2 APIでサポートされるフィールド

| フィールド | 型 | 説明 |
|-----------|------|------|
| `name` | string | フィーチャー名（必須） |
| `key` | string | フィーチャーキー（必須） |
| `description` | string | フィーチャーの説明 |
| `type` | string | フィーチャータイプ: release, experiment, permission, ops |
| `tags` | string[] | フィーチャーを整理するためのタグ |
| `variables` | object[] | 変数定義 |
| `variations` | object[] | 変数値を含むバリエーション定義 |
| `controlVariation` | string | コントロールバリエーションのキー |
| `configurations` | object | 環境固有のターゲティング設定 |
| `sdkVisibility` | object | SDK可視性設定（mobile, client, server） |
| `settings` | object | フィーチャー設定（publicName, publicDescription, optInEnabled） |

### 注意事項

- フィーチャーキーはプロジェクト内で一意である必要があります
- フィーチャーキーには小文字、数字、ハイフン、アンダースコアを含めることができます
- 指定しない場合、デフォルトのフィーチャータイプは`release`です
- 標準入力からJSONを読み込むには`--from-file -`を使用します
- フィーチャーを作成せずに設定を検証するには`--dry-run`を使用します

---

## update

既存のフィーチャーを更新します。

### 使用方法

```bash
dvcx features update <feature-key> [flags]
```

### 引数

| 引数 | 説明 |
|------|------|
| `feature-key` | 更新するフィーチャーの一意のキー |

### フラグ

| フラグ | 短縮形 | 説明 | 必須 |
|------|-------|------|------|
| `--project` | `-p` | プロジェクトキー | はい（または設定で指定） |
| `--name` | `-n` | 新しいフィーチャー名 | いいえ |
| `--description` | `-d` | 新しいフィーチャーの説明 | いいえ |
| `--output` | `-o` | 出力形式 (table, json, yaml) | いいえ |

### 例

```bash
# フィーチャー名を更新
$ dvcx features update dark-mode -p my-app -n "Dark Theme"

# フィーチャーの説明を更新
$ dvcx features update dark-mode -p my-app -d "Enable dark theme for the application"

# 名前と説明の両方を更新
$ dvcx features update dark-mode -p my-app -n "Dark Theme" -d "Enable dark theme"
```

### 注意事項

- 指定されたフィールドのみが更新されます
- フィーチャーキーとタイプは作成後に変更できません

---

## delete

プロジェクトからフィーチャーを削除します。

### 使用方法

```bash
dvcx features delete <feature-key> [flags]
```

### 引数

| 引数 | 説明 |
|------|------|
| `feature-key` | 削除するフィーチャーの一意のキー |

### フラグ

| フラグ | 短縮形 | 説明 | 必須 |
|------|-------|------|------|
| `--project` | `-p` | プロジェクトキー | はい（または設定で指定） |
| `--force` | `-f` | 確認プロンプトをスキップ | いいえ |

### 例

```bash
# フィーチャーを削除（確認プロンプトあり）
$ dvcx features delete dark-mode -p my-app
Are you sure you want to delete feature 'dark-mode'? [y/N]: y
Feature 'dark-mode' deleted successfully

# 確認なしでフィーチャーを削除
$ dvcx features delete dark-mode -p my-app --force
Feature 'dark-mode' deleted successfully
```

### 注意事項

- フィーチャーを削除すると、関連するすべての変数とターゲティングルールも削除されます
- この操作は元に戻せません
- 自動化スクリプトでは`--force`フラグを使用して確認をスキップできます
