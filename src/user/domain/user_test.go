package user

import (
	"testing"
	"time"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser(
		"Name",
		"mock@email.com",
		"OUTRO",
		time.Now(),
	)

	if err != nil {
		t.Error(err)
	}
	if user.Status != IN_ANALYSIS {
		t.Errorf("expected %+v, got %+v", IN_ANALYSIS, user.Status)
	}
}

func TestNewUserInvalidInputs(t *testing.T) {
	tests := []struct {
		tName, name, email, gender string
	}{
		{"invalid name", "", "test@email.com", "OUTRO"},
		{"invalid email", "Name", "test@email", "OUTRO"},
		{"invalid gender", "Name", "test@email.com", "ANY_STRING"},
	}

	for _, tt := range tests {
		t.Run(tt.tName, func(t *testing.T) {
			if _, err := NewUser(
				tt.name,
				tt.email,
				Gender(tt.gender),
				time.Now(),
			); err == nil {
				t.Errorf("expected error for %s", tt.tName)
			}
		})
	}
}

func TestValidate(t *testing.T) {
	tests := []struct {
		tName, name, status, gender string
	}{
		{"invalid name", "", "ATIVO", "OUTRO"},
		{"invalid status", "Name", "ANY_STRING", "OUTRO"},
		{"invalid gender", "Name", "ATIVO", "ANY_STRING"},
	}

	for _, tt := range tests {
		t.Run(tt.tName, func(t *testing.T) {
			user, _ := NewUser(
				"Name",
				"test@email.com",
				OTHER,
				time.Now(),
			)

			user.Name = tt.name
			user.Status = Status(tt.status)
			user.Gender = Gender(tt.gender)

			if err := user.Validate(); err == nil {
				t.Errorf("expected error for %s", tt.tName)
			}
		})
	}
}

func TestEnable(t *testing.T) {
	user, _ := NewUser(
		"Name",
		"mock@email.com",
		"OUTRO",
		time.Now(),
	)

	if err := user.Enable(); err != nil {
		t.Error(err)
	}
	if err := user.Enable(); err == nil {
		t.Error("expected error when enabling active user")
	}
}

func TestDisable(t *testing.T) {
	user, _ := NewUser(
		"Name",
		"mock@email.com",
		"OUTRO",
		time.Now(),
	)

	if err := user.Disable(); err != nil {
		t.Error(err)
	}
	if err := user.Disable(); err == nil {
		t.Error("expected error when disabling already inactive user")
	}
}

func TestGetters(t *testing.T) {
	user, _ := NewUser(
		"Name",
		"mock@email.com",
		"OUTRO",
		time.Now(),
	)

	if user.ID != user.GetID() {
		t.Errorf("expected %+v, got %+v", user.ID, user.GetID())
	}
	if user.Status != user.GetStatus() {
		t.Errorf("expected %+v, got %+v", user.Status, user.GetStatus())
	}
	if user.Name != user.GetName() {
		t.Errorf("expected %+v, got %+v", user.Name, user.GetName())
	}
	if user.Email != user.GetEmail() {
		t.Errorf("expected %+v, got %+v", user.Email, user.GetEmail())
	}
	if user.Gender != user.GetGender() {
		t.Errorf("expected %+v, got %+v", user.Gender, user.GetGender())
	}
	if user.BirthDate != user.GetBirthDate() {
		t.Errorf("expected %+v, got %+v", user.BirthDate, user.GetBirthDate())
	}
}
