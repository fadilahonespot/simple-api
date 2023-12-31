test-coverage:
	@echo "=================================================================================="
	@echo "Coverage Test"
	@echo "=================================================================================="
	go test -v -coverprofile coverage.cov ./...
	@echo "\n"
	@echo "=================================================================================="
	@echo "All Package Coverage"
	@echo "=================================================================================="
	go tool cover -func coverage.cov

build:
	@echo "=================================================================================="
	@echo "Stop And Delete Service"
	@echo "=================================================================================="
	docker stop simple-api || true 
	docker rm simple-api || true 
	@echo "\n"
	@echo "=================================================================================="
	@echo "Build And Run Service"
	@echo "=================================================================================="
	docker build -t fadilahonespot/simple-api:1.0.0 . 
	docker run -d -p 7690:7690 --name simple-api fadilahonespot/simple-api:1.0.0
	@echo "=================================================================================="
	@echo "Delete Image Not Tag"
	@echo "=================================================================================="
	docker images -f "dangling=true" -q | xargs docker rmi