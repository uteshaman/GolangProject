package main

import (
	"io"
	"net/http"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
+	 _ "github.com/go-sql-driver/mysql"
+	 "github.com/jinzhu/gorm"
+	 _ "github.com/jinzhu/gorm/dialects/mysql"
)

+var db, _ := gorm.Open("mysql", "root:root@/todolist?charset=utf8&parseTime=True&loc=Local")

...

func main() {
	+ defer db.Close()
	log.Info("Starting Todolist API server")
	router := mux.NewRouter()
	router.HandleFunc("/healthz", Healthz).Methods("GET")
	http.ListenAndServe(":8000", router)
}