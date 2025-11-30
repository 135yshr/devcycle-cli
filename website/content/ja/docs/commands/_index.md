---
title: "コマンド"
weight: 2
bookCollapseSection: true
---

# コマンド

dvcxの全コマンドリファレンス。

## グローバルフラグ

すべてのコマンドで使用可能なフラグ：

| フラグ | 短縮形 | 説明 | デフォルト |
|------|-------|------|---------|
| `--output` | `-o` | 出力形式 (table, json, yaml) | table |
| `--config` | | 設定ファイルのパス | .devcycle/config.yaml |
| `--help` | `-h` | ヘルプを表示 | |

## コマンドカテゴリ

### 認証

| コマンド | 説明 |
|---------|------|
| [auth login]({{< relref "/docs/commands/auth#login" >}}) | DevCycleで認証 |
| [auth logout]({{< relref "/docs/commands/auth#logout" >}}) | 保存された認証情報を削除 |

### プロジェクト

| コマンド | 説明 |
|---------|------|
| [projects list]({{< relref "/docs/commands/projects#list" >}}) | 全プロジェクトを一覧表示 |
| [projects get]({{< relref "/docs/commands/projects#get" >}}) | プロジェクトの詳細を取得 |

### フィーチャー

| コマンド | 説明 |
|---------|------|
| [features list]({{< relref "/docs/commands/features#list" >}}) | プロジェクト内のフィーチャーを一覧表示 |
| [features get]({{< relref "/docs/commands/features#get" >}}) | フィーチャーの詳細を取得 |

### 変数

| コマンド | 説明 |
|---------|------|
| [variables list]({{< relref "/docs/commands/variables#list" >}}) | プロジェクト内の変数を一覧表示 |
| [variables get]({{< relref "/docs/commands/variables#get" >}}) | 変数の詳細を取得 |

### 環境

| コマンド | 説明 |
|---------|------|
| [environments list]({{< relref "/docs/commands/environments#list" >}}) | プロジェクト内の環境を一覧表示 |
| [environments get]({{< relref "/docs/commands/environments#get" >}}) | 環境の詳細を取得 |

### オーディエンス

| コマンド | 説明 |
|---------|------|
| [audiences list]({{< relref "/docs/commands/audiences#list" >}}) | プロジェクト内のオーディエンスを一覧表示 |
| [audiences get]({{< relref "/docs/commands/audiences#get" >}}) | オーディエンスの詳細を取得 |
| [audiences create]({{< relref "/docs/commands/audiences#create" >}}) | 新しいオーディエンスを作成 |
| [audiences update]({{< relref "/docs/commands/audiences#update" >}}) | 既存のオーディエンスを更新 |
| [audiences delete]({{< relref "/docs/commands/audiences#delete" >}}) | オーディエンスを削除 |

### オーバーライド

| コマンド | 説明 |
|---------|------|
| [overrides list]({{< relref "/docs/commands/overrides#list" >}}) | フィーチャーの全オーバーライドを一覧表示 |
| [overrides get]({{< relref "/docs/commands/overrides#get" >}}) | 自分のオーバーライドを取得 |
| [overrides set]({{< relref "/docs/commands/overrides#set" >}}) | 自分用のオーバーライドを設定 |
| [overrides delete]({{< relref "/docs/commands/overrides#delete" >}}) | 自分のオーバーライドを削除 |
| [overrides list-mine]({{< relref "/docs/commands/overrides#list-mine" >}}) | 自分の全オーバーライドを一覧表示 |
| [overrides delete-mine]({{< relref "/docs/commands/overrides#delete-mine" >}}) | 自分の全オーバーライドを削除 |

### 監査ログ

| コマンド | 説明 |
|---------|------|
| [audit list]({{< relref "/docs/commands/audit#list" >}}) | プロジェクトの監査ログを一覧表示 |
| [audit feature]({{< relref "/docs/commands/audit#feature" >}}) | フィーチャーの監査ログを表示 |

### メトリクス

| コマンド | 説明 |
|---------|------|
| [metrics list]({{< relref "/docs/commands/metrics#list" >}}) | プロジェクト内のメトリクスを一覧表示 |
| [metrics get]({{< relref "/docs/commands/metrics#get" >}}) | メトリクスの詳細を取得 |
| [metrics create]({{< relref "/docs/commands/metrics#create" >}}) | 新しいメトリクスを作成 |
| [metrics update]({{< relref "/docs/commands/metrics#update" >}}) | 既存のメトリクスを更新 |
| [metrics delete]({{< relref "/docs/commands/metrics#delete" >}}) | メトリクスを削除 |
| [metrics results]({{< relref "/docs/commands/metrics#results" >}}) | メトリクスの結果を取得 |

### Webhook

| コマンド | 説明 |
|---------|------|
| [webhooks list]({{< relref "/docs/commands/webhooks#list" >}}) | プロジェクト内の Webhook を一覧表示 |
| [webhooks get]({{< relref "/docs/commands/webhooks#get" >}}) | Webhook の詳細を取得 |
| [webhooks create]({{< relref "/docs/commands/webhooks#create" >}}) | 新しい Webhook を作成 |
| [webhooks update]({{< relref "/docs/commands/webhooks#update" >}}) | 既存の Webhook を更新 |
| [webhooks delete]({{< relref "/docs/commands/webhooks#delete" >}}) | Webhook を削除 |

### カスタムプロパティ

| コマンド | 説明 |
|---------|------|
| [custom-properties list]({{< relref "/docs/commands/custom-properties#list" >}}) | プロジェクト内のカスタムプロパティを一覧表示 |
| [custom-properties get]({{< relref "/docs/commands/custom-properties#get" >}}) | カスタムプロパティの詳細を取得 |
| [custom-properties create]({{< relref "/docs/commands/custom-properties#create" >}}) | 新しいカスタムプロパティを作成 |
| [custom-properties update]({{< relref "/docs/commands/custom-properties#update" >}}) | 既存のカスタムプロパティを更新 |
| [custom-properties delete]({{< relref "/docs/commands/custom-properties#delete" >}}) | カスタムプロパティを削除 |

### その他

| コマンド | 説明 |
|---------|------|
| [version]({{< relref "/docs/commands/version" >}}) | バージョン情報を表示 |
