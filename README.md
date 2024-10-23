# meowmorize
Flashcard app

rm -rf docs

swag init --dir ./cmd/main --parseDependency --parseInternal --output ./docs --verbose
 
 
run from cmd/main
swag init --parseDependency --parseInternal --output ../../docs
