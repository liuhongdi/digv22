-- phpMyAdmin SQL Dump
-- version 5.0.4
-- https://www.phpmyadmin.net/
--
-- 主机： localhost
-- 生成日期： 2021-01-12 18:51:52
-- 服务器版本： 8.0.22-0ubuntu0.20.10.2
-- PHP 版本： 7.4.9

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- 数据库： `business`
--

-- --------------------------------------------------------

--
-- 表的结构 `m_goods`
--

CREATE TABLE `m_goods` (
  `goodsId` bigint UNSIGNED NOT NULL COMMENT '商品id',
  `goodsName` varchar(200) NOT NULL DEFAULT '' COMMENT '商品名称',
  `stock` int UNSIGNED NOT NULL DEFAULT '0' COMMENT '库存数量'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='商品表';

--
-- 转存表中的数据 `m_goods`
--

INSERT INTO `m_goods` (`goodsId`, `goodsName`, `stock`) VALUES
(1, '施丹兰蜂蜜牛奶手工皂', 10);

-- --------------------------------------------------------

--
-- 表的结构 `m_order`
--

CREATE TABLE `m_order` (
  `orderId` bigint UNSIGNED NOT NULL COMMENT '订单id',
  `userId` bigint UNSIGNED NOT NULL DEFAULT '0' COMMENT '用户id',
  `salePrice` decimal(10,0) NOT NULL DEFAULT '0' COMMENT '订单金额'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='订单表';

-- --------------------------------------------------------

--
-- 表的结构 `m_order_goods`
--

CREATE TABLE `m_order_goods` (
  `ogId` bigint UNSIGNED NOT NULL COMMENT 'id',
  `orderId` bigint UNSIGNED NOT NULL COMMENT '订单id',
  `goodsId` bigint UNSIGNED NOT NULL COMMENT '商品id',
  `buyNum` int UNSIGNED NOT NULL COMMENT '购买数量'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='订单商品表';

-- --------------------------------------------------------
--
-- 转储表的索引
--

--
-- 表的索引 `m_goods`
--
ALTER TABLE `m_goods`
  ADD PRIMARY KEY (`goodsId`);

--
-- 表的索引 `m_order`
--
ALTER TABLE `m_order`
  ADD PRIMARY KEY (`orderId`),
  ADD KEY `userId` (`userId`);

--
-- 表的索引 `m_order_goods`
--
ALTER TABLE `m_order_goods`
  ADD PRIMARY KEY (`ogId`),
  ADD UNIQUE KEY `orderId_2` (`orderId`,`goodsId`);


--
-- 在导出的表使用AUTO_INCREMENT
--

--
-- 使用表AUTO_INCREMENT `m_goods`
--
ALTER TABLE `m_goods`
  MODIFY `goodsId` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '商品id', AUTO_INCREMENT=2;

--
-- 使用表AUTO_INCREMENT `m_order`
--
ALTER TABLE `m_order`
  MODIFY `orderId` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '订单id';

--
-- 使用表AUTO_INCREMENT `m_order_goods`
--
ALTER TABLE `m_order_goods`
  MODIFY `ogId` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'id';



/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
