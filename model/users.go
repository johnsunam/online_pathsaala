package model

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type RegisterPayload struct {
	Email           string `form:"email" binding:"required,email"`
	Password        string `form:"password" binding:"required,min=8,max=32,alphanum"`
	ConfirmPassword string `form:"confirmPassword" binding:"required,eqfield=Password"`
	UserRole        string `json:"userRole" db:"user_role" binding:"required,oneof=student admin superadmin"`
	UserName        string `json:"userName" db:"username"`
}

type LoginPayload struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

type User struct {
	ID          string    `json:"id" db:"id"`
	Email       string    `json:"email" db:"email"`
	Password    string    `json:"password" db:"password"`
	Activated   bool      `json:"activated" db:"activated"`
	ActivatedAt time.Time `json:"activated_at" db:"activated_at"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
	UpdateAt    time.Time `json:"updatedAt" db:"updated_at"`
	UserName    string    `json:"userName" db:"username"`
	Roles       []string  `json:"roles" db:"roles"`
}

type SignedDetails struct {
	Email string   `json:"email"`
	Id    string   `json:"id"`
	Roles []string `json:"roles"`
	jwt.RegisteredClaims
}

type LoginResponse struct {
	Id    string `json:"UserId"`
	Email string `json:"email"`
	Token string `json:"token"`
}

type Profile struct {
	DeviceIds []string               `json:"-" db:"device_ids"`
	FirstName *string                `json:"firstName,omitempty"  db:"first_name"`
	Birthdate *string                `json:"birthdate" binding:"omitnil,datetime=2006-01-02" db:"birth_date"`
	Language  string                 `json:"language,omitempty" db:"language"`
	Goal      *string                `json:"goal,omitempty" db:"goal"`
	UserId    string                 `json:"-" db:"user_id"`
	CreatedAt string                 `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt string                 `json:"updated_at,omitempty" db:"updated_at"`
	Email     string                 `json:"email"`
	Scores    []UserWheelOfLifeScore `json:"scores"`
}

type WheelOfLifeScore struct {
	WheelOfLifeId string `json:"wheelOfLifeCode" binding:"required,uuid"`
	Score         int    `json:"score" binding:"required,gte=1,lte=5"`
}

type WheelOfLifeScorePayload struct {
	Scores []WheelOfLifeScore `json:"scores" binding:"required,min=1"`
}

type UserWheelOfLifeScore struct {
	WheelOfLifeId   string `json:"wheelOfLifeId"`
	WheelOfLifeName string `json:"wheelOfLifeName"`
	Score           int    `json:"score"`
	Language        string `json:"-"`
}

// func (c wols) Value() (driver.Value, error) {
// 	return json.Marshal(c)
// }

// func (c *wols) Scan(value interface{}) error {
// 	b, ok := value.([]byte)
// 	if !ok {
// 		return errors.New("type assertion to []byte failed")
// 	}
// 	return json.Unmarshal(b, &c)
// }
