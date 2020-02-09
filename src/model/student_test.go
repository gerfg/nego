package model_test

import (
	"github.com/InsideCI/nego/src/model"
	"testing"
)

func TestNewStudent(t *testing.T) {
	got := model.NewStudent(1010, "Student", 1010)
	if err := got.Valid(); err != nil {
		t.Error(err)
	}
}
