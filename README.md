# mongodb-view-vs-pipeline

MongoDB において、View で検索するのとパイプラインで検索するのとではどちらがより高速か検証する

## 実行方法

MongoDB の起動。

```zsh
docker run -d --name mongodb_view_vs_pipeline -p 27017:27017 mongo
```

実行

```zsh
go run main.go
```

## 実行結果

```
mongodb-view-vs-pipeline % go run main.go
Inserting sample data...
Sample data inserted successfully.
Creating MongoDB View...
View created successfully.
View Find fetched 50000 documents.
Aggregation Find fetched 50000 documents.
View find time: 48.502167ms
Aggregation find time: 43.696917ms
```
