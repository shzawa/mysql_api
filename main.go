package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	empdao "./pkg/dao/emp"
)

func main() {
	r := mux.NewRouter()

	// localhost:8080/emp/で従業員の一覧を取得できるようにする
	r.HandleFunc("/emp/", showEmpIndex)
	log.Fatal(http.ListenAndServe(":8080", r))

}

func showEmpIndex(w http.ResponseWriter, r *http.Request) {
	emp := empdao.FetchIndex()

	// json形式に変換する
	bytes, err := json.Marshal(emp)
	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte(string(bytes)))
}
