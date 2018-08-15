# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 5.7.22)
# Database: test
# Generation Time: 2018-08-15 09:17:06 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table weather_day_tomorrow
# ------------------------------------------------------------

DROP TABLE IF EXISTS `weather_day_tomorrow`;

CREATE TABLE `weather_day_tomorrow` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `header_id` int(11) DEFAULT NULL,
  `day` varchar(30) DEFAULT NULL,
  `temperature` varchar(10) DEFAULT NULL,
  `img` varchar(5) DEFAULT NULL,
  `situation` varchar(100) DEFAULT NULL,
  `desc` varchar(100) DEFAULT NULL,
  `rain` varchar(5) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table weather_header
# ------------------------------------------------------------

DROP TABLE IF EXISTS `weather_header`;

CREATE TABLE `weather_header` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '流水號',
  `current_temperature` varchar(6) DEFAULT NULL,
  `cwb_update_time` datetime DEFAULT NULL,
  `json_update_time` datetime DEFAULT NULL,
  `information` varchar(300) DEFAULT NULL COMMENT '天氣消息',
  `city_name` varchar(20) DEFAULT NULL COMMENT '城市名稱',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table weather_week
# ------------------------------------------------------------

DROP TABLE IF EXISTS `weather_week`;

CREATE TABLE `weather_week` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `header_id` int(11) NOT NULL,
  `day` varchar(10) DEFAULT NULL,
  `morning_temperature` varchar(10) DEFAULT NULL,
  `morning_situation` varchar(30) DEFAULT NULL,
  `morning_img` varchar(5) DEFAULT NULL,
  `night_temperature` varchar(10) DEFAULT NULL,
  `nitght_situation` varchar(30) DEFAULT NULL,
  `night_img` varchar(5) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;




/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
