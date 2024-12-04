docker build -t bloomblog-api -f cmd/api/Dockerfile .
docker build -t bloomblog-user -f cmd/user/Dockerfile .
docker build -t bloomblog-relation -f cmd/relation/Dockerfile .
docker build -t bloomblog-favorite -f cmd/favorite/Dockerfile .
docker build -t bloomblog-feed -f cmd/feed/Dockerfile .
docker build -t bloomblog-comment -f cmd/comment/Dockerfile .
docker build -t bloomblog-publish -f cmd/publish/Dockerfile .