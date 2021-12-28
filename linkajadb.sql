-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: Dec 28, 2021 at 05:30 AM
-- Server version: 10.4.21-MariaDB
-- PHP Version: 7.3.30

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `linkajadb`
--

DELIMITER $$
--
-- Procedures
--
CREATE DEFINER=`root`@`localhost` PROCEDURE `spTransactionDetail` (IN `param_accounts_id` INT(100), IN `param_customer_number` INT(100), IN `param_debit` INT(100), IN `param_credit` INT(100), IN `param_total_balance` INT(100), IN `param_usr_crt` VARCHAR(100))  BEGIN
        INSERT INTO transaction_detail 
                    (accounts_id, customer_number, trx_debit, 					trx_credit, total_balance, usr_crt) 
        VALUES (param_accounts_id,  param_customer_number,  param_debit, param_credit, param_total_balance, param_usr_crt);
        
        
        
     UPDATE accounts SET balance = param_total_balance, dtm_upd = CURRENT_TIMESTAMP() WHERE id = param_accounts_id;
 
END$$

DELIMITER ;

-- --------------------------------------------------------

--
-- Table structure for table `accounts`
--

CREATE TABLE `accounts` (
  `id` int(3) NOT NULL,
  `account_number` int(6) NOT NULL,
  `customer_number` int(4) NOT NULL,
  `balance` int(50) NOT NULL,
  `dtm_upd` timestamp(6) NOT NULL DEFAULT current_timestamp(6),
  `dtm_crt` timestamp(6) NOT NULL DEFAULT current_timestamp(6)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `accounts`
--

INSERT INTO `accounts` (`id`, `account_number`, `customer_number`, `balance`, `dtm_upd`, `dtm_crt`) VALUES
(1, 555001, 1001, 2600, '2021-09-08 04:18:08.777219', '2021-09-08 04:18:32.029730'),
(2, 555002, 1002, 22400, '2021-09-08 04:18:08.777219', '2021-09-08 04:18:32.029730'),
(4, 555003, 1003, 11500, '2021-09-09 14:42:40.000000', '2021-09-08 04:20:24.000000'),
(5, 555004, 1004, 48500, '2021-09-09 14:42:40.000000', '2021-09-08 04:20:24.000000');

--
-- Triggers `accounts`
--
DELIMITER $$
CREATE TRIGGER `dtm_upd` BEFORE UPDATE ON `accounts` FOR EACH ROW BEGIN
	IF NEW.dtm_upd = '0000-00-00 00:00:00' THEN
    SET NEW.dtm_upd = CURRENT_TIMESTAMP();
    END IF;
END
$$
DELIMITER ;

-- --------------------------------------------------------

--
-- Table structure for table `customers`
--

CREATE TABLE `customers` (
  `id` int(3) NOT NULL,
  `customer_number` int(4) NOT NULL,
  `name` varchar(25) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `customers`
--

INSERT INTO `customers` (`id`, `customer_number`, `name`) VALUES
(1, 1001, 'Bob Martin'),
(2, 1002, 'Linus Torvalds'),
(3, 1003, 'Steve Jobs'),
(4, 1004, 'Fariz');

-- --------------------------------------------------------

--
-- Table structure for table `transaction_detail`
--

CREATE TABLE `transaction_detail` (
  `id` int(100) NOT NULL,
  `accounts_id` int(100) NOT NULL,
  `customer_number` int(100) NOT NULL,
  `trx_debit` int(100) NOT NULL,
  `trx_credit` int(100) NOT NULL,
  `total_balance` int(100) NOT NULL,
  `usr_crt` varchar(100) NOT NULL,
  `dtm_crt` timestamp(6) NOT NULL DEFAULT current_timestamp(6) ON UPDATE current_timestamp(6)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `transaction_detail`
--

INSERT INTO `transaction_detail` (`id`, `accounts_id`, `customer_number`, `trx_debit`, `trx_credit`, `total_balance`, `usr_crt`, `dtm_crt`) VALUES
(7, 4, 1003, 0, 1000, 9000, 'Farizree', '2021-09-09 14:14:48.479773'),
(8, 5, 1004, 1000, 0, 51000, 'Farizree', '2021-09-09 14:14:48.485132'),
(9, 4, 1003, 0, 500, 8500, 'Farizree', '2021-09-09 14:15:16.357766'),
(10, 5, 1004, 500, 0, 51500, 'Farizree', '2021-09-09 14:15:16.372052'),
(11, 5, 1004, 0, 1000, 50500, 'Farizree', '2021-09-09 14:15:53.001617'),
(12, 4, 1003, 1000, 0, 9500, 'Farizree', '2021-09-09 14:15:53.007978'),
(13, 5, 1004, 0, 1000, 49500, 'Farizree', '2021-09-09 14:39:39.519739'),
(14, 4, 1003, 1000, 0, 10500, 'Farizree', '2021-09-09 14:39:39.534176'),
(15, 5, 1004, 0, 1000, 48500, 'Farizree', '2021-09-09 14:42:40.929988'),
(16, 4, 1003, 1000, 0, 11500, 'Farizree', '2021-09-09 14:42:40.938298');

-- --------------------------------------------------------

--
-- Stand-in structure for view `view_customeraccount`
-- (See below for the actual view)
--
CREATE TABLE `view_customeraccount` (
`accounts_id` int(3)
,`account_number` int(6)
,`customer_name` varchar(25)
,`balance` int(50)
,`customer_number` int(4)
);

-- --------------------------------------------------------

--
-- Structure for view `view_customeraccount`
--
DROP TABLE IF EXISTS `view_customeraccount`;

CREATE ALGORITHM=UNDEFINED DEFINER=`root`@`localhost` SQL SECURITY DEFINER VIEW `view_customeraccount`  AS SELECT `accounts`.`id` AS `accounts_id`, `accounts`.`account_number` AS `account_number`, `customers`.`name` AS `customer_name`, `accounts`.`balance` AS `balance`, `accounts`.`customer_number` AS `customer_number` FROM (`customers` join `accounts` on(`accounts`.`customer_number` = `customers`.`customer_number`)) ;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `accounts`
--
ALTER TABLE `accounts`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `account_number` (`account_number`),
  ADD KEY `customer_number` (`customer_number`);

--
-- Indexes for table `customers`
--
ALTER TABLE `customers`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `customer_number` (`customer_number`);

--
-- Indexes for table `transaction_detail`
--
ALTER TABLE `transaction_detail`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `accounts`
--
ALTER TABLE `accounts`
  MODIFY `id` int(3) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT for table `customers`
--
ALTER TABLE `customers`
  MODIFY `id` int(3) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT for table `transaction_detail`
--
ALTER TABLE `transaction_detail`
  MODIFY `id` int(100) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=17;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `accounts`
--
ALTER TABLE `accounts`
  ADD CONSTRAINT `accounts_ibfk_1` FOREIGN KEY (`customer_number`) REFERENCES `customers` (`customer_number`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
