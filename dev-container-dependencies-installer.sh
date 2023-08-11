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

echo "install python"
apt install build-essential zlib1g-dev libncurses5-dev libgdbm-dev libnss3-dev libssl-dev libreadline-dev libffi-dev libsqlite3-dev wget libbz2-dev pkg-config -y
wget https://www.python.org/ftp/python/3.11.4/Python-3.11.4.tar.xz
tar -xvf Python-3.11.4.tar.xz
cd Python-3.11.4/
./configure --enable-optimizations
make -j 4
make altinstall
python3.11 --version
cd ..
echo "******************************"
echo "install pip"
curl -sS https://bootstrap.pypa.io/get-pip.py | python3.11 
pip3.11 -V
cd ./usefulApps/executors/python
pip3.11 install -r requirements.txt
echo "******************************"

rm Python-3.11.4.tar.xz
rm -rf Python-3.11.4
echo "******************************"

ehco "install golang"
curl -O https://dl.google.com/go/go1.21.0.linux-amd64.tar.gz
tar -C /usr/local -xvf go1.21.0.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
go version
rm go1.21.0.linux-amd64.tar.gz
echo "******************************"