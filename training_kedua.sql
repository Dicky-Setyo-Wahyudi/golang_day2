-- phpMyAdmin SQL Dump
-- version 4.8.3
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 10 Bulan Mei 2022 pada 17.15
-- Versi server: 10.1.36-MariaDB
-- Versi PHP: 7.2.11

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `training_kedua`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `products`
--

CREATE TABLE `products` (
  `id` int(11) NOT NULL,
  `nama_produk` varchar(255) DEFAULT NULL,
  `deskripsi` varchar(255) DEFAULT NULL,
  `status` tinyint(1) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data untuk tabel `products`
--

INSERT INTO `products` (`id`, `nama_produk`, `deskripsi`, `status`) VALUES
(1, 'produk01', 'Deskripsi Produk 001', 0),
(2, 'produk02', 'Deskripsi Produk 002', 0),
(3, 'produk03', 'Deskripsi Produk 003', 0),
(4, 'produk04', 'Deskripsi Produk 004', 0),
(5, 'produk05', 'Deskripsi Produk 005', 0),
(6, 'produk06', 'Deskripsi Produk 006', 0),
(7, 'produk07', 'Deskripsi Produk 007', 0),
(8, 'produk08', 'Deskripsi Produk 008', 0),
(9, 'produk09', 'Deskripsi Produk 009', 0),
(10, 'produk10', 'Deskripsi Produk 010', 0),
(11, 'produk11', 'Deskripsi Produk 011', 0),
(12, 'produk12', 'Deskripsi Produk 012', 0),
(13, 'produk13', 'Deskripsi Produk 013', 0),
(14, 'produk14', 'Deskripsi Produk 014', 0),
(15, 'produk15', 'Deskripsi Produk 015', 0);

-- --------------------------------------------------------

--
-- Struktur dari tabel `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `username` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `no_telepon` varchar(255) DEFAULT NULL,
  `role` varchar(255) DEFAULT NULL,
  `status` tinyint(1) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data untuk tabel `users`
--

INSERT INTO `users` (`id`, `username`, `email`, `no_telepon`, `role`, `status`) VALUES
(1, 'admin', 'admin@gmail.com', '0218569822', 'admin', 0),
(2, 'super_admin1', 'superadmin@gmail.com', '0217895203', 'superadmin', 0),
(3, 'admin2', 'admin@gmail.com', '0217895203', 'admin', 0),
(4, 'admin3', 'admin@gmail.com', '0217895203', 'admin', 0),
(5, 'admin4', 'admin@gmail.com', '0217895203', 'admin', 0),
(6, 'admin5', 'admin@gmail.com', '0217895203', 'admin', 0),
(7, 'admin6', 'admin@gmail.com', '0217895203', 'admin', 0),
(8, 'admin7', 'admin@gmail.com', '0217895203', 'admin', 0),
(9, 'admin8', 'admin@gmail.com', '0217895203', 'admin', 0),
(10, 'admin9', 'admin@gmail.com', '0217895203', 'admin', 0),
(11, 'admin10', 'admin@gmail.com', '0217895203', 'admin', 0),
(12, 'admin11', 'admin@gmail.com', '0217895203', 'admin', 0),
(13, 'admin12', 'admin@gmail.com', '0217895203', 'admin', 0),
(14, 'admin13', 'admin@gmail.com', '0217895203', 'admin', 0),
(15, 'admin20', 'admin20@gmail.com', '021789520320', 'admin20', 0);

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `products`
--
ALTER TABLE `products`
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
-- AUTO_INCREMENT untuk tabel `products`
--
ALTER TABLE `products`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=17;

--
-- AUTO_INCREMENT untuk tabel `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=18;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
