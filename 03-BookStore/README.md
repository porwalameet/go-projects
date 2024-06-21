Prerequisites:
1. Install mysql on Ubuntu or WSL using the command
`sudo apt-get install mysql-server`
2. Start mysql service
`sudo service mysql start`
3. Run mysql security script and do the neccesary configuration.
`sudo mysql_secure_installation`
Refer: https://pen-y-fan.github.io/2021/08/08/How-to-install-MySQL-on-WSL-2-Ubuntu/
4. Create database and user which will be passed in to config/app.go to connect to do.
```
sudo mysql
create database books;
CREATE USER 'admin' IDENTIFIED WITH authentication_plugin BY 'admin123';
CREATE USER 'admin' IDENTIFIED BY 'admin123';
GRANT ALL PRIVILEGES ON *.* TO 'admin' WITH GRANT OPTION;
```

Then run the main go program
```
cd cmd/main
go run main.go
```