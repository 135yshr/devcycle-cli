---
title: "はじめに"
weight: 1
---

# はじめに

このガイドでは、dvcxをシステムにインストールして設定する方法を説明します。

## 前提条件

dvcxをインストールする前に、以下が必要です：

1. **DevCycleアカウント**: [devcycle.com](https://devcycle.com)でサインアップ
2. **API認証情報**: DevCycleダッシュボードからClient IDとClient Secretを取得

### API認証情報の取得

1. DevCycleダッシュボードにログイン
2. **Settings** → **API Credentials**に移動
3. 新しい認証情報を作成するか、既存のものを使用
4. **Client ID**と**Client Secret**をメモ

## インストール

### Go Installを使用

Goがインストールされている場合、最も簡単な方法です：

```bash
go install github.com/135yshr/devcycle-cli@latest
```

### バイナリをダウンロード

[リリースページ](https://github.com/135yshr/devcycle-cli/releases)からプラットフォーム用のバイナリをダウンロードします。

{{< tabs "installation" >}}
{{< tab "macOS (Apple Silicon)" >}}

```bash
curl -L https://github.com/135yshr/devcycle-cli/releases/latest/download/dvcx_darwin_arm64.tar.gz | tar xz
sudo mv dvcx /usr/local/bin/
```

{{< /tab >}}
{{< tab "macOS (Intel)" >}}

```bash
curl -L https://github.com/135yshr/devcycle-cli/releases/latest/download/dvcx_darwin_amd64.tar.gz | tar xz
sudo mv dvcx /usr/local/bin/
```

{{< /tab >}}
{{< tab "Linux" >}}

```bash
curl -L https://github.com/135yshr/devcycle-cli/releases/latest/download/dvcx_linux_amd64.tar.gz | tar xz
sudo mv dvcx /usr/local/bin/
```

{{< /tab >}}
{{< tab "Windows" >}}
リリースページから`.zip`ファイルをダウンロードし、展開したフォルダをPATHに追加してください。
{{< /tab >}}
{{< /tabs >}}

### ソースからビルド

```bash
git clone https://github.com/135yshr/devcycle-cli.git
cd devcycle-cli
make build
```

バイナリは`bin/dvcx`に作成されます。

## 認証

インストール後、DevCycle認証情報で認証します：

```bash
dvcx auth login
```

以下の入力を求められます：

- **Client ID**: DevCycle APIクライアントID
- **Client Secret**: DevCycle APIクライアントシークレット

認証情報はプロジェクトディレクトリの`.devcycle/token.json`に安全に保存されます。

## インストールの確認

dvcxが正しくインストールされていることを確認：

```bash
dvcx version
```

バージョン情報が表示されます：

```
dvcx version v0.1.0
  commit: abc1234
  built at: 2025-01-15T10:00:00Z
  go version: go1.24
  platform: darwin/arm64
```

## 最初のコマンド

セットアップが完了したら、以下のコマンドを試してください：

```bash
# すべてのプロジェクトを一覧表示
dvcx projects list

# 特定のプロジェクトの詳細を取得
dvcx projects get my-project-key

# プロジェクト内のフィーチャーを一覧表示
dvcx features list -p my-project-key
```

## 次のステップ

- すべての[コマンド]({{< relref "/docs/commands" >}})について学ぶ
- [デフォルト設定]({{< relref "/docs/configuration" >}})を構成する
- よくある質問は[FAQ]({{< relref "/docs/faq" >}})を参照
