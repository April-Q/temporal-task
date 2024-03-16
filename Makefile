build:
	go mod tidy
	go build -o bin/starter-1 starter-1/main.go
	go build -o bin/starter-5 starter-5/main.go
	go build -o bin/worker worker/main.go

deploy:
	go run worker/main.go &
	go run starter-1/main.go
	go run starter-5/main.go

docker-build:
	docker build . -t shujiangdocker/task1-workflow:0.1 

	docker push shujiangdocker/task1-workflow:0.1

clean:
	temporal workflow delete  --query 'WorkflowType="SampleParentWorkflow"' --reason "Testing"
	temporal workflow delete  --query 'WorkflowType="SampleChildWorkflow"' --reason "Testing"	
