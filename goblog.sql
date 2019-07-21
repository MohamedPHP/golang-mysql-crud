-- phpMyAdmin SQL Dump
-- version 4.6.6deb5
-- https://www.phpmyadmin.net/
--
-- Host: localhost:3306
-- Generation Time: Jul 21, 2019 at 06:08 PM
-- Server version: 5.7.26-0ubuntu0.18.04.1
-- PHP Version: 7.1.30-1+ubuntu18.04.1+deb.sury.org+1

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `goblog`
--

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `email` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `password` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `rememberToken` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT '123',
  `createdAt` timestamp NULL DEFAULT NULL,
  `updatedAt` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `name`, `email`, `password`, `rememberToken`, `createdAt`, `updatedAt`) VALUES
(1, 'mohamed', 'mohamedzayed709@yahoo.com', '123123', '123', '2019-06-18 12:10:27', '2019-06-18 12:10:27'),
(2, 'Mohamed', 'mohamed@yahoo.com', '123123', '123123', '2019-07-04 22:00:00', '2019-07-04 22:00:00'),
(3, 'Mohamed', 'mohamed.zayed@uflare.io', '123456', '123', '2019-07-05 14:55:43', '2019-07-05 14:55:43'),
(6, 'Mohamed', 'mohamed.zayed@uflare.io', '123456', '123', '2019-07-05 14:59:33', '2019-07-05 14:59:33'),
(7, 'Mohamed Alaa', 'mohamedzayed709@yahoo.com', '123456', '123', '2019-07-05 14:59:49', '2019-07-05 14:59:49'),
(9, 'Ash3al', 'mohamed.zayed@uflare.io22', 'Ash3al123', '123', '2019-07-05 15:03:02', '2019-07-21 15:06:48');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=12;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
