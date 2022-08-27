Use project;

CREATE TABLE `demo_db` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `demo_db` varchar(50) NOT NULL,
  `deleted_at` timestamp NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

INSERT INTO `demo_db`(`id`, `demo_db`) VALUES (1,'demo');
