package howdah

import (
	//"github.com/troyk/sqlx"
	"database/sql"
	"github.com/kisielk/sqlstruct"
	_ "github.com/lib/pq"
	"log"
	"time"
)

type Model struct {
	Id int `sql:"id"`
}

//type DB struct {
//sqlx.DB
//}

type User struct {
	Model
	Foo   int       `sql:"foo"`
	Name  string    `sql:"name"`
	Email string    `sql:"email"`
	Cr    time.Time `sql:"cr"`
}

//func Connect() (*sqlx.DB, error) {
//db, err := sqlx.Connect("postgres://troy@localhost/blit?sslmode=disable")
//return db, err
//}

func Users(db *sql.DB) []*User {
	var users []*User
	rows, err := db.Query("SELECT id,email,name,null as foo, created_at::timestamptz FROM users limit 5")
	if err != nil {
		//log.Println(err)
		return nil
	}
	defer rows.Close()
	for rows.Next() {
		var user User
		var id int
		var email string
		var name string
		var foo string
		var cr time.Time
		rows.Scan(&id, &email, &name, &foo, &cr)
		sqlstruct.Scan(&user, rows)
		//rows.Scan(user.Id)
		log.Println(cr)
		log.Println(id)
		users = append(users, &user)
	}
	//log.Println(users)
	return users
}

//func (db *DB) Query(query string, args ...interface{}) (*sql.Rows, error) {
//log.Printf("sql : %s\nargs: ", query)
//log.Println(args)
//return db.DB.Query(query, args)
//}

//func Troy() string {
//return "pecker"
/*}*/
