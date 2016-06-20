#!/bin/sh
wget -qO- https://test.docker.com/ | sh
sudo usermod -aG docker vagrant
