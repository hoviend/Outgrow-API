DATABASE="$(DB_DRIVER)://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)"

db.create: check-env
	@echo "creating database $(DB_NAME)..."
	@PGPASSWORD=${DB_PASSWORD} createdb -h $(DB_HOST) -U $(DB_USER) -Eutf8 $(DB_NAME)

db.drop: check-env
	@echo "dropping database $(DB_NAME)..."
	@PGPASSWORD=${DB_PASSWORD} dropdb --if-exists -h $(DB_HOST) -U $(DB_USER) $(DB_NAME)

db.seed.events: check-env-test
	PGPASSWORD=${DB_PASSWORD} psql -h $(DB_HOST) -U $(DB_USER) $(DB_NAME) < db/seeds/masterevent.sql

migrate.apply: check-env
	atlas migrate apply --url $(DATABASE) --dir "file://migrations"

migrate.status: check-env
	atlas migrate status --dir "file://migrations" --url $(DATABASE) --revisions-schema "atlas_schema_revisions"

migrate.clean: check-env ## not recommended
	atlas schema clean -u $(DATABASE)

migrate.new: ## It takes parameter file as migration file name
	go run -mod=mod ent/migrate/main.go $(file)

migrate.manual.new: ## It takes parameter file as migration file name
	atlas migrate new --dir "file://migrations" $(file)
	@echo "new migration files successfully created..."	


check-env:
ifndef DB_USER
	$(error DB_USER is undefined)
endif
ifndef DB_HOST
	$(error DB_HOST is undefined)
endif
ifndef DB_NAME
	$(error DB_NAME is undefined)
endif