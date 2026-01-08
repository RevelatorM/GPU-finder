package main

import (
	"database/sql" //this library to add sqlite
	"fmt"
	"os"

	"bufio"
	"strings"

	_ "modernc.org/sqlite" //this command to add sqlite - "go get modernc.org/sqlite"
)

func main() {
	database, err := sql.Open("sqlite", "./database.db") //creating db
	if err != nil {
		fmt.Println(err)
		return
	}
	//===================
	stmt, _ := database.Prepare("CREATE TABLE IF NOT EXISTS gpus (id INTEGER PRIMARY KEY, gpuname TEXT, vram TEXT, MHz TEXT, vmhz TEXT)")
	stmt.Exec() //creating table
	//filling db
	stmt, _ = database.Prepare("INSERT INTO gpus(gpuname,vram,MHz,vmhz) VALUES (?,?,?,?)") //SQL command to fill
	stmt.Exec("RTX 3060", "12GB", "1867 MHz chip", "15000 MHz VRAM")
	stmt.Exec("RTX 3070", "8GB", "1815 MHz chip", "14000 MHz VRAM")
	stmt.Exec("RTX 3080", "10GB", "1710 MHz chip", "19000 MHz VRAM")
	stmt.Exec("RTX 4060", "8GB", "2535 MHz chip", "17000 MHz VRAM")
	stmt.Exec("RTX 4070", "12GB", "2550 MHz chip", "21000 MHz VRAM")
	stmt.Exec("RX 7600", "8GB", "2755 MHz chip", "18000 MHz VRAM")
	stmt.Exec("RX 9070 XT", "16GB", "3030 MHz chip", "20000 MHz VRAM")
	stmt.Exec("Arc B570", "10GB", "2660 MHz chip", "19000 MHz VRAM")
	//===================
	for {
		fmt.Print("Welcome to GPU finder,type in GPU`s name to find: \n")
		//===================
		var choice string
		fmt.Print("or type 0 to close: \n")
		reader := bufio.NewReader(os.Stdin)
		choice, _ = reader.ReadString('\n')
		choice = strings.TrimSpace(choice)
		switch choice {
		case "0":
			fmt.Println("Exit\n")
			os.Exit(0)

		default:
			query := "SELECT * FROM gpus WHERE gpuname LIKE ?" //SQL request - "SELECT * FROM gpus WHERE name LIKE ?"
			rows, err := database.Query(query, "%"+choice+"%") //rows variabale equals database,query is being taken from database variable
			if err != nil {                                    //"%"+choice+"%" means LIKE from SQL "%" "%" = {}
				fmt.Println(err)
				return
			}
			defer rows.Close()
			//===================
			for rows.Next() {
				//rows.Next() - moves each string and searches
				//after we got rows from here - query := "SELECT * FROM gpus WHERE gpuname LIKE ?"
				//rows, err := database.Query(query, "%"+choice+"%")
				var id int
				var gpuname string
				var vram string
				var mhz string
				var vmhz string
				err := rows.Scan(&id, &gpuname, &vram, &mhz, &vmhz)
				if err != nil {
					fmt.Println("Scan error:", err)
					return
				}
				fmt.Printf("Name: %s  VRAM: %s  MHz: %s  VMHz: %s\n", gpuname, vram, mhz, vmhz) //%s for strings
			}
			//===================
			if err = rows.Err(); err != nil {
				fmt.Println("Rows error:", err)
			}
			//===================
		}

	}
}
