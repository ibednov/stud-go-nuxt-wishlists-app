-- Создание базы данных и пользователя
CREATE DATABASE IF NOT EXISTS wishlists_db;
CREATE USER IF NOT EXISTS 'wishlists_user'@'%' IDENTIFIED BY 'wishlists_password';
GRANT ALL PRIVILEGES ON wishlists_db.* TO 'wishlists_user'@'%';
FLUSH PRIVILEGES;

-- Используем созданную базу данных
USE wishlists_db;

-- Устанавливаем кодировку
SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";
SET NAMES utf8mb4;

-- --------------------------------------------------------
-- Структура таблицы `users`
-- --------------------------------------------------------

CREATE TABLE `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password_hash` varchar(255) NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------
-- Структура таблицы `wishlists`
-- --------------------------------------------------------

CREATE TABLE `wishlists` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint unsigned NOT NULL,
  `name` varchar(255) NOT NULL,
  `description` text,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `wishlists_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------
-- Структура таблицы `wishes`
-- --------------------------------------------------------

CREATE TABLE `wishes` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `wishlist_id` bigint unsigned NOT NULL,
  `name` varchar(255) NOT NULL,
  `description` text,
  `price` decimal(10,2) DEFAULT NULL,
  `source_url` varchar(255) DEFAULT NULL,
  `image_url` varchar(255) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `status` enum('available','purchased','archived') DEFAULT 'available',
  PRIMARY KEY (`id`),
  KEY `wishlist_id` (`wishlist_id`),
  CONSTRAINT `wishes_ibfk_1` FOREIGN KEY (`wishlist_id`) REFERENCES `wishlists` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------
-- Структура таблицы `wishlist_collaborators`
-- --------------------------------------------------------

CREATE TABLE `wishlist_collaborators` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `wishlist_id` bigint unsigned NOT NULL,
  `user_id` bigint unsigned NOT NULL,
  `role` enum('owner','editor','viewer') NOT NULL,
  `permissions` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `wishlist_id` (`wishlist_id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `wishlist_collaborators_ibfk_1` FOREIGN KEY (`wishlist_id`) REFERENCES `wishlists` (`id`),
  CONSTRAINT `wishlist_collaborators_ibfk_2` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------
-- Структура таблицы `wishlist_changes`
-- --------------------------------------------------------

CREATE TABLE `wishlist_changes` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `wishlist_id` bigint unsigned NOT NULL,
  `change_description` text,
  `changed_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `user_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`id`),
  KEY `wishlist_id` (`wishlist_id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `wishlist_changes_ibfk_1` FOREIGN KEY (`wishlist_id`) REFERENCES `wishlists` (`id`),
  CONSTRAINT `wishlist_changes_ibfk_2` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------
-- Структура таблицы `tags`
-- --------------------------------------------------------

CREATE TABLE `tags` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`(255))
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------
-- Структура таблицы `wishlist_tags`
-- --------------------------------------------------------

CREATE TABLE `wishlist_tags` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `wishlist_id` bigint unsigned NOT NULL,
  `tag_id` bigint unsigned NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `wishlist_id` (`wishlist_id`),
  KEY `tag_id` (`tag_id`),
  CONSTRAINT `wishlist_tags_ibfk_1` FOREIGN KEY (`wishlist_id`) REFERENCES `wishlists` (`id`),
  CONSTRAINT `wishlist_tags_ibfk_2` FOREIGN KEY (`tag_id`) REFERENCES `tags` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------
-- Вставка тестовых данных
-- --------------------------------------------------------

-- Данные для таблицы `users`
INSERT INTO `users` (`username`, `email`, `password_hash`, `created_at`) VALUES
('user1', 'user1@example.com', 'password1', '2024-11-12 13:53:45'),
('user2', 'user2@example.com', 'password2', '2024-11-12 13:53:45'),
('user3', 'user3@example.com', 'password3', '2024-11-12 13:53:45'),
('user4', 'user4@example.com', 'password4', '2024-11-12 13:53:45'),
('user5', 'user5@example.com', 'password5', '2024-11-12 13:53:45'),
('alice', 'alice@example.com', 'hash1', '2024-11-12 13:59:16'),
('bob', 'bob@example.com', 'hash2', '2024-11-12 13:59:16'),
('charlie', 'charlie@example.com', 'hash3', '2024-11-12 13:59:16'),
('david', 'david@example.com', 'hash4', '2024-11-12 13:59:16'),
('eve', 'eve@example.com', 'hash5', '2024-11-12 13:59:16');

