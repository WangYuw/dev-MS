#!/bin/bash
sudo docker stop auth1 reg1 postgresql
sudo docker rm auth1 reg1 postgresql
sudo docker rmi auth_img reg_img postgres_img
sudo docker network rm ms_network

sudo docker build -t auth_img ./dev-AuthMS
sudo docker build -t reg_img ./dev-SR
sudo docker build -t postgres_img ./dev-DB/postgresql

sudo docker network create -d bridge --subnet 172.25.0.0/16 ms_network --attachable
sudo docker run --network=ms_network -itd --name auth1 -p 8081:8080 auth_img 
sudo docker run --network=ms_network -itd --name reg1 -p 8080:8080 reg_img
sudo docker run --network=ms_network -itd --name postgresql -p 5432:8080 postgres_img

sudo docker images

sudo docker ps -a