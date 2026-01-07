package main

import (
	"database/sql"
	"fmt"

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
	stmt, _ := database.Prepare("CREATE IF NOT EXISTS gpus (id INTEGER PRIMARY KEY, gpuname TEXT, vram INTEGER, MHz TEXT, vmhz TEXT)")
	//===================
	for {
		fmt.Print("Welcome to GPU finder,type in GPU`s name to find: \n")

	}
}
