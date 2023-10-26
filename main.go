db, err := gorm.Open("mysql", "user:password@/database?charset=utf8mb4&parseTime=True&loc=Local")
if err != nil {
    log.Fatal(err)
}

defer db.Close()
