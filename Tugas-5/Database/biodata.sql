-- phpMyAdmin SQL Dump
-- version 4.8.5
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Sep 30, 2019 at 07:21 AM
-- Server version: 10.1.40-MariaDB
-- PHP Version: 7.3.5

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `bio`
--

-- --------------------------------------------------------

--
-- Table structure for table `biodata`
--

CREATE TABLE `biodata` (
  `NPM` varchar(12) COLLATE utf8_unicode_ci NOT NULL,
  `NAMA` char(50) COLLATE utf8_unicode_ci DEFAULT 'rather not to say',
  `EMAIL` char(50) COLLATE utf8_unicode_ci DEFAULT 'rather not to say',
  `PHONE` char(13) COLLATE utf8_unicode_ci DEFAULT 'None',
  `ALAMAT` char(100) COLLATE utf8_unicode_ci DEFAULT 'rather not to say',
  `GENERAL` text COLLATE utf8_unicode_ci,
  `ORG` char(100) COLLATE utf8_unicode_ci NOT NULL DEFAULT 'rather not to say',
  `ORG_DESC` text COLLATE utf8_unicode_ci
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

--
-- Dumping data for table `biodata`
--

INSERT INTO `biodata` (`NPM`, `NAMA`, `EMAIL`, `PHONE`, `ALAMAT`, `GENERAL`, `ORG`, `ORG_DESC`) VALUES
('17081010054', 'Fajar Andhika Putra', 'fajarhandika12@gmail.com', '085854574670', '', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Praesent posuere ornare ipsum, ut ornare diam mollis a. Mauris consectetur ornare viverra. Suspendisse potenti. Mauris scelerisque justo non libero accumsan, iaculis porttitor leo rutrum. Proin imperdiet nisl id tortor elementum facilisis. Nunc ornare risus eget ultrices dapibus. Integer sagittis nisi a molestie pharetra. Phasellus sit amet finibus purus.', 'FA\'s', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Praesent posuere ornare ipsum, ut ornare diam mollis a. Mauris consectetur ornare viverra. Suspendisse potenti. Mauris scelerisque justo non libero accumsan, iaculis porttitor leo rutrum. Proin imperdiet nisl id tortor elementum facilisis. Nunc ornare risus eget ultrices dapibus. Integer sagittis nisi a molestie pharetra. Phasellus sit amet finibus purus.');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `biodata`
--
ALTER TABLE `biodata`
  ADD PRIMARY KEY (`NPM`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
