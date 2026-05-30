# go-ogen

`ogen` を使って Go の HTTP API を実装してみる個人プロジェクトです。

OpenAPI 仕様を先に書き、`ogen` でサーバーコードを生成し、手書きのアプリケーションロジックを `internal` 配下に実装します。

## Current Direction

- API contract: `api/openapi.yaml`
- Generated server package: `internal/gen/todoapi`
- Application logic: `internal/service`
- Entrypoint: `cmd/server`

最初の実装対象は小さな TODO API です。これは外部向けデモではなく、`ogen` の生成フロー、生成コードとの付き合い方、通常の Go プロジェクト構成を確認するための最小ドメインとして扱います。
