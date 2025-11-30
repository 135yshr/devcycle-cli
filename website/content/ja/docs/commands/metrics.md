---
title: "metrics"
weight: 21
---

# metrics

DevCycle メトリクスを管理するコマンドです。メトリクスを使用すると、フィーチャーフラグの効果を測定できます。

## list

プロジェクト内の全メトリクスを一覧表示します。

### 使用方法

```bash
dvcx metrics list [flags]
```

### フラグ

| フラグ | 短縮形 | 説明 | 必須 |
|-------|--------|------|------|
| `--project` | `-p` | プロジェクトキー | はい |
| `--output` | `-o` | 出力形式 | いいえ |

### 例

```bash
$ dvcx metrics list -p my-app
```

---

## get

特定のメトリクスの詳細を取得します。

### 使用方法

```bash
dvcx metrics get <metric-key> [flags]
```

### 例

```bash
$ dvcx metrics get conversion-rate -p my-app -o json
```

---

## create

新しいメトリクスを作成します。

### 使用方法

```bash
dvcx metrics create [flags]
```

### フラグ

| フラグ | 短縮形 | 説明 | 必須 |
|-------|--------|------|------|
| `--project` | `-p` | プロジェクトキー | はい |
| `--key` | `-k` | メトリクスキー | はい |
| `--name` | `-n` | メトリクス名 | はい |
| `--type` | `-t` | タイプ（count, conversion, sum, average） | はい |
| `--event-type` | | イベントタイプ | はい |
| `--optimize-for` | | 最適化目標（increase, decrease） | はい |
| `--description` | `-d` | 説明 | いいえ |

### 例

```bash
# カウントメトリクスを作成
$ dvcx metrics create -p my-app \
  --key page-views \
  --name "ページビュー" \
  --type count \
  --event-type pageview \
  --optimize-for increase

# コンバージョンメトリクスを作成
$ dvcx metrics create -p my-app \
  --key signup-rate \
  --name "サインアップ率" \
  --type conversion \
  --event-type signup \
  --optimize-for increase
```

### メトリクスタイプ

| タイプ | 説明 |
|--------|------|
| `count` | イベント数をカウント |
| `conversion` | コンバージョン率を測定 |
| `sum` | イベント値の合計 |
| `average` | イベント値の平均 |

---

## update

既存のメトリクスを更新します。

### 使用方法

```bash
dvcx metrics update <metric-key> [flags]
```

### 例

```bash
$ dvcx metrics update page-views -p my-app --name "総ページビュー"
```

---

## delete

メトリクスを削除します。

### 使用方法

```bash
dvcx metrics delete <metric-key> [flags]
```

### 例

```bash
$ dvcx metrics delete page-views -p my-app --force
```

---

## results

メトリクスの結果を取得します。

### 使用方法

```bash
dvcx metrics results <metric-key> [flags]
```

### フラグ

| フラグ | 短縮形 | 説明 | 必須 |
|-------|--------|------|------|
| `--project` | `-p` | プロジェクトキー | はい |
| `--environment` | `-e` | 環境でフィルタ | いいえ |
| `--feature` | `-f` | フィーチャーでフィルタ | いいえ |
| `--start-date` | | 開始日（YYYY-MM-DD） | いいえ |
| `--end-date` | | 終了日（YYYY-MM-DD） | いいえ |

### 例

```bash
# メトリクス結果を取得
$ dvcx metrics results conversion-rate -p my-app

# フィルタを組み合わせ
$ dvcx metrics results conversion-rate -p my-app \
  --environment production \
  --feature new-checkout \
  --start-date 2024-01-01 \
  --end-date 2024-01-31
```
