package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/zceuiv/gorest/dba"
	"net/http"
	"strings"
)

type CategoryObj map[string]string

func CategoryListHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		PostHandler(w, r)
	case "GET":
		GetHandler(w, r)

	}
}

func CategoryHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		GetHandler(w, r)
	case "PUT":
		PutHandler(w, r)
	case "DELETE":
		DeleteHandler(w, r)
	}
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	categoryName := vars["category"]
	//w.Write([]byte("List: "))
	//w.Write([]byte(categoryName))
	//w.Write([]byte("\n"))
	query := "INSERT INTO " + categoryName
	r.ParseForm()
	colAdd := ""
	valAdd := ""
	for k, v := range r.Form {
		colAdd += fmt.Sprintf("%s,", k)
		valAdd += fmt.Sprintf("'%s',", strings.Replace(v[0], "'", "''", -1))
	}
	colAdd = strings.TrimRight(colAdd, ",")
	valAdd = strings.TrimRight(valAdd, ",")
	query = fmt.Sprintf("%s(%s) VALUES(%s) RETURNING id", query, colAdd, valAdd)
	//fmt.Println(query)
	id := dba.CreateOrReplace(query)
	query2 := fmt.Sprintf("SELECT * FROM %s WHERE id=%d", categoryName, id)
	fmt.Println(query2)
	cols, result := dba.Query(query2)
	//fmt.Println(cols)
	//fmt.Println(result)
	jsonRet := JsonResult(cols, result)
	w.Write(jsonRet)
	w.Write([]byte("\n"))
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	categoryName := vars["category"]
	//w.Write([]byte("List: "))
	//w.Write([]byte(categoryName))
	//w.Write([]byte("\n"))
	id := vars["id"]
	query := "SELECT * FROM " + categoryName + " WHERE 1=1 "
	r.ParseForm()
	queryAdd := ""
	if id != "0" && id != "" {
		queryAdd += " AND id = " + id
	}
	for k, v := range r.Form {
		queryAdd += fmt.Sprintf(" AND %s = '%s'", k, strings.Replace(v[0], "'", "''", -1))
	}
	query += queryAdd
	//fmt.Println(query)
	cols, result := dba.Query(query)
	//fmt.Println(cols)
	//fmt.Println(result)
	jsonRet := JsonResult(cols, result)
	w.Write(jsonRet)
	w.Write([]byte("\n"))
}

func PutHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	categoryName := vars["category"]
	id := vars["id"]
	//w.Write([]byte("List: "))
	//w.Write([]byte(categoryName))
	//w.Write([]byte("\n"))
	r.ParseForm()
	setStr := ""
	for k, v := range r.Form {
		setStr += fmt.Sprintf(" %s = '%s',", k, strings.Replace(v[0], "'", "''", -1))
	}
	setStr = strings.TrimRight(setStr, ",")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = %s", categoryName, setStr, id)
	dba.CreateOrReplace(query)
	//fmt.Println(query)
	query2 := fmt.Sprintf("SELECT * FROM %s WHERE id=%s", categoryName, id)
	fmt.Println(query2)
	cols, result := dba.Query(query2)
	//fmt.Println(cols)
	//fmt.Println(result)
	jsonRet := JsonResult(cols, result)
	w.Write(jsonRet)
	w.Write([]byte("\n"))
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	categoryName := vars["category"]
	id := vars["id"]

	query := fmt.Sprintf("DELETE FROM %s WHERE id=%s", categoryName, id)
	dba.Delete(query)
	w.Write([]byte("OK\n"))
}

func JsonResult(cols []string, result [][]string) (ret []byte) {
	cats := make([]CategoryObj, 0)
	for i := 0; i < len(result); i++ {
		cat := make(CategoryObj)
		line := result[i]
		for j := 0; j < len(cols); j++ {
			cat[cols[j]] = line[j]
		}
		cats = append(cats, cat)
		fmt.Println("catslength", len(cats))
	}
	ret, _ = json.Marshal(cats)
	return ret
}
