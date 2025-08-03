# 🥃 Whisky Web API

Go言語 + Ginフレームワークで作成したRESTful API

## 📖 概要

ウイスキーの情報を管理するWeb APIです。CRUD操作、ヘルスチェック機能、ユニットテスト完備の本格的なAPIです。

## 🚀 機能

- ✅ **ウイスキー一覧取得** `GET /whiskies`
- ✅ **ウイスキー詳細取得** `GET /whiskies/:id`
- ✅ **ウイスキー新規作成** `POST /whiskies`
- ✅ **ウイスキー削除** `DELETE /whiskies/:id`
- ✅ **ヘルスチェック** `GET /health`

## 🛠️ 技術スタック

- **言語**: Go 1.21+
- **フレームワーク**: Gin
- **テスト**: Go標準テスト + testify
- **データ形式**: JSON

## 🏃‍♂️ 実行方法

### 1. 依存関係のインストール
```bash
go mod download
```

### 2. サーバー起動
```bash
go run main.go
```

サーバーは `http://localhost:8000` で起動します。

## 📋 API使用例

### ウイスキー一覧取得
```bash
curl http://localhost:8000/whiskies
```

### 特定のウイスキー取得
```bash
curl http://localhost:8000/whiskies/1
```

### 新しいウイスキー追加
```bash
curl -X POST http://localhost:8000/whiskies \
  -H "Content-Type: application/json" \
  -d '{
    "name": "山崎",
    "region": "日本",
    "type": "シングルモルト",
    "abv": 43.0,
    "price": 15000
  }'
```

### ウイスキー削除
```bash
curl -X DELETE http://localhost:8000/whiskies/3
```

### ヘルスチェック
```bash
curl http://localhost:8000/health
```

## 🧪 テスト

### テスト実行
```bash
go test
```

### 詳細なテスト結果
```bash
go test -v
```

### カバレッジ確認
```bash
go test -cover
```

### HTMLでカバレッジ確認
```bash
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## 📊 レスポンス例

### ウイスキー一覧
```json
[
  {
    "id": 1,
    "name": "デュワーズ",
    "region": "スコットランド",
    "type": "ブレンデッド",
    "abv": 40,
    "price": 12000
  },
  {
    "id": 2,
    "name": "駒ヶ岳",
    "region": "日本",
    "type": "シングルモルト",
    "abv": 52,
    "price": 9460
  }
]
```

### ヘルスチェック
```json
{
  "status": "OK",
  "message": "Whisky API is running",
  "whisky_count": 3
}
```

## 🏗️ プロジェクト構成

```
whisky-web-api/
├── main.go          # API実装
├── main_test.go     # ユニットテスト
├── go.mod          # モジュール定義
├── go.sum          # 依存関係のハッシュ
├── coverage.out    # カバレッジ結果
└── README.md       # このファイル
```

## 🎯 特徴

- **高速レスポンス**: マイクロ秒レベルの応答時間
- **型安全**: Go言語の型システムによる安全性
- **テストカバレッジ**: 包括的なユニットテスト
- **RESTful設計**: 標準的なHTTPメソッドとステータスコード
- **エラーハンドリング**: 適切なエラーレスポンス
- **運用監視**: ヘルスチェック機能

## 🚀 今後の拡張予定

- [ ] データベース連携（PostgreSQL/SQLite）
- [ ] 認証・認可機能
- [ ] Docker化
- [ ] CI/CD パイプライン
- [ ] AWS デプロイ
- [ ] テイスティング記録機能

## 👨‍💻 作者

**UhRhythm**

## 📄 ライセンス

MIT License
