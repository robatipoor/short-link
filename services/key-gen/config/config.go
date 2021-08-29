package config

import (
	"log"
	"strconv"

	"github.com/gocql/gocql"
	"github.com/robatipoor/short-link/pkg/configs"
	"github.com/robatipoor/short-link/pkg/utils"
)

const (
	KeyPoolSize              = 10
	KeyLen                   = 6
	KeyJobScheduleTimeSecond = 10
	ServerPort               = "8081"
	ServerAddr               = "127.0.0.1"
)

var SessionDB *gocql.Session

var cassandraConfig = configs.CassandraConfig{
	Host:        utils.GetEnv("CASSANDRA_HOST", "127.0.0.1"),
	Port:        utils.GetEnv("CASSANDRA_PORT", "9042"),
	Keyspace:    utils.GetEnv("CASSANDRA_KEYSPACE", "key_gen_space"),
	Consistancy: utils.GetEnv("CASSANDRA_CONSISTANCY", "LOCAL_QUORUM"),
}

func init() {
	initSessionDB()
}

func initSessionDB() {
	port := func(p string) int {
		i, err := strconv.Atoi(p)
		if err != nil {
			return 9042
		}
		return i
	}

	consistancy := func(c string) gocql.Consistency {
		gc, err := gocql.MustParseConsistency(c)
		if err != nil {
			return gocql.All
		}
		return gc
	}

	cluster := gocql.NewCluster(cassandraConfig.Host)
	cluster.Port = port(cassandraConfig.Port)
	cluster.Keyspace = cassandraConfig.Keyspace
	cluster.Consistency = consistancy(cassandraConfig.Consistancy)
	var err error
	SessionDB, err = cluster.CreateSession()
	if err != nil {
		log.Panic("failed create session cassendra details error messsage : ", err)
	}
}
