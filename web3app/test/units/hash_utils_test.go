package testing

import (
	"SimpleId/internal/utils"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashing(t *testing.T) {
	results, err := utils.HashAndEncodePassword("123")
	assert.NoError(t, err)
	log.Default().Print(results)
}

func TestHashCheckPassword(t *testing.T) {
	pass := "JDJhJDEwJG5CL0RjeEhtaExTOTRhMUluTXgvRC4yVTJETlQ5MzNKZEtDTjMyUUVTSlJXTTdLMlh4QWVt"
	results, err := utils.CheckPasswordHash("1235", pass)
	assert.Error(t, err)
	log.Default().Print(results)

}
