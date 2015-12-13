/*
MARIADB Backup
Source Server Version: 10.1.9
Source Database: db_injuv_backoffice
Date: 12/13/2015 14:42:02
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
--  Table structure for `adjuntos`
-- ----------------------------
DROP TABLE IF EXISTS `adjuntos`;
CREATE TABLE `adjuntos` (
  `id` varchar(25) NOT NULL,
  `idTarea` varchar(25) NOT NULL,
  `nombre` varchar(45) NOT NULL,
  `extension` varchar(45) NOT NULL,
  `fechaCreacion` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `comentarios`
-- ----------------------------
DROP TABLE IF EXISTS `comentarios`;
CREATE TABLE `comentarios` (
  `id` int(25) NOT NULL AUTO_INCREMENT,
  `fechaCreacion` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `ultimaModificacion` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `comentarioscol` varchar(45) NOT NULL,
  `tareas_id` int(22) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `id_tarea` (`tareas_id`),
  CONSTRAINT `id_tarea` FOREIGN KEY (`tareas_id`) REFERENCES `tareas` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `estado`
-- ----------------------------
DROP TABLE IF EXISTS `estado`;
CREATE TABLE `estado` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `nombre` varchar(45) NOT NULL,
  `value` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `miembros`
-- ----------------------------
DROP TABLE IF EXISTS `miembros`;
CREATE TABLE `miembros` (
  `id` varchar(25) NOT NULL,
  `idTarea` varchar(25) NOT NULL,
  `miembroscol` varchar(45) DEFAULT NULL,
  `Usuario_id` varchar(25) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `pioridad`
-- ----------------------------
DROP TABLE IF EXISTS `pioridad`;
CREATE TABLE `pioridad` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `nombre` varchar(45) DEFAULT NULL,
  `color` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `tareas`
-- ----------------------------
DROP TABLE IF EXISTS `tareas`;
CREATE TABLE `tareas` (
  `id` int(22) NOT NULL AUTO_INCREMENT,
  `titulo` varchar(45) NOT NULL,
  `descripcion` varchar(140) NOT NULL,
  `fechaTermino` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `miembros_id` int(22) NOT NULL,
  `estado` varchar(45) NOT NULL,
  `adjuntos_id` varchar(45) NOT NULL,
  `pioridad` varchar(45) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=40 DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `usuario`
-- ----------------------------
DROP TABLE IF EXISTS `usuario`;
CREATE TABLE `usuario` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `nombre` varchar(45) NOT NULL,
  `apellidos` varchar(45) NOT NULL,
  `userName` varchar(45) NOT NULL,
  `email` varchar(45) NOT NULL,
  `contrasena` varchar(45) NOT NULL,
  `fechaCreacion` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `admin` tinyint(1) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

-- ----------------------------
--  Records 
-- ----------------------------
INSERT INTO `comentarios` VALUES ('1','2015-12-12 18:38:59','2015-12-12 18:38:59','hola','2'), ('2','2015-12-30 22:21:00','2015-12-30 22:21:00','hola Insert correcto','2'), ('3','2015-12-30 22:21:00','2015-12-30 22:21:00','hola Insert correcto','2');
INSERT INTO `estado` VALUES ('1','Impedidas','0'), ('2','Pendientes','0'), ('3','EnProceso','0'), ('4','Terminadas','0');
INSERT INTO `tareas` VALUES ('1','example','example','2015-12-30 22:21:00','2','pendiente','1','alta'), ('2','titulo example','descripcion example','2015-12-30 22:21:00','2','impedida','0','alta'), ('18','testing05','testing05- descripcion','0000-00-00 00:00:00','1','Pendientes','0','Alta'), ('19','testing05','testing05- descripcion','0000-00-00 00:00:00','1','Pendientes','0','Alta'), ('20','testing-06','testing-06 description','0000-00-00 00:00:00','1','Pendientes','0','Media'), ('22','testing-07','testing-07description','0000-00-00 00:00:00','1','Pendientes','0','Media'), ('23','test-08','test-08 - description','0000-00-00 00:00:00','1','EnProceso','0','Media'), ('24','dasdasdas','lashdsalfdhsadlkf','0000-00-00 00:00:00','1','Pendientes','0','Media'), ('25','dasdasdas','lashdsalfdhsadlkf','0000-00-00 00:00:00','1','Pendientes','0','Media'), ('26','test-10','test-10    -  description','0000-00-00 00:00:00','1','Pendientes','0','Media'), ('27','test-11','test-11    -  description','0000-00-00 00:00:00','1','Pendientes','0','Alta'), ('28','test-12','test-12   description','0000-00-00 00:00:00','1','EnProceso','0','Media'), ('29','test-13','test-13 description','0000-00-00 00:00:00','1','EnProceso','0','Baja'), ('30','test-15','test-15 description','0000-00-00 00:00:00','2','Pendientes','0','Media'), ('31','test-16','test-16 description','0000-00-00 00:00:00','1','Pendientes','0','Media'), ('32','test-17','test-17 description','0000-00-00 00:00:00','1','Pendientes','0','Baja'), ('33','test-17','test-17- descriptrion','0000-00-00 00:00:00','1','Pendientes','0','Alta'), ('34','test20','test-20','0000-00-00 00:00:00','1','Pendientes','0','Alta'), ('35','test-23','test-23 description','0000-00-00 00:00:00','1','Pendientes','0','Media'), ('36','test-nn','test-nn description','0000-00-00 00:00:00','1','Pendientes','0','Media'), ('37','test-22','test-22 description','0000-00-00 00:00:00','1','Terminados','0','Baja'), ('38','tets23-','test-23 descrip','2015-12-09 19:00:00','1','Terminados','0','Baja'), ('39','test30','testr30','2015-12-25 19:00:00','1','Terminados','0','Baja');
INSERT INTO `usuario` VALUES ('1','ignacio','carvajal','ic','ignaciocarvajald@gmail.com','123456','2015-12-12 23:48:16','1'), ('2','juan','perez','@jp','jp@injuv.cl','123456','2015-12-06 18:55:22','0');
