package controllers

import (
	"database/sql"
	"fmt"
	"github.com/revel/revel"
	"log"
	"regexp"
	"revelTest/app/pkg/dao"
)

type App struct {
	*revel.Controller
}

func (c App) Error() revel.Result  {
	return c.Render()
}

func (c App) Index() revel.Result {
	message := "Revel"
	return c.Render(message)
}

func (c App) DeptList() revel.Result {
	db, err := sql.Open("mysql", "scott:tiger@tcp(127.0.0.1:8889)/wp32scott")//通常：ポート番号３３０６、＊manp:8889
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("データベース接続失敗")
		return c.Redirect(App.Error)
	}

	defer db.Close()
	deptList := dao.FindAll(db)
	if deptList == nil {
		fmt.Println("から")
	}
	return c.Render(deptList)

}

func (c App) ConfirmDeptDelete(deleteDeptDeptno string) revel.Result {
	db, err := sql.Open("mysql", "scott:tiger@tcp(127.0.0.1:8889)/wp32scott")//通常：ポート番号３３０６、＊manp:8889
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("データベース接続失敗")
		return c.Redirect(App.Error)
	}

	defer db.Close()
	deptList := dao.FindByPk(db, deleteDeptDeptno)
	if deptList == nil {
		fmt.Println("空")
	}
	deptno := deptList.Deptno
	dname := deptList.Dname
	loc := deptList.Loc
	return c.Render(deptno, dname, loc)
}

func (c App) DeptDelete(deleteDeptDeptno string) revel.Result {
	db, err := sql.Open("mysql", "scott:tiger@tcp(127.0.0.1:8889)/wp32scott")//通常：ポート番号３３０６、＊manp:8889
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("データベース接続失敗")
		return c.Redirect(App.Error)
	}
	defer db.Close()


	dao.Delete(db,deleteDeptDeptno)
	c.Flash.Success("部門番号：" + deleteDeptDeptno + "を削除しました。")
	return c.Redirect(App.DeptList)
}

func (c App) GoDeptAdd() revel.Result {
	return c.Render()
}

func (c App) DeptAdd(addDeptDeptno string, addDeptDname string, addDeptLoc string) revel.Result {
	db, err := sql.Open("mysql", "scott:tiger@tcp(127.0.0.1:8889)/wp32scott")//通常：ポート番号３３０６、＊manp:8889
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("データベース接続失敗")
		return c.Redirect(App.Error)
	}
	defer db.Close()
	c.Validation.Required(addDeptDeptno).Message("部門番号は必須です。")
	c.Validation.Match(addDeptDeptno, regexp.MustCompile("^[0-9]{2.}")).Message("部門番号は数字二桁で入力して下さい。")
	//c.Validation.MinSize(addDeptDeptno, 2).Message("部門番号は数字二桁で入力して下さい。")

	c.Validation.Required(addDeptDname).Message("部門名は必須です。")

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.GoDeptAdd)
	}

	dao.Insert(db, addDeptDeptno, addDeptDname, addDeptLoc)
	c.Flash.Success("部門番号：" + addDeptDeptno + "を追加しました。")

	return c.Redirect(App.DeptList)
}

func (c App) GoDeptEdit(editDeptDeptno string) revel.Result {
	db, err := sql.Open("mysql", "scott:tiger@tcp(127.0.0.1:8889)/wp32scott")//通常：ポート番号３３０６、＊manp:8889
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("データベース接続失敗")
		return c.Redirect(App.Error)
	}

	defer db.Close()
	deptList := dao.FindByPk(db, editDeptDeptno)
	if deptList == nil {
		fmt.Println("空")
	}
	deptno := deptList.Deptno
	dname := deptList.Dname
	loc := deptList.Loc
	return c.Render(deptno, dname, loc)
}

func (c App) DeptEdit(editDeptDeptno string, editDeptDname string, editDeptLoc string) revel.Result {
	db, err := sql.Open("mysql", "scott:tiger@tcp(127.0.0.1:8889)/wp32scott")//通常：ポート番号３３０６、＊manp:8889
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("データベース接続失敗")
		return c.Redirect(App.Error)
	}
	defer db.Close()
	c.Validation.Required(editDeptDname).Message("部門名は必須です。")

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()


		return c.Redirect("/App/GoDeptEdit?editDeptDeptno=" + editDeptDeptno)
	}


	dao.Update(db, editDeptDname, editDeptLoc, editDeptDeptno)
	c.Flash.Success("部門番号：" + editDeptDeptno + "を更新しました。")

	return c.Redirect(App.DeptList)
}



