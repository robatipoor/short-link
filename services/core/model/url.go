package model

import (
	"errors"
	"log"
	"time"

	"github.com/gocql/gocql"
	"github.com/robatipoor/short-link/services/core/config"
)

type Url struct {
	OriginalUrl string    `cql:"original_url"`
	ShortUrl    string    `cql:"short_url"`
	CreateDate  time.Time `cql:"create_date"`
	// UserId      gocql.UUID
}

func InsertUrl(url *Url, exp int) error {
	log.Println(url)
	q := `
		INSERT INTO urls (
			original_url,
			short_url,
		    create_date
		)
		VALUES (?, ? ,?) IF NOT EXISTS USING TTL ?;
   	`
	err := config.SessionDB.Query(q,
		url.OriginalUrl,
		url.ShortUrl,
		url.CreateDate,
		exp,
	).Exec()
	log.Println(url)
	if err != nil {
		log.Printf("ERROR: fail create urls , %s", err.Error())
	}
	return err
}

func FindUrlByShortUrl(url string) (*Url, error) {
	q := `SELECT * FROM urls WHERE short_url = ? ;`
	m := map[string]interface{}{}
	iter := config.SessionDB.Query(q, url).Consistency(gocql.One).Iter()
	if iter.MapScan(m) {
		u := Url{}
		u.ShortUrl = url
		u.OriginalUrl = m["original_url"].(string)
		u.CreateDate = m["create_date"].(time.Time)
		return &u, nil
	}
	return nil, errors.New("find url by short url failed !!!")
}

func DeleteUrlByShortUrl(url string) error {
	q := ` DELETE FROM urls WHERE short_url = ?;`
	err := config.SessionDB.Query(q, url).Exec()
	if err != nil {
		log.Printf("ERROR: fail delete urls , %s", err.Error())
	}
	return err
}
