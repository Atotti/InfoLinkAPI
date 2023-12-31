# 開発手順

## 環境構築

Dockerをインストールして起動しておきます。(Docker Desktopが良いと思います)
Docker以外に環境構築する必要はありません。Dockerの導入方法は適当に調べてください！

Dockerコンテナを起動します。
```
docker-compose up --build -d
```

データベースのdumpファイルを`schema/dump.sql`に配置してコンテナ上にコピーします。
```
docker cp ./schema/dump.sql mysql-container:/tmp/dump.sql
```

mysqlのコンテナに入ります
```
docker exec -it mysql-container bash
```

dumpされたデータをデータベースに流し込みます。
```
mysql -u root -p demo < /tmp/dump.sql
```

コンテナを再起動します。
```
docker-compose up --build -d
```
[Swagger UI](http://localhost:8080/swagger/index.html)にアクセスして正常に動作しているか確認します。
データベースのコンテナがたってないとgoアプリケーションのコンテナが落ちることがあるので、その時はコンテナを再起動してください。

## 開発時の手順
コードを書いたら、以下の手順を実行して、テストと実行をします。

### パッケージを追加したら、go.modに反映させます。
```
docker-compose run tidy
```
コンテナの中で以下を実行します。
```
go mod tidy
```
コンテナから抜けます。
```
exit
```
### swaggerドキュメントを生成します。
```
docker-compose run swag
```
コンテナ中で以下を実行します。
```
swag init
```
コンテナから抜けます。
```
exit
```
#### テストを実行します。
```
docker-compose up --build test
```
### 問題が無かったら、アプリケーションを実行します。
```
docker-compose up --build -d
```
[Swagger UI](http://localhost:8080/swagger/index.html)にアクセスして正常に動作しているか確認します。

正常に動作していて、完成したらプルリクエスを発行します。

# その他
docker un したプロセスが溜まるっていう問題があります。

## テーブル名は必ず複数形
GORMのドキュメントをよく読むこと