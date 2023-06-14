package usecase

// get gender label
var getGenderLabel = func(gender int) string {
	switch gender {
	case 0:
		return "Laki-laki"
	case 1:
		return "Perempuan"
	default:
		return ""
	}
}

// get religion label
var getReligionLabel = func(religion int) string {
	switch religion {
	case 1:
		return "Islam"
	case 2:
		return "Kristen Protestan"
	case 3:
		return "Kristen Katolik"
	case 4:
		return "Hindu"
	case 5:
		return "Budha"
	case 6:
		return "Konghucu"
	default:
		return ""
	}
}

// get marital status label
var getMaritalStatusLabel = func(maritalStatus int) string {
	switch maritalStatus {
	case 1:
		return "Belum Kawin"
	case 2:
		return "Kawin"
	case 3:
		return "Cerai Hidup"
	case 4:
		return "Cerai Mati"
	default:
		return ""
	}
}
