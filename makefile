install:
	@go get .
	@cd website	
	@yarn

dev-front:
	@cd website && yarn build:watch

dev-back:
	@go run .&

dev: 
	npx concurrently "make dev-back" "make dev-front"

