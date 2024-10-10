--
-- Table structure for table `factors`
--

CREATE TABLE `factors` (
                           `id` bigint(20) NOT NULL,
                           `customer_name` varchar(50) COLLATE utf8_persian_ci DEFAULT NULL,
                           `customer_mobile` char(15) COLLATE utf8_persian_ci DEFAULT NULL,
                           `seller` varchar(50) COLLATE utf8_persian_ci NOT NULL,
                           `created_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_persian_ci;

-- --------------------------------------------------------

--
-- Table structure for table `factor_details`
--

CREATE TABLE `factor_details` (
                                  `id` bigint(20) NOT NULL,
                                  `factor_id` bigint(20) NOT NULL,
                                  `product_id` bigint(20) NOT NULL,
                                  `sale_count` int(11) NOT NULL,
                                  `sale_price` bigint(20) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_persian_ci;

-- --------------------------------------------------------

--
-- Table structure for table `products`
--

CREATE TABLE `products` (
                            `id` bigint(20) NOT NULL,
                            `name` varchar(50) COLLATE utf8_persian_ci NOT NULL,
                            `brand` varchar(50) COLLATE utf8_persian_ci DEFAULT NULL,
                            `model` varchar(50) COLLATE utf8_persian_ci DEFAULT NULL,
                            `init_number` int(11) NOT NULL,
                            `present_number` int(11) NOT NULL,
                            `buy_price` bigint(20) NOT NULL,
                            `buy_date` date NOT NULL,
                            `sale_price` bigint(20) DEFAULT NULL,
                            `created_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_persian_ci;

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
                         `username` varchar(50) COLLATE utf8_persian_ci NOT NULL,
                         `password` varchar(255) COLLATE utf8_persian_ci NOT NULL,
                         `full_name` varchar(50) COLLATE utf8_persian_ci NOT NULL,
                         `mobile` char(15) COLLATE utf8_persian_ci NOT NULL,
                         `password_changed_at` timestamp NULL DEFAULT NULL,
                         `created_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_persian_ci;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `factors`
--
ALTER TABLE `factors`
    ADD PRIMARY KEY (`id`),
    ADD KEY `fk_seller` (`seller`);

--
-- Indexes for table `factor_details`
--
ALTER TABLE `factor_details`
    ADD PRIMARY KEY (`id`),
    ADD KEY `fk_factor` (`factor_id`),
    ADD KEY `fk_product` (`product_id`);

--
-- Indexes for table `products`
--
ALTER TABLE `products`
    ADD PRIMARY KEY (`id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
    ADD PRIMARY KEY (`username`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `factors`
--
ALTER TABLE `factors`
    MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `factor_details`
--
ALTER TABLE `factor_details`
    MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `products`
--
ALTER TABLE `products`
    MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `factors`
--
ALTER TABLE `factors`
    ADD CONSTRAINT `fk_seller` FOREIGN KEY (`seller`) REFERENCES `users` (`username`) ON UPDATE CASCADE;

--
-- Constraints for table `factor_details`
--
ALTER TABLE `factor_details`
    ADD CONSTRAINT `fk_factor` FOREIGN KEY (`factor_id`) REFERENCES `factors` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    ADD CONSTRAINT `fk_product` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`) ON UPDATE CASCADE;