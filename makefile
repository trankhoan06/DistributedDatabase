mysql:
	docker run --name project_personal -e MYSQL_ROOT_PASSWORD=123456 -d -p 3309:3306 mysql:8.0.31