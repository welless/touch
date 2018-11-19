package main

import (
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  "time"
  "log"
)

func main() {
  db, err := sql.Open("mysql", upaswd + "@tcp(gz-cdb-nvmjxb1y.sql.tencentcdb.com:62823)/information_schema?charset=utf8")
  if err != nil {
    panic(err.Error())
  }
  defer db.Close()

  db.SetMaxOpenConns(10)
  db.SetMaxIdleConns(5)   
  db.SetConnMaxLifetime(8*60*time.Second)   

  err = db.Ping()
  if err != nil {
    panic(err.Error())
  }

  rows, err := db.Query("select substring_index(host,':',1) as ip , count(*) from information_schema.processlist group by ip");
  if err != nil {
    panic(err.Error())
  }
  
  defer rows.Close()
  var ip string
  var count int
  for rows.Next() {
    err := rows.Scan(&ip, &count)
    if err != nil {
      panic(err.Error())
    }
    
    if ip != "" {
      log.Println(ip, count)
    }
  }
  
  err = rows.Err()
  if err != nil {
    panic(err.Error())
  }
}

