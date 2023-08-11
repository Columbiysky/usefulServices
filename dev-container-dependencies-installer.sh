echo "apt update and upgrade"
apt update -y
apt upgrade -y
echo "******************************"

echo "install nodeJs 20 and dependencies for client app"
curl -fsSL https://deb.nodesource.com/setup_20.x | sudo -E bash - && sudo apt-get install -y nodejs
cd ./usefulApps/frontend
npm i
cd ../..
node --version
echo "******************************"

echo "install python, pip and dependencies"
apt install -y python3-pip
pip3 install -r ./usefulApps/executors/python/requirements.txt
echo "******************************"

ehco "install golang"
curl -O https://dl.google.com/go/go1.21.0.linux-amd64.tar.gz
tar -C /usr/local -xvf go1.21.0.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
go version
rm go1.21.0.linux-amd64.tar.gz
echo "******************************"