Requirements
 - go
 - docker
 - minikube
 - kubectl
 - hey


Create a binary from the go file
```
go build main.go
```
Run the go file wihtout creating a binary
```
go run main.go
```

Create a docker image and run it
```
docker images
docker build -t random-number-server .
docker run -d -p 8080:8080 random-number-server
docker ps
```

Start local kubernetes cluster
```
minikube start
minikube kubectl -- get pods -A
# In a different terminal
minikube dashboard
# In a different terminal
minikube tunnel
```

Deploy the service to the newly created cluster
```
eval $(minikube docker-env)
docker build -t random-number-server .
kubectl apply -f deployment.yaml
# To expose the deployed service (so we can call it from outside of the k8s cluster)
kubectl expose deployment random-number-server --type=LoadBalancer --port=8080
# To get the IP of the deployed service
minikube kubectl get svc
```


---

Test scalling
```
---
---
apiVersion: autoscaling/v2
kind: HorizontalPodAutoscaler
metadata:
  name: random-number-server-hpa
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: random-number-server
  minReplicas: 1
  maxReplicas: 10
  metrics:
  - type: Resource
    resource:
      name: cpu
      target:
        type: Utilization
        averageUtilization: 3

```


```
kubectl get hpa -w
hey -n 20000 -c 1000 http://10.99.39.200:8080/random
```
