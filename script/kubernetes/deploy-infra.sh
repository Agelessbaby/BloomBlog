kubectl apply -f config/kubernetes/infra/mysql.yaml
kubectl apply -f config/kubernetes/infra/rabbitmq.yaml
sleep 30
kubectl apply -f config/kubernetes/infra/jobs/create_database.yaml
