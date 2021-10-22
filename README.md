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
# DB作成
CREATE DATABASE `wp32scott` CHARACTER SET utf8;

# 部門テーブル
CREATE TABLE `dept` (
  `deptno` int(2) NOT NULL DEFAULT '0',
  `dname` text,
  `loc` text,
  PRIMARY KEY (`deptno`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO (deptno, dname, loc) VALUES (10, "ACCOUNTING", "NEW YORK");
INSERT INTO (deptno, dname, loc) VALUES (20, "RESEARCH", "DALLAS");
INSERT INTO (deptno, dname, loc) VALUES (30, "SALES", "CHICAGO");
INSERT INTO (deptno, dname, loc) VALUES (40, "OPERATIONS", "BOSTON");

# 従業員テーブル
CREATE TABLE `emp` (
  `empno` int(4) NOT NULL DEFAULT '0',
  `ename` text,
  `job` text,
  `mgr` int(4) DEFAULT NULL,
  `hiredate` date DEFAULT NULL,
  `sal` decimal(7,2) DEFAULT NULL,
  `comm` decimal(7,2) DEFAULT NULL,
  `deptno` int(2) DEFAULT NULL,
  PRIMARY KEY (`empno`),
  KEY `deptno` (`deptno`),
  CONSTRAINT `emp_ibfk_1` FOREIGN KEY (`deptno`) REFERENCES `dept` (`deptno`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO emp (empno, ename, job, mgr, hiredate, sal, comm, deptno) VALUES (7369,'SMITH','CLERK',7902,'1980-12-17',800,null,20);
INSERT INTO emp (empno, ename, job, mgr, hiredate, sal, comm, deptno) VALUES (7499,'ALLEN','SALESMAN',7698,'1981-02-20',1600,300,30);
INSERT INTO emp (empno, ename, job, mgr, hiredate, sal, comm, deptno) VALUES (7521,'WARD','SALESMAN',7698,'1981-02-22',1250,500,30);
INSERT INTO emp (empno, ename, job, mgr, hiredate, sal, comm, deptno) VALUES (7566,'JONES','MANAGER',7839,'1981-04-02',2975,null,20);
INSERT INTO emp (empno, ename, job, mgr, hiredate, sal, comm, deptno) VALUES (7654,'MARTIN','SALESMAN',7698,'1981-09-28',1250,1400,30);
INSERT INTO emp (empno, ename, job, mgr, hiredate, sal, comm, deptno) VALUES (7698,'BLAKE','MANAGER',7839,'1981-05-01',2850,null,30);
INSERT INTO emp (empno, ename, job, mgr, hiredate, sal, comm, deptno) VALUES (7782,'CLARK','MANAGER',7839,'1981-06-09',2450,null,10);
INSERT INTO emp (empno, ename, job, mgr, hiredate, sal, comm, deptno) VALUES (7788,'SCOTT','ANALYST',7566,'1987-04-19',3000,null,20);
INSERT INTO emp (empno, ename, job, mgr, hiredate, sal, comm, deptno) VALUES (7839,'KING','PRESIDENT',null,'1981-11-17',5000,null,10);
INSERT INTO emp (empno, ename, job, mgr, hiredate, sal, comm, deptno) VALUES (7844,'TURNER','SALESMAN',7698,'1981-09-08',1500,0,30);
INSERT INTO emp (empno, ename, job, mgr, hiredate, sal, comm, deptno) VALUES (7876,'ADAMS','CLERK',7788,'1987-05-23',1100,null,20);
INSERT INTO emp (empno, ename, job, mgr, hiredate, sal, comm, deptno) VALUES (7900,'JAMES','CLERK',7698,'1981-12-03',950,null,30);
INSERT INTO emp (empno, ename, job, mgr, hiredate, sal, comm, deptno) VALUES (7902,'FORD','ANALYST',7566,'1981-12-03',3000,null,20);
INSERT INTO emp (empno, ename, job, mgr, hiredate, sal, comm, deptno) VALUES (7934,'MILLER','CLERK',7782,'1982-01-23',1300,null,10);
```


## Help

* The [Getting Started with Revel](http://revel.github.io/tutorial/gettingstarted.html).
* The [Revel guides](http://revel.github.io/manual/index.html).
* The [Revel sample apps](http://revel.github.io/examples/index.html).
* The [API documentation](https://godoc.org/github.com/revel/revel).

