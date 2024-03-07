-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Mar 07, 2024 at 04:03 PM
-- Server version: 10.4.32-MariaDB
-- PHP Version: 8.1.25

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `company`
--

-- --------------------------------------------------------

--
-- Table structure for table `karyawans`
--

CREATE TABLE `karyawans` (
  `id` bigint(20) NOT NULL,
  `nama` longtext NOT NULL,
  `nip` longtext DEFAULT NULL,
  `tempat_lahir` longtext DEFAULT NULL,
  `tanggal_lahir` datetime(3) DEFAULT NULL,
  `umur` bigint(20) UNSIGNED DEFAULT NULL,
  `alamat` longtext DEFAULT NULL,
  `agama` longtext DEFAULT NULL,
  `jenis_kelamin` longtext DEFAULT NULL,
  `no_hp` longtext DEFAULT NULL,
  `email` longtext DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `karyawans`
--

INSERT INTO `karyawans` (`id`, `nama`, `nip`, `tempat_lahir`, `tanggal_lahir`, `umur`, `alamat`, `agama`, `jenis_kelamin`, `no_hp`, `email`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 'Arif Widiasan', '123456789', 'Malang', '1999-10-06 07:00:00.000', 24, 'JL.Kenjeran No.2 KotaSurabaya', 'Islam', 'Laki - Laki', '1234567890', 'arifw@example.com', '2024-03-06 22:28:07.416', '2024-03-07 00:35:09.991', NULL),
(3, 'Budi Nugroho', '987654321', 'Bandung', '2001-02-12 07:00:00.000', 23, 'JL.Rungkut No.2 Kota Surabaya', 'Islam', 'Laki - Laki', '1234567890', 'budi@example.com', '2024-03-07 00:11:59.953', '2024-03-07 00:36:17.285', NULL),
(4, 'indah permata', '987654321', 'Jakarta', '2002-02-12 07:00:00.000', 22, 'JL.Kedung Kandang No.2 Kota Surabaya', 'Islam', 'Perempuan', '1234567890', 'indah@example.com', '2024-03-07 00:25:18.906', NULL, NULL),
(5, 'q', 'we', 'wqe', '2023-02-08 07:00:00.000', 1, 'asdasd', 'islam', 'perempuan', '123', 'ads@adas.v', '2024-03-07 19:46:49.472', NULL, '2024-03-07 21:49:57.666'),
(6, 'qadsad', 'we', 'wqe', '2023-02-08 07:00:00.000', 1, 'asdasd', 'islam', 'perempuan', '123', 'ads@adas.v', '2024-03-07 19:48:06.123', NULL, NULL),
(7, 'Jok', '123123', 'Njo', '2023-12-06 07:00:00.000', 0, 'adsasd', 'islam', 'perempuan', '123123', 'ada@sa.d', '2024-03-07 19:55:07.232', NULL, NULL),
(8, 'Jokasd', '123123', 'Njo', '2023-12-06 07:00:00.000', 0, 'adsasd', 'islam', 'perempuan', '123123', 'ada@sa.d', '2024-03-07 19:55:53.580', NULL, '2024-03-07 20:44:25.408'),
(9, 'New', 'qweq', 'qweqe', '2021-06-02 07:00:00.000', 2, 'dqw', 'katolik', 'perempuan', 'd21', 'dad@s.s', '2024-03-07 19:57:51.553', '2024-03-07 20:38:47.075', NULL),
(10, 'qwew', 'e', 't', '2020-02-04 07:00:00.000', 4, 'dqd', 'islam', 'perempuan', 'wqd', 'qwdd@s', '2024-03-07 20:01:47.800', NULL, '2024-03-07 20:43:28.323');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` bigint(20) NOT NULL,
  `username` varchar(32) NOT NULL,
  `password` longtext NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `username`, `password`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 'arif', '$2a$10$p14a3MxkXL.BiQBRHhamieaw9UstW97fag9v1RedXtT/Ht03IPOje', '2024-03-07 04:02:45.972', '2024-03-06 21:02:46.075', '0000-00-00 00:00:00.000'),
(2, 'budi', '$2a$10$bqHiDu5WkT.vxpIofUvrQ.IjsqY9Xpb5UPLURSYa8/iORCHX9rVTS', '2024-03-06 21:08:26.135', '2024-03-06 21:08:26.232', '0000-00-00 00:00:00.000'),
(3, 'indah', '$2a$10$XLCu1uz9f1zs1dsi.gXf0e3uz2lFdcB2fLLSZsiXFly09Bq10Ufb2', '2024-03-06 21:15:00.092', '2024-03-06 21:15:00.092', '0000-00-00 00:00:00.000'),
(4, 'nando', '$2a$10$s/UAGF91Z.n4NxnbG1sF6u9nYImdY.n15W3HTn0HHv/WEMks1BAQq', '2024-03-07 18:41:41.909', '2024-03-07 18:41:41.909', '0000-00-00 00:00:00.000'),
(5, 'lukman', '$2a$10$nW9d2alM.LC3Wk1yTuGTqe/6MqUg51o9lOMziipD0wru3PaoGA.06', '2024-03-07 18:42:35.053', '2024-03-07 18:42:35.053', '0000-00-00 00:00:00.000');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `karyawans`
--
ALTER TABLE `karyawans`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_karyawans_deleted_at` (`deleted_at`),
  ADD KEY `idx_karyawans_updated_at` (`updated_at`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `idx_users_username` (`username`),
  ADD KEY `idx_users_deleted_at` (`deleted_at`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `karyawans`
--
ALTER TABLE `karyawans`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
