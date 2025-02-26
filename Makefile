.PHONY: db db-stop benchmark test-benchmark

db:
	docker run -d --name mongodb_view_vs_pipeline -p 27017:27017 mongo

db-stop:
	docker stop mongodb_view_vs_pipeline || true
	docker rm mongodb_view_vs_pipeline || true

benchmark: db-stop
	@$(MAKE) db
	@echo "Running benchmark tests..."
	@go test -bench . -benchmem | tee benchmark_results.txt || true
	@$(MAKE) db-stop
