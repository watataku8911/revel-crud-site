package deptDAO

import (
	"revelTest/app/pkg/dept"

	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func FindByPk(db *sql.DB, deptno int) *dept.Dept {
	stmt, err := db.Prepare("SELECT * FROM dept WHERE deptno = ?")
	if err != nil {
		log.Fatal(err)
	}

	var d dept.Dept
	err = stmt.QueryRow(deptno).Scan(&d.Deptno, &d.Dname, &d.Loc)
	if err != nil {
		log.Fatal(err)
	}

	return &d
}

func FindByPkCount(db *sql.DB, deptno string) int{
	var count int
	stmt, err := db.Prepare("SELECT COUNT(*) FROM dept WHERE deptno = ?")
	if err != nil {
		log.Fatal(err)
	}
	err = stmt.QueryRow(deptno).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	return count
}


func FindAll(db *sql.DB) *dept.DeptList {
	rows, err := db.Query("SELECT * FROM dept")
	if err != nil {
		log.Fatal(err);
	}

	defer rows.Close()
	var d dept.Dept
	var deptList dept.DeptList
	for rows.Next() {
		err := rows.Scan(&d.Deptno, &d.Dname, &d.Loc)
		if err != nil {
			log.Fatal(err)
		}
		deptList = append(deptList, d)
	}
	return &deptList
}
/**
 *削除
 *引数：dbオブジェクト、部門番号
 *戻り値なし
 */
func Delete(db *sql.DB, deptno string) {
	tr,err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := db.Prepare("DELETE FROM dept WHERE deptno=?")
	if err != nil {
		log.Fatal(err)
		tr.Rollback()
	}
	res, err := stmt.Exec(deptno)
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

func Update(db *sql.DB, dname, loc, deptno string) {
	tr,err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := db.Prepare("UPDATE dept SET dname=?,loc=? WHERE deptno=?")
	if err != nil {
		log.Fatal(err)
		tr.Rollback()
	}
	res, err := stmt.Exec(dname,loc, deptno)
	if err != nil {
		log.Fatal(err)
		tr.Rollback()
	}
	tr.Commit()
	fmt.Println(res)
	fmt.Println("部門番号" + deptno + "を更新しました")

}
/**
 *追加
 *引数：dbオブジェクト、部門番号、部門名、所在地
 *戻り値なし
 */
func Insert(db *sql.DB, deptno string, dname string, loc string) {
	tr, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := db.Prepare("INSERT INTO dept(deptno, dname, loc) VALUES(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
		tr.Rollback()
	}
	res, err := stmt.Exec(deptno, dname, loc)
	if err != nil {
		log.Fatal(err)
		tr.Rollback()
	}
	fmt.Println(res)
}