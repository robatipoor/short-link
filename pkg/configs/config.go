package configs

const ExpireTimeUrl = 100

type CassandraConfig struct {
	Host        string
	Port        string
	Keyspace    string
	Consistancy string
}
