package main

import (
	"io"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

var data, _ = gorm.Open("mysql", "root:root@/todo?charset=utf8&parseTime=True&loc=Local")

type TodoItemModel struct {
	Id   int `gorm:"primary_key"`
	Info string
	Done bool
}

func t_alive(w http.ResponseWriter, r *http.Request) {
	log.Info("test is okey")
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"test": true}`)
}

func init() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetReportCaller(true)
}

func main() {
	defer data.Close()

	data.Debug().AutoMigrate(&TodoItemModel{})
	data.Debug().DropTableIfExists(&TodoItemModel{})

	log.Info("run")
	rout := mux.NewRouter()
	rout.HandleFunc("/t_alive", t_alive).Methods("GET")
	http.ListenAndServe(":8000", rout)
}
