psql -U postgres -d godb << "EOSQL"
    CREATE SCHEMA gogo;
    CREATE TABLE gogo.authorities
        (
            authority varchar(20) primary key,
            name varchar(20) UNIQUE NOT NULL
        );
    CREATE TABLE gogo.users
        (
            id serial primary key,
            user_name varchar(20) UNIQUE NOT NULL,
            password varchar(20) NOT NULL,
            authority varchar(20) references gogo.authorities(authority) NOT NULL,
            invalid_flg boolean NOT NULL
        );
    CREATE TABLE gogo.praises
        (
            id serial primary key,
            content varchar(100) NOT NULL,
            target_user_id int references gogo.users(id) NOT NULL,
            registered_user_id int references gogo.users(id) NOT NULL,
            has_approved boolean NOT NULL
        );
    BEGIN;
    INSERT INTO gogo.authorities (authority, name) VALUES ('general', '一般ユーザ');
    INSERT INTO gogo.authorities (authority, name) VALUES ('administrator', '管理者');
    COMMIT;
EOSQL
psql -l