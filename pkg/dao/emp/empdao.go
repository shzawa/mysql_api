package empdao

import (
	"github.com/tk-ozawa/mysql_api/pkg/db"
)

// Emp型の構造体を用意
// あとでjson形式にする為、jsonのタグを予めつけておく。
type Emp struct {
	ID         int    `json:"id"`
	EmNo       int    `json"em_no"`
	EmName     string `json"em_name"`
	EmJob      string `json"em_job"`
	EmMgr      int    `json"em_mgr"`
	EmHiredate string `json"em_hiredate"`
	DeptID     int    `json"dept_id"`
}

// empsを全行取得する関数
func FetchIndex() []Emp {
	db := db.Connect()
	defer db.Close()

	// rowを取得
	rows, err := db.Query("SELECT * FROM emps")
	if err != nil {
		panic(err.Error())
	}

	// emp型のスライスに格納する
	empArgs := make([]Emp, 0) // empArgsの初期化
	for rows.Next() {
		var emp Emp
		err = rows.Scan(&emp.ID, &emp.EmNo, &emp.EmName, &emp.EmJob, &emp.EmMgr, &emp.EmHiredate, &emp.DeptID)
		if err != nil {
			panic(err.Error())
		}
		empArgs = append(empArgs, emp) // empArgsの最後尾の要素にempを追加
	}
	return empArgs
}
