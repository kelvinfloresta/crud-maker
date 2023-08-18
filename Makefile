default:
	go run main.go

clean:
	rm -rf frameworks adapters usecases
	git checkout HEAD -- go.mod go.sum