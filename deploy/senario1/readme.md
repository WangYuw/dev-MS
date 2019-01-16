./start_docker.sh
(Terminal 1) sudo docker exec -ti auth1 "/bin/bash"
  go test main_test.go 
  [ http.Get() failed with 'Get http://172.25.0.3:8080/services/all: dial tcp 172.25.0.3:8080: connect: connection refused' ] - change connection ip in src/main.go

(Terminal 2) sudo docker exec -ti reg1 "/bin/bash"
  [ client.Do() failed with 'Post http://auth1:8080/services/quality: dial tcp 172.25.0.2:8080: connect: connection refused' ] - change connection ip in src/main.go


* Deploy dev-AuthMS (an example of microservices)
  - cd ../../dev-AuthMS [ ok ]
  - create image : sudo docker build -t auth_img . [ ok ]
  - consult images' status : sudo docker images [ ok ]
  - create container : sudo docker run -d --name auth1 -p8080:8080 auth_img [ ok ]
  - consult containers' status : sudo docker ps -a [ ok ] (listen 8080)
  - go into container : sudo docker exec -ti auth1 "/bin/bash"
  - Delete containers : sudo docker rm auth1 [ ok ]
  - Delete images : sudo docker rmi auth_img [ ok ]
  
* Deploy dev-SR (service discovery)
  - cd ../../dev-SR [ ok ]
  - create image : sudo docker build -t reg_img . [ ok ]
  - consult images' status : sudo docker images [ ok ]
  - create container : sudo docker run -d --name sr1 -p8081:8080 reg_img [ ok ]
  - consult containers' status : sudo docker ps -a [ ok ] (listen 8080)
  - go into container : sudo docker exec -ti sr1 "/bin/bash"
  - Delete containers : sudo docker rm sr1 [ ok ]
  - Delete images : sudo docker rmi registry_img [ ok ]

create own bridge network : 
  sudo docker network create -d bridge --subnet 172.25.0.0/16 ms_network --attachable
  sudo docker run --network=ms_network --ip=172.25.0.2 -itd --name auth1 -p 8080:8080 auth_img

