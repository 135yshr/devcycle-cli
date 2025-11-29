---
title: "version"
weight: 10
---

# version

dvcxのバージョン情報を表示します。

## 使用方法

```bash
dvcx version
```

## 出力

このコマンドは以下を表示します：

- **Version**: セマンティックバージョン番号
- **Commit**: ビルドのGitコミットハッシュ
- **Built at**: ビルド日時
- **Go version**: 使用されたGoコンパイラのバージョン
- **Platform**: オペレーティングシステムとアーキテクチャ

## 例

```bash
$ dvcx version
dvcx version v0.1.0
  commit: abc1234
  built at: 2025-01-15T10:00:00Z
  go version: go1.24
  platform: darwin/arm64
```

## 短縮バージョン

ルートコマンドで`--version`フラグを使用することもできます：

```bash
$ dvcx --version
dvcx version v0.1.0
```

## 注意事項

- バージョン情報はGoのldflagsを使用してビルド時に注入されます
- コミットハッシュは正確なソースコードバージョンを識別するのに役立ちます
- デバッグや問題報告に便利です
