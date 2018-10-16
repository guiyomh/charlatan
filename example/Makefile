
run: clean build/go-faker-fixture-osx-x64 ## Build an Run osx
	LOG=* ./build/go-faker-fixture-osx-x64 load ./fixtures -u fixtures_user -p fixtures_pass -d fixtures

mariadb/start:
	@docker build -f docker/mariadb/Dockerfile -t gofixtures docker/mariadb
	@docker run -d -p 3306:3306 --name=gofixtures gofixtures

mariadb/stop:
	@docker rm -f gofixtures