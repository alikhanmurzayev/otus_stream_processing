# otus_stream_processing


docker run --rm -d -p 5432:5432 --name postgres --env POSTGRES_DB=mydb --env POSTGRES_USER=myuser --env POSTGRES_PASSWORD=mypassword postgres:latest
psql -h localhost -p 5432 -U myuser -W mydb

create table users (id bigserial primary key, first_name varchar, last_name varchar, login varchar, password varchar, salt varchar);
create table accounts (id bigserial primary key, user_id int, amount numeric, constraint fk_account_user foreign key(user_id) references users);
create table orders (id bigserial primary key, user_id int, price numeric, status varchar, constraint fk_order_user foreign key(user_id) references users);
create table notifications (id bigserial primary key, user_id int, email varchar, text varchar);