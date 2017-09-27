## linux安装mysql
```sh
sudo apt-get update
sudo apt-get install mysql-server
sudo mysql_secure_installation
```

``` sh
https://www.digitalocean.com/community/tutorials/how-to-install-mysql-on-ubuntu-16-04?comment=53320
If it occurs error because of broken MySQL package on Ubuntu 16.04. Just do this trick

# Purge all MySQL packages
sudo apt purge mysql*
sudo rm -rf /var/lib/mysql
sudo rm -rf /etc/mysql

# Reinstall MySQL
sudo apt install mysql-server mysql-client
```

## 命令行启动mssql：
```sh
mysql -u root -p
```

## 参考资料：
- [How To Install MySQL on Ubuntu 16.04](error 1045 access denied for user using password no)
- https://stackoverflow.com/questions/10299148/mysql-error-1045-28000-access-denied-for-user-billlocalhost-using-passw
