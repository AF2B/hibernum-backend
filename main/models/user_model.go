package models

type User struct {
	ID                        string          `json:"id"`
	Name                      string          `json:"name"`
	Email                     string          `json:"email"`
	Password                  string          `json:"password"`
	Address                   Address         `json:"address"`
	Phone                     string          `json:"phone"`
	RegistrationDate          string          `json:"registration_date"`
	LastUpdate                string          `json:"last_update"`
	Role                      string          `json:"role"`
	Permissions               []string        `json:"permissions"`
	Token                     string          `json:"token"`
	ActivityLog               []ActivityLog   `json:"activity_log"`
	ProfilePicture            string          `json:"profile_picture"`
	LanguagePreference        string          `json:"language_preference"`
	EmailVerified             string          `json:"email_verified"`
	PasswordStrengthIndicator string          `json:"password_strength_indicator"`
	ActiveSessions            []ActiveSession `json:"active_sessions"`
	PasswordReset             PasswordReset   `json:"password_reset"`
	Blocked                   string          `json:"blocked"`
}

type ActiveSession struct {
	Device     string `json:"device"`
	IPAddress  string `json:"ip_address"`
	LastAccess string `json:"last_access"`
}

type ActivityLog struct {
	Timestamp string `json:"timestamp"`
	Action    string `json:"action"`
}

type Address struct {
	Street       string `json:"street"`
	Number       int    `json:"number"`
	Complement   string `json:"complement"`
	Neighborhood string `json:"neighborhood"`
	City         string `json:"city"`
	State        string `json:"state"`
	ZipCode      string `json:"zip_code"`
}

type PasswordReset struct {
	VerificationCode string `json:"verification_code"`
	ExpirationTime   string `json:"expiration_time"`
}
