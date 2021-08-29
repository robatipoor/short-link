package model

import (
	"log"
	"time"

	"github.com/gocql/gocql"
	"github.com/robatipoor/short-link/services/key-gen/config"
)

type Key struct {
	Value      string    `cql:"value"`
	CreateDate time.Time `cql:"create_date"`
}

func NewKey(value string) Key {
	return Key{
		Value:      value,
		CreateDate: time.Now(),
	}
}

func InsertUnusedKey(key *Key) error {
	q := `
		INSERT INTO unused_keys (
		    value,
			create_date
		)
		VALUES (?, ?) IF NOT EXISTS;
    	`
	err := config.SessionDB.Query(q,
		key.Value,
		key.CreateDate).Exec()
	if err != nil {
		log.Printf("ERROR: InsertUnusedKey , %s", err.Error())
	}
	return err
}

func InsertUsedKey(key *Key, exp int) error {
	query := `
		INSERT INTO used_keys (
		    value,
			create_date
		)
		VALUES (?, ?) IF NOT EXISTS USING TTL ? ;
    	`
	err := config.SessionDB.Query(query,
		key.Value,
		key.CreateDate,
		exp,
	).Exec()
	if err != nil {
		log.Printf("ERROR: InsertUsedKey , %s", err.Error())
	}
	return err

}

func ReInsertUsedKey(key *Key, exp int) error {
	query := `
		INSERT INTO used_keys (
		    value,
			create_date
		)
		VALUES (?, ?) USING TTL ?;
    	`
	err := config.SessionDB.Query(query,
		key.Value,
		key.CreateDate,
		exp,
	).Exec()
	if err != nil {
		log.Printf("ERROR: ReInsertUsedKey , %s", err.Error())
	}
	return err

}

func FindListUnUsedKey(n int) []*Key {
	q := `SELECT * FROM unused_keys LIMIT ?`
	m := map[string]interface{}{}
	itr := config.SessionDB.Query(q, n).Consistency(gocql.One).Iter()
	var keys []*Key
	for itr.MapScan(m) {
		key := &Key{}
		key.Value = m["value"].(string)
		key.CreateDate = m["create_date"].(time.Time)
		keys = append(keys, key)
	}
	return keys
}

func FindListUsedKeyByValue(key string) []*Key {
	q := `SELECT * FROM used_keys WHERE value = ?`
	m := map[string]interface{}{}
	itr := config.SessionDB.Query(q, key).Consistency(gocql.All).Iter()
	var keys []*Key
	for itr.MapScan(m) {
		key := &Key{}
		key.Value = m["value"].(string)
		key.CreateDate = m["create_date"].(time.Time)
		keys = append(keys, key)
	}
	return keys
}

func FindListUsedKey(n int) []*Key {
	q := `SELECT * FROM used_keys LIMIT ?`
	m := map[string]interface{}{}
	itr := config.SessionDB.Query(q, n).Consistency(gocql.All).Iter()
	var keys []*Key
	for itr.MapScan(m) {
		key := &Key{}
		key.Value = m["value"].(string)
		key.CreateDate = m["create_date"].(time.Time)
		keys = append(keys, key)
	}
	return keys
}

func DeleteUnUsedKey(key *Key) error {
	q := `DELETE FROM unused_keys WHERE value = ?`
	err := config.SessionDB.Query(q, key.Value).Exec()
	if err != nil {
		log.Printf("ERROR: fail delete key, %s", err.Error())
		return err
	}

	return nil
}

func DeleteUsedKey(key *Key) error {
	q := `DELETE FROM used_keys WHERE value = ?`
	err := config.SessionDB.Query(q, key.Value).Exec()
	if err != nil {
		log.Printf("ERROR: fail delete key, %s", err.Error())
		return err
	}

	return nil
}
