build:
	make build_auth
	make build_user
	make build_billing
	make build_order
	make build_notification

build_auth:
	docker build -t murzayev/otu_stream_processing_auth:latest apps/auth/
	docker push murzayev/otu_stream_processing_auth:latest

build_user:
	docker build -t murzayev/otu_stream_processing_user:latest apps/user/
	docker push murzayev/otu_stream_processing_user:latest

build_billing:
	docker build -t murzayev/otu_stream_processing_billing:latest apps/billing/
	docker push murzayev/otu_stream_processing_billing:latest

build_order:
	docker build -t murzayev/otu_stream_processing_order:latest apps/order/
	docker push murzayev/otu_stream_processing_order:latest

build_notification:
	docker build -t murzayev/otu_stream_processing_notification:latest apps/notification/
	docker push murzayev/otu_stream_processing_notification:latest

build_rabbit_isready:
	docker build -t murzayev/otu_stream_processing_rabbit_isready:latest apps/rabbit_isready/
	docker push murzayev/otu_stream_processing_rabbit_isready:latest
