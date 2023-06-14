package mysql

const (
	selectQuery = `
		SELECT
			id,
			nik,
			name,
			placeOfBirth,
			dateOfBirth,
			gender,
			bloodType,
			address,
			religion,
			maritalStatus,
			createdAt,
			updatedAt
		FROM
			employees
	`

	QueryGet = selectQuery + `
		ORDER BY id ASC
	`

	QueryCreate = `
		INSERT INTO userAddress(
			guid,
			firstName,
			lastName,
			email,
			phoneNumber,
			province,
			city,
			district,
			village,
			postalCode,
			address,
			isDefault,
			createdAt,
			updatedAt
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	QueryUpdate = `
		UPDATE userAddress
		SET
			firstName = ?,
			lastName = ?,
			email = ?,
			phoneNumber = ?,
			province = ?,
			city = ?,
			district = ?,
			village = ?,
			postalCode = ?,
			address = ?,
			updatedAt = ?
		WHERE
			id = ?
			AND guid = ?
	`

	QueryDelete = `
		DELETE FROM userAddress
		WHERE
			id = ?
			AND guid = ?
			AND isDefault = false
	`

	QueryUpdateIsDefault = `
		UPDATE userAddress
		SET
			isDefault = ?,
			updatedAt = ?
		WHERE
			id = ?
			AND guid = ?
	`

	QueryGetIsDefaultTrue = selectQuery + `
		WHERE
			guid = ?
			AND isDefault = 1
	`
)
