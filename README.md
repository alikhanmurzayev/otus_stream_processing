# otus_stream_processing

minikube start --cpus=4 --memory=4g --vm-driver=virtualbox

kubectl apply -f k8s/namespaces.yaml
helm install -f k8s/rabbit.yaml my-rabbit bitnami/rabbitmq -n rabbit
kubectl apply -f k8s/config.yaml -f k8s/postgres.yaml -n postgres
kubectl apply -f k8s/config.yaml -f k8s/auth-app.yaml -n auth
kubectl apply -f k8s/config.yaml -f k8s/user-app.yaml -n user
kubectl apply -f k8s/config.yaml -f k8s/billing-app.yaml -n billing
kubectl apply -f k8s/config.yaml -f k8s/order-app.yaml -n order





docker run -d --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:management
docker run --rm -d -p 5432:5432 --name postgres --env POSTGRES_DB=mydb --env POSTGRES_USER=myuser --env POSTGRES_PASSWORD=mypassword postgres:latest
psql -h localhost -p 5432 -U myuser -W mydb

create table users (id bigserial primary key, first_name varchar, last_name varchar, login varchar, password varchar, salt varchar);
create table accounts (id bigserial primary key, user_id int, amount numeric, constraint fk_account_user foreign key(user_id) references users);
create table orders (id bigserial primary key, user_id int, price numeric, status varchar, constraint fk_order_user foreign key(user_id) references users);
create table notifications (id bigserial primary key, user_id int, text varchar, constraint fk_notification_user foreign key(user_id) references users);