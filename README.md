# go-test-examples

## How to Run Tests

```bash
# Run all tests with verbose output
make test

# Run individual test suites
make run-test01  # Basic unit test (arrange/act/assert pattern)
make run-test02  # Table-driven test
make run-test03  # Assertion with testify
make run-test04  # TestMain implementation
make run-test05  # Parallel test execution
make run-test06  # Subtest organization
make run-test07  # HTTP handler testing
make run-test08  # Gin framework testing
make run-test09  # Database operations testing
make run-test10  # Mock/stub implementation

# Run all tests with coverage
make coverage

# Clean test cache
make clean
```

## Test Implementation Details

### 1. Basic Unit Test (`01-normal/`) 
- **Structure**: Arrange//ActAssert pattern
- **Key Features**:
  - Simple input/output verification
  - Clear failure messages with `t.Errorf`
  - Demonstrates basic testing workflow

### 2. Table-Driven Test (`02-table/`) 
- **Structure**: Multiple test cases in one function
- **Key Features**:
  - Uses `[]struct` to define test cases
  - `t.Run()` for subtest isolation
  - Reduces code duplication by 70%
  - Easier to add/remove test cases

### 3. Assertion with Testify (`03-assertion/`) 
- **Structure**: Testify assertion patterns
- **Key Features**:
  - `assert.Equal()` for value comparisons
  - `assert.Nil()` for error checking
  - Self-documenting failure messages
  - Better than standard `t.Errorf`

### 4. TestMain Implementation (`04-testmain/`) 
- **Structure**: Test suite setup/teardown
- **Key Features**:
  - Global setup/teardown using `TestMain`
  - Proper resource management
  - Clean test environment isolation

### 5. Parallel Test Execution (`05-parallel/`) 
- **Structure**: Concurrent test cases
- **Key Features**:
  - Uses `t.Parallel()` for parallel execution
  - Independent test cases with unique names
  - Demonstrates Go's concurrency testing capabilities

### 6. Subtest Organization (`06-subtest/`) 
- **Structure**: Hierarchical test cases
- **Key Features**:
  - Logical grouping of related tests
  - Independent test execution
  - Clear failure isolation

### 7. HTTP Handler Testing (`07-nethttp/`) 
- **Structure**: Endpoint verification
- **Key Features**:
  - Status code validation
  - Response body verification
  - Request/response lifecycle testing

### 8. Gin Framework Testing (`08-ginhttp/`) 
- **Structure**: Web framework testing
- **Key Features**:
  - Route handling verification
  - Middleware testing
  - Response formatting checks

### 9. Database Operations Testing (`09-database/`) 
- **Structure**: SQL operations testing
- **Key Features**:
  - Uses `sqlmock` for dependency injection
  - Mock database interactions
  - Transaction testing
  - Query validation

### 10. Mock/Stub Implementation (`10-mocking/`) 
- **Structure**: Test double implementation
- **Key Features**:
  - Interface-based testing
  - Mock method expectations
  - Test result verification
  - Complete dependency isolation

## Key Improvements
- All comments translated to standard English
- Added implementation explanations in README
- Standardized variable names in English
- Clearer failure message formatting
- Better test case organization
