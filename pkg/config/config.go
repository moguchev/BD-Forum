package config

import ()

const (
	HostAddress = "0.0.0.0:5000"
	User        = "forum"
	Password    = "forum"
	DBName      = "forum"
	SSLMode     = "disable"
	MaxConn     = 1000
	DBSchema    = "forum.sql"
	DBPath      = "postgresql://forum:forum@localhost:5432/forum"
)

// const (
// 	HostAddress = "0.0.0.0:5001"
// 	User        = "postgres"
// 	Password    = "forum"
// 	DBName      = "dbtechno"
// 	SSLMode     = "disable"
// 	MaxConn     = 1000
// 	DBSchema    = "forum.sql"
// 	DBPath      = "postgresql://postgres:forum@localhost:5432/dbtechno"
// )
