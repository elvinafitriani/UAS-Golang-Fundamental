-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Dec 08, 2022 at 06:38 PM
-- Server version: 10.4.22-MariaDB
-- PHP Version: 8.0.14

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `akbar_futsal`
--

-- --------------------------------------------------------

--
-- Table structure for table `bookings`
--

CREATE TABLE `bookings` (
  `id_booking` bigint(20) NOT NULL,
  `nama_team` longtext NOT NULL,
  `no_hp` longtext NOT NULL,
  `no_lap` bigint(20) NOT NULL,
  `tanggal` longtext NOT NULL,
  `harga_lap` bigint(20) DEFAULT NULL,
  `tanggal_main` datetime(3) NOT NULL,
  `tanggal_pesan` datetime(3) NOT NULL,
  `kode_transaksi` bigint(20) DEFAULT NULL,
  `dp_status` longtext NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `bookings`
--

INSERT INTO `bookings` (`id_booking`, `nama_team`, `no_hp`, `no_lap`, `tanggal`, `harga_lap`, `tanggal_main`, `tanggal_pesan`, `kode_transaksi`, `dp_status`) VALUES
(308, 'Saintama', '08938329909', 1, '2022-12-21 8:00 AM', 200000, '2022-12-21 08:00:00.000', '2022-12-08 17:28:51.185', 8, 'Sudah Dibayar'),
(309, 'DarmaFc', '08938329909', 1, '2022-12-21 9:00 AM', 200000, '2022-12-21 09:00:00.000', '2022-12-08 17:29:17.041', 9, 'Sudah Dibayar'),
(310, 'DarmaFc', '08938329909', 1, '2022-12-21 10:00 AM', 200000, '2022-12-21 10:00:00.000', '2022-12-08 17:29:27.969', 10, 'Sudah Dibayar'),
(333, 'Ciffo', '0893832932', 1, '2022-12-22 8:00 AM', 200000, '2022-12-22 08:00:00.000', '2022-12-08 17:18:50.675', 33, 'Sudah Dibayar'),
(334, 'IwanFc', '0893832932', 2, '2022-12-22 8:00 AM', 110000, '2022-12-22 08:00:00.000', '2022-12-08 17:20:02.956', 34, 'Sudah Dibayar'),
(335, 'Ciffo', '0893832932', 1, '2022-12-22 9:00 AM', 200000, '2022-12-22 09:00:00.000', '2022-12-08 17:20:22.710', 35, 'Sudah Dibayar'),
(336, 'Ciffo', '0893832932', 2, '2022-12-22 10:00 AM', 110000, '2022-12-22 10:00:00.000', '2022-12-08 17:20:37.563', 36, 'Sudah Dibayar'),
(337, 'Maranata', '0893832932', 3, '2022-12-22 12:00 AM', 90000, '2022-12-22 12:00:00.000', '2022-12-08 17:21:24.486', 37, 'Sudah Dibayar'),
(338, 'Ciffo', '0893832932', 1, '2022-12-22 12:00 AM', 200000, '2022-12-22 12:00:00.000', '2022-12-08 17:22:56.177', 38, 'Sudah Dibayar'),
(339, 'Maranata', '0893832932', 2, '2022-12-23 12:00 AM', 110000, '2022-12-23 12:00:00.000', '2022-12-08 17:23:11.521', 39, 'Sudah Dibayar'),
(340, 'CardiFc', '0893832932', 1, '2022-12-23 8:00 AM', 200000, '2022-12-23 08:00:00.000', '2022-12-08 17:11:49.383', 40, 'Sudah Dibayar');

-- --------------------------------------------------------

--
-- Table structure for table `lapangans`
--

CREATE TABLE `lapangans` (
  `no_lapangan` bigint(20) NOT NULL,
  `harga` bigint(20) NOT NULL,
  `images` longtext NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `lapangans`
--

INSERT INTO `lapangans` (`no_lapangan`, `harga`, `images`) VALUES
(1, 200000, 'Gambar/image1.jpg'),
(2, 110000, 'Gambar/image2.jpg'),
(3, 90000, 'Gambar/image3.jpg\r\n');

-- --------------------------------------------------------

--
-- Table structure for table `riwayats`
--

CREATE TABLE `riwayats` (
  `id_booking` bigint(20) NOT NULL,
  `nama_team` longtext NOT NULL,
  `no_hp` longtext NOT NULL,
  `no_lap` bigint(20) NOT NULL,
  `tanggal` longtext NOT NULL,
  `harga_lap` bigint(20) DEFAULT NULL,
  `tanggal_main` datetime(3) NOT NULL,
  `tanggal_pesan` datetime(3) NOT NULL,
  `kode_transaksi` bigint(20) DEFAULT NULL,
  `dp_status` longtext NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `riwayats`
--

INSERT INTO `riwayats` (`id_booking`, `nama_team`, `no_hp`, `no_lap`, `tanggal`, `harga_lap`, `tanggal_main`, `tanggal_pesan`, `kode_transaksi`, `dp_status`) VALUES
(301, 'Carban', '0893832932', 1, '2022-12-23 8:00 PM', 200000, '2022-12-23 20:00:00.000', '0000-00-00 00:00:00.000', 1, 'Sudah Dibayar'),
(302, 'Gibran', '089383292', 1, '2022-12-23 9:00 PM', 200000, '2022-12-23 21:00:00.000', '0000-00-00 00:00:00.000', 2, 'Sudah Dibayar'),
(303, 'Magenta', '08938329232', 1, '2022-12-23 10:00 AM', 200000, '2022-12-23 10:00:00.000', '0000-00-00 00:00:00.000', 3, 'Sudah Dibayar\r\n'),
(304, 'Krisna', '08938329232', 1, '2022-12-20 11:00 AM', 200000, '2022-12-20 11:00:00.000', '0000-00-00 00:00:00.000', 4, 'Sudah Dibayar'),
(305, 'Tisna', '08938329232', 1, '2022-12-20 12:00 AM', 200000, '2022-12-20 12:00:00.000', '0000-00-00 00:00:00.000', 5, 'Sudah Dibayar'),
(306, 'EntisTisna', '08938329232', 1, '2022-12-20 8:00 AM', 200000, '2022-12-20 08:00:00.000', '0000-00-00 00:00:00.000', 6, 'Sudah Dibayar'),
(307, 'EntisTisnaSasmita', '0893832934', 1, '2022-12-20 9:00 AM', 200000, '2022-12-20 09:00:00.000', '0000-00-00 00:00:00.000', 7, 'Sudah Dibayar');

-- --------------------------------------------------------

--
-- Table structure for table `transaksis`
--

CREATE TABLE `transaksis` (
  `id_transaksi` bigint(20) NOT NULL,
  `no_lap` bigint(20) NOT NULL,
  `dp_status` longtext NOT NULL,
  `tanggal_transaksi` datetime(3) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `transaksis`
--

INSERT INTO `transaksis` (`id_transaksi`, `no_lap`, `dp_status`, `tanggal_transaksi`) VALUES
(1, 1, 'Sudah Dibayar', '2022-12-08 10:27:22.574'),
(2, 1, 'Sudah Dibayar', '2022-12-08 10:27:22.574'),
(3, 1, 'Sudah Dibayar\r\n', '2022-12-08 10:27:22.574'),
(4, 1, 'Sudah Dibayar', '2022-12-08 10:27:22.574'),
(5, 1, 'Sudah Dibayar', '2022-12-08 10:27:22.574'),
(6, 1, 'Sudah Dibayar', '2022-12-08 10:27:22.574'),
(7, 1, 'Sudah Dibayar', '2022-12-08 10:27:22.574'),
(8, 1, 'Sudah Dibayar', '2022-12-08 10:27:22.574'),
(9, 1, 'Sudah Dibayar', '2022-12-08 10:27:22.574'),
(10, 1, 'Sudah Dibayar', '2022-12-08 10:27:22.574'),
(11, 1, 'Sudah Dibayar', '2022-12-08 10:27:22.574'),
(12, 1, 'Sudah Dibayar', '2022-12-08 10:27:22.574'),
(13, 1, 'Sudah Dibayar', '2022-12-08 10:27:22.574'),
(14, 1, 'Sudah Dibayar', '2022-12-08 10:27:22.574'),
(15, 1, 'Sudah Dibayar', '2022-12-08 10:27:22.574'),
(16, 1, 'Sudah Dibayar', '2022-12-08 10:27:22.574'),
(17, 1, 'Sudah Dibayar', '2022-12-08 10:27:22.574'),
(18, 1, 'Sudah Dibayar', '2022-12-08 10:27:22.574'),
(19, 3, 'Sudah Dibayar', '2022-12-08 10:27:22.574'),
(20, 2, 'Sudah Dibayar', '2022-12-08 10:27:22.574'),
(21, 1, 'Sudah Dibayar', '2022-12-08 10:27:22.574'),
(22, 2, 'Sudah Dibayar', '2022-12-08 10:27:22.574'),
(24, 2, 'Sudah Dibayar', '2022-12-08 10:27:22.574'),
(25, 2, 'Sudah Dibayar', '2022-12-08 10:27:22.574'),
(26, 2, 'Sudah Dibayar', '0000-00-00 00:00:00.000'),
(27, 3, 'Sudah Dibayar', '0000-00-00 00:00:00.000'),
(28, 1, 'Sudah Dibayar', '0000-00-00 00:00:00.000'),
(29, 2, 'Sudah Dibayar', '0000-00-00 00:00:00.000'),
(30, 1, 'Sudah Dibayar', '2022-12-08 16:37:42.871'),
(31, 2, 'Sudah Dibayar', '0000-00-00 00:00:00.000'),
(32, 3, 'Sudah Dibayar', '0000-00-00 00:00:00.000'),
(33, 1, 'Sudah Dibayar', '0000-00-00 00:00:00.000'),
(34, 2, 'Sudah Dibayar', '0000-00-00 00:00:00.000'),
(35, 1, 'Sudah Dibayar', '0000-00-00 00:00:00.000'),
(36, 2, 'Sudah Dibayar', '0000-00-00 00:00:00.000'),
(37, 3, 'Sudah Dibayar', '0000-00-00 00:00:00.000'),
(38, 1, 'Sudah Dibayar', '0000-00-00 00:00:00.000'),
(39, 2, 'Sudah Dibayar', '0000-00-00 00:00:00.000'),
(40, 1, 'Sudah Dibayar', '2022-12-08 16:43:40.989');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `bookings`
--
ALTER TABLE `bookings`
  ADD PRIMARY KEY (`id_booking`),
  ADD KEY `fk_lapangans_booking` (`no_lap`),
  ADD KEY `fk_transaksis_booking` (`kode_transaksi`);

--
-- Indexes for table `lapangans`
--
ALTER TABLE `lapangans`
  ADD PRIMARY KEY (`no_lapangan`);

--
-- Indexes for table `riwayats`
--
ALTER TABLE `riwayats`
  ADD PRIMARY KEY (`id_booking`),
  ADD KEY `fk_lapangans_riwayat` (`no_lap`);

--
-- Indexes for table `transaksis`
--
ALTER TABLE `transaksis`
  ADD PRIMARY KEY (`id_transaksi`),
  ADD KEY `fk_lapangans_transaksi` (`no_lap`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `bookings`
--
ALTER TABLE `bookings`
  MODIFY `id_booking` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=341;

--
-- AUTO_INCREMENT for table `lapangans`
--
ALTER TABLE `lapangans`
  MODIFY `no_lapangan` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `riwayats`
--
ALTER TABLE `riwayats`
  MODIFY `id_booking` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=308;

--
-- AUTO_INCREMENT for table `transaksis`
--
ALTER TABLE `transaksis`
  MODIFY `id_transaksi` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=41;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `bookings`
--
ALTER TABLE `bookings`
  ADD CONSTRAINT `fk_lapangans_booking` FOREIGN KEY (`no_lap`) REFERENCES `lapangans` (`no_lapangan`),
  ADD CONSTRAINT `fk_transaksis_booking` FOREIGN KEY (`kode_transaksi`) REFERENCES `transaksis` (`id_transaksi`);

--
-- Constraints for table `riwayats`
--
ALTER TABLE `riwayats`
  ADD CONSTRAINT `fk_lapangans_riwayat` FOREIGN KEY (`no_lap`) REFERENCES `lapangans` (`no_lapangan`);

--
-- Constraints for table `transaksis`
--
ALTER TABLE `transaksis`
  ADD CONSTRAINT `fk_lapangans_transaksi` FOREIGN KEY (`no_lap`) REFERENCES `lapangans` (`no_lapangan`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
