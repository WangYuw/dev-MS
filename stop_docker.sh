#!/bin/bash
sudo docker stop auth1 reg1 pq_db
sudo docker rm auth1 reg1 pq_db
sudo docker rmi auth_img reg_img postgres_img
sudo docker network rm ms_network