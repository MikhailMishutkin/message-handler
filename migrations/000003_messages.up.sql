CREATE TABLE IF NOT EXISTS messages (
                                     uuid            serial not null PRIMARY KEY,
                                     author_uuid integer,
                                     message varchar,
                                     recieved_at timestamp,
                                     handled boolean,
                                     FOREIGN KEY (author_uuid) REFERENCES authors (uuid)
)