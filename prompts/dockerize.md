
I am running an app that is a react frontend and go backend.
There is a sqllite db being used for memory.  

the golang has a standard directory struct at the root
/cmd
/internal

the react project is under
/meowmorize-frontend

Create the docker compose I need to run it all

Here is a helper file I use to run the front and back

```bash

#!/bin/bash

# Helper Script for Managing the Go Application

# Function to run the main application
run_main() {
    echo "Running the main application..."
    go run cmd/main/main.go
}


# Function to initialize swagger
init_swagger() {
    echo "Initializing swagger documentation..."
    rm -rf docs
    swag init --dir ./cmd/main --parseDependency --parseInternal --output ./docs --verbose
}

# Function to redo swagger from project root
redo_swagger() {
    echo "Redoing swagger documentation from cmd/main..."
    swag init  --dir ./cmd/main  --parseDependency --parseInternal --output ./docs
}

# Function to generate mocks
generate_mocks() {
    echo "Generating mocks..."
    mockery --dir=internal/domain  --name=MeowDomain --output=internal/domain/mocks --outpkg=mocks --case=underscore
    mockery --dir=internal/adapters/repositories  --name=CardRepository --output=internal/adapters/repositories/mocks --outpkg=mocks --case=underscore
    mockery --dir=internal/adapters/repositories  --name=DeckRepository --output=internal/adapters/repositories/mocks --outpkg=mocks --case=underscore
}

# Function to run build npm in meowmorize directory
run_npm_build() {
    echo "Running npm in meowmorize directory..."
    cd meowmorize-frontend && npm run build && cd ..
}

# Function to run npm in meowmorize directory
run_npm_start() {
    echo "Running npm in meowmorize directory..."
    cd meowmorize-frontend && npm run start && cd ..
}

# Function to run tests
run_tests() {
    echo "Running tests..."
    go test ./...
}

# Function to display usage information
usage() {
    echo "Usage: $0 {run|swagger|redoswagger|mocks|npm|test|help}"
    echo ""
    echo "Commands:"
    echo "  run         Run the main application"
    echo "  swagger     Initialize swagger docs"
    echo "  redoswagger Redo swagger docs from cmd/main directory"
    echo "  mocks       Generate mock files"
    echo "  npm-start   Run npm in meowmorize directory"
    echo "  npm-build   Run build npm in meowmorize directory"
    echo "  test        Run tests"
    echo "  help        Display this help message"
}

# Check if at least one argument is provided
if [ $# -lt 1 ]; then
    usage
    exit 1
fi

# Parse the command
case "$1" in
    run)
        run_main
        ;;
    swagger)
        init_swagger
        ;;
    redoswagger)
        redo_swagger
        ;;
    mocks)
        generate_mocks
        ;;
    npm-build)
        run_npm_build
        ;;
    npm-start)
        run_npm_start
        ;;
    test)
        run_tests
        ;;
    help|--help|-h)
        usage
        ;;
    *)
        echo "Error: Unknown command '$1'"
        usage
        exit 1
        ;;
esac