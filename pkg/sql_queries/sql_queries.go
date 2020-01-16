package sql_queries

const (
	InsertUser = `INSERT INTO users(about, email, fullname, nickname)
		VALUES($1,$2,$3,$4) RETURNING id;`
	SelectUserByNickname = "SELECT p.about, p.email, p.fullname, p.nickname FROM users as p WHERE p.nickname = $1"
	SelectUserByEmail    = "SELECT p.about, p.email, p.fullname, p.nickname FROM users as p WHERE p.email = $1"
)
