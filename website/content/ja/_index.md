---
title: "dvcx - DevCycle CLI Extended"
type: docs
---

# dvcx

**DevCycle Management API用の非公式CLIツール**

dvcxは、公式CLIでは利用できない、または制限されているDevCycleの機能にコマンドラインからアクセスできるようにします。

## 特徴

- **フルManagement APIアクセス**: すべてのDevCycle Management APIエンドポイントにアクセス
- **複数の出力フォーマット**: テーブル、JSON、YAML形式で出力
- **プロジェクトスコープの設定**: ディレクトリごとにデフォルトプロジェクト設定を保存
- **クロスプラットフォーム**: macOS、Linux、Windowsで動作

## クイックスタート

```bash
# インストール
go install github.com/135yshr/devcycle-cli@latest

# DevCycle認証情報でログイン
dvcx auth login

# プロジェクト一覧を表示
dvcx projects list

# プロジェクト内のフィーチャー一覧を表示
dvcx features list -p your-project-key
```

## ドキュメント

{{< columns >}}

### [はじめに]({{< relref "/docs/getting-started" >}})

dvcxのインストールと設定方法を学びます。

<--->

### [コマンド]({{< relref "/docs/commands" >}})

すべてのCLIコマンドの完全なリファレンス。

<--->

### [設定]({{< relref "/docs/configuration" >}})

設定オプションについて学びます。

{{< /columns >}}

## GitHub

- [リポジトリ](https://github.com/135yshr/devcycle-cli)
- [Issue](https://github.com/135yshr/devcycle-cli/issues)
- [リリース](https://github.com/135yshr/devcycle-cli/releases)
