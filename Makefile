default:
	go run main.go

clean:
	rm -rf frameworks adapters usecases factories
	git checkout HEAD -- go.mod go.sum

e2e:
	make clean
	make