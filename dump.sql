-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Хост: db
-- Время создания: Ноя 12 2024 г., 14:04
-- Версия сервера: 9.1.0
-- Версия PHP: 8.2.25

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- База данных: `Zhuk`
--

-- --------------------------------------------------------

--
-- Структура таблицы `users`
--

CREATE TABLE `users` (
  `id` int NOT NULL,
  `username` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password_hash` varchar(255) NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Дамп данных таблицы `users`
--

INSERT INTO `users` (`id`, `username`, `email`, `password_hash`, `created_at`) VALUES
(1, 'user1', 'user1@example.com', 'password1', '2024-11-12 13:53:45'),
(2, 'user2', 'user2@example.com', 'password2', '2024-11-12 13:53:45'),
(3, 'user3', 'user3@example.com', 'password3', '2024-11-12 13:53:45'),
(4, 'user4', 'user4@example.com', 'password4', '2024-11-12 13:53:45'),
(5, 'user5', 'user5@example.com', 'password5', '2024-11-12 13:53:45'),
(6, 'alice', 'alice@example.com', 'hash1', '2024-11-12 13:59:16'),
(7, 'bob', 'bob@example.com', 'hash2', '2024-11-12 13:59:16'),
(8, 'charlie', 'charlie@example.com', 'hash3', '2024-11-12 13:59:16'),
(9, 'david', 'david@example.com', 'hash4', '2024-11-12 13:59:16'),
(10, 'eve', 'eve@example.com', 'hash5', '2024-11-12 13:59:16');

-- --------------------------------------------------------

--
-- Структура таблицы `wishes`
--

CREATE TABLE `wishes` (
  `id` int NOT NULL,
  `wishlist_id` int NOT NULL,
  `name` varchar(255) NOT NULL,
  `description` text,
  `price` decimal(10,2) DEFAULT NULL,
  `source_url` varchar(255) DEFAULT NULL,
  `image_url` varchar(255) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Дамп данных таблицы `wishes`
--

INSERT INTO `wishes` (`id`, `wishlist_id`, `name`, `description`, `price`, `source_url`, `image_url`, `created_at`, `updated_at`) VALUES
(1, 1, 'Camera', 'High-resolution DSLR', 500.00, 'http://example.com/camera', 'http://example.com/camera.jpg', '2024-11-12 14:01:12', '2024-11-12 14:01:12'),
(2, 2, 'Smartphone', 'Latest model smartphone', 800.00, 'http://example.com/smartphone', 'http://example.com/smartphone.jpg', '2024-11-12 14:01:12', '2024-11-12 14:01:12'),
(3, 3, 'E-book', 'Best-selling novel', 15.00, 'http://example.com/ebook', 'http://example.com/ebook.jpg', '2024-11-12 14:01:12', '2024-11-12 14:01:12'),
(4, 4, 'Vase', 'Decorative glass vase', 50.00, 'http://example.com/vase', 'http://example.com/vase.jpg', '2024-11-12 14:01:12', '2024-11-12 14:01:12'),
(5, 5, 'Yoga Mat', 'High-quality yoga mat', 30.00, 'http://example.com/yogamat', 'http://example.com/yogamat.jpg', '2024-11-12 14:01:12', '2024-11-12 14:01:12');

-- --------------------------------------------------------

--
-- Структура таблицы `wishlists`
--

CREATE TABLE `wishlists` (
  `id` int NOT NULL,
  `user_id` int NOT NULL,
  `name` varchar(255) NOT NULL,
  `description` text,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Дамп данных таблицы `wishlists`
--

INSERT INTO `wishlists` (`id`, `user_id`, `name`, `description`, `created_at`, `updated_at`) VALUES
(1, 1, 'Holiday Wishlist', 'Things I want for my vacation', '2024-11-12 14:00:39', '2024-11-12 14:00:39'),
(2, 2, 'Tech Gadgets', 'Wishlist for tech enthusiasts', '2024-11-12 14:00:39', '2024-11-12 14:00:39'),
(3, 3, 'Books', 'Books I plan to read this year', '2024-11-12 14:00:39', '2024-11-12 14:00:39'),
(4, 4, 'Home Decor', 'Items for decorating my apartment', '2024-11-12 14:00:39', '2024-11-12 14:00:39'),
(5, 5, 'Fitness Gear', 'Equipment for staying fit', '2024-11-12 14:00:39', '2024-11-12 14:00:39');

-- --------------------------------------------------------

--
-- Структура таблицы `wishlist_changes`
--

CREATE TABLE `wishlist_changes` (
  `id` int NOT NULL,
  `wishlist_id` int NOT NULL,
  `change_description` text NOT NULL,
  `changed_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Дамп данных таблицы `wishlist_changes`
--

INSERT INTO `wishlist_changes` (`id`, `wishlist_id`, `change_description`, `changed_at`) VALUES
(1, 1, 'Added a new camera', '2024-11-12 14:01:44'),
(2, 2, 'Updated smartphone details', '2024-11-12 14:01:44'),
(3, 3, 'Removed an old book', '2024-11-12 14:01:44'),
(4, 4, 'Added a new vase', '2024-11-12 14:01:44'),
(5, 5, 'Changed yoga mat description', '2024-11-12 14:01:44');

-- --------------------------------------------------------

--
-- Структура таблицы `wishlist_collaborators`
--

CREATE TABLE `wishlist_collaborators` (
  `id` int NOT NULL,
  `wishlist_id` int NOT NULL,
  `user_id` int NOT NULL,
  `role` enum('owner','editor','viewer') NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Дамп данных таблицы `wishlist_collaborators`
--

INSERT INTO `wishlist_collaborators` (`id`, `wishlist_id`, `user_id`, `role`) VALUES
(1, 1, 1, 'owner'),
(2, 2, 2, 'owner'),
(3, 3, 3, 'owner'),
(4, 4, 4, 'owner'),
(5, 5, 5, 'owner');

--
-- Индексы сохранённых таблиц
--

--
-- Индексы таблицы `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `username` (`username`),
  ADD UNIQUE KEY `email` (`email`);

--
-- Индексы таблицы `wishes`
--
ALTER TABLE `wishes`
  ADD PRIMARY KEY (`id`),
  ADD KEY `wishlist_id` (`wishlist_id`);

--
-- Индексы таблицы `wishlists`
--
ALTER TABLE `wishlists`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_id` (`user_id`);

--
-- Индексы таблицы `wishlist_changes`
--
ALTER TABLE `wishlist_changes`
  ADD PRIMARY KEY (`id`),
  ADD KEY `wishlist_id` (`wishlist_id`);

--
-- Индексы таблицы `wishlist_collaborators`
--
ALTER TABLE `wishlist_collaborators`
  ADD PRIMARY KEY (`id`),
  ADD KEY `wishlist_id` (`wishlist_id`),
  ADD KEY `user_id` (`user_id`);

--
-- AUTO_INCREMENT для сохранённых таблиц
--

--
-- AUTO_INCREMENT для таблицы `users`
--
ALTER TABLE `users`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;

--
-- AUTO_INCREMENT для таблицы `wishes`
--
ALTER TABLE `wishes`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT для таблицы `wishlists`
--
ALTER TABLE `wishlists`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT для таблицы `wishlist_changes`
--
ALTER TABLE `wishlist_changes`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT для таблицы `wishlist_collaborators`
--
ALTER TABLE `wishlist_collaborators`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- Ограничения внешнего ключа сохраненных таблиц
--

--
-- Ограничения внешнего ключа таблицы `wishes`
--
ALTER TABLE `wishes`
  ADD CONSTRAINT `wishes_ibfk_1` FOREIGN KEY (`wishlist_id`) REFERENCES `wishlists` (`id`);

--
-- Ограничения внешнего ключа таблицы `wishlists`
--
ALTER TABLE `wishlists`
  ADD CONSTRAINT `wishlists_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

--
-- Ограничения внешнего ключа таблицы `wishlist_changes`
--
ALTER TABLE `wishlist_changes`
  ADD CONSTRAINT `wishlist_changes_ibfk_1` FOREIGN KEY (`wishlist_id`) REFERENCES `wishlists` (`id`);

--
-- Ограничения внешнего ключа таблицы `wishlist_collaborators`
--
ALTER TABLE `wishlist_collaborators`
  ADD CONSTRAINT `wishlist_collaborators_ibfk_1` FOREIGN KEY (`wishlist_id`) REFERENCES `wishlists` (`id`),
  ADD CONSTRAINT `wishlist_collaborators_ibfk_2` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
