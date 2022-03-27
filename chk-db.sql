CREATE TABLE `chk`.`price_histories`
(
    `id`     bigint NOT NULL AUTO_INCREMENT,
    `unix`   bigint      DEFAULT NULL,
    `symbol` varchar(45) DEFAULT NULL,
    `open`   double      DEFAULT NULL,
    `high`   double      DEFAULT NULL,
    `low`    double      DEFAULT NULL,
    `close`  double      DEFAULT NULL,
    PRIMARY KEY (`id`),
    INDEX    `symbol` USING BTREE (`symbol` ASC),
    INDEX    `unix` (`unix` ASC)
) ENGINE=InnoDB DEFAULT CHARSET=utf8
