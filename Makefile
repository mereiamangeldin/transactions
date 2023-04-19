start:
	docker-compose up --build

stop:
	docker-compose down
	docker rmi transaction-transaction

