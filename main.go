package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "modernc.org/sqlite"
)

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
		//===================
		var choice string
		fmt.Print("or type 0 to close: \n")
		fmt.Scanln(&choice)
		switch choice {
		case "0":
			fmt.Println("Exit\n")
			os.Exit(0)

		default:
			query := "SELECT * FROM gpus WHERE name LIKE ?"    //SQL request - "SELECT * FROM gpus WHERE name LIKE ?"
			rows, err := database.Query(query, "%"+choice+"%") //rows variabale equals database,query is being taken from database variable
			if err != nil {                                    //"%"+choice+"%" means LIKE from SQL "%" "%" = {}
				fmt.Println(err)
				return
			}
			defer rows.Close()

		}

	}
}
