#!/bin/sh
sudo apt-get update
sudo apt-get install -y linux-generic-lts-vivid
wget -qO- https://test.docker.com/ | sh
sudo usermod -aG docker vagrant
