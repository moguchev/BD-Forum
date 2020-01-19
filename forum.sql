CREATE EXTENSION IF NOT EXISTS citext;

DROP TABLE IF EXISTS users;
CREATE UNLOGGED TABLE users(
    about TEXT,
    email CITEXT NOT NULL UNIQUE CONSTRAINT email_right CHECK(email ~ '^.*@[A-Za-z0-9\-_\.]*$'),
    fullname TEXT NOT NULL,
    nickname CITEXT PRIMARY KEY COLLATE "POSIX" CONSTRAINT nick_right CHECK(nickname ~ '^[A-Za-z0-9_\.]*$')
);

CREATE INDEX idx_users_nick ON users (nickname);
CREATE INDEX idx_users_email ON users (email);

DROP TABLE IF EXISTS forums;
CREATE UNLOGGED TABLE forums (
    posts INTEGER DEFAULT 0,
    slug CITEXT PRIMARY KEY CONSTRAINT slug_correct CHECK(slug ~ '^(\d|\w|-|_)*(\w|-|_)(\d|\w|-|_)*$'),,
    threads INTEGER DEFAULT 0,  
    title TEXT NOT NULL,
    user_nick CITEXT REFERENCES users(nickname) ON DELETE RESTRICT ON UPDATE RESTRICT
);

CREATE INDEX idx_forum_slug ON forums (slug);

DROP TABLE IF EXISTS threads;
CREATE UNLOGGED TABLE threads (
    id SERIAL PRIMARY KEY,
    slug CITEXT UNIQUE CONSTRAINT slug_correct CHECK(slug ~ '^(\d|\w|-|_)*(\w|-|_)(\d|\w|-|_)*$'),
    forum CITEXT REFERENCES forums (slug) ON DELETE CASCADE ON UPDATE RESTRICT NOT NULL,
    author CITEXT REFERENCES users (nickname) ON DELETE CASCADE ON UPDATE CASCADE NOT NULL,
    created TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    message TEXT NOT NULL,
    title TEXT NOT NULL,
    votes INTEGER NOT NULL DEFAULT 0
);


-- CREATE SEQUENCE threads_id_seq
--     AS integer
--     START WITH 1
--     INCREMENT BY 1
--     NO MINVALUE
--     NO MAXVALUE
--     CACHE 1;

-- ALTER SEQUENCE threads_id_seq OWNED BY threads.id;

-- ALTER TABLE ONLY threads ALTER COLUMN id SET DEFAULT nextval('threads_id_seq'::regclass);

-- SELECT pg_catalog.setval('threads_id_seq', 1, true);

-- CREATE SEQUENCE persons_id_seq
--     AS integer
--     START WITH 1
--     INCREMENT BY 1
--     NO MINVALUE
--     NO MAXVALUE
--     CACHE 1;


-- ALTER SEQUENCE persons_id_seq OWNED BY persons.id;

-- ALTER TABLE ONLY persons ALTER COLUMN id SET DEFAULT nextval('persons_id_seq'::regclass);

-- SELECT pg_catalog.setval('persons_id_seq', 1, true);



-- INSERT INTO persons(about, email, fullname, nickname)VALUES('asdas', 'email', 'vasya', 'nick');