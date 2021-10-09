package empDAO

import (
	"revelTest/app/pkg/emp"

	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
	"log"

)

func FindByPk(db *sql.DB, empno int) *emp.Emp {
	stmt, err := db.Prepare("SELECT * FROM emp WHERE empno = ?")
	if err != nil {
		log.Fatal(err)
	}

	var e emp.Emp
	err = stmt.QueryRow(empno).Scan(&e.Empno, &e.Ename, &e.Job, &e.Mgr, &e.Hiredate, &e.Sal, &e.Comm, &e.Deptno)
	if err != nil {
		log.Fatal(err)
	}

	return &e
}

func FindByPkCount(db *sql.DB, empno int) int{
	var count int
	stmt, err := db.Prepare("SELECT COUNT(*) FROM emp WHERE empno = ?")
	if err != nil {
		log.Fatal(err)
	}
	err = stmt.QueryRow(empno).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	return count
}

func FindByMgr(db *sql.DB, mgr *int) *emp.Emp {
	stmt, err := db.Prepare("SELECT ename FROM emp WHERE empno = ?")
	if err != nil {
		log.Fatal(err)
	}

	var e emp.Emp
	err = stmt.QueryRow(&mgr).Scan(&e.Ename)
	if err != nil {
		log.Fatal(err)
	}

	return &e
}


func FindAll(db *sql.DB) *emp.EmpList {
	rows, err := db.Query("SELECT * FROM emp")
	if err != nil {
		log.Fatal(err);
	}

	defer rows.Close()
	var e emp.Emp
	var empList emp.EmpList
	for rows.Next() {
		if &e.Mgr != nil {
			err := rows.Scan(&e.Empno, &e.Ename, &e.Job, &e.Mgr, &e.Hiredate, &e.Sal, &e.Comm, &e.Deptno)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			err := rows.Scan(&e.Empno, &e.Ename, &e.Job, &e.Mgr, &e.Hiredate, &e.Sal, &e.Comm, &e.Deptno)
			if err != nil {
				log.Fatal(err)
			}
		}
		empList = append(empList, e)
	}
	return &empList
}
/**
 *削除
 *引数：dbオブジェクト、部門番号
 *戻り値なし
 */
func Delete(db *sql.DB, empno string) {
	tr,err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := db.Prepare("DELETE FROM emp WHERE empno=?")
	if err != nil {
		log.Fatal(err)
		tr.Rollback()
	}
	res, err := stmt.Exec(empno)
	if err != nil {
		log.Fatal(err)
		tr.Rollback()
	}
	fmt.Println(res)


}
/**
 *更新
 *引数：dbオブジェクト、部門名、所在地、部門番号
 *戻り値なし
 */

func Update(db *sql.DB, empno, ename, job, mgr, hiredate, sal, comm, deptno string) {
	tr,err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := db.Prepare("UPDATE emp SET dname=?,loc=? WHERE empno=?")
	if err != nil {
		log.Fatal(err)
		tr.Rollback()
	}
	res, err := stmt.Exec(empno, ename, job, mgr, hiredate, sal, comm, deptno)
	if err != nil {
		log.Fatal(err)
		tr.Rollback()
	}
	tr.Commit()
	fmt.Println(res)
	fmt.Println("部門番号" + empno + "を更新しました")

}
/**
 *追加
 *引数：dbオブジェクト、部門番号、部門名、所在地
 *戻り値なし
 */
func Insert(db *sql.DB, empno, ename, job, mgr, hiredate, sal, comm, deptno string) {
	tr, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := db.Prepare("INSERT INTO emp (empno, ename, job, mgr, hiredate, sal, comm, deptno) VALUES(?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
		tr.Rollback()
	}
	res, err := stmt.Exec(empno, ename, job, mgr, hiredate, sal, comm, deptno)
	if err != nil {
		log.Fatal(err)
		tr.Rollback()
	}
	fmt.Println(res)
}