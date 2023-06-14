package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetGenderLabel(t *testing.T) {
	assert.Equal(t, "Laki-laki", getGenderLabel(0))
	assert.Equal(t, "Perempuan", getGenderLabel(1))
}

func TestGetReligionLabel(t *testing.T) {
	assert.Equal(t, "Islam", getReligionLabel(1))
	assert.Equal(t, "Kristen Protestan", getReligionLabel(2))
	assert.Equal(t, "Kristen Katolik", getReligionLabel(3))
	assert.Equal(t, "Hindu", getReligionLabel(4))
	assert.Equal(t, "Budha", getReligionLabel(5))
	assert.Equal(t, "Konghucu", getReligionLabel(6))
}

func TestGetMaritalStatusLabel(t *testing.T) {
	assert.Equal(t, "Belum Kawin", getMaritalStatusLabel(1))
	assert.Equal(t, "Kawin", getMaritalStatusLabel(2))
	assert.Equal(t, "Cerai Hidup", getMaritalStatusLabel(3))
	assert.Equal(t, "Cerai Mati", getMaritalStatusLabel(4))
}
