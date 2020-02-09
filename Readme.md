# Web App
docker run --name=mysql -p 3306:3306 -d mysql/mysql-server:5.7
docker exec -it mysql mysql -uroot -p
docker logs mysql 2>&1 | grep GENERATED
docker exec -it mysql mysql -uroot --password="HiLIvUdzOHBEkeL3lIgOPahicT3m"
ALTER USER 'root'@'localhost' IDENTIFIED BY 'test@123';

Grant All Privileges ON *.* to 'root'@'%' Identified By 'test@123'; 
