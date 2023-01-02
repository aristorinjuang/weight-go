CREATE TABLE IF NOT EXISTS `weights` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `max` float NOT NULL,
  `min` float NOT NULL,
  `date` date NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `date` (`date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;