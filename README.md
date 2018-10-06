# Taller Práctico de Kubernetes 
<img height="300" width="500" src="https://raw.githubusercontent.com/twogg-git/k8s-workshop/master/src/kubernetes.png">

Trabajaremos los conceptos claves al momento de realizar despliegues con Kubernetes, sus objetos, los comandos mas usados, los trucos, y también incluye un repaso de Docker y Docker Registry. 

Todos los recursos utilizados hacen parte del repositorio y los mencionare a medida que los usemos. Este taller esta pensado para que al final cuentes con tu repositorio de images de docker desplegadas en kubernetes, asi que en lo posible no te saltes la parte de Docker.

## Recursos a utilizar
- https://github.com/
- https://hub.docker.com/
- https://www.katacoda.com/

## Slides taller

Los conceptos de Kubernetes a manejar y de Docker esta en los siguientes slides, este repositorio es la base para los ejercicios explicados en los slides.

`https://github.com/twogg-git/talks/blob/master/Kubernetes_DuraznoConf.pdf`

# 1. Docker

## 1.1. Desplegar una muy sencilla pagina web

Docker play ground para crear y ejecutar la imagen:    
https://www.katacoda.com/courses/docker/persisting-data-using-volumes   

- Ejecutar el primer contenedor
```sh 
docker container run --rm -d -p 80:80 --name nginx nginx:1.8-alpine
``` 

- Creando archivos recursos para cargar nuestra propia pagina
```sh
mkdir website

cd website

vim index.html
```

- Contenido HTML pagina web
``` sh
<!DOCTYPE html>
<html>
<head>
  <title>Docker</title>
</head>
<body><center>
  <img src="https://raw.githubusercontent.com/twogg-git/k8s-workshop/master/src/0.0.png">
  <h1 style="color:blue">Baby steps with docker!</h1>   
</center></body>
</html>
``` 

- Cargar nuestra pagina con dos tipos de servidores: Nginx y Httpd
``` sh
docker container run --rm -d -p 80:80 --name nginx -v /root/website:/usr/share/nginx/html nginx:1.8-alpine

docker container run --rm -d -p 8080:80 --name httpd -v /root/website/:/usr/local/apache2/htdocs/ httpd:2.4-alpine
``` 

- Crear archivo Dockerfile
``` sh
cd ..

vim Dockerfile
``` 

- Contenido archivo para construir nuestra imagen
``` sh
FROM httpd:2.4-alpine

ADD website/ /usr/local/apache2/htdocs/

EXPOSE 80
``` sh

- Construir archivo Dockerfile y ejecutar el contenedor
``` sh 
docker build -t httpd .

docker container run --rm -p 80:80 --name httpd httpd
```

## 1.2. Repositorios en GitHub

