-- phpMyAdmin SQL Dump
-- version 5.0.2
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: Aug 30, 2020 at 06:50 AM
-- Server version: 10.4.11-MariaDB
-- PHP Version: 7.4.5

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `bonjour`
--

-- --------------------------------------------------------

--
-- Table structure for table `messages`
--

CREATE TABLE `messages` (
  `id` int(11) NOT NULL,
  `sender_id` int(11) NOT NULL,
  `receiver_id` int(11) NOT NULL,
  `message` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `messages`
--

INSERT INTO `messages` (`id`, `sender_id`, `receiver_id`, `message`, `created_at`) VALUES
(117, 70, 2, 'À plus tard', '2020-08-17 07:48:03'),
(118, 68, 2, 'A tout a l\'heur', '2020-08-17 07:49:20'),
(119, 67, 2, 'Ceci est un journal anglais', '2020-08-17 07:49:50'),
(120, 65, 2, 'Je prends du riz frit au déjeuner', '2020-08-17 07:51:21'),
(121, 6, 2, 'Je viens d\'Allemagne', '2020-08-17 07:52:13'),
(122, 64, 2, 'Aimez-vous la nourriture?', '2020-08-17 07:53:39'),
(123, 45, 2, 'Oui, j\'ai un petit frère', '2020-08-17 07:54:07'),
(124, 44, 2, 'Comment vous appelez vous?', '2020-08-17 07:55:00'),
(125, 3, 2, 'Il coûte dix dollars', '2020-08-17 07:56:57'),
(126, 2, 64, 'Oui, j\'aime beacoup', '2020-08-17 14:59:14'),
(129, 64, 2, 'Ce n\'est pas trop épicé, si?', '2020-08-17 15:04:56'),
(130, 2, 64, 'Non, c\'est juste comme il faut', '2020-08-17 15:11:27'),
(131, 64, 2, 'Je vous remercie!', '2020-08-17 15:11:59');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `name` varchar(40) NOT NULL,
  `email` varchar(60) NOT NULL,
  `password` varchar(60) NOT NULL,
  `profile_img` varchar(255) NOT NULL DEFAULT 'profile-placeholder.png',
  `about` varchar(60) NOT NULL DEFAULT 'Bonjour!',
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `name`, `email`, `password`, `profile_img`, `about`, `created_at`, `updated_at`) VALUES
(2, 'Daniel Saputra', 'danielwetan.io@gmail.com', '$2b$10$VnHFPN9I9usAA/VouW8Fc.40nXNt2A60JAsXna7d6OzuGD2FHyGaG', 'profile.jpg', 'Je parle français', '2020-07-26 08:11:14', '2020-07-26 08:11:14'),
(3, 'Ahmad', 'ahmad@gmail.com', '$2b$10$NKB2ABYr88YIzOGj2fK7g..F6juGjnEfgmdF7ii3TXOwyAXE6/JZm', '9.jpg', 'Bonjour!', '2020-07-26 09:46:02', '2020-07-26 09:46:02'),
(44, 'Idris', 'idris@gmail.com', '$2b$10$KIZ7Fo1qwJ.NUjFa/mZ/eeHk7GJqYL1RwiTFz/hMHRJok6uRM4HEm', '13.jpg', 'Bonjour!', '2020-08-07 02:43:09', '2020-08-07 02:43:09'),
(45, 'John', 'john@mail.com', '$2b$10$DydfOlzYAWMmJRLR1FoijujraxMB/iT6Y.3pLQNYtgLTCCoea7506', '10.jpg', 'Bonjour!', '2020-08-07 02:46:44', '2020-08-07 02:46:44'),
(64, 'Adam', 'adam@mail.com', '$2b$10$DydfOlzYAWMmJRLR1FoijujraxMB/iT6Y.3pLQNYtgLTCCoea7506', '11.jpg', 'Bonjour!', '2020-08-07 02:46:44', '2020-08-07 02:46:44'),
(65, 'David', 'david@gmail.com', '$2b$10$KIZ7Fo1qwJ.NUjFa/mZ/eeHk7GJqYL1RwiTFz/hMHRJok6uRM4HEm', '12.jpg', 'Bonjour!', '2020-08-07 02:43:09', '2020-08-07 02:43:09'),
(67, 'Pedro', 'pedro@gmail.com', '$2b$10$NKB2ABYr88YIzOGj2fK7g..F6juGjnEfgmdF7ii3TXOwyAXE6/JZm', '3.jpg', 'Bonjour!', '2020-07-26 09:46:02', '2020-07-26 09:46:02'),
(68, 'Ronald', 'ronald@gmail.com', '$2b$10$NKB2ABYr88YIzOGj2fK7g..F6juGjnEfgmdF7ii3TXOwyAXE6/JZm', '15.jpg', 'Bonjour!', '2020-07-26 09:46:02', '2020-07-26 09:46:02'),
(70, 'Bryan', 'bryan@gmail.com', '$2b$10$NKB2ABYr88YIzOGj2fK7g..F6juGjnEfgmdF7ii3TXOwyAXE6/JZm', '14.jpg', 'Bonjour!', '2020-07-26 09:46:02', '2020-07-26 09:46:02');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `messages`
--
ALTER TABLE `messages`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `messages`
--
ALTER TABLE `messages`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=132;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=71;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
