CREATE TABLE `demspirals`.`teams`
( 
    `id` VARCHAR(255) NOT NULL UNIQUE,
    `slug` VARCHAR(255) NOT NULL ,
    `name` VARCHAR(100) NOT NULL ,
    `nickname` VARCHAR(100) NOT NULL ,
    `color` VARCHAR(20) NOT NULL ,
    `hashtag` VARCHAR(30) NOT NULL ,
    `lastupdated` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`(255))
) ENGINE = InnoDB COMMENT = 'Team information table';
