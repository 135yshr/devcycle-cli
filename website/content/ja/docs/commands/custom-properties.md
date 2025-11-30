---
title: "custom-properties"
weight: 23
---

# custom-properties

カスタムプロパティを管理するコマンドです。カスタムプロパティは、ターゲティングに使用できる追加のユーザー属性を定義します。

**エイリアス:** `cp`

## list

プロジェクト内の全カスタムプロパティを一覧表示します。

### 使用方法

```bash
dvcx custom-properties list [flags]
# または
dvcx cp list [flags]
```

### フラグ

| フラグ | 短縮形 | 説明 | 必須 |
|-------|--------|------|------|
| `--project` | `-p` | プロジェクトキー | はい |
| `--output` | `-o` | 出力形式 | いいえ |

### 例

```bash
# 全カスタムプロパティを一覧表示
$ dvcx cp list -p my-app
```

---

## get

特定のカスタムプロパティの詳細を取得します。

### 使用方法

```bash
dvcx custom-properties get <property-key> [flags]
# または
dvcx cp get <property-key> [flags]
```

### 例

```bash
$ dvcx cp get user-type -p my-app -o json
```

---

## create

新しいカスタムプロパティを作成します。

### 使用方法

```bash
dvcx custom-properties create [flags]
# または
dvcx cp create [flags]
```

### フラグ

| フラグ | 短縮形 | 説明 | 必須 |
|-------|--------|------|------|
| `--project` | `-p` | プロジェクトキー | はい |
| `--key` | `-k` | プロパティキー（SDKのプロパティ名と一致必須） | はい |
| `--display-name` | | 表示名 | はい |
| `--type` | `-t` | タイプ（Boolean, Number, String） | はい |
| `--description` | | 説明 | いいえ |

### 例

```bash
# String プロパティを作成
$ dvcx cp create -p my-app \
  --key user-type \
  --display-name "ユーザータイプ" \
  --type String \
  --description "ユーザーアカウントの種類"

# Boolean プロパティを作成
$ dvcx cp create -p my-app \
  --key is-premium \
  --display-name "プレミアム会員" \
  --type Boolean \
  --description "プレミアム会員かどうか"

# Number プロパティを作成
$ dvcx cp create -p my-app \
  --key account-age \
  --display-name "アカウント年齢" \
  --type Number \
  --description "アカウント作成からの日数"
```

### プロパティタイプ

| タイプ | 説明 | 例 |
|--------|------|-----|
| `Boolean` | 真偽値 | `true`, `false` |
| `Number` | 数値 | `42`, `3.14`, `-10` |
| `String` | 文字列 | `"premium"`, `"basic"` |

---

## update

既存のカスタムプロパティを更新します。

### 使用方法

```bash
dvcx custom-properties update <property-key> [flags]
# または
dvcx cp update <property-key> [flags]
```

### フラグ

| フラグ | 短縮形 | 説明 | 必須 |
|-------|--------|------|------|
| `--project` | `-p` | プロジェクトキー | はい |
| `--display-name` | | 新しい表示名 | いいえ |
| `--description` | | 新しい説明 | いいえ |

### 例

```bash
# 表示名を更新
$ dvcx cp update user-type -p my-app \
  --display-name "ユーザーアカウントタイプ"
```

### 注意事項

- プロパティキーとタイプは作成後に変更できません
- 表示名と説明のみ更新可能です

---

## delete

カスタムプロパティを削除します。

### 使用方法

```bash
dvcx custom-properties delete <property-key> [flags]
# または
dvcx cp delete <property-key> [flags]
```

### フラグ

| フラグ | 短縮形 | 説明 | 必須 |
|-------|--------|------|------|
| `--project` | `-p` | プロジェクトキー | はい |
| `--force` | | 確認プロンプトをスキップ | いいえ |

### 例

```bash
$ dvcx cp delete user-type -p my-app --force
```

### 注意事項

- 削除するとこのプロパティを使用するターゲティングルールに影響する可能性があります
- この操作は元に戻せません

## SDK での使用

カスタムプロパティはアプリケーションコードから渡す必要があります：

```javascript
// JavaScript SDK の例
const user = {
  user_id: "user-123",
  customData: {
    "user-type": "premium",
    "is-premium": true,
    "account-age": 365
  }
};
```

SDK のプロパティキーが DevCycle で定義したキーと一致していることを確認してください。
