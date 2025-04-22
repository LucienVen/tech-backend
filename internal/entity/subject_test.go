package entity

import (
	"github.com/LucienVen/tech-backend/internal/utils"
	"testing"
)

func TestGetAllSubject(t *testing.T) {
	utils.FirstInit()

	subjects, _ := GetAllSubject()
	t.Log(subjects)
}
