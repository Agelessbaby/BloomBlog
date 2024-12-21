docker tag bloomblog-api registry.digitalocean.com/bloomblog/bloomblog-api:latest
docker tag bloomblog-user registry.digitalocean.com/bloomblog/bloomblog-user:latest
docker tag bloomblog-relation registry.digitalocean.com/bloomblog/bloomblog-relation:latest
docker tag bloomblog-feed registry.digitalocean.com/bloomblog/bloomblog-feed:latest
docker tag bloomblog-comment registry.digitalocean.com/bloomblog/bloomblog-comment:latest
docker tag bloomblog-publish registry.digitalocean.com/bloomblog/bloomblog-publish:latest
docker tag bloomblog-favorite registry.digitalocean.com/bloomblog/bloomblog-favorite:latest
docker push  registry.digitalocean.com/bloomblog/bloomblog-api:latest
docker push  registry.digitalocean.com/bloomblog/bloomblog-user:latest
docker push  registry.digitalocean.com/bloomblog/bloomblog-feed:latest
docker push  registry.digitalocean.com/bloomblog/bloomblog-comment:latest
docker push  registry.digitalocean.com/bloomblog/bloomblog-publish:latest
docker push  registry.digitalocean.com/bloomblog/bloomblog-favorite:latest
docker push  registry.digitalocean.com/bloomblog/bloomblog-relation:latest