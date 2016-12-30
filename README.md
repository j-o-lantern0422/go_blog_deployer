# github_blog
githubでブログの記事管理してアレコレしてくれるすごいやつ。

# 簡単な使い方

  1. githubからWebhokを受け取ることのできるサーバなどで `go run server.go hatena.go github.go` などとするか、ビルドしたバイナリを置きます
  1. サーバの3000番ポートで受け取るようになっているので開けてあげてください
  1. `config.json` にはてなブログのAPIを利用するための情報と、そのはてなブログに紐づくリポジトリの名前を設定します
  1. `githuconfig.json` にgithubのアカウント名やトークンを設定します
  1. githubにブログ用のリポジトリを作り、記事ごとにブランチを切ります
  1. そこに記事をプッシュします
  1. プルリクを出します
    1. ラベルがカテゴリになるのでカテゴリを設定したい場合プルリクにラベルをつけます
  1. マージしたらそれをトリガーにはてなブログに下書きステータスで投稿されます
    1. これは、今のところ設定やマージ時のコメント等で公開ステータスにできないので今後変更予定


# Feature

 - [ ] おひとりさまなブログであることが前提みたいな設計になっているところがあるので変更する必要を感じている
 - [ ] Wordpressに対応したい
 - [ ] 作りが少々雑なので整えてちゃんとしたドキュメントを作る
