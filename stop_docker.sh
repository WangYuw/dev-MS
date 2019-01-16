#!/bin/bash
sudo docker stop auth1
sudo docker stop reg1
sudo docker rm auth1
sudo docker rm reg1
sudo docker rmi auth_img
sudo docker rmi reg_img
sudo docker network rm ms_network