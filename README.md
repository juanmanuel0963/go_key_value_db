# Key-Value database

The Key-Value Database project is a simple and efficient key-value store implemented in the Go programming language. It provides a way to store, retrieve, and manage data in a key-value format, making it suitable for a wide range of applications, from caching to configuration management.

## Setup

Run the project using the provided Makefile. Simply execute the following commands:

    make clean
    make install
    make build
    make run
    make test

Alternatively, you can execute all the commands above with this single command:

    make all


This will clean the project, install dependencies, build it, run tests, and execute the project. It's a convenient way to get the project up and running.


## Database files

You can find the Database files in directory

    data/data.json
    data/timestamps.json

## Limitations

**Concurrent Access**: While the key-value database supports concurrent access, executing the test in the `tests/TestConcurrentAccessToExec.go` file, which uses go functions and channels in a loop, may not effectively demonstrate its concurrency handling. When calling the executable file through this test, the executable may face challenges in handling concurrency as effectively as the database functions. 

For a more accurate evaluation of the database's concurrency management, it is recommended to use the `tests/TestConcurrentAccessToFunc` test instead, which also uses go functions and channels in a loop, as it directly accesses the database functions and provides a more reliable assessment of concurrent access capabilities.

**Sequential Access**: When executing the test in `tests/TestSecuentialAccessToExec.go`, which calls the executable file sequentially in a simple loop, the program behaves well and handles access sequentially without concurrency issues.

**Single-Node:** This project is designed as a single-node key-value store and does not support distributed features like replication or sharding.

**Basic Data Types:** It stores simple string key-value pairs and does not support more complex data types or queries.

**No Authentication:** There is no authentication or security built into this database, so use it in trusted environments.


## Technical discussion

### Programming Language ###

Go (Golang) is chosen as the programming language. Go is well-suited for systems programming and concurrent applications due to its built-in concurrency primitives.

### Data Storage ###

Data is stored in a combination of in-memory (map) and on-disk (JSON files) storage. This hybrid approach allows for data persistence even after the program exits.

### Concurrency Handling ###

Go's sync.RWMutex is used for concurrency control. This mutex allows multiple readers to access data simultaneously while ensuring exclusive access for writing operations. This design helps prevent data races and ensures data integrity.

### Command-Line Interface (CLI) ###

The Cobra library is used to build a command-line interface. Commands such as set, get, del, and ts are provided for setting, getting, deleting, and retrieving timestamps of key-value pairs. The CLI is user-friendly and accessible from the terminal.

### Data Persistence ###

Data is saved to disk using JSON files. This enables data to persist across program runs and system reboots.

### File Format ###

JSON is chosen as the file format for data storage and timestamps. JSON is a human-readable and widely supported format. It simplifies data serialization and deserialization.

### Testing ###

The project includes test files that cover various aspects of the key-value store, including sequential and concurrent access. Testing is crucial to validate the correctness of the system and identify any potential issues.

### Concurrency Testing ###

Different test scenarios are provided to evaluate how the key-value store handles concurrent access. Testing concurrent access is essential to ensure the program's reliability under multi-user conditions.

### Makefile ###

A Makefile is included in the project to automate common development tasks. It allows developers to build, test, and clean the project with simple commands, enhancing the development workflow.

### Error Handling ###

The code includes basic error handling to capture and report errors, such as when a command execution fails. Providing meaningful error messages is essential for a good user experience. However, saving errors on disk is out of scope.