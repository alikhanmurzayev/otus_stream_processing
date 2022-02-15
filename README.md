# otus_stream_processing

minikube start --cpus=4 --memory=4g --vm-driver=virtualbox

kubectl apply -f k8s/namespaces.yaml
helm install -f k8s/rabbit.yaml my-rabbit bitnami/rabbitmq -n rabbit
kubectl apply -f k8s/config.yaml -f k8s/postgres.yaml -n postgres
kubectl apply -f k8s/config.yaml -f k8s/auth-app.yaml -n auth
kubectl apply -f k8s/config.yaml -f k8s/user-app.yaml -n user
kubectl apply -f k8s/config.yaml -f k8s/billing-app.yaml -n billing
kubectl apply -f k8s/config.yaml -f k8s/order-app.yaml -n order
kubectl apply -f k8s/config.yaml -f k8s/notification-app.yaml -n notification

helm install -n ambassador --set licenseKey.value=eyJhbGciOiJQUzUxMiIsInR5cCI6IkpXVCJ9.eyJsaWNlbnNlX2tleV92ZXJzaW9uIjoidjIiLCJjdXN0b21lcl9pZCI6ImplZG93YXQ3MDlAdmViMzQuY29tLTE2NDE3NDM0ODUiLCJjdXN0b21lcl9lbWFpbCI6ImplZG93YXQ3MDlAdmViMzQuY29tIiwiZW5hYmxlZF9mZWF0dXJlcyI6WyIiLCJmaWx0ZXIiLCJyYXRlbGltaXQiLCJ0cmFmZmljIiwiZGV2cG9ydGFsIl0sImVuZm9yY2VkX2xpbWl0cyI6W3sibCI6ImRldnBvcnRhbC1zZXJ2aWNlcyIsInYiOjV9LHsibCI6InJhdGVsaW1pdC1zZXJ2aWNlIiwidiI6NX0seyJsIjoiYXV0aGZpbHRlci1zZXJ2aWNlIiwidiI6NX0seyJsIjoidHJhZmZpYy11c2VycyIsInYiOjV9XSwibWV0YWRhdGEiOnt9LCJleHAiOjE2NzMyNzk0ODUsImlhdCI6MTY0MTc0MzQ4NSwibmJmIjoxNjQxNzQzNDg1fQ.BhQi32pQiNR9KWO8iPblts39iU2asq7x7yrWHpIV4_n7wdpVrDHTjAmRRowu_FDGXMkCrUVfDJuNheMrmHjMoA1avCKYlL9E-xv1oQwrg6kFuTaS3xHL6rP9VEX7aawobtoybOpXZiorr0W3W6hvER1yihwbxjCW1dcpYjD1lZ9I-5qd9na1fWrKD1L37Oxm5kTixPKF7usbAfxpZN-cqVRPIpivHsqsPbsh-AB44hPWDKnnh3qZwUQ39Y0lw3z1R7cWrqet4G_yYcIU-KQZFQVXRMlL9FCCwgS9g_xHe8Xx65k3xxXIjKt-RKwZuRc7q6S-ACeWPNwcfdvW0vVH4w \
-f k8s/ambassador.yaml ambassador datawire/ambassador

kubectl apply -f k8s/ambassador-routes.yaml -f k8s/ambassador-auth.yaml

newman run  \
--global-var "baseUrl=$(minikube service list|grep 'ambassador'|grep 'http'|grep -Eo 'http://[^ >]+'|head -1)" \
OTUS.API_Gateway.postman_collection.json













docker run -d --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:management
docker run --rm -d -p 5432:5432 --name postgres --env POSTGRES_DB=mydb --env POSTGRES_USER=myuser --env POSTGRES_PASSWORD=mypassword postgres:latest
psql -h localhost -p 5432 -U myuser -W mydb

create table users (id bigserial primary key, first_name varchar, last_name varchar, login varchar, password varchar, salt varchar);
create table accounts (id bigserial primary key, user_id int, amount numeric, constraint fk_account_user foreign key(user_id) references users);
create table orders (id bigserial primary key, user_id int, price numeric, status varchar, constraint fk_order_user foreign key(user_id) references users);
create table notifications (id bigserial primary key, user_id int, text varchar, constraint fk_notification_user foreign key(user_id) references users);