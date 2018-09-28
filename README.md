# k8s-workshop
Workshop repository to play with Docker and Kubernetes

### Links

- https://www.katacoda.com/courses/docker/deploying-first-container
- https://github.com/
- https://hub.docker.com/
- https://www.katacoda.com/courses/kubernetes/launch-single-node-cluster

### Docker

```sh 
docker run -dit --name web -p 8080:80 -v /home/user/website/:/usr/local/apache2/htdocs/ httpd:2.4-alpine
```

### GitHub

Resources
- Dockerfile
- index.html


```sh 
docker build -t web:1.0 .
```

```sh 
docker run -dit --name web -p 8080:80 twogghub/k8s-workshop:1.0
```

```sh 
localhost:8080
```

### Kubernetes 

```sh 
docker pull twogghub/k8s-workshop
``` 
```sh 
kubectl run web --image=twogghub/k8s-workshop:1.0
```

```sh 
kubectl expose deployment web --port=80 --external-ip=$(minikube ip) --type=LoadBalancer
```

```sh
vim deployment.yaml
```

```sh
{
  "apiVersion": "extensions/v1beta1",
  "kind": "Deployment",
  "metadata": {
    "name": "webyml"
  },
  "spec": {
    "replicas": 2,
    "template": {
      "metadata": {
        "labels": {
          "app": "webyml"
        }
      },
      "spec": {
        "containers": [
          {
            "name": "webyml",
            "image": "twogghub/k8s-workshop:1.0",
            "ports": [
              {
                "containerPort": 80
              }
            ]
          }
        ]
      }
    }
  }
}
```

```sh
kubectl create -f deployment.yaml
```

```sh 
kubectl expose deployment webyml --port=80 --external-ip=$(minikube ip) --type=LoadBalancer
```