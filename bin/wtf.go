package main

import (
	"database/sql"
	"fmt"
	"github.com/kisielk/sqlstruct"
	_ "github.com/lib/pq"
	"log"
	"strings"
)

type EmailAddress string

func (g *EmailAddress) Scan(src interface{}) error {
	bytes, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("Expected []byte, got %T", src)
	}
	str := string(bytes)
	log.Printf("%s is %d len", str, len(str))
	switch len(str) {
	case 10:
		str = fmt.Sprintf("(%s)%s-%s", str[0:3], str[3:6], str[6:10])
	}
	*g = EmailAddress(strings.ToLower(str))
	return nil
}

type User struct {
	Id           int
	MobilePhone  EmailAddress `sql:"mobile_phone"`
	Email        string
	StripeCardId string `sql:"stripe_card_id"`
}

func main() {
	db, _ := sql.Open("postgres", "host=localhost dbname=blit sslmode=disable")
	defer db.Close()

	rows, _ := db.Query("SELECT id, mobile_phone, COALESCE(email,'') as email FROM users limit 5")

	for rows.Next() {
		var t User
		_ = sqlstruct.Scan(&t, rows)
		log.Printf("%+v\n", t)
	}
}
