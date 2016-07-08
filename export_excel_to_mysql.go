package main

import (
        "database/sql"
        _ "github.com/go-sql-driver/mysql"
        "log"
        "github.com/tealeg/xlsx"
)

func main() {
        db, err := sql.Open("mysql","username:passwd@tcp(hostname:port)/dbname?charset=utf8")
        if err != nil {
                log.Fatalf("Open database error: %s\n", err)
        }
        defer db.Close()

        err = db.Ping()
        if err != nil {
                log.Fatal(err)
        }

        rows, err := db.Query("select * from tablename limit 1000000")
        if err != nil {
                log.Println(err)
        }
        defer rows.Close()

        var file *xlsx.File
        var sheet *xlsx.Sheet
        var row *xlsx.Row
        var cell *xlsx.Cell
        //var err error

        file = xlsx.NewFile()
        sheet, err = file.AddSheet("Sheet1")
        if err != nil {
                log.Printf(err.Error())
        } 
        cols, _ := rows.Columns()
        row = sheet.AddRow()
        for i:=0;i<len(cols);i++ {
                cell = row.AddCell()
                cell.Value = cols[i]
        }   
        buff := make([]interface{}, len(cols)) 
        data := make([]string, len(cols))  
        for i, _ := range buff {
             buff[i] = &data[i]  
        }
        for rows.Next() {
                row = sheet.AddRow()
                rows.Scan(buff...)
                
                for _, col := range data {
                   cell = row.AddCell()
                   cell.Value = col
                }
                if err != nil {
                        log.Fatal(err)
                }


        }

        err = rows.Err()
        if err != nil {
                log.Fatal(err)
        }

        err = file.Save("out.xlsx")
        if err != nil {
                log.Printf(err.Error())
        }
}