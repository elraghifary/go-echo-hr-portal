CREATE TABLE `units`(
  `id` bigint unsigned AUTO_INCREMENT PRIMARY KEY,
  `name` varchar(191) NOT NULL,
  `createdAt` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP(),
  `updatedAt` timestamp NULL
);

INSERT INTO `units`
	(name)
VALUES
  ('IT'),
  ('Finance'),
  ('Sales & Marketing');

CREATE TABLE `positions`(
  `id` bigint unsigned AUTO_INCREMENT PRIMARY KEY,
  `unitId` bigint NOT NULL,
  `name` varchar(191) NOT NULL,
  `createdAt` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP(),
  `updatedAt` timestamp NULL
);

INSERT INTO `positions`
	(unitId, name)
VALUES
  (1, 'Software Engineer'),
  (1, 'Product Manager'),
  (1, 'Software Engineer in Test'),
  (1, 'Quality Assurance'),
  (2, 'Finance Clerk'),
  (2, 'Payroll Clerk'),
  (3, 'Marketing Analyst'),
  (3, 'Sales Executive');

CREATE TABLE `assignments`(
  `id` bigint unsigned AUTO_INCREMENT PRIMARY KEY,
  `employeeId` bigint NOT NULL,
  `positionId` bigint NOT NULL,
  `startDate` date NOT NULL,
  `endDate` date NOT NULL,
  `status` tinyint(1) NOT NULL COMMENT '1 = Pekerja Lepas, 2 = Magang, 3 = Kontrak, 4 = Permanen',
  `createdAt` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP(),
  `updatedAt` timestamp NULL
);

CREATE TABLE `employees`(
  `id` bigint unsigned AUTO_INCREMENT PRIMARY KEY,
  `nik` varchar(9) NOT NULL,
  `name` varchar(191) NOT NULL,
  `placeOfBirth` varchar(191) NOT NULL,
  `dateOfBirth` date NOT NULL,
  `gender` tinyint(1) NOT NULL,
  `bloodType` varchar(2) NULL,
  `address` text NOT NULL,
  `religion` tinyint(1) NOT NULL COMMENT '1 = Islam, 2 = Kristen Protestan, 3 = Kristen Katolik, 4 = Hindu, 5 = Budha, 6 = Konghucu',
  `maritalStatus` tinyint(1) NOT NULL COMMENT '1 = Belum Kawin, 2 = Kawin, 3 = Cerai Hidup, 4 = Cerai Mati',
  `createdAt` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP(),
  `updatedAt` timestamp NULL
);
