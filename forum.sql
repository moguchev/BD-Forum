-- DROP TABLE threads;
-- DROP TABLE forums;

-- DROP TABLE posts;

CREATE EXTENSION IF NOT EXISTS citext;

DROP TABLE IF EXISTS users;
CREATE TABLE users(
    about TEXT,
    email citext NOT NULL UNIQUE,
    fullname TEXT NOT NULL,
    nickname citext PRIMARY KEY COLLATE "POSIX"
);

CREATE INDEX idx_users_nick ON users (nickname);
CREATE INDEX idx_users_email ON users (email);

-- CREATE TABLE forums(
--     id serial,
--     posts integer DEFAULT 0 NOT NULL,
--     slug text PRIMARY KEY,
--     threads integer DEFAULT 0 NOT NULL,
--     title text NOT NULL,
--     person text NOT NULL REFERENCES persons(nickname) ON DELETE CASCADE NOT NULL
-- );

-- CREATE TABLE threads(
--     id SERIAL PRIMARY KEY,
--     author text NOT NULL REFERENCES persons(nickname) ON DELETE CASCADE NOT NULL,
--     created timestamptz DEFAULT now(),
--     forum text REFERENCES forums(slug) ON DELETE CASCADE NOT NULL,
--     message text NOT NULL,
--     slug text,
--     title text NOT NULL,
--     votes integer DEFAULT 0 NOT NULL
-- );

-- CREATE TABLE posts(
--     id SERIAL PRIMARY KEY, --maybe need to switch serial to int
--     author text NOT NULL REFERENCES persons(nickname) ON DELETE CASCADE NOT NULL,
--     created timestamp with time zone DEFAULT '1970-01-01 03:00:00+03'::timestamp with time zone NOT NULL,
--     forum text REFERENCES forums(slug) ON DELETE CASCADE NOT NULL,
--     is_edited boolean DEFAULT false NOT NULL,
--     message text NOT NULL,
--     parent integer DEFAULT 0 NOT NULL,
--     thread integer REFERENCES threads(id) ON DELETE CASCADE NOT NULL
-- );

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