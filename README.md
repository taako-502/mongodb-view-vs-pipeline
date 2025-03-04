# mongodb-view-vs-pipeline

MongoDB において、View で検索するのとパイプラインで検索するのとではどちらがより高速か検証する

記事：
https://ap-ep.com/mongodb-view-vs-pipeline/

## **テスト内容**
このリポジトリでは、以下の 5 つのテストパターンを実施し、
**MongoDB View と Aggregation Pipeline のパフォーマンス比較** を行う。

### **データ量の増減による影響**
- **データ量を 10,000件 / 100,000件 / 1,000,000件 / 10,000,000件 に増減** させ、
  **データサイズが大きくなったときの View と Aggregation の速度変化** を検証。

※
　本来は以下の項目も検証するべきだが、今回は割愛
- インデックスの有無による影響
  - `score` フィールドに **インデックスあり / なし** の 2 パターンで、検索速度がどのように変化するかを検証。
- クエリの複雑さによる影響
  - 以下のようなクエリを使用し、**シンプルな検索と複雑な検索での速度比較** を行う。
    - **シンプルな `$match` のみ**
    - **フィルタ + ソート (`$match` + `$sort` + `$limit`)**
    - **集計 (`$match` + `$group` + `$sort`)**
- 取得件数による影響
  - **全件取得 vs 最初の 10 件のみ取得 (`Find().Limit(10)`)** で、それぞれ View と Aggregation の速度を比較。
- データ追加・更新による影響
  - `InsertMany()` で 10,000 件のデータを追加し、その直後に View と Aggregation で検索を実行。
    **ビューのリアルタイム更新がパフォーマンスに及ぼす影響** を検証。

## 実行方法

```zsh
$ make benchmark
```

### 考察
**ViewよりもAggregation Pipelineのほうがわずかに高速**

（1）**ビュー (View) の仕組み**

MongoDB の View は、内部的に Aggregation Pipeline によりデータを生成しており、実行時に同様の処理が走る。<br>
そのため、Aggregation Pipeline とほぼ同等のパフォーマンスとなる。<br>
ただし、View は仮想コレクションとして扱われるため、特定の状況下では内部で追加のオーバーヘッドが発生する可能性がある。

（2）**Aggregation Pipeline の利点**

Aggregation Pipeline は柔軟性が高く、パイプラインの各段階で細かい最適化が行われる場合がある。<br>
また、Aggregation の実行計画のキャッシュや最適化がうまく働くと、View よりも若干速くなるケースが見られる可能性がある。


**データ量が増加するにつれて、Aggregation Pipeline のほうが高速**

特に10,000,000件の場合、Aggregation のほうが約200〜250ミリ秒程度短縮できている点から、大規模データでの高速性を求める場合は Aggregation Pipeline を採用するメリットがあるかもしれない。
