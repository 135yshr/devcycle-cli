# 使い方ガイド

このガイドでは、`dvcx` の全コマンドの詳細な使用例を説明します。

## 目次

- [グローバルフラグ](#グローバルフラグ)
- [認証](#認証)
- [プロジェクト](#プロジェクト)
- [フィーチャー](#フィーチャー)
- [変数](#変数)
- [環境](#環境)
- [ターゲティング](#ターゲティング)
- [バリエーション](#バリエーション)
- [オーディエンス](#オーディエンス)
- [オーバーライド](#オーバーライド)
- [監査ログ](#監査ログ)
- [メトリクス](#メトリクス)
- [Webhook](#webhook)
- [カスタムプロパティ](#カスタムプロパティ)

## グローバルフラグ

全コマンドで以下のグローバルフラグが使用できます：

```bash
--output, -o    出力形式: table, json, yaml（デフォルト: table）
--project, -p   プロジェクトキー（設定ファイルのデフォルトを上書き）
--help, -h      コマンドのヘルプを表示
```

## 認証

### ログイン

OAuth2 認証情報を使用して DevCycle に認証します。

```bash
# 設定ファイルの認証情報でログイン
dvcx auth login

# 明示的に認証情報を指定
dvcx auth login --client-id YOUR_CLIENT_ID --client-secret YOUR_CLIENT_SECRET
```

### ログアウト

保存された認証トークンを削除します。

```bash
dvcx auth logout
```

## プロジェクト

### プロジェクト一覧

```bash
# 全プロジェクトを一覧表示
dvcx projects list

# JSON で出力
dvcx projects list -o json
```

### プロジェクト詳細

```bash
# プロジェクトの詳細を取得
dvcx projects get my-project
```

### プロジェクト作成

```bash
# 新しいプロジェクトを作成
dvcx projects create --name "My Project" --key my-project --description "プロジェクトの説明"
```

### プロジェクト更新

```bash
# プロジェクト名を更新
dvcx projects update my-project --name "更新された名前"
```

## フィーチャー

### フィーチャー一覧

```bash
# プロジェクト内の全フィーチャーを一覧表示
dvcx features list -p my-project

# ステータスでフィルタ
dvcx features list -p my-project --status active
```

### フィーチャー詳細

```bash
# フィーチャーの詳細を取得
dvcx features get my-feature -p my-project
```

### フィーチャー作成

```bash
# 新しいフィーチャーを作成
dvcx features create -p my-project \
  --key my-feature \
  --name "My Feature" \
  --description "フィーチャーの説明" \
  --type release
```

### フィーチャー更新

```bash
# フィーチャーを更新
dvcx features update my-feature -p my-project --name "更新されたフィーチャー名"
```

### フィーチャー削除

```bash
# フィーチャーを削除（確認あり）
dvcx features delete my-feature -p my-project

# 確認なしで削除
dvcx features delete my-feature -p my-project --force
```

## 変数

### 変数一覧

```bash
# 全変数を一覧表示
dvcx variables list -p my-project
```

### 変数詳細

```bash
# 変数の詳細を取得
dvcx variables get my-variable -p my-project
```

### 変数作成

```bash
# Boolean 変数を作成
dvcx variables create -p my-project \
  --key my-variable \
  --name "My Variable" \
  --type Boolean

# String 変数を作成
dvcx variables create -p my-project \
  --key string-var \
  --name "String Variable" \
  --type String
```

### 変数更新

```bash
# 変数を更新
dvcx variables update my-variable -p my-project --name "更新された名前"
```

### 変数削除

```bash
# 変数を削除
dvcx variables delete my-variable -p my-project --force
```

## 環境

### 環境一覧

```bash
# 全環境を一覧表示
dvcx environments list -p my-project
```

### 環境詳細

```bash
# 環境の詳細を取得
dvcx environments get development -p my-project
```

## ターゲティング

### ターゲティング取得

```bash
# 環境内のフィーチャーのターゲティング設定を取得
dvcx targeting get -p my-project -f my-feature -e development
```

### ターゲティング更新

```bash
# ターゲティングルールを更新
dvcx targeting update -p my-project -f my-feature -e development \
  --status active \
  --rules '[{"audience":"all-users","variation":"on"}]'
```

### フィーチャーの有効化/無効化

```bash
# 環境内でフィーチャーを有効化
dvcx targeting enable -p my-project -f my-feature -e development

# 環境内でフィーチャーを無効化
dvcx targeting disable -p my-project -f my-feature -e development
```

## バリエーション

### バリエーション一覧

```bash
# フィーチャーの全バリエーションを一覧表示
dvcx variations list -p my-project -f my-feature
```

### バリエーション詳細

```bash
# バリエーションの詳細を取得
dvcx variations get variation-on -p my-project -f my-feature
```

### バリエーション作成

```bash
# 新しいバリエーションを作成
dvcx variations create -p my-project -f my-feature \
  --key variation-new \
  --name "新しいバリエーション" \
  --variables '{"my-variable": true}'
```

### バリエーション更新

```bash
# バリエーションを更新
dvcx variations update variation-on -p my-project -f my-feature \
  --name "更新されたバリエーション"
```

### バリエーション削除

```bash
# バリエーションを削除
dvcx variations delete variation-old -p my-project -f my-feature --force
```

## オーディエンス

オーディエンスを使用すると、ターゲティング用の再利用可能なユーザーセグメントを定義できます。

### オーディエンス一覧

```bash
# 全オーディエンスを一覧表示
dvcx audiences list -p my-project

# JSON で出力
dvcx audiences list -p my-project -o json
```

### オーディエンス詳細

```bash
# オーディエンスの詳細を取得
dvcx audiences get beta-users -p my-project
```

### オーディエンス作成

```bash
# フィルター付きで新しいオーディエンスを作成
dvcx audiences create -p my-project \
  --key beta-users \
  --name "ベータユーザー" \
  --description "ベータプログラムのユーザー" \
  --filters '[{"type":"user","subType":"email","comparator":"contain","values":["@beta.example.com"]}]'
```

### オーディエンス更新

```bash
# オーディエンスの名前と説明を更新
dvcx audiences update beta-users -p my-project \
  --name "ベータテスター" \
  --description "更新された説明"
```

### オーディエンス削除

```bash
# オーディエンスを削除（確認あり）
dvcx audiences delete beta-users -p my-project

# 確認なしで削除
dvcx audiences delete beta-users -p my-project --force
```

## オーバーライド

オーバーライド（セルフターゲティング）を使用すると、開発中に自分用の特定のバリエーションを設定できます。

### オーバーライド一覧

```bash
# フィーチャーの全オーバーライドを一覧表示
dvcx overrides list -p my-project -f my-feature

# プロジェクト内の自分の全オーバーライドを一覧表示
dvcx overrides list-mine -p my-project
```

### オーバーライド取得

```bash
# フィーチャーの自分の現在のオーバーライドを取得
dvcx overrides get -p my-project -f my-feature -e development
```

### オーバーライド設定

```bash
# 特定のバリエーションにオーバーライドを設定
dvcx overrides set -p my-project -f my-feature -e development \
  --variation variation-on
```

### オーバーライド削除

```bash
# フィーチャーの自分のオーバーライドを削除
dvcx overrides delete -p my-project -f my-feature -e development

# プロジェクト内の自分の全オーバーライドを削除
dvcx overrides delete-mine -p my-project
```

## 監査ログ

プロジェクト内の変更を追跡するための監査ログを表示します。

### プロジェクト監査ログ一覧

```bash
# プロジェクトの全監査ログを一覧表示
dvcx audit list -p my-project

# 詳細情報を JSON で出力
dvcx audit list -p my-project -o json
```

### フィーチャー監査ログ一覧

```bash
# 特定のフィーチャーの監査ログを一覧表示
dvcx audit feature my-feature -p my-project
```

## メトリクス

メトリクスを使用すると、フィーチャーフラグの効果を測定できます。

### メトリクス一覧

```bash
# 全メトリクスを一覧表示
dvcx metrics list -p my-project
```

### メトリクス詳細

```bash
# メトリクスの詳細を取得
dvcx metrics get conversion-rate -p my-project
```

### メトリクス作成

```bash
# カウントメトリクスを作成
dvcx metrics create -p my-project \
  --key page-views \
  --name "ページビュー" \
  --type count \
  --event-type pageview \
  --optimize-for increase \
  --description "ページビューイベントを追跡"

# コンバージョンメトリクスを作成
dvcx metrics create -p my-project \
  --key signup-rate \
  --name "サインアップ率" \
  --type conversion \
  --event-type signup \
  --optimize-for increase
```

### メトリクス更新

```bash
# メトリクスを更新
dvcx metrics update page-views -p my-project \
  --name "総ページビュー" \
  --description "更新された説明"
```

### メトリクス削除

```bash
# メトリクスを削除
dvcx metrics delete page-views -p my-project --force
```

### メトリクス結果取得

```bash
# メトリクス結果を取得
dvcx metrics results conversion-rate -p my-project

# 環境でフィルタ
dvcx metrics results conversion-rate -p my-project --environment production

# フィーチャーでフィルタ
dvcx metrics results conversion-rate -p my-project --feature my-feature

# 日付範囲でフィルタ
dvcx metrics results conversion-rate -p my-project \
  --start-date 2024-01-01 \
  --end-date 2024-01-31

# フィルタを組み合わせ
dvcx metrics results conversion-rate -p my-project \
  --environment production \
  --feature my-feature \
  --start-date 2024-01-01 \
  --end-date 2024-01-31
```

## Webhook

Webhook を使用すると、プロジェクト内でイベントが発生したときに通知を受け取ることができます。

### Webhook 一覧

```bash
# 全 Webhook を一覧表示
dvcx webhooks list -p my-project
```

### Webhook 詳細

```bash
# Webhook の詳細を取得
dvcx webhooks get webhook-id -p my-project
```

### Webhook 作成

```bash
# 有効な Webhook を作成
dvcx webhooks create -p my-project \
  --url "https://example.com/webhook" \
  --description "本番 Webhook" \
  --enabled

# 無効な Webhook を作成
dvcx webhooks create -p my-project \
  --url "https://example.com/webhook" \
  --description "テスト Webhook"
```

### Webhook 更新

```bash
# Webhook URL を更新
dvcx webhooks update webhook-id -p my-project \
  --url "https://new-url.example.com/webhook"

# Webhook を有効化
dvcx webhooks update webhook-id -p my-project --enabled

# Webhook を無効化
dvcx webhooks update webhook-id -p my-project --disabled
```

### Webhook 削除

```bash
# Webhook を削除
dvcx webhooks delete webhook-id -p my-project --force
```

## カスタムプロパティ

カスタムプロパティは、ターゲティング用の追加ユーザー属性を定義します。

### カスタムプロパティ一覧

```bash
# 全カスタムプロパティを一覧表示
dvcx custom-properties list -p my-project

# エイリアスを使用
dvcx cp list -p my-project
```

### カスタムプロパティ詳細

```bash
# カスタムプロパティの詳細を取得
dvcx custom-properties get user-type -p my-project
```

### カスタムプロパティ作成

```bash
# String プロパティを作成
dvcx custom-properties create -p my-project \
  --key user-type \
  --display-name "ユーザータイプ" \
  --type String \
  --description "ユーザーアカウントの種類"

# Boolean プロパティを作成
dvcx custom-properties create -p my-project \
  --key is-premium \
  --display-name "プレミアム会員" \
  --type Boolean \
  --description "ユーザーがプレミアム会員かどうか"

# Number プロパティを作成
dvcx custom-properties create -p my-project \
  --key account-age \
  --display-name "アカウント年齢" \
  --type Number \
  --description "アカウント作成からの日数"
```

### カスタムプロパティ更新

```bash
# カスタムプロパティを更新
dvcx custom-properties update user-type -p my-project \
  --display-name "ユーザーアカウントタイプ" \
  --description "更新された説明"
```

### カスタムプロパティ削除

```bash
# カスタムプロパティを削除
dvcx custom-properties delete user-type -p my-project --force
```

## 出力形式

全コマンドで複数の出力形式がサポートされています：

```bash
# テーブル形式（デフォルト）
dvcx features list -p my-project

# JSON 形式
dvcx features list -p my-project -o json

# YAML 形式
dvcx features list -p my-project -o yaml
```

## 設定

### 設定ファイル

プロジェクトルートに `.devcycle/config.yaml` を作成：

```yaml
client_id: your-client-id
client_secret: your-client-secret
project: default-project-key
```

### 環境変数

環境変数も使用できます：

```bash
export DVCX_CLIENT_ID=your-client-id
export DVCX_CLIENT_SECRET=your-client-secret
export DVCX_PROJECT=default-project-key
```

環境変数は設定ファイルの値より優先されます。
