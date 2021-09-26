# Welcome to Revel

> RevelはGoのフレームワーク

[参照](https://qiita.com/watataku8911/items/4762e727085aeb8fc2fe)

## Revelのインストール
*事前に$GOPATHがと打っていることを確認する。*

[確認方法](https://qiita.com/watataku8911/items/4762e727085aeb8fc2fe#1revelのインストール準備編)

```
$ go get github.com/revel/revel

$ go get github.com/revel/cmd/revel
```

## プロジェクトの作成

```
$ cd $GOPATH

$ revel new アプリケーション名
```

## 実行

```
$ revel run アプリケーション名
```
<hr></hr>

```
CREATE TABLE `dept` (
  `deptno` int(2) NOT NULL DEFAULT '0',
  `dname` text,
  `loc` text,
  PRIMARY KEY (`deptno`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 |

INSERT INTO (deptno, dname, loc) VALUES (10, "ACCOUNTING", "NEW YORK");
INSERT INTO (deptno, dname, loc) VALUES (20, "RESEARCH", "DALLAS");
INSERT INTO (deptno, dname, loc) VALUES (30, "SALES", "CHICAGO");
INSERT INTO (deptno, dname, loc) VALUES (40, "OPERATIONS", "BOSTON");
```


## Help

* The [Getting Started with Revel](http://revel.github.io/tutorial/gettingstarted.html).
* The [Revel guides](http://revel.github.io/manual/index.html).
* The [Revel sample apps](http://revel.github.io/examples/index.html).
* The [API documentation](https://godoc.org/github.com/revel/revel).

