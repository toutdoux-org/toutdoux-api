package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gofrs/uuid"
)

// TodoEntryLabel is used by pop to map your todo_entry_labels database table to your go code.
type TodoEntryLabel struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`

	TodoEntryID uuid.UUID     `json:"todo_entry_id" db:"todo_entry_id"`
	LabelID     uuid.UUID     `json:"todo_list_label_id" db:"todo_list_label_id"`
	TodoEntry   TodoEntry     `belongs_to:"todo_entries" db:"-"`
	Label       TodoListLabel `belongs_to:"todo_list_labels" db:"-"`
}

// String is not required by pop and may be deleted
func (t TodoEntryLabel) String() string {
	jt, _ := json.Marshal(t)
	return string(jt)
}

// TodoEntryLabels is not required by pop and may be deleted
type TodoEntryLabels []TodoEntryLabel

// String is not required by pop and may be deleted
func (t TodoEntryLabels) String() string {
	jt, _ := json.Marshal(t)
	return string(jt)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (t *TodoEntryLabel) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (t *TodoEntryLabel) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (t *TodoEntryLabel) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
