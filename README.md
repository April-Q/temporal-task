### Steps to run this sample:
1) You need a Temporal service running. See details in README.md
2) Run the following command to start the worker
```
go run child-workflow/worker/main.go
```
3) Run the following command to start the example
```
go run child-workflow/starter-1/main.go
go run child-workflow/starter-5/main.go
```
