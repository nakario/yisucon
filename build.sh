cd /var/www/webapp
export GOPATH=/var/www/webapp
#ln -sfv go src
cd /var/www/webapp/go/isutomo
glide install
go build
cd /var/www/webapp/go/isuwitter
glide install
go build
