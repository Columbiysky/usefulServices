echo "apt update and upgrade"
apt update -y
apt upgrade -y

echo "install nodeJs 19 and dependencies for client app"
curl -fsSL https://deb.nodesource.com/setup_19.x | -E bash - && sudo apt-get install -y nodejs
cd ./usefulApps/frontend
npm i
cd ../..

echo "install python"
apt install python3 -y

ehco "install golang"
curl -O https://dl.google.com/go/go1.20.6.linux-amd64.tar.gz
tar -C /usr/local -xvf go1.20.6.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
go version
rm go1.20.6.linux-amd64.tar.gz