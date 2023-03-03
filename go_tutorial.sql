-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 02 Mar 2023 pada 17.42
-- Versi server: 10.4.27-MariaDB
-- Versi PHP: 8.0.25

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `go_tutorial`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `photos`
--

CREATE TABLE `photos` (
  `id` bigint(20) NOT NULL,
  `title` varchar(300) DEFAULT NULL,
  `caption` varchar(300) DEFAULT NULL,
  `photo_url` varchar(300) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `photos`
--

INSERT INTO `photos` (`id`, `title`, `caption`, `photo_url`) VALUES
(1, 'Danau Toba', 'pemandangan alam di danau toba', 'https://awsimages.detik.net.id/community/media/visual/2022/04/30/tipang_169.jpeg?w=650&q=90'),
(2, 'Gunung Sibayak', 'Spot Camping Anak Medan', 'https://cdns.klimg.com/merdeka.com/i/w/news/2020/07/21/1200376/content_images/670x335/20200721101646-1-favorit-di-kalangan-pendaki-pemula-ini-4-fakta-menarik-gunung-sibayak-003-fatimah-rahmawati.png'),
(3, 'Penatapan Berastagi', 'Menikmati Udara Sejuk Sembari Bakar Jagung', 'https://awsimages.detik.net.id/community/media/visual/2021/02/05/dev-santai-dulu-makan-jagung-bakar-di-kawasan-penatapan.jpeg?w=600&q=90');

-- --------------------------------------------------------

--
-- Struktur dari tabel `users`
--

CREATE TABLE `users` (
  `id` bigint(20) NOT NULL,
  `nama_lengkap` longtext DEFAULT NULL,
  `email` longtext DEFAULT NULL,
  `username` longtext DEFAULT NULL,
  `password` longtext DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `users`
--

INSERT INTO `users` (`id`, `nama_lengkap`, `email`, `username`, `password`) VALUES
(1, 'jhon doe', 'jhondoe@gmail.com', 'jhon', '$2a$10$LBAMpmniNPkbLCXP73TV2.TxA1Uz1aFqyZgVoXjCiCg8/GvWVeUgq');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `photos`
--
ALTER TABLE `photos`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `photos`
--
ALTER TABLE `photos`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT untuk tabel `users`
--
ALTER TABLE `users`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
