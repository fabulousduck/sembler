CXX=go
BUILD_FLAGS=build
TEST_FLAGS=test
ENTRY=cmd/sembler/main.go
TEST_ENTRY=github.com/fabulousduck/sembler/parser
make:
	$(CXX) $(BUILD_FLAGS) $(ENTRY)
test:
	$(CXX) $(TEST_FLAGS) $(TEST_ENTRY)