```sh
https://github.com/
```
Repositorios para el codigo fuente, yaml, y Dockerfile, tambien lo usaremos para construir imagenes en DockerHub.      
![](https://raw.githubusercontent.com/twogg-git/k8s-workshop/master/src/github_signup.png)

- Ahora a crear los repositorios de las dos imagenes, Go y Httpd   
<img height="190" width="150" align="center" src="https://raw.githubusercontent.com/twogg-git/k8s-workshop/master/src/new_repo.png">

- Tener en cuenta el nombre del repositorio lo usaremos luego!, tambien agregar el archivo readme   
<img height="190" width="150" src="https://raw.githubusercontent.com/twogg-git/k8s-workshop/master/src/new_repo2.png">

- Repositorio *Httpd*, archivos Dockerfile y website-index.html. Los fuentes a crear esta en los siguientes links:     
```sh
https://repl.it/@twogg_git/index   
https://repl.it/@twogg_git/DockerHttpd   
```
<img src="https://raw.githubusercontent.com/twogg-git/k8s-workshop/master/src/repo_files_html.png">

- Repositorio *Go*, archivos Dockerfile y k8s.go 
```sh
https://repl.it/@twogg_git/k8s-10   
https://repl.it/@twogg_git/DockerGO 
```
<img src="https://raw.githubusercontent.com/twogg-git/k8s-workshop/master/src/repo_files_go.png">

## 1.3. Docker Registry

```sh
https://hub.docker.com/
```
Este es el repositorio para las imagenes Docker que vamos a crear.
![](https://raw.githubusercontent.com/twogg-git/k8s-workshop/master/src/dockerhub_signup.png)

- Ahora a crear nuestra construccion automatizada, *este paso para ambos repositorios de GitHub Go y Httpd   
<img height="190" width="150" src="https://raw.githubusercontent.com/twogg-git/k8s-workshop/master/src/hub_new_repo.png">

- Ahora a crear nuestra construccion automatizada, *este paso para ambos repositorios Go y Httpd   
<img height="190" width="150" src="https://raw.githubusercontent.com/twogg-git/k8s-workshop/master/src/hub_new_repo.png">

- Seleccionar repo tipo GitHub y luego buscar el que vamos a crear, ahora agregar la descripcion y crear 
<img src="https://raw.githubusercontent.com/twogg-git/k8s-workshop/master/src/aut_hub.png">

- Ahora, en *Build Settings* damos una version a nuestra imagen ejemplo 1.0, luego click en save y al final click trigger. 
<img src="https://raw.githubusercontent.com/twogg-git/k8s-workshop/master/src/build_settings.png">

- Revisamos que se construya la imagen en *Build Details*. 
<img src="https://raw.githubusercontent.com/twogg-git/k8s-workshop/master/src/build_details.png">

- El tag de nuestra imagen. 
<img src="https://raw.githubusercontent.com/twogg-git/k8s-workshop/master/src/tags.png">

# 2. Kubernetes

## 2.1. Comandos básicos 

Codigo fuente en ejecución de la pagina HTML: [index.html](https://github.com/twogg-git/k8s-workshop/blob/1.0-baby/docker/website/index.html)

Katacoda Minikube playground: https://www.katacoda.com/courses/kubernetes/launch-single-node-cluster#

```sh 
kubectl get all

dashboard

kubectl get pods --output wide --show-labels --watch

kubectl run k8sjr --image=twogghub/k8s-workshop:1.0-httpd

clear && kubectl get all

kubectl expose deployment k8sjr --port=80 --external-ip=$(minikube ip) --type=LoadBalancer

kubectl get all

kubectl describe pod k8sjr-6d9fd6c5c-htm9f

kubectl label pod k8sjr-6d9fd6c5c-htm9f version="jrV1"

kubectl scale --replicas=5 deployment k8sjr 

kubectl scale --replicas=1 deployment k8sjr

kubectl run k8sqa --image=twogghub/k8s-workshop:1.0-httpd --replicas=2 --labels="version=qa,dev=twogg"

kubectl expose deployment k8sqa --port=90 --target-port=80 --external-ip=$(minikube ip) --type=LoadBalancer
 
kubectl get pods -L run,version

kubectl get pod k8sqa-794f9b449c-hnbwg -o json

kubectl get pods -o=custom-columns="APP:.spec.containers[*].name,POD:.metadata.name,LABELS:.metadata.labels,IMAGE:.spec.containers[*].image,IP:.status.podIP"

kubectl get pods -l run=k8sjr

kubectl delete pod k8sjr-6d9fd6c5c-htm9f  

kubectl delete deployment k8sjr

kubectl delete deployments --all

kubectl delete services --all
```

## 2.2. Selectores y Outputs

Codigo fuente en ejecucion version [k8s-1.0-baby](https://repl.it/@twogg_git/k8s-10), [k8s-1.1-rolling](https://repl.it/@twogg_git/k8s-11)

```sh 
clear && kubectl get all

clear && kubectl get pods -o=custom-columns="APP:.spec.containers[*].name,POD:.metadata.name,LABELS:.metadata.labels,IMAGE:.spec.containers[*].image,IP:.status.podIP,PHASE:.status.phase"

kubectl run k8sgo --image=twogghub/k8s-workshop:1.0-baby --image-pull-policy=Always --replicas=3 --labels="deploy=baby"

kubectl expose deployment k8sgo --port=8080 --external-ip=$(minikube ip) --type=LoadBalancer

kubectl set image deployment k8sgo k8sgo=twogghub/k8s-workshop:1.1-rolling

kubectl delete pod k8sgo-84d9f99df8-d7trj

kubectl scale --replicas=5 deployment k8sgo

kubectl get pods,services,deployments --output wide

kubectl delete deployments,services,pods,replicasets --selector="deploy=baby"
``` 

## 2.3. Deplieges con archivos YAML

Fuente archivo YAML de despliegue [k8slatest.yaml](https://github.com/twogg-git/k8s-workshop/blob/master/yamls/k8slatest.yaml)

```sh 
apiVersion: extensions/v1beta1
kind: Pod
metadata:
  name: k8slatest
spec:
  replicas: 1
  template:  
    metadata:  
      labels:  
        env: k8slatest
    spec:
      containers:
        - image: twogghub/k8s-workshop:1.1-rolling
          imagePullPolicy: Always
          name: k8slatest
          ports:
          - name: http
            containerPort: 8080
``` 

```sh 
clear && kubectl get all

clear && kubectl get pods -o=custom-columns="APP:.spec.containers[*].name,POD:.metadata.name,LABELS:.metadata.labels,IMAGE:.spec.containers[*].image,IP:.status.podIP,PHASE:.status.phase"

vim deployment.yaml

kubectl create -f deployment.yaml

kubectl expose deployment k8slatest --port=8080 --external-ip=$(minikube ip) --type=LoadBalancer

kubectl describe deployment k8slatest

kubectl delete pods,services,deployments,replicaset --all
``` 

## 2.4. Politicas de despliegue

Codigo fuente en ejecucion version [k8s-1.2-yaml](https://repl.it/@twogg_git/k8s-12)

```sh
...
  strategy:
    type: RollingUpdate
    rollingUpdate:
      maxSurge: 1
      maxUnavailable: 1
      ...
```

```sh
...
  strategy:
    type: Recreate
    ...
```

```sh
clear && kubectl get all

clear && kubectl get pods -o=custom-columns="APP:.spec.containers[*].name,POD:.metadata.name,LABELS:.metadata.labels,IMAGE:.spec.containers[*].image,IP:.status.podIP,PHASE:.status.phase"

kubectl create -f https://raw.githubusercontent.com/twogg-git/k8s-workshop/master/yamls/k8sqa.yaml

kubectl expose deployment k8sqa --port=9090 --external-ip=$(minikube ip) --type=LoadBalancer

kubectl describe deployment k8sqa

kubectl create -f https://raw.githubusercontent.com/twogg-git/k8s-workshop/master/yamls/k8sprod.yaml

kubectl expose deployment k8sprod --port=8080 --external-ip=$(minikube ip) --type=LoadBalancer

kubectl get pods --output wide --show-labels --watch

kubectl set image deployment/k8sqa k8sqa=twogghub/k8s-workshop:1.1-qaonly

kubectl set image deployment/k8sprod k8sprod=twogghub/k8s-workshop:1.1-rolling

kubectl delete deployments,services,pods,replicasets --selector="env=qa"

kubectl scale --replicas=3 deployment k8sprod

kubectl set image deployment k8sprod k8sprod=twogghub/k8s-workshop:1.2-yaml

kubectl rollout undo deployment/k8sprod

kubectl delete deployments,services,pods,replicasets --selector="env=prod"
```

## 5. Validacion del despliegue mediante endpoints  

Codigo fuente en ejecucion version [k8s-1.3-livenes](https://repl.it/@twogg_git/k8s-13)

```sh
...
spec:
  ...
  template:  
    metadata:  
      ...
      annotations:
        kubernetes.io/change-cause: "HttpGet /health return error!" 
    spec:
      containers:
        - image: twogghub/k8s-workshop:1.3-liveness
          name: k8sdp
          livenessProbe:
            httpGet:
              path: /health
              port: 8080
            initialDelaySeconds: 5
            timeoutSeconds: 1
            periodSeconds: 2
            failureThreshold: 1
          ...
```

```sh
clear && kubectl get all

clear && kubectl get pods -o=custom-columns="APP:.spec.containers[*].name,POD:.metadata.name,LABELS:.metadata.labels,IMAGE:.spec.containers[*].image,IP:.status.podIP,PHASE:.status.phase"

kubectl create -f https://raw.githubusercontent.com/twogg-git/k8s-workshop/1.3-liveness/yamls/k8sdp.yaml

kubectl get pods -L env --output wide --watch

kubectl describe deployment k8sdp

kubectl expose deployment k8sdp --port=8080 --external-ip=$(minikube ip) --type=LoadBalancer

kubectl set image deployment k8sdp k8sdp=twogghub/k8s-workshop:1.2-yaml --record

kubectl rollout undo deployment/k8sdp

kubectl rollout history deployments k8sdp

kubectl delete pods,services,deployments,replicaset --all

```
