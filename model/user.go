package model

import (
	"database/sql"
	"encoding/json"
)

type User struct {
	BaseModel
	Email    string         `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	Password string         `gorm:"type:varchar(255);not null" json:"-"`
	Name     sql.NullString `gorm:"type:varchar(255)" json:"name"`
}

func (u User) MarshalJSON() ([]byte, error) {
	type Alias User // Avoid recursion by creating an alias type
	return json.Marshal(&struct {
		Alias
		Name *string `json:"name"`
	}{
		Alias: Alias(u),
		Name: func() *string {
			if u.Name.Valid {
				return &u.Name.String
			}
			return nil
		}(),
	})
}
