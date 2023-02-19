package main



var db, _ := gorm.Open("mysql", "root:root@/todolist?charset=utf8&parseTime=True&loc=Local")

+type TodoItemModel struct{
+	 Id int `gorm:"primary_key"`
+	 Description string
+	 Completed bool
+}

...

func main() {
     defer db.Close()
+	 db.Debug().DropTableIfExists(&TodoItemModel{})
+	 db.Debug().AutoMigrate(&TodoItemModel{})
+

	log.Info("Starting Todolist API server")
	router := mux.NewRouter()
	router.HandleFunc("/healthz", Healthz).Methods("GET")
	http.ListenAndServe(":8000", router)
}