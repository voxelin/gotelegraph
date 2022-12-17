#!/bin/bash -e
if [ -f /usr/bin/telegraph ]; then
    echo "telegraph is already installed..."
    echo "reinstalling..."
    sudo rm /usr/bin/telegraph
fi

# Installation
cd /usr/bin
sudo wget https://github.com/voxelin/gotelegraph/releases/download/v0.0.0/telegraph -q --show-progress
sudo chmod +x telegraph

# Check if installation was successful
if [ -f /usr/bin/telegraph ]; then
    echo "telegraph was successfully installed"
    exit 0
else
    echo "telegraph was not successfully installed"
    exit 1
fi