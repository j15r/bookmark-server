# bookmark server

Simple self hosted bookmark server to sync browser bookmarks with WebExtension.

Provides api to fetch and store bookmarks from browser extension or phone app.

## build

```bash=
docker build -t bookmark-server:latest -f Dockerfile .
```

## start dependencies

```bash=
docker-compose up -d
```