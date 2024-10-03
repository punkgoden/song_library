up_db:
	docker compose -f docker-compose.db.yml up --build -d
logs_db:
	docker compose -f docker-compose.db.yml logs -f
down_db:
	docker compose -f docker-compose.db.yml down
down_vol_db:
	docker compose -f docker-compose.db.yml down -v

up_song:
	docker compose -f docker-compose.yml up --build -d
down_song:
	docker compose -f docker-compose.yml down
logs_song:
	docker compose -f docker-compose.yml logs -f