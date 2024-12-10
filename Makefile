start:
	air -c .air.toml

generate:
	templ generate -watch -proxy=http://localhost:8080 --proxyport=3000 --proxybind="localhost"

tailwind:
	npx --yes tailwindcss -i ./server/public/input.css -o ./server/public/output.css --minify --watch

notify:
	air -c .assets.toml

migrateup:
	@echo "Migrating up"
	migrate -path sqlx/migration -database "$(url)" -verbose up
	@echo "Migrate up completed"

migratedown:
	@echo "Migrating down"
	migrate -path sqlx/migration -database "$(url)" -verbose down "$(count)"
	@echo "Migrate down completed"

forcedown:
	@echo "Forcing down migration"
	migrate -path sqlx/migration -database "$(url)" force "$(version)"
	@echo "Forced down migration one step completed"

sqlc:
	@echo "Generating sqlc"
	sqlc generate
	@echo "Sqlc generated"

generate_migration:
	migrate create -ext sql -dir sqlx/migration $(name)

dev:
	make -j4 start generate tailwind notify
