package model

// import (
// 	"database/sql"
// 	"log"

// 	_ "github.com/go-sql-driver/mysql"
// )

// type Pic struct {
// 	Id   int    `json:"id"`
// 	Name string `json:"name" form:"name"`
// 	Date []byte `json:"date" form:"date"`
// }

// var db *sql.DB

// //Select One
// func (p Pic) get() (pic Pic, err error) {
// 	row := db.QueryRow("SELECT id, name, date FROM pic WHERE id =?", p.Id)
// 	err = row.Scan(&p.Id, &p.Name, &p.Date)
// 	if err != nil {
// 		log.Printf("Select Error : %s", err.Error())
// 		return
// 	}
// 	return
// }

// //Select All
// func (p Pic) getAll() (pics []Pic, err error) {
// 	rows, err := db.Query("SELECT id, name, date FROM pic")
// 	if err != nil {
// 		return
// 	}
// 	for rows.Next() {
// 		var pic Pic
// 		rows.Scan(&pic.Id, &pic.Name, &pic.Date)
// 		pics = append(pics, pic)
// 	}
// 	defer rows.Close()
// 	return
// }

// //Insert
// // func (pic *Pic) Creat() int {
// // 	rs, err := db.QueryRow("INSERT init user(name,date) value(?,?)", pic.Name, pic.Date)
// // 	if err != nil {
// // 		log.Printf("Insert Error : %s", err)
// // 	}
// // 	id, err := rs.LastInsertId()
// // 	if err != nil {
// // 		log.Printf("LastId Error : %s", err)
// // 	}
// // 	return id
// // }
