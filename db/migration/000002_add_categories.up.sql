--
-- Table structure for table `categories`
--

CREATE TABLE `categories` (
                              `id` int(11) NOT NULL,
                              `name` varchar(50) COLLATE utf8_persian_ci NOT NULL,
                              `active` tinyint(1) NOT NULL DEFAULT 1,
                              `created_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_persian_ci;

-- --------------------------------------------------------

--
-- Table structure for table `product_categories`
--

CREATE TABLE `product_categories` (
                                      `id` bigint(20) NOT NULL,
                                      `product_id` bigint(20) NOT NULL,
                                      `category_id` int(11) NOT NULL,
                                      `created_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_persian_ci;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `categories`
--
ALTER TABLE `categories`
    ADD PRIMARY KEY (`id`),
    ADD UNIQUE KEY `name` (`name`);

--
-- Indexes for table `product_categories`
--
ALTER TABLE `product_categories`
    ADD PRIMARY KEY (`id`),
    ADD KEY `fk_product_cat` (`product_id`),
    ADD KEY `fk_category_cat` (`category_id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `categories`
--
ALTER TABLE `categories`
    MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `product_categories`
--
ALTER TABLE `product_categories`
    MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `product_categories`
--
ALTER TABLE `product_categories`
    ADD CONSTRAINT `fk_category_cat` FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`) ON UPDATE CASCADE,
    ADD CONSTRAINT `fk_product_cat` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;