# init
  gcloud init --skip-diagnostics

# kubecluster
gcloud container clusters create sopes1-proyecto2-cluster --zone us-central1-c --machine-type e2-medium --num-nodes 1

# build docker image
gcloud builds submit --tag gcr.io/[PROJECT_ID]/api
gcloud builds submit --tag gcr.io/[PROJECT_ID]/frontend

# deploy
gcloud run deploy api --image gcr.io/[PROJECT_ID]/api --platform managed
gcloud run deploy frontend --image gcr.io/[PROJECT_ID]/frontend --platform managed

# setenv
gcloud run deploy api --image gcr.io/[PROJECT_ID]/api --platform managed --set-env-vars MONGO_URI=mongodb://your-mongodb-ip:port,PORT=3000

# credentials
gcloud container clusters get-credentials NOMBRE_DEL_CLUSTER --zone ZONA_DEL_CLUSTER --project NOMBRE_DEL_PROYECTO

# ingress
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/cloud/deploy.yaml 
kubectl get pods -n ingress-nginx

# create image
docker build -t username/api .
docker push username/api

# create service
gcloud run deploy node-log-sopes1-proy2 --image=docker.io/xhuniktzi/node-log:1 --set-env-vars MONGO_URI="mongodb://104.197.89.206:27017" --platform managed --allow-unauthenticated
gcloud run deploy vue-log-sopes1-proy2 --image=docker.io/xhuniktzi/vue-log:1 --platform managed --allow-unauthenticated --port 80
