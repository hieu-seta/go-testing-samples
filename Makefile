.PHONY: test test-all run-test01 run-test02 run-test03 run-test04 run-test05 run-test06 run-test07 run-test08 run-test09 run-test10

# Run all tests with verbose output
test:
	go test ./... -v

# Alias for test
test-all: test

# Run individual test suites with verbose output
run-test01:
	go test ./01-normal -v

run-test02:
	go test ./02-table -v

run-test03:
	go test ./03-assertion -v

run-test04:
	go test ./04-testmain -v

run-test05:
	go test ./05-parallel -v

run-test06:
	go test ./06-subtest -v

run-test07:
	go test ./07-nethttp -v

run-test08:
	go test ./08-ginhttp -v

run-test09:
	go test ./09-database -v

run-test10:
	go test ./10-mocking -v

# Run all tests with coverage
coverage:
	go test ./... -cover -v

# Clean test cache
clean:
	go clean -testcache
