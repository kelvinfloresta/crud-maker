default:
	go run main.go

clean:
	rm -rf frameworks adapters usecases factories
	git checkout HEAD -- go.mod go.sum

e2e:
	make clean
	make

build-template:
	go-bindata -o utils/bindata.go generators/templates/
	@echo ""
	@echo "Rename the package name at utils/bindata.go"