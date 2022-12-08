package entities

type LineAuthorization struct {
	ID                    int64  `gorm:"column:id;type:bigint;not null;primary_key"`
	AccountID             int64  `gorm:"column:account_id;type:bigint;not null"`
	EncryptedAccessToken  string `gorm:"column:encrypted_access_token;type:varchar(255);not null"`
	EncryptedRefreshToken string `gorm:"column:encrypted_refresh_token;type:varchar(255);not null"`
	CreatedAt             int64  `gorm:"column:created_at;type:bigint;not null"`
	UpdatedAt             int64  `gorm:"column:updated_at;type:bigint;not null"`
}
