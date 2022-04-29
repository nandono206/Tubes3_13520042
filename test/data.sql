-- MariaDB dump 10.19  Distrib 10.6.5-MariaDB, for Win64 (AMD64)
--
-- Host: localhost    Database: datatubes3stima
-- ------------------------------------------------------
-- Server version	10.6.5-MariaDB

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `diseases`
--

DROP TABLE IF EXISTS `diseases`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `diseases` (
  `nama` varchar(191) NOT NULL,
  `dna_seq` varchar(191) DEFAULT NULL,
  PRIMARY KEY (`nama`),
  UNIQUE KEY `dna_seq` (`dna_seq`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `diseases`
--

LOCK TABLES `diseases` WRITE;
/*!40000 ALTER TABLE `diseases` DISABLE KEYS */;
INSERT INTO `diseases` VALUES ('dead','AAAAAAAAAAAAAAAAAA'),('kurang tidur2','AAAAAAAAAAAAAAAAAATTTTTTTTTTTTTTTTTTTTTTTTTTTT'),('DEad3','AAAAAAAAAAAAAAAAAATTTTTTTTTTTTTTTTTTTTTTTTTTTTGGGGGGGGG'),('dead5','AAAAAAAATGTGTGT'),('newdead','AAAAAGGGGCCCTTCCCCC'),('flooooj','AAAAAGGGGGGGTC'),('dead2','AAAAAGGGGGGGTCC'),('flooosoj','AAAAAGGGGGGT'),('FlueH','AAAGGGTTTCCCT'),('sdfasfasdfaef','AAGGTTGGAAA'),('HOO','AAGTCTGTGGGT'),('NEW','AAGTCTTGGGGTGGGTTTTT'),('HIUHUHUHHUH','AATATATATATTATAATATATTATATTAGATAGAT'),('fsdfasfsadfefsfdae','ATTAAG'),('Deadliner','GGGGGGGGGG'),('kurang tidur','TTTAAGG'),('HIHIHIllll','TTTAAGGGATTTGAGTAGATAG'),('asfdfrg','TTTAAGGGATTTGAGTAGATAGTTGAAA');
/*!40000 ALTER TABLE `diseases` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `records`
--

DROP TABLE IF EXISTS `records`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `records` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `tanggal` longtext DEFAULT NULL,
  `nama_pasien` longtext DEFAULT NULL,
  `nama` longtext DEFAULT NULL,
  `status` longtext DEFAULT NULL,
  `presentase` double DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `records`
--

LOCK TABLES `records` WRITE;
/*!40000 ALTER TABLE `records` DISABLE KEYS */;
INSERT INTO `records` VALUES (1,'2022-04-29','IKMAL','dead','True',100),(2,'2022-04-29','nanod','dead','False',44.44444444444444),(3,'2022-04-29','SD','dead3','False',21.818181818181817),(4,'2022-04-29','dead2','dead','False',44.44444444444444),(5,'2022-04-29','dead2','dead','True',100),(6,'2022-04-29','NANO','dead','False',44.44444444444444),(7,'2022-04-29','Nando','DEAD','True',100),(8,'2022-04-29','Nando','DEAD','False',27.77777777777778);
/*!40000 ALTER TABLE `records` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-04-29 23:11:21
