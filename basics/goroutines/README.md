# goroutines

## Modules
- [Basic](./basic.go)
- [NoSync](./nosync.go)
- [Mutex](./mutex.go)
- [Channels](./channels.go)
- [ProducerConsumer](./producer_consumer.go)
## Run 
go run main.go -run={parent_module} -routine={module}

<!--- identify race condition --->
go run -race main.go -run={parent_module} -routine={module} 
