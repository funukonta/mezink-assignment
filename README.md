# Mezink-Assignment
This Assignment uses stdlib from golang to develop REST API

## External library used :
- mux
- pq
- godotenv

## How To Set Up:
 1. Clone the Repo `git clone https://github.com/funukonta/mezink-assignment.git`
 2. Download docker images for postgresDb `docker run --name some-postgres --env-file .env -d   postgres`
 3. Build docker image from Dockerfile `docker build --tag docker-mezink .`
 4. Run the API Server `docker run -p 8080:8080 docker-mezink`
 5. Test endpoint "/getData" with method `GET` with command
 `curl -X GET -H "Content-Type: application/json" -d '{
  "startDate": "2024-01-01",
  "endDate": "2024-01-09",
  "minCount": 100,
  "maxCount": 300
}' http://localhost:8080/getData`
