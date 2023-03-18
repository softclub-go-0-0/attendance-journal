package main

import (
	"fmt"
	"github.com/softclub-go-0-0/attendance-journal/pkg/exporter"
	"github.com/softclub-go-0-0/database"
	"log"
)

func main() {
	fmt.Println("This is the Go Programming Language Course's students attendance journal.")
	for {
		fmt.Println("Enter a command you want to run:")
		fmt.Println("0 - close the journal")
		fmt.Println("1 - complete student attendance journal")
		fmt.Println("2 - export attendance journal data to a file")
		var command int
		fmt.Print("Your command: ")
		fmt.Scan(&command)
		switch command {

		case 0:
			fmt.Println("Bye!")
			return

		case 1:
			fmt.Println("\nFilling up attendance...")
		enterDayNumber:
			fmt.Print("Enter day number (13-18): ")
			var dayNumber int
			fmt.Scan(&dayNumber)
			if dayNumber < 13 || dayNumber > 18 {
				fmt.Println("Wrong day number! Enter a number between 13 and 18.")
				goto enterDayNumber
			}
			fmt.Println("Attendance of students on day", dayNumber, "\n")
			for i := 0; i < len(database.Students); i++ {
				fmt.Printf("Is %s %s in lass? (1/0): ", database.Students[i].Name, database.Students[i].Surname)
				var presence int
				fmt.Scan(&presence)
				database.Students[i].Attendance[dayNumber] = presence
			}
			fmt.Println("\nAttendance of", dayNumber, "day accepted.\n")

		case 2:
			fmt.Println("\nExporting students attendance to a file:")
			fmt.Print("Enter file name: ")
			var filename string
			fmt.Scan(&filename)
			numberOfStudents, err := exporter.ExportAttendanceToFile(filename)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("\nExported", numberOfStudents, "students' attendance to", filename, "file.\n")
		default:
			fmt.Println("Wrong command!")
		}
	}
}
