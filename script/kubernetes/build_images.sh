docker buildx build --platform linux/amd64 -t registry.digitalocean.com/bloomblog/bloomblog-api:latest -f cmd/api/Dockerfile --push .
docker buildx build --platform linux/amd64 -t registry.digitalocean.com/bloomblog/bloomblog-user:latest -f cmd/user/Dockerfile --push .
docker buildx build --platform linux/amd64 -t registry.digitalocean.com/bloomblog/bloomblog-relation:latest -f cmd/relation/Dockerfile --push .
docker buildx build --platform linux/amd64 -t registry.digitalocean.com/bloomblog/bloomblog-favorite:latest -f cmd/favorite/Dockerfile --push .
docker buildx build --platform linux/amd64 -t registry.digitalocean.com/bloomblog/bloomblog-feed:latest -f cmd/feed/Dockerfile --push .
docker buildx build --platform linux/amd64 -t registry.digitalocean.com/bloomblog/bloomblog-comment:latest -f cmd/comment/Dockerfile --push .
docker buildx build --platform linux/amd64 -t registry.digitalocean.com/bloomblog/bloomblog-publish:latest -f cmd/publish/Dockerfile --push .