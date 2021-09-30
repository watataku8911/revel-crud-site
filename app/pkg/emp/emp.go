package emp

type Emp struct {
	Empno int
  	Ename string
	Job string
	Mgr **int
	Hiredate string
	Sal string
	Comm **string
	Deptno int
}

type EmpList []Emp