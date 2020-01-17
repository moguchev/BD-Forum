package sql_queries

const (
	InsertUser = `INSERT INTO users(about, email, fullname, nickname)
				    VALUES($1,$2,$3,$4);`

	UpdateUser = `UPDATE users SET about = $1, email = $2, fullname = $3
					WHERE nickname = $4`

	SelectUserByNickname = `SELECT u.about, u.email, u.fullname, u.nickname 
							  FROM users as u
							  WHERE u.nickname = $1`

	SelectUserByEmail = `SELECT u.about, u.email, u.fullname, u.nickname 
						   FROM users as u
						   WHERE u.email = $1`

	SelectUsersByNicknameAndEmail = `SELECT u.about, u.email, u.fullname, u.nickname 
									   FROM users as u
									   WHERE u.nickname = $1
									   OR u.email = $2`
)
