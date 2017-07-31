CREATE TABLE `demspirals`.`players`
( 
    `pid` VARCHAR(255) NOT NULL UNIQUE,
    `slug` VARCHAR(255) NOT NULL ,
    `name` VARCHAR(200) NOT NULL ,
    `position` VARCHAR(100) NOT NULL ,
    `lastupdated` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`pid`(255))
) ENGINE = InnoDB COMMENT = 'Player information table';