-- Данные для таблицы `wishlists`
INSERT INTO `wishlists` (`user_id`, `name`, `description`, `created_at`, `updated_at`) VALUES
(1, 'Holiday Wishlist', 'Things I want for my vacation', '2024-11-12 14:00:39', '2024-11-12 14:00:39'),
(2, 'Tech Gadgets', 'Wishlist for tech enthusiasts', '2024-11-12 14:00:39', '2024-11-12 14:00:39'),
(3, 'Books', 'Books I plan to read this year', '2024-11-12 14:00:39', '2024-11-12 14:00:39'),
(4, 'Home Decor', 'Items for decorating my apartment', '2024-11-12 14:00:39', '2024-11-12 14:00:39'),
(5, 'Fitness Gear', 'Equipment for staying fit', '2024-11-12 14:00:39', '2024-11-12 14:00:39');

-- Данные для таблицы `wishes`
INSERT INTO `wishes` (`wishlist_id`, `name`, `description`, `price`, `source_url`, `image_url`, `created_at`, `updated_at`) VALUES
(1, 'Camera', 'High-resolution DSLR', 500.00, 'http://example.com/camera', 'http://example.com/camera.jpg', '2024-11-12 14:01:12', '2024-11-12 14:01:12'),
(2, 'Smartphone', 'Latest model smartphone', 800.00, 'http://example.com/smartphone', 'http://example.com/smartphone.jpg', '2024-11-12 14:01:12', '2024-11-12 14:01:12'),
(3, 'E-book', 'Best-selling novel', 15.00, 'http://example.com/ebook', 'http://example.com/ebook.jpg', '2024-11-12 14:01:12', '2024-11-12 14:01:12'),
(4, 'Vase', 'Decorative glass vase', 50.00, 'http://example.com/vase', 'http://example.com/vase.jpg', '2024-11-12 14:01:12', '2024-11-12 14:01:12'),
(5, 'Yoga Mat', 'High-quality yoga mat', 30.00, 'http://example.com/yogamat', 'http://example.com/yogamat.jpg', '2024-11-12 14:01:12', '2024-11-12 14:01:12');

-- Данные для таблицы `wishlist_collaborators`
INSERT INTO `wishlist_collaborators` (`wishlist_id`, `user_id`, `role`) VALUES
(1, 1, 'owner'),
(2, 2, 'owner'),
(3, 3, 'owner'),
(4, 4, 'owner'),
(5, 5, 'owner');

-- Данные для таблицы `wishlist_changes`
INSERT INTO `wishlist_changes` (`wishlist_id`, `change_description`, `changed_at`, `user_id`) VALUES
(1, 'Added a new camera', '2024-11-12 14:01:44', 1),
(2, 'Updated smartphone details', '2024-11-12 14:01:44', 2),
(3, 'Removed an old book', '2024-11-12 14:01:44', 3),
(4, 'Added a new vase', '2024-11-12 14:01:44', 4),
(5, 'Changed yoga mat description', '2024-11-12 14:01:44', 5);

-- Данные для таблицы `tags`
INSERT INTO `tags` (`name`, `created_at`) VALUES
('electronics', '2024-11-12 14:02:00'),
('books', '2024-11-12 14:02:00'),
('home', '2024-11-12 14:02:00'),
('fitness', '2024-11-12 14:02:00'),
('travel', '2024-11-12 14:02:00');

-- Данные для таблицы `wishlist_tags`
INSERT INTO `wishlist_tags` (`wishlist_id`, `tag_id`, `created_at`) VALUES
(1, 5, '2024-11-12 14:02:30'), -- Holiday Wishlist -> travel
(2, 1, '2024-11-12 14:02:30'), -- Tech Gadgets -> electronics
(3, 2, '2024-11-12 14:02:30'), -- Books -> books
(4, 3, '2024-11-12 14:02:30'), -- Home Decor -> home
(5, 4, '2024-11-12 14:02:30'); -- Fitness Gear -> fitness
