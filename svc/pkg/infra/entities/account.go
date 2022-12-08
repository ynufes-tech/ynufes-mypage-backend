package entities

type (
	ID      int64
	Account struct {
		ID                ID                `gorm:"column:id;type:bigint;not null;primary_key"`
		LineAuthorization LineAuthorization `gorm:"foreignKey:AccountID";references:ID`
	}
)
