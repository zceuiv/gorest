package main

import (
	//"fmt"
	"github.com/gorilla/mux"
	"github.com/zceuiv/gorest/configure"
	"github.com/zceuiv/gorest/godbc"
	"github.com/zceuiv/gorest/handler"
	"net/http"
	"os"
)

/*
 * 初始化pg.Options和http.Server
 **/
type Config interface {
	Server() (server *http.Server)
	DbOptions() (dbOpt *godbc.Options)
}

var (
	// 读取服务器配置和数据库配置的接口
	__Config Config

	// http.Server实例
	__Server *http.Server
)

func init() {
	__Config = configure.SingletonConfig()
	if __Config == nil {
		os.Exit(1)
	}
	__Server = __Config.Server()
}

func main() {

	categoryListRouter := mux.NewRouter()
	categoryListRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome!\n"))
	})
	categoryListRouter.HandleFunc("/{category}", handler.CategoryListHandler).
		Methods("GET", "POST")
	categoryRouter := categoryListRouter.PathPrefix("/{category}").Subrouter()
	categoryRouter.HandleFunc("/{id:[0-9]+}", handler.CategoryHandler).
		Methods("GET", "PUT", "DELETE")

	http.ListenAndServe(__Server.Addr, categoryListRouter)
	http.ListenAndServe(__Server.Addr, categoryRouter)
}
