---
title: "audit"
weight: 20
---

# audit

監査ログを表示するコマンドです。監査ログはプロジェクト内の全ての変更を追跡します。

## list

プロジェクトの全監査ログを一覧表示します。

### 使用方法

```bash
dvcx audit list [flags]
```

### フラグ

| フラグ | 短縮形 | 説明 | 必須 |
|-------|--------|------|------|
| `--project` | `-p` | プロジェクトキー | はい |
| `--output` | `-o` | 出力形式（table, json, yaml） | いいえ |

### 例

```bash
# プロジェクトの監査ログを一覧表示
$ dvcx audit list -p my-app

# JSON形式で出力
$ dvcx audit list -p my-app -o json
```

### 監査ログタイプ

| タイプ | 説明 |
|--------|------|
| `feature.created` | フィーチャーが作成された |
| `feature.updated` | フィーチャーが更新された |
| `feature.deleted` | フィーチャーが削除された |
| `variable.created` | 変数が作成された |
| `variable.updated` | 変数が更新された |
| `variable.deleted` | 変数が削除された |
| `targeting.updated` | ターゲティングルールが更新された |

---

## feature

特定のフィーチャーの監査ログを一覧表示します。

### 使用方法

```bash
dvcx audit feature <feature-key> [flags]
```

### 引数

| 引数 | 説明 |
|------|------|
| `feature-key` | フィーチャーの一意キー |

### フラグ

| フラグ | 短縮形 | 説明 | 必須 |
|-------|--------|------|------|
| `--project` | `-p` | プロジェクトキー | はい |
| `--output` | `-o` | 出力形式 | いいえ |

### 例

```bash
# 特定のフィーチャーの監査ログを表示
$ dvcx audit feature dark-mode -p my-app
```

### 注意事項

- 監査ログは読み取り専用で変更できません
- ログには変更内容、変更者、変更日時の詳細が含まれます
