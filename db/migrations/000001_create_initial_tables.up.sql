CREATE TABLE `users` (
    `id` varchar(36) PRIMARY KEY,
    `username` varchar(255),
    `password` varchar(255),
    `email` varchar(360),
    `phone` varchar(10),
    `address` varchar(255)
);

CREATE TABLE `products` (
    `id` varchar(36) PRIMARY KEY,
    `name` varchar(255) NOT NULL ,
    `image` varchar(255) NOT NULL ,
    `description` varchar(255),
    `price` float NOT NULL ,
    `label_name` varchar(255) NOT NULL,
    `created_at` timestamp
);

CREATE TABLE `order_details` (
    `order_id` varchar(36) NOT NULL ,
    `product_id` varchar(36) NOT NULL ,
    `quantity` integer NOT NULL
);

CREATE TABLE `orders` (
    `id` varchar(36) PRIMARY KEY,
    `amount` float NOT NULL,
    `status` int NOT NULL,
    `user_id` varchar(36) NOT NULL,
    `created_at` timestamp NOT NULL
);

CREATE TABLE `labels` (
   `name` varchar(255) PRIMARY KEY ,
   `image` varchar(255) NOT NULL
);