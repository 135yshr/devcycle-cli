---
title: "FAQ"
weight: 4
---

# よくある質問

## インストール

### dvcxをインストールするには？

Goを使用するのが最も簡単です：

```bash
go install github.com/135yshr/devcycle-cli@latest
```

または[リリースページ](https://github.com/135yshr/devcycle-cli/releases)からバイナリをダウンロードしてください。

### DevCycle API認証情報はどこで見つけられますか？

1. [DevCycleダッシュボード](https://app.devcycle.com)にログイン
2. **Settings** → **API Credentials**に移動
3. 新しい認証情報を作成するか、既存のものをコピー

## 認証

### 認証トークンが期限切れになりました。どうすればいいですか？

`dvcx auth login`を再度実行して新しいトークンを取得してください。

### トークンはどこに保存されますか？

トークンはカレントプロジェクトディレクトリの`.devcycle/token.json`に保存されます。

### CI/CDでdvcxを使用するには？

`DVCX_CLIENT_ID`と`DVCX_CLIENT_SECRET`環境変数を設定します：

```bash
export DVCX_CLIENT_ID=your-client-id
export DVCX_CLIENT_SECRET=your-client-secret
dvcx auth login
```

## コマンド

### デフォルトプロジェクトを設定するには？

`.devcycle/config.yaml`ファイルを作成します：

```yaml
project: your-project-key
```

### JSON出力を取得するには？

`-o json`フラグを使用します：

```bash
dvcx projects list -o json
```

### dvcxと公式DevCycle CLIの違いは？

dvcxは以下を提供する非公式ツールです：

- フルManagement APIアクセス
- 複数の出力形式（テーブル、JSON、YAML）
- プロジェクトスコープの設定
- 公式CLIでは利用できない追加機能

## トラブルシューティング

### 「Project key is required」エラー

以下のいずれかを行ってください：

1. `-p`フラグでプロジェクトを指定：`dvcx features list -p my-project`
2. `.devcycle/config.yaml`でデフォルトプロジェクトを設定

### 「Unauthorized」エラー

トークンが期限切れになっている可能性があります。`dvcx auth login`を実行して再認証してください。

### コマンドが見つからない

バイナリがPATHにあることを確認してください：

```bash
# dvcxがアクセス可能か確認
which dvcx

# go installを使用している場合、GOPATH/binがPATHにあることを確認
export PATH=$PATH:$(go env GOPATH)/bin
```

## コントリビューション

### どのように貢献できますか？

[コントリビューションガイド](https://github.com/135yshr/devcycle-cli/blob/main/docs/contributing.md)を参照してください：

- 開発環境のセットアップ
- プルリクエストの提出
- コーディング規約

### バグはどこに報告しますか？

[GitHub Issues](https://github.com/135yshr/devcycle-cli/issues)でイシューを開いてください。
