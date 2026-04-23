# go-test-examples

## How to Run Tests

```bash
# Run all tests from root directory
go test ./...

# Run individual test directories
go test ./01-normal      # Basic unit test (arrange/act/assert pattern)
  # Simple function testing with clear failure messages

# go test ./02-table       # Table-driven test
  # Multiple test cases in single function using t.Run()
  # Reduces code duplication through test case struct

# go test ./03-assertion  # Assertion with testify
  # Uses testify/assert for clearer checks
  # Better failure messages with descriptive error texts

# go test ./10-mocking    # Mock/stub example
  # Interface-based dependency injection
  # Mock implementation for testing without real dependencies

# Run with verbose output (shows subtests)
go test -v ./02-table

# Run specific test
go test -run TestSum ./01-normal
```

## Test Implementation Details

### 1. Basic Unit Test (`01-normal/`)
- **Structure**: Arrange/Act/Assert pattern
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

### 4. Mock/Stub Example (`10-mocking/`)
- **Structure**: Interface-based mocking
- **Key Features**:
  - `MathService` interface defines behavior
  - `MockRepository` implements the interface
  - `On()` sets expectations
  - `AssertExpectations()` verifies mock usage
  - Complete dependency isolation

## Key Improvements
- All comments translated to standard English
- Added implementation explanations in README
- Standardized variable names (e.g., `want` instead of Turkish equivalents)
- Clearer failure message formatting
- Better test case organization