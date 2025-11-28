# コントリビューションガイド

`dvcx` への貢献に興味を持っていただきありがとうございます！このガイドでは、参加方法を説明します。

[English](../contributing.md)

## 貢献方法

### Issue の報告

- バグ報告や機能リクエストには [GitHub Issues](https://github.com/135yshr/devcycle-cli/issues) を使用してください
- 新しい Issue を作成する前に、既存の Issue を確認してください
- 用意されている Issue テンプレートを使用してください

### Pull Request の提出

1. リポジトリをフォーク
2. `main` から機能ブランチを作成
3. 変更を実装
4. 必要に応じてテストを作成/更新
5. すべてのテストがパスすることを確認
6. Pull Request を提出

## 開発ワークフロー

### ブランチ命名規則

```
feature/<phase>-<task-name>
```

例：
- `feature/phase1-auth` - 認証機能
- `feature/phase1-projects-list` - プロジェクト一覧コマンド
- `fix/token-refresh` - トークンリフレッシュのバグ修正

### コミットメッセージ規則

コミットメッセージには [Gitmoji](https://gitmoji.dev/) を使用します：

```
<gitmoji> <type>: <description>

[optional body]
```

**よく使う Gitmoji:**

| 絵文字 | コード | 説明 |
|--------|--------|------|
| ✨ | `:sparkles:` | 新機能 |
| 🐛 | `:bug:` | バグ修正 |
| 📝 | `:memo:` | ドキュメント |
| ♻️ | `:recycle:` | リファクタリング |
| ✅ | `:white_check_mark:` | テストの追加/更新 |
| 🔧 | `:wrench:` | 設定 |
| ⬆️ | `:arrow_up:` | 依存関係のアップグレード |

**例：**
```
✨ feat: プロジェクト一覧コマンドを追加
🐛 fix: トークン期限切れの処理を修正
📝 docs: API リファレンスを更新
```

### Pull Request プロセス

1. PR テンプレートを完全に記入
2. 関連する Issue をリンク
3. CI チェックがパスすることを確認
4. メンテナにレビューを依頼
5. レビューのフィードバックに対応
6. 要求された場合はコミットをスカッシュ

## コードスタイル

### Go コード

- 標準的な Go の規約に従う
- コミット前に `make fmt` を実行
- `make lint` で問題をチェック
- 新機能にはテストを作成

### ドキュメント

- 英語（プライマリ）で記述
- 日本語翻訳は `docs/ja/` に配置
- README.md は簡潔に保つ
- 詳細情報は `docs/` に記載

## テスト

```bash
# すべてのテストを実行
make test

# カバレッジ付きでテストを実行
make test-coverage
```

## 質問がありますか？

お気軽に：
- 質問用の Issue を開く
- GitHub Discussions でディスカッションを始める

## 行動規範

[行動規範](../CODE_OF_CONDUCT.md)を読み、遵守してください。
