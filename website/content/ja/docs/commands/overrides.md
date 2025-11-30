---
title: "overrides"
weight: 11
---

# overrides

セルフターゲティングオーバーライドを管理するコマンドです。オーバーライドを使用すると、開発やテスト中に自分用の特定のバリエーションを設定できます。

## list

フィーチャーの全オーバーライドを一覧表示します。

### 使用方法

```bash
dvcx overrides list [flags]
```

### フラグ

| フラグ | 短縮形 | 説明 | 必須 |
|-------|--------|------|------|
| `--project` | `-p` | プロジェクトキー | はい |
| `--feature` | `-f` | フィーチャーキー | はい |
| `--output` | `-o` | 出力形式 | いいえ |

### 例

```bash
$ dvcx overrides list -p my-app -f dark-mode
```

---

## get

特定の環境におけるフィーチャーの自分のオーバーライドを取得します。

### 使用方法

```bash
dvcx overrides get [flags]
```

### フラグ

| フラグ | 短縮形 | 説明 | 必須 |
|-------|--------|------|------|
| `--project` | `-p` | プロジェクトキー | はい |
| `--feature` | `-f` | フィーチャーキー | はい |
| `--environment` | `-e` | 環境キー | はい |

### 例

```bash
$ dvcx overrides get -p my-app -f dark-mode -e development
```

---

## set

自分用のオーバーライドを設定します。

### 使用方法

```bash
dvcx overrides set [flags]
```

### フラグ

| フラグ | 短縮形 | 説明 | 必須 |
|-------|--------|------|------|
| `--project` | `-p` | プロジェクトキー | はい |
| `--feature` | `-f` | フィーチャーキー | はい |
| `--environment` | `-e` | 環境キー | はい |
| `--variation` | `-v` | バリエーションキー | はい |

### 例

```bash
# フィーチャーを有効にするオーバーライドを設定
$ dvcx overrides set -p my-app -f dark-mode -e development --variation on
```

---

## delete

自分のオーバーライドを削除します。

### 使用方法

```bash
dvcx overrides delete [flags]
```

### 例

```bash
$ dvcx overrides delete -p my-app -f dark-mode -e development
```

---

## list-mine

プロジェクト内の自分の全オーバーライドを一覧表示します。

### 使用方法

```bash
dvcx overrides list-mine [flags]
```

### 例

```bash
$ dvcx overrides list-mine -p my-app
```

---

## delete-mine

プロジェクト内の自分の全オーバーライドを削除します。

### 使用方法

```bash
dvcx overrides delete-mine [flags]
```

### フラグ

| フラグ | 短縮形 | 説明 | 必須 |
|-------|--------|------|------|
| `--project` | `-p` | プロジェクトキー | はい |
| `--force` | | 確認プロンプトをスキップ | いいえ |

### 例

```bash
$ dvcx overrides delete-mine -p my-app --force
```
