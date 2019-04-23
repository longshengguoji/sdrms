/*
SQLyog v10.2 
MySQL - 5.6.14 : Database - sdrms
*********************************************************************
*/


/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
/*CREATE DATABASE IF NOT EXISTS `orderdb` DEFAULT CHARACTER SET utf8 */;

USE orderdb;

/*Table structure for table `rms_backend_user' */

DROP TABLE IF EXISTS tb_backend_user;

CREATE TABLE tb_backend_user (
  id serial primary key ,
  real_name varchar(255) NOT NULL DEFAULT '',
  user_name varchar(255) NOT NULL DEFAULT '',
  user_pwd varchar(255) NOT NULL DEFAULT '',
  is_super int NOT NULL DEFAULT 0,
  status int NOT NULL DEFAULT 0,
  mobile varchar(16) NOT NULL DEFAULT '',
  email varchar(256) NOT NULL DEFAULT '',
  avatar varchar(256) NOT NULL DEFAULT ''
);

/*Data for the table tb_backend_user */

insert  into tb_backend_user(id,real_name,user_name,user_pwd,is_super,status,mobile,email,avatar) values (1,'zhuwei','admin','123456',0,1,'18612348765','lhtzbj18@126.com','/static/upload/1.jpg');

/*Table structure for table 'tb_backend_user_rms_roles' */

DROP TABLE IF EXISTS tb_backend_user_tb_roles;

CREATE TABLE rb_backend_user_tb_roles (
  id serial primary key ,
  tb_backend_user_id int NOT NULL,
  tb_role_id int NOT NULL
);

/*Data for the table tb_backend_user_tb_roles */

/*Table structure for table tb_resource */

DROP TABLE IF EXISTS tb_resource;

CREATE TABLE tb_resource (
  id serial primary key,
  rtype int NOT NULL DEFAULT 0,
  name varchar(64) NOT NULL DEFAULT '',
  parent_id int DEFAULT NULL,
  seq int NOT NULL DEFAULT 0,
  icon varchar(32) NOT NULL DEFAULT '',
  url_for varchar(256) NOT NULL DEFAULT ''
);
insert  into tb_resource(id,rtype,name,parent_id,seq,icon,url_for) values (7,1,'面料',NULL,100,'fa fa-balance-scale',''),(12,1,'内销',7,100,'','RoleController.Index'),(13,1,'外销',7,100,'','BackendUserController.Index'),(30,2,'编辑',12,100,'fa fa-pencil','RoleController.Edit'),(31,2,'删除',12,100,'fa fa-trash','RoleController.Delete'),(32,2,'分配资源',12,100,'fa fa-th','RoleController.Allocate'),(33,2,'编辑',13,100,'fa fa-pencil','RoleController.Edit'),(34,2,'删除',13,100,'fa fa-trash','RoleController.Delete'),(35,2,'分配资源',13,100,'fa fa-th','RoleController.Allocate');

/*Table structure for table 'tb_role' */
DROP TABLE IF EXISTS tb_role;

CREATE TABLE tb_role (
  id serial primary key,
  name varchar(255) NOT NULL DEFAULT '',
  seq int NOT NULL DEFAULT 0
);

/*Data for the table tb_role */

insert  into tb_role(id,name,seq) values (22,'超级管理员',20),(24,'角色管理员',10),(25,'课程资源管理员',5);

/*Table structure for table 'tb_role_backenduser_rel' */

DROP TABLE IF EXISTS tb_role_backenduser_rel;

CREATE TABLE tb_role_backenduser_rel (
  id serial primary key,
  role_id int NOT NULL,
  backend_user_id int NOT NULL,
  created timestamp NOT NULL
);

/*Data for the table tb_role_backenduser_rel */

insert  into tb_role_backenduser_rel(id,role_id,backend_user_id,created) values (61,22,1,'2017-12-18 07:35:58'),(66,25,3,'2017-12-19 06:40:34');

/*Table structure for table 'tb_role_resource_rel' */

DROP TABLE IF EXISTS tb_role_resource_rel;

CREATE TABLE tb_role_resource_rel (
  id serial primary key,
  role_id int NOT NULL,
  resource_id int NOT NULL,
  created timestamp NOT NULL
);

/*Data for the table tb_role_resource_rel */

insert  into tb_role_resource_rel(id,role_id,resource_id,created) values (429,25,21,'2017-12-19 06:40:04'),(430,25,22,'2017-12-19 06:40:04'),(431,22,35,'2017-12-19 06:40:07'),(432,22,21,'2017-12-19 06:40:07'),(433,22,22,'2017-12-19 06:40:07'),(434,22,8,'2017-12-19 06:40:07'),(435,22,14,'2017-12-19 06:40:07'),(436,22,23,'2017-12-19 06:40:07'),(437,22,7,'2017-12-19 06:40:07'),(438,22,9,'2017-12-19 06:40:07'),(439,22,25,'2017-12-19 06:40:07'),(440,22,27,'2017-12-19 06:40:07'),(441,22,12,'2017-12-19 06:40:07'),(442,22,30,'2017-12-19 06:40:07'),(443,22,31,'2017-12-19 06:40:07'),(444,22,32,'2017-12-19 06:40:07'),(445,22,13,'2017-12-19 06:40:07'),(446,22,26,'2017-12-19 06:40:07'),(447,22,29,'2017-12-19 06:40:07'),(448,24,8,'2017-12-19 06:40:16'),(449,24,14,'2017-12-19 06:40:16'),(450,24,23,'2017-12-19 06:40:16');


DROP TABLE IF EXISTS tb_fabric_list;
CREATE TABLE tb_fabric_list (
  id serial PRIMARY KEY,
  label varchar(255),
	catalogue varchar(255) NOT NULL,
	sub_catalogue varchar(255) NOT NULL,
	bar_code varchar(255) NOT NULL,
	chinese_name varchar(255) NOT NULL,
	sample_number varchar(255) NOT NULL,
	description varchar(255),
	color_pattern varchar(255) NOT NULL,
	style_number varchar(255),
	specification varchar(255) NOT NULL,
	greige_supplier varchar(255),
	product_category varchar(255) NOT NULL,
	supplier varchar(255) NOT NULL,
	unit_price varchar(255) NOT NULL,
	element varchar(255) NOT NULL,
	weight varchar(255)
);


insert into tb_fabric_list values(1001,null,'面料','面料','12050033','珍珠梭','SC-S-12-0000001',null,'绿色',null,'15.59MM',null,'梭织','客供','元/米','100%丝',null);
insert into tb_fabric_list values(1002,null,'面料','面料','22050033','珍珠梭','SC-S-13-0000001',null,'绿色',null,'15.59MM',null,'梭织','客供','元/米','100%丝',null);
insert into tb_fabric_list values(1003,null,'面料','面料','22050034','珍珠梭','SC-S-13-0000001',null,'绿色',null,'15.59MM',null,'机织','天供','元/米','100%丝',null);



