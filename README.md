# gogo-postgres

## 動かし方

- ohomebox-by-goルートディレクトリにて以下を実行する

```sh
docker-compose build
```
```sh
docker-compose up
```

- `localhost:8080/home` にアクセスする

## 作ろうとしているもの
OHOMEBOXというWEBアプリ
- 人への褒め言葉を登録
- 人を選んでその人に向けて登録された褒め言葉をランダム表示

### 使い方の流れ
1. お褒め/お褒め登録から「ユーザーさん」を選んで褒め言葉を入力して「Add OHOME」
1. お褒め/お褒め承認で↑で登録した褒め言葉にチェックを入れて「チェックしたお褒めを承認する」
1. OHOME BOXクリックで遷移するhome画面にて「ユーザーさん」を選んで「お褒めをランダム表示」
1. 登録した褒め言葉が表示される

