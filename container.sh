docker run -d -p 80:8000 -e SERVER_PORT="8000" -e DB_USERNAME="admin" -e DB_PASSWORD="Heribudiyana11" -e DB_HOST="database-1.ck5ejk7nunv8.us-east-1.rds.amazonaws.com" -e DB_PORT="3306" -e DB_NAME="project" --name project heribudiyana/project:latest