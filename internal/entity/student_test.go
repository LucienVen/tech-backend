package entity

import (
	"github.com/LucienVen/tech-backend/internal/utils"
	"testing"
)

func TestGetAllStudent(t *testing.T) {
	cleanFunc := utils.FirstInit()
	defer cleanFunc()

	res, err := GetAllStudent()
	if err != nil {
		t.Error(err)
	}

	utils.StructPrintf(res[0])
}
