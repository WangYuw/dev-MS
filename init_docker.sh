sudo docker exec -it pq_db psql --command "ALTER USER postgres WITH ENCRYPTED PASSWORD 'postgres';"
sudo docker exec -it pq_db createdb registry
sudo docker exec -it pq_db psql -U postgres -f /sql/init.sql -d registry
sudo docker exec -ti auth1 go run test/main.go
sudo docker exec -ti reg1 go run test/main.go