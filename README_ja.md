# Typetalk Google 翻訳ボット

このリポジトリは Typetalk で Webhook 機能を使ったデモを含めています。

Typetalk ユーザーが英語でないメッセージを投稿したときに、このボットはそのメッセージを英語に翻訳を試みます。翻訳が成功した場合は、そのボットは英語翻訳したメッセージを返信します。

## 始め方
### 1. Google から API キーを取得

Bot 用の API キーを取得する方法は、次の Google ドキュメントを参照してください: https://support.google.com/cloud/answer/6158862?hl=ja

### 2. リポジトリをクローン

このリポジトリをアクセス可能なサーバーの場所にクローンしてください。

```
# via SSH:
$ git clone git@github.com:nulab/typetalk-google-translate-bot.git

# via HTTPS:
$ git clone https://github.com/nulab/typetalk-google-translate-bot.git
```

### 3. アプリケーションを実行

Go ライブラリを取得し、先ほど取得した Google API キーを環境変数に設定後、アプリケーションを実行します。

```
$ go get all
$ export $GOOGLE_API_KEY=????????????????
$ go run app.go
```

### 4. Typetalk ボット作成

Typetalk ボットを英語翻訳させたいトピックに作成してください。

作成に関しては、 Typetalk の Webhook ドキュメントページを参照してください: https://developer.nulab-inc.com/ja/docs/typetalk/#webhook

Webhook の URL はリポジトリを複製したホスト名、ポートは `12345` にしてください。

例） https://example.com:12345/

もしも実際に公開して使用したい場合は、リバースプロキシを使うなどして別のポートを使用することを強く推奨します（HTTPS を使うことを推奨します）

### 5. 動作確認

実際に英語でないメッセージ "こんにちは" を投稿してみましょう。うまく動作すれば、そのボットは "Hello" と返信してくれるはずです！
