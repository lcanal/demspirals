CREATE TABLE `demspirals`.`playerstats` 
( 
    `pid` VARCHAR(255) NOT NULL UNIQUE , 
    `runs` INT NULL DEFAULT NULL , 
    `passes` INT NULL DEFAULT NULL , 
    `receptions` INT NULL , 
    `lastupdated` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`pid`(255))
) ENGINE = InnoDB COMMENT = 'Player collected stats';
