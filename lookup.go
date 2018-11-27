package main

import (
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  "time"
  "log"
  "os"
)

func main() {
  if len(os.Args) > 2 {
    return
  }
  
  var mode string = "client"
  if len(os.Args) > 1 {
    mode = os.Args[1]
  }
  
  db, err := sql.Open("mysql", "user:123456@tcp(gz-cdb-nvmjxb1y.sql.tencentcdb.com:62823)/test?charset=utf8")
  if err != nil {
    panic(err.Error())
  }
  defer db.Close()

  db.SetMaxOpenConns(3)
  db.SetMaxIdleConns(2)
  db.SetConnMaxLifetime(60*time.Second)

  err = db.Ping()
  if err != nil {
    panic(err.Error())
  }
  
  if mode == "server" {
    stmtIns, err := db.Query("insert into test.lookup(ip) SELECT SUBSTRING_INDEX(USER(), '@', -1) AS ip")
    if err != nil {
      panic(err.Error())
    }
    
    defer stmtIns.Close()
    
    stmtDel, err := db.Query("DELETE FROM a USING test.lookup a INNER JOIN test.lookup b WHERE a.ip = b.ip AND a.id < b.id OR a.createTime < NOW() - INTERVAL 12 hour")
    if err != nil {
      panic(err.Error())
    }
    
    defer stmtDel.Close()
  }
  
  rows, err := db.Query("select ip, max(createTime) ct from test.lookup group by ip order by createTime desc")
  if err != nil {
    panic(err.Error())
  }
  
  var ip string
  var ct string
  for rows.Next() {
    err := rows.Scan(&ip, &ct)
    if err != nil {
      panic(err.Error())
    }
    
    if ip != "" {
      log.Println(ct, ip)
    }
  }
  
  err = rows.Err()
  if err != nil {
    panic(err.Error())
  }
}
