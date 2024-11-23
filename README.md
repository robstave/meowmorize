# meowmorize
Flashcard app

rm -rf docs

swag init --dir ./cmd/main --parseDependency --parseInternal --output ./docs --verbose
 
 
run from cmd/main
swag init --parseDependency --parseInternal --output ../../docs


# Mock tests
to generate
 
mockery --dir=internal/domain  --name=MeowDomain --output=internal/domain/mocks --outpkg=mocks --case=underscore

mockery --dir=internal/adapters/repositories  --name=CardRepository --output=internal/adapters/repositories/mocks --outpkg=mocks --case=underscore

mockery --dir=internal/adapters/repositories  --name=DeckRepository --output=internal/adapters/repositories/mocks --outpkg=mocks --case=underscore

npm
in meowmorize dir
  npm run