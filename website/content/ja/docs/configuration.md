---
title: "設定"
weight: 3
---

# 設定

dvcxは設定とデフォルト値を保存するためにYAML設定ファイルを使用します。

## 設定ファイルの場所

設定ファイルはプロジェクトディレクトリの`.devcycle/config.yaml`にあります。

```
your-project/
├── .devcycle/
│   ├── config.yaml    # 設定ファイル
│   └── token.json     # 認証トークン（自動生成）
├── src/
└── ...
```

## 設定オプション

### 設定例

```yaml
# .devcycle/config.yaml

# デフォルトプロジェクトキー
project: my-app

# デフォルト出力形式 (table, json, yaml)
output: table
```

### 利用可能なオプション

| オプション | タイプ | 説明 | デフォルト |
|----------|------|------|---------|
| `project` | string | コマンドのデフォルトプロジェクトキー | (なし) |
| `output` | string | デフォルト出力形式 | `table` |

## デフォルトプロジェクトの設定

すべてのコマンドで`--project`を指定する代わりに、デフォルトを設定できます：

```yaml
# .devcycle/config.yaml
project: my-app
```

これで`-p`フラグなしでコマンドを実行できます：

```bash
# 以前：プロジェクトフラグが必要
dvcx features list -p my-app

# 以後：設定のデフォルトを使用
dvcx features list
```

## 出力形式

希望するデフォルト出力形式を設定します：

```yaml
output: json  # または table, yaml
```

## 環境変数

設定オプションは`DVCX_`プレフィックス付きの環境変数でも設定できます：

| 環境変数 | 設定オプション |
|---------|--------------|
| `DVCX_PROJECT` | `project` |
| `DVCX_OUTPUT` | `output` |

### 例

```bash
export DVCX_PROJECT=my-app
export DVCX_OUTPUT=json

dvcx features list  # 環境変数を使用
```

## 優先順位

設定値は以下の順序で解決されます（優先度が高い順）：

1. **コマンドラインフラグ** (`--project`, `--output`)
2. **環境変数** (`DVCX_PROJECT`, `DVCX_OUTPUT`)
3. **設定ファイル** (`.devcycle/config.yaml`)
4. **デフォルト値**

## 認証トークン

認証トークンは`.devcycle/token.json`に別途保存されます：

```json
{
  "access_token": "...",
  "expires_at": "2025-01-15T12:00:00Z"
}
```

{{< hint warning >}}
**セキュリティ注意**: 認証情報を誤ってコミットしないように、`.devcycle/token.json`を`.gitignore`に追加してください。
{{< /hint >}}

### 推奨.gitignore

```gitignore
# DevCycle CLI
.devcycle/token.json
```

## カスタム設定ファイルパス

`--config`フラグを使用してカスタム設定ファイルを指定できます：

```bash
dvcx --config /path/to/custom-config.yaml projects list
```
