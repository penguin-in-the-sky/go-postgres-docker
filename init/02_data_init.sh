psql -U postgres -d godb << "EOSQL"
    CREATE SCHEMA gogo;
    CREATE TABLE gogo.users
        (
            id serial,
            user_name varchar(20) UNIQUE,
            invalid_flg boolean
        );
    CREATE TABLE gogo.praises
        (
            id serial,
            content varchar(600),
            user_id int,
            has_approved boolean
        );
EOSQL
psql -l