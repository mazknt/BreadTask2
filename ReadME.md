# BreadTask2

##確認手順

- `.env`をルートに配置し、以下を設定する
  - PROJECT_ID(Firestore の project id)
  - CREDENTIAL_OPTION(Firestore の credential option)
- `go run ./server.go`でサーバーを起動
- 任意のリクエストを送ってレスポンスを確認

###レスポンス例

- a. 1 で保存したデータを一覧で取得

```
query {
  getAllBreads {
    id
    breadInfo {
      name
      createdAt
    }
  }
}

```

- b. id を指定して、1 で保存したデータ中の任意のデータを取得

```
query {
  getBread(id: "特定のパンのID") {
    id
    breadInfo {
      name
      createdAt
    }
  }
}
```
