DRO TABLE IF EXISTS `iotdata`;

CREATE TABLE `iotdata` (
    `no` int(11) unsigned NOT NULL AUTO_INCREMENT,
    `machineid` varchar(3) NOT NULL,
    `gettime` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `temp` varchar(10) DEFAULT NULL,
    `humid` varchar(10) DEFAULT NULL,
    `soil_humid` varchar(10) DEFAULT NULL,
    `co2` varchar(10) DEFAULT NULL,
    `wavelength` varchar(10) DEFAULT NULL,
    `illuminance` varchar(10) DEFAULT NULL,
    PRIMARY KEY (`no`)
)   ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `iotdata` WRITE;

UNLOCK TABLES;
