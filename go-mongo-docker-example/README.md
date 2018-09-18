# Go-Mongo-Docker-Example

An example Docker Compose setup with MongoDB, Nginx and Go.



在 minikube 上搭建自己的 MongoDB 。

kubectl create -f mongo-deployment.yaml

➭ kubectl get pods
NAME                             READY     STATUS    RESTARTS   AGE
mongodb-65f7986d99-ll8tm         1/1       Running   0          12m

kubectl exec -it mongodb-65f7986d99-ll8tm -c mongodb bash
