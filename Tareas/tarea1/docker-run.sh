cd ./goiang-backend
docker build -t tarea1-backend .
docker run -d -p 8080:8080 tarea1-backend
cd ..
cd ./react-frontend
docker build -t tarea1-frontend .
docker run -d -p 80:80 tarea1-frontend
cd ..
