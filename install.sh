# Check if program is installed
if [ -f /usr/local/bin/telegraph ]; then
    echo "telegraph is already installed"
    exit 0
fi

# Installation
cd /usr/local/bin
wget https://github.com/voxelin/gotelegraph/releases/download/v0.0.0/telegraph

# Check if installation was successful
if [ -f /usr/local/bin/telegraph ]; then
    echo "telegraph was successfully installed"
    exit 0
else
    echo "telegraph was not successfully installed"
    exit 1
fi