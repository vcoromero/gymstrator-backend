CREATE DATABASE IF NOT EXISTS gymstration DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;
USE gymstration;
GRANT ALL PRIVILEGES ON gymstration.* TO 'gymstrator'@'%';
FLUSH PRIVILEGES;