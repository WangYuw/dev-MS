#!/bin/bash
sudo docker stop auth1 reg1
sudo docker rm auth1 reg1
sudo docker rmi auth_img reg_img
sudo docker network rm ms_network