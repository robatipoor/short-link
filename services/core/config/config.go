package config

import (
	"log"
	"strconv"

	"github.com/gocql/gocql"
	"github.com/robatipoor/short-link/pkg/utils"
)

const (
	KeyPoolSize              = 10
	KeyLen                   = 6
	KeyJobScheduleTimeSecond = 10
	ServerPort = "8080"
	ServerAddr = "127.0.0.1"
	Http = "http://"
)

var SessionDB *gocql.Session

type CassandraConfig struct {
	host        string
	port        string
	keyspace    string
	consistancy string
}

var cassandraConfig = CassandraConfig{
	host:        utils.GetEnv("CASSANDRA_HOST", "127.0.0.1"),
	port:        utils.GetEnv("CASSANDRA_PORT", "9042"),
	keyspace:    utils.GetEnv("CASSANDRA_KEYSPACE", "core_space"),
	consistancy: utils.GetEnv("CASSANDRA_CONSISTANCY", "LOCAL_QUORUM"),
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

	cluster := gocql.NewCluster(cassandraConfig.host)
	cluster.Port = port(cassandraConfig.port)
	cluster.Keyspace = cassandraConfig.keyspace
	cluster.Consistency = consistancy(cassandraConfig.consistancy)
	var err error
	SessionDB, err = cluster.CreateSession()
	if err != nil {
		log.Panic("failed create session cassendra details error messsage : ", err)
	}
}

