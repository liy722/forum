-- MySQL dump 10.13  Distrib 5.6.21, for Win64 (x86_64)
--
-- Host: localhost    Database: bbs
-- ------------------------------------------------------
-- Server version	5.6.21-log

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `subthread`
--

DROP TABLE IF EXISTS `subthread`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `subthread` (
  `stid` int(11) NOT NULL AUTO_INCREMENT COMMENT 'subthread id',
  `tid` int(11) DEFAULT NULL COMMENT 'thread id',
  `revert` varchar(200) DEFAULT NULL COMMENT 'reply',
  `ruid` int(11) DEFAULT NULL COMMENT 'id of reply user',
  `tuid` int(11) DEFAULT NULL COMMENT 'id of thread owner',
  PRIMARY KEY (`stid`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `subthread`
--

LOCK TABLES `subthread` WRITE;
/*!40000 ALTER TABLE `subthread` DISABLE KEYS */;
INSERT INTO `subthread` VALUES (7,3,'ab dd  ff',1,2),(8,3,'ab dd  ff',1,2),(9,3,'yy uu',1,2),(10,5,'aa  bb cc',1,3),(11,8,'abzz',1,4),(12,4,'zz',1,2),(13,3,'hh',1,2),(14,3,'abzz',1,2),(15,8,'oopp ppoo',1,4);
/*!40000 ALTER TABLE `subthread` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `thread`
--

DROP TABLE IF EXISTS `thread`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `thread` (
  `tid` int(11) NOT NULL AUTO_INCREMENT COMMENT 'thread id',
  `uid` int(11) DEFAULT NULL COMMENT 'id of user',
  `topic` text COMMENT 'topic',
  PRIMARY KEY (`tid`)
) ENGINE=InnoDB AUTO_INCREMENT=76 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `thread`
--

LOCK TABLES `thread` WRITE;
/*!40000 ALTER TABLE `thread` DISABLE KEYS */;
INSERT INTO `thread` VALUES (3,2,'test21'),(4,2,'test22'),(5,3,'test31'),(6,3,'test32'),(7,4,'test41'),(8,4,'test42'),(9,5,'test51'),(68,1,'dd'),(71,1,'jjkk'),(72,1,'hello'),(73,1,'hello2'),(74,1,'hello3');
/*!40000 ALTER TABLE `thread` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `users` (
  `uid` int(11) NOT NULL AUTO_INCREMENT COMMENT 'user id',
  `email` varchar(100) DEFAULT NULL COMMENT 'email',
  `username` varchar(100) DEFAULT NULL COMMENT 'username',
  PRIMARY KEY (`uid`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'liy@qq.com','liy'),(2,'lcw@qq.com','lcw'),(3,'xiaoqiu@qq.com','xiaoqiu'),(4,'value','test01'),(5,'test02@qq.com','test02'),(6,'test06@qq.com','test06'),(7,'ss','fff'),(8,'lzj@qq.com','lzj'),(9,'bb','aa'),(10,'pp','kk');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-01-23 18:59:43
