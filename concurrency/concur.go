package main

import (
	"fmt"
	"sync"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)

const numGoroutines = 4

func sqlQuery(wg *sync.WaitGroup) {
	defer wg.Done()
	username := "root"
	password := ""
	dbName := "lab_bookings"
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s",username,password,dbName)

	db,err := sql.Open("mysql",dsn)

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	insert, err := db.Query("UPDATE students SET Amount_Due = Amount_Due +1 WHERE Roll_Number = '1' ;")

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

}
func main() {
	fmt.Println("Go SQL Tutorial")
	
	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	for i := 0; i<numGoroutines; i++ {
		go sqlQuery(&wg)
	}

	wg.Wait()

	fmt.Println("All Tasks Completed")
}
