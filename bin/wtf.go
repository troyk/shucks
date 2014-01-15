package main

import (
	"database/sql"
	"fmt"
	"github.com/kisielk/sqlstruct"
	_ "github.com/lib/pq"
	"log"
	"strings"
	"time"
)

type DateTime time.Time

func (d *DateTime) Scan(src interface{}) error {
	bytes, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("Expected []byte, got %T", src)
	}
	log.Println(bytes)
	log.Println(src)
	return nil
}

type EmailAddress string

func (g *EmailAddress) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	bytes, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("Expected []byte, got %T", src)
	}
	str := string(bytes)
	//log.Printf("%s is %d len", str, len(str))
	switch len(str) {
	case 10:
		str = fmt.Sprintf("(%s)%s-%s", str[0:3], str[3:6], str[6:10])
	}
	*g = EmailAddress(strings.ToLower(str))
	return nil
}

type User struct {
	Id           int
	CreatedAt    time.Time    `sql:"created_at"`
	MobilePhone  EmailAddress `sql:"mobile_phone"`
	Email        string
	StripeCardId string `sql:"stripe_card_id"`
}

func List(db *sql.DB) {
	rows, _ := db.Query("SELECT id, mobile_phone, email, created_at FROM users where id>0 order by id limit 6")

	for rows.Next() {
		var t User
		err := sqlstruct.Scan(&t, rows)
		if err != nil {
			log.Printf("ERROR %v\n", err)
		} else {
			log.Printf("%+v\n", t)
		}
	}
}

type StructChanges map[string]interface{}

func Changes(l interface{}, r interface{}) StructChanges {
	ch := make(StructChanges)
	v := reflect.ValueOf(s)
	fields := getFieldInfo(v.Type())
	for f := range fields {
		//TODO how to set struct values???
	}

}

func GetUser(db *sql.DB, id int) User {
	var u User
	rows, _ := db.Query("SELECT * FROM users WHERE id=$1 LIMIT 1", id)
	rows.Next()
	sqlstruct.Scan(&u, rows)
	return u
}

func main() {
	db, _ := sql.Open("postgres", "host=localhost dbname=blit sslmode=disable")
	defer db.Close()
	List(db)
	u1 := GetUser(db, 3)
	u2 := GetUser(db, 3)
	u2.Email = "tkruthoff@gmail.com"
	u2.MobilePhone = "9167496969"
	log.Println(u1)
}
