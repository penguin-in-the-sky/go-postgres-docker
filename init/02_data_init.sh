psql -U postgres -d godb << "EOSQL"
    CREATE SCHEMA gogo;
    CREATE TABLE gogo.users
        (
            id serial,
            user_name varchar(20),
            invalid_flg boolean
        );
EOSQL
psql -l