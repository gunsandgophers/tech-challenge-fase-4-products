# tech-challenge

Tech Challenge...

Let's go tech challenge!!!

Links:

<https://miro.com/app/board/uXjVKQtHwOA=/>

## Architecture


![TechChallengeApp Kubernetes](https://github.com/user-attachments/assets/5032e83b-b429-4d6b-a41a-eb916e71fac5)

## Run project

To run the application it is necessary to execute the command `make start`

### Aplication

### Migration

All migrations are executed as soon as the `make start` or `make build` command is executed

#### Create

To create a migration, you need to run the `make migrate/create` command passing the file name

example:

```bash
make migrate/create name=add_user
```

to create a migration to add a user

### Swagger

URL to access running Swagger is `/api/v1/swagger/index.html`

## Kubernetes

> [!IMPORTANT]  
> [Minikube](https://minikube.sigs.k8s.io/docs?target=_blank) must be installed.

```bash
minikube start
eval $(minikube docker-env)
minikube addons enable volumesnapshots
minikube addons enable csi-hostpath-driver
docker buildx build -t tech-challenge-fase-4-products .
docker buildx build -t tech-challenge-fase-4-products-migration ./migrations/
kubectl apply -f k8s/configmap.yaml
kubectl apply -f k8s/secrets.yaml
kubectl apply -f k8s/database.yaml
kubectl apply -f k8s/deployment.yaml
kubectl expose deployment/tech-challenge-fase-4-products-deployment --port=80 --target-port=8080
kubectl apply -f k8s/nodeport.yaml
kubectl apply -f k8s/hpa.yaml
kubectl apply -f k8s/loadbalancer.yaml

#wait for postgres pod to finish
kubectl apply -f k8s/migration-job.yaml
minikube service tech-challenge-fase-4-products-nodeport --url
```
