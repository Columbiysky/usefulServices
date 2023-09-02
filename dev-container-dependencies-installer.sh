echo "apt update and upgrade"
apt update -y
apt upgrade -y
echo "******************************"

echo "install nodeJs 20 and dependencies for client app"
sudo apt-get update
sudo apt-get install -y ca-certificates curl gnupg
sudo mkdir -p /etc/apt/keyrings
curl -fsSL https://deb.nodesource.com/gpgkey/nodesource-repo.gpg.key | sudo gpg --dearmor -o /etc/apt/keyrings/nodesource.gpg
NODE_MAJOR=20
echo "deb [signed-by=/etc/apt/keyrings/nodesource.gpg] https://deb.nodesource.com/node_$NODE_MAJOR.x nodistro main" | sudo tee /etc/apt/sources.list.d/nodesource.list
sudo apt-get update
sudo apt-get install nodejs -y
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