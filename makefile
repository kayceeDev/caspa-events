run:
	docker-compose up --build -d
log:
	docker logs -f gnl-ai-ps
gen:
	@go generate ./ent-go

update: ## run update after you run create_migration
	@docker-compose down
	@go mod tidy -e -v
	@go generate ./ent-go
	$(MAKE) run

test:
	@go test ./... -v

create_migration: ## Generates a sql migration file based on the defined Ent models locally. Usage is `make create_migration name=Foobar`
	@go run -mod=mod entgo.io/ent/cmd/ent new --target "ent-go/schema" $(name)
	# $(MAKE) update
check:


.PHONY:run down gen update create_migration
