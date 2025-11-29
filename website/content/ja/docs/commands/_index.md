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

### その他

| コマンド | 説明 |
|---------|------|
| [version]({{< relref "/docs/commands/version" >}}) | バージョン情報を表示 |
