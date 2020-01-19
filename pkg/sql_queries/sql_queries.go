package sql_queries

const (
	InsertUser = `INSERT INTO users(about, email, fullname, nickname)
				    VALUES($1,$2,$3,$4);`

	UpdateUser = `UPDATE users SET about = $1, email = $2, fullname = $3
					WHERE nickname = $4;`

	SelectUserByNickname = `SELECT u.about, u.email, u.fullname, u.nickname 
							  FROM users as u
							  WHERE u.nickname = $1;`

	SelectUserByEmail = `SELECT u.about, u.email, u.fullname, u.nickname 
						   FROM users as u
						   WHERE u.email = $1;`

	SelectUsersByNicknameAndEmail = `SELECT u.about, u.email, u.fullname, u.nickname 
									   FROM users as u
									   WHERE u.nickname = $1
									   OR u.email = $2;`

	Truncate = "TRUNCATE users, forums, threads RESTART IDENTITY CASCADE;"

	InsertForum = `INSERT INTO forums(slug, title, user_nick)
					 VALUES($1,$2,$3);`

	SelectForum = `SELECT posts, slug, threads, title, user_nick
					 FROM forums
					 WHERE slug = $1;`

	InsertThread = `INSERT INTO threads (author, forum, message, created, title, slug)
						VALUES ($1, (SELECT slug FROM forums WHERE slug = $2), $3,
						COALESCE($4, NOW())::timestamptz, $5, NULLIF($6, ''))
						RETURNING id, slug, 0, forum`

	SelectThreadBySlug = `SELECT slug, title, message, forum, author, created, votes, id
							FROM threads 
							WHERE slug = $1`

	SelectThreads = `SELECT t.slug, t.title, t.message, t.forum, t.author, t.created, t.votes, t.id
						FROM threads AS t
						RIGHT OUTER JOIN forums AS f
						ON t.forum = f.slug
						WHERE f.slug = $1`
)
