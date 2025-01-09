package main

import (
	uuid "github.com/satori/go.uuid"
	"github.com/xuri/excelize/v2"
	"log"
	"os"
)

func main() {
	fin, err := excelize.OpenFile("pkg/school.xlsx")
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer fin.Close()
	rows, err := fin.GetRows("sheet1")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fout, err := os.OpenFile("pkg/sql/init.sql", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	defer fout.Close()
	for _, row := range rows {
		if len(row) >= 6 {
			id := uuid.NewV4().String()
			sql := "insert into school(id,name,code,address,create_time,update_time) values (" + id + ",'" + row[0] + "','" + row[1] + "','" + row[2] + "',now(),now());\n"
			_, err := fout.WriteString(sql)
			if err != nil {
				log.Fatalf(err.Error())
			}
		}
	}

}
