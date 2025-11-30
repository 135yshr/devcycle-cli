# dvcx - DevCycle CLI Extended

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/badge/Go-1.24+-00ADD8?logo=go)](https://go.dev/)

[DevCycle Management API](https://docs.devcycle.com/management-api/) の非公式コマンドラインツールです。

[English](../../README.md)

## なぜ dvcx？

DevCycleには[公式CLI](https://docs.devcycle.com/cli/)がありますが、高度な操作には機能が限られています。`dvcx` は Management API への包括的なアクセスを提供し、以下を可能にします：

- すべてのリソース（Projects、Features、Variables など）の完全な CRUD 操作
- 一括操作と自動化
- 詳細な監査ログへのアクセス
- 高度なターゲティングとオーディエンス管理

## インストール

### ソースから

```bash
git clone https://github.com/135yshr/devcycle-cli.git
cd devcycle-cli
make build
```

バイナリは `bin/dvcx` に生成されます。

### Go Install

```bash
go install github.com/135yshr/devcycle-cli@latest
```

## クイックスタート

1. **認証情報を設定**

   プロジェクトルートに `.devcycle/config.yaml` を作成：

   ```yaml
   client_id: your-client-id
   client_secret: your-client-secret
   project: your-project-key
   ```

2. **ログイン**

   ```bash
   dvcx auth login
   ```

3. **プロジェクト一覧を表示**

   ```bash
   dvcx projects list
   ```

## コマンド

> 注: このプロジェクトは開発中です。実装状況は[ロードマップ](../roadmap.md)を参照してください。

### コアコマンド

| コマンド | 説明 | ステータス |
|---------|------|----------|
| `auth login/logout` | 認証 | ✅ 完了 |
| `projects list/get/create/update` | プロジェクト管理 | ✅ 完了 |
| `features list/get/create/update/delete` | フィーチャーフラグ | ✅ 完了 |
| `variables list/get/create/update/delete` | 変数 | ✅ 完了 |
| `environments list/get/create/update/delete` | 環境 | ✅ 完了 |
| `targeting get/update/enable/disable` | フィーチャーターゲティング | ✅ 完了 |
| `variations list/get/create/update/delete` | フィーチャーバリエーション | ✅ 完了 |

### オーディエンス & オーバーライド (Phase 4)

| コマンド | 説明 | ステータス |
|---------|------|----------|
| `audiences list/get/create/update/delete` | オーディエンス管理 | ✅ 完了 |
| `overrides list/get/set/delete` | セルフターゲティングオーバーライド | ✅ 完了 |

### 運用 & モニタリング (Phase 5)

| コマンド | 説明 | ステータス |
|---------|------|----------|
| `audit list/feature` | 監査ログ | ✅ 完了 |
| `metrics list/get/create/update/delete/results` | メトリクス管理 | ✅ 完了 |
| `webhooks list/get/create/update/delete` | Webhook 管理 | ✅ 完了 |
| `custom-properties list/get/create/update/delete` | カスタムプロパティ | ✅ 完了 |

### 環境管理 (Phase 6)

| コマンド | 説明 | ステータス |
|---------|------|----------|
| `keys list/rotate` | SDK キー管理 | ✅ 完了 |

詳細な使用例は [使い方ガイド](usage.md) を参照してください。

## ドキュメント

- [使い方ガイド](usage.md) - コマンドの詳細な使用例
- [API リファレンス](../api-reference.md) - DevCycle Management API エンドポイント
- [ロードマップ](../roadmap.md) - 実装フェーズと進捗
- [コントリビューションガイド](contributing.md) - 貢献方法
- [開発ガイド](../development.md) - 開発環境セットアップとアーキテクチャ

## 開発

### 前提条件

- Go 1.24+
- [pre-commit](https://pre-commit.com/) - Git hooks 用
- [markdownlint-cli2](https://github.com/DavidAnson/markdownlint-cli2) - Markdown lint 用

### セットアップ

```bash
# pre-commit と markdownlint-cli2 をインストール
brew install pre-commit markdownlint-cli2

# pre-commit hooks をインストール
pre-commit install
```

詳細なセットアップ手順は [開発ガイド](../development.md) を参照してください。

## コントリビューション

コントリビューションを歓迎します！Pull Request を送る前に[コントリビューションガイド](contributing.md)をお読みください。

## ライセンス

このプロジェクトは MIT ライセンスの下で公開されています - 詳細は [LICENSE](../../LICENSE) ファイルを参照してください。

## 免責事項

これは非公式ツールであり、DevCycle とは関係ありません。ご自身の責任でご使用ください。
