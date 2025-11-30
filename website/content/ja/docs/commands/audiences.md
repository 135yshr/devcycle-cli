---
title: "audiences"
weight: 10
---

# audiences

DevCycle オーディエンスを管理するコマンドです。オーディエンスを使用すると、ターゲティング用の再利用可能なユーザーセグメントを定義できます。

## list

プロジェクト内の全オーディエンスを一覧表示します。

### 使用方法

```bash
dvcx audiences list [flags]
```

### フラグ

| フラグ | 短縮形 | 説明 | 必須 |
|-------|--------|------|------|
| `--project` | `-p` | プロジェクトキー | はい（または設定ファイルで指定） |
| `--output` | `-o` | 出力形式（table, json, yaml） | いいえ |

### 例

```bash
# 全オーディエンスを一覧表示
$ dvcx audiences list -p my-app

# JSON形式で出力
$ dvcx audiences list -p my-app -o json
```

---

## get

特定のオーディエンスの詳細を取得します。

### 使用方法

```bash
dvcx audiences get <audience-key> [flags]
```

### 引数

| 引数 | 説明 |
|------|------|
| `audience-key` | オーディエンスの一意キー |

### フラグ

| フラグ | 短縮形 | 説明 | 必須 |
|-------|--------|------|------|
| `--project` | `-p` | プロジェクトキー | はい（または設定ファイルで指定） |
| `--output` | `-o` | 出力形式（table, json, yaml） | いいえ |

### 例

```bash
# オーディエンスの詳細を取得
$ dvcx audiences get beta-users -p my-app -o json
```

---

## create

新しいオーディエンスを作成します。

### 使用方法

```bash
dvcx audiences create [flags]
```

### フラグ

| フラグ | 短縮形 | 説明 | 必須 |
|-------|--------|------|------|
| `--project` | `-p` | プロジェクトキー | はい |
| `--key` | `-k` | オーディエンスキー | はい |
| `--name` | `-n` | オーディエンス名 | はい |
| `--description` | `-d` | 説明 | いいえ |
| `--filters` | | フィルターJSON | はい |
| `--output` | `-o` | 出力形式 | いいえ |

### 例

```bash
# ベータユーザー用オーディエンスを作成
$ dvcx audiences create -p my-app \
  --key beta-users \
  --name "ベータユーザー" \
  --description "ベータプログラムのユーザー" \
  --filters '[{"type":"user","subType":"email","comparator":"contain","values":["@beta.example.com"]}]'
```

---

## update

既存のオーディエンスを更新します。

### 使用方法

```bash
dvcx audiences update <audience-key> [flags]
```

### 例

```bash
# オーディエンス名を更新
$ dvcx audiences update beta-users -p my-app --name "ベータテスター"
```

---

## delete

オーディエンスを削除します。

### 使用方法

```bash
dvcx audiences delete <audience-key> [flags]
```

### フラグ

| フラグ | 短縮形 | 説明 | 必須 |
|-------|--------|------|------|
| `--project` | `-p` | プロジェクトキー | はい |
| `--force` | | 確認プロンプトをスキップ | いいえ |

### 例

```bash
# 確認なしで削除
$ dvcx audiences delete beta-users -p my-app --force
```
