DROP TABLE IF EXISTS `article`;
CREATE TABLE `article` (
  `Id` int(18) unsigned NOT NULL AUTO_INCREMENT,
  `Title` varchar(50) NOT NULL,
  `Content` varchar(50) NOT NULL,
  `Summary` varchar(360) NOT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;