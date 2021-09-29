package controllers

import (
	"database/sql"
	"fmt"
	"github.com/revel/revel"
	"log"
	"regexp"
	"revelTest/app/pkg/deptDAO"
	"revelTest/app/pkg/empDAO"
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

// -------------------------------------------------------------------------------------------------------------------------------------------
// ------------------------------------------------------------- dept ----------------------------------------------------------------------------
// -------------------------------------------------------------------------------------------------------------------------------------------

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
	deptList := deptDAO.FindAll(db)
	if deptList == nil {
		fmt.Println("から")
	}
	return c.Render(deptList)

}

// ----------------------------------------------------------------------------------------------------------------------------
// --------------------------------------------------- dept add ---------------------------------------------------------------------
// ----------------------------------------------------------------------------------------------------------------------

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

	c.Validation.Required(addDeptDname).Message("部門名は必須です。")

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.GoDeptAdd)
	}

	deptDAO.Insert(db, addDeptDeptno, addDeptDname, addDeptLoc)
	c.Flash.Success("部門番号：" + addDeptDeptno + "を追加しました。")

	return c.Redirect(App.DeptList)
}

// ------------------------------------------------------------------------------------------------------------
// --------------------------------- dept edit ---------------------------------------------------------
// --------------------------------------------------------------------------------------------------------------

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
	dept := deptDAO.FindByPk(db, editDeptDeptno)
	if dept == nil {
		fmt.Println("空")
	}
	deptno := dept.Deptno
	dname := dept.Dname
	loc := dept.Loc
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


	deptDAO.Update(db, editDeptDname, editDeptLoc, editDeptDeptno)
	c.Flash.Success("部門番号：" + editDeptDeptno + "を更新しました。")

	return c.Redirect(App.DeptList)
}

// ----------------------------------------------------------------------------------------------
// -------------------- dept delete ------------------------------------------------------------------
// ---------------------------------------------------------------------------------

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
	deptList := deptDAO.FindByPk(db, deleteDeptDeptno)
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


	deptDAO.Delete(db,deleteDeptDeptno)
	c.Flash.Success("部門番号：" + deleteDeptDeptno + "を削除しました。")
	return c.Redirect(App.DeptList)
}


// -------------------------------------------------------------------------------------------------------------------------------------------
// ------------------------------------------------------------- Emp ----------------------------------------------------------------------------
// -------------------------------------------------------------------------------------------------------------------------------------------

func (c App) EmpList() revel.Result {
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
	empList := empDAO.FindAll(db)
	if empList == nil {
		fmt.Println("から")
	}
	return c.Render(empList)

}

// ----------------------------------------------------------------------------------------------------------------------------------------------------
// ----------------------------------------------------------------- emp add --------------------------------------------------------------------------
// ------------------------------------------------------------------------------------------------------------------------------------------

func (c App) GoEmpAdd() revel.Result {
	var year [] int
	for i := 1980;i < 2023;i++ {
		year = append(year, i)
	}

	fmt.Println(year)

	var month [] int
	for i := 1;i < 13;i++ {
		month = append(month, i)
	}

	var day [] int
	for i := 1;i < 32;i++ {
		day = append(day, i)
	}

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
	deptList := deptDAO.FindAll(db)
	if deptList == nil {
		fmt.Println("から")
	}
	empList := empDAO.FindAll(db)
	if empList == nil {
		fmt.Println("から")
	}

	return c.Render(deptList, empList, year, month, day)
}

func (c App) EmpAdd(addEmpEmpno string, addEmpEname string, addEmpJob string, addEmpMgr string, addEmpHiredate string, addEmpSal string, addEmpComm string, addEmpDeptno string) revel.Result {
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
	c.Validation.Required(addEmpEmpno).Message("従業員番号は必須です。")
	c.Validation.Match(addEmpEmpno, regexp.MustCompile("[0-9]")).Message("従業員番号は数値で入力して下さい。")
	c.Validation.Match(addEmpEmpno, regexp.MustCompile("^[0-9]{4.}")).Message("従業員番号は数字4桁で入力して下さい。")

	c.Validation.Required(addEmpEname).Message("従業員名は必須です。")

	c.Validation.Required(addEmpJob).Message("役職は必須です。")





	c.Validation.Required(addEmpSal).Message("給与は必須です。")
	c.Validation.Match(addEmpSal, regexp.MustCompile("^[0-9,，.．]+")).Message("給与は数値で入力して下さい。")

	c.Validation.Required(addEmpComm).Message("歩合は必須です。")
	c.Validation.Match(addEmpComm, regexp.MustCompile("^[0-9,，.．]+")).Message("歩合は数値で入力して下さい。")

	c.Validation.Required(addEmpDeptno).Message("部門番号を選択してください")

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.GoEmpAdd)
	}

	empDAO.Insert(db, addEmpEmpno, addEmpEname, addEmpJob, addEmpMgr, addEmpHiredate, addEmpSal, addEmpComm, addEmpDeptno)
	c.Flash.Success("従業員番号：" + addEmpEmpno + "を追加しました。")

	return c.Redirect(App.DeptList)
}
// ---------------------------------------------------------------------------------------------------------------------------
// -------------------------------------------------- emp edit ------------------------------------------------------------
// ----------------------------------------------------------------------------------------------------------------------------

func (c App) GoEmpEdit(editEmpEmpno string) revel.Result {
	fmt.Println(editEmpEmpno)
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
	deptList := deptDAO.FindAll(db)
	if deptList == nil {
		fmt.Println("から")
	}
	empList := empDAO.FindAll(db)
	if deptList == nil {
		fmt.Println("から")
	}
	emp := empDAO.FindByPk(db, editEmpEmpno)
	if empList == nil {
		fmt.Println("空")
	}
	empno := emp.Empno
	ename := emp.Ename
	job := emp.Job
	mgr := emp.Mgr
	hiredate := emp.Hiredate
	sal := emp.Sal
	comm := emp.Comm
	deptno := emp.Deptno
	
	return c.Render(deptList, empList, empno, ename, job, mgr, hiredate, sal, comm, deptno)
}

// ---------------------------------------------------------------------------------------------------------------------------
// -------------------------------------------------- emp delete ------------------------------------------------------------
// ----------------------------------------------------------------------------------------------------------------------------

func (c App) ConfirmEmpDelete(deleteEmpEmpno string) revel.Result {
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

	empList := empDAO.FindByPk(db, deleteEmpEmpno)
	if empList == nil {
		fmt.Println("空")
	}
	empno := empList.Empno
	ename := empList.Ename
	job := empList.Job
	mgr := empList.Mgr
	hiredate := empList.Hiredate
	sal := empList.Sal
	comm := empList.Comm
	deptno := empList.Deptno
	
	return c.Render(empno, ename, job, mgr, hiredate, sal, comm, deptno)
}

func (c App) EmpDelete(deleteEmpEmpno string) revel.Result {
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


	empDAO.Delete(db,deleteEmpEmpno)
	c.Flash.Success("従業員番号：" + deleteEmpEmpno + "を削除しました。")
	return c.Redirect(App.EmpList)
}

