# go-basic-2-webinar-10

Linux

IP сервера 194.190.152.152

apt install nginx
apt install golang-1.21-go
apt install certbot python3-certbot-nginx

git clone git@github.com:vmarunin/go-basic-2-webinar-10.git

sudo cp config/systemd/simpleback.service /etc/systemd/system/

systemctl status simpleback

journalctl -feu simpleback.service

Работает http://vmarunin2.viewdns.net/hello/

https://letsencrypt.org/

