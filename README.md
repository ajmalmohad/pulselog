# Pulselog

Pulse log is an event tracking application. You can create a project in pulselog and use the sdk to trigger events from various applications. You can use it for testing, logging, tracking, and various other things. It's up to your imagination

# Run docker compose

## Development environment
```sh
docker-compose -f docker-compose.dev.yaml --env-file .env build
docker-compose -f docker-compose.dev.yaml --env-file .env up
```

## Production environment
```sh
docker-compose -f docker-compose.prod.yaml --env-file .env build
docker-compose -f docker-compose.prod.yaml --env-file .env up -d
```

## LSP
As node_modules is a volume you have to install it inside the directory as well for your LSP to work properly during development