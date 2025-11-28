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

| コマンド | 説明 | ステータス |
|---------|------|----------|
| `auth login/logout` | 認証 | 予定 |
| `projects list/get` | プロジェクト管理 | 予定 |
| `features list/get/create/update/delete` | フィーチャーフラグ | 予定 |
| `variables list/get/create/update/delete` | 変数 | 予定 |
| `environments list/get` | 環境 | 予定 |

## ドキュメント

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
