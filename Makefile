build:
	@echo Building Application...
	@go build -o htmlanalyzer ./cmd
	@echo Finished!

test:
	@echo Running Tests...
	@go test ./service
	@go test ./util
	@echo Done!

run:
	@echo Running Application...
	./htmlanalyzer

clean:
	@echo Removing Binary...
	rm -fr htmlanalyzer
	@echo Done!