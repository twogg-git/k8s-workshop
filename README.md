# k8s-workshop
Workshop repository to play with Docker and Kubernetes


```sh 
docker run -dit --name web -p 8080:80 -v /home/ccruz/website/:/usr/local/apache2/htdocs/ httpd:2.4-alpine


docker build -t web:1.0 .

docker run -dit --name web -p 8080:80 web:1.0


docker pull twogghub/k8s-workshop

kubectl run webk8s --image=twogghub/k8s-workshop:1.0

kubectl expose deployment webk8s --port=8080 --external-ip=$(minikube ip) --type=LoadBalancer

```

