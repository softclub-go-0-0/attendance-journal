package exporter

import (
	"fmt"
	"github.com/softclub-go-0-0/database"
	"os"
)

func ExportAttendanceToFile(filename string) (totalNumber int, err error) {
	file, err := os.OpenFile(filename, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return
	}

	for _, student := range database.Students {
		studentData := fmt.Sprintf("%d. %s %s - %d:%d, %d:%d, %d:%d, %d:%d, %d:%d, %d:%d\n",
			student.ID,
			student.Name,
			student.Surname,
			13,
			student.Attendance[13],
			14,
			student.Attendance[14],
			15,
			student.Attendance[15],
			16,
			student.Attendance[16],
			17,
			student.Attendance[17],
			18,
			student.Attendance[18],
		)
		_, err = file.WriteString(studentData)
		if err != nil {
			return
		}
	}

	return len(database.Students), err
}
