package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "modernc.org/sqlite"
)

// ===================
type GPU struct {
	Name string
	VRAM int
	MHz  float32
	vmhz float32
}

//===================

func main() {
	database, err := sql.Open("sqlite", "./database.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	//===================
	stmt, _ := database.Prepare("CREATE TABLE IF NOT EXISTS gpus (id INTEGER PRIMARY KEY, gpuname TEXT, vram TEXT, MHz TEXT, vmhz TEXT)")
	stmt.Exec()

	stmt, _ = database.Prepare("INSERT INTO gpus(gpuname,vram,MHz,vmhz) VALUES (?,?)")
	//===================
	for {
		fmt.Print("Welcome to GPU finder,type in GPU`s name to find: \n")
		fmt.Print("or type 0 to cose: \n")
		//===================
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			addTask(database)
		case 2:
			showTask(database)
		case 3:
			deleteTask(database)
		case 0:
			os.Exit(0)
		default:
			fmt.Println("error")
		}
	}
}
