package contracts

import "golang.org/x/crypto/bcrypt"

// Authentication interface represents a backend flow authenticate user to SMTP for store message, or to dashboard for
// manage stored messages.
type Authentication interface {

	// SMTP returns object what process authentication for SMTP protocol.
	SMTP() SmtpAuthentication

	// Dashboard returns object what process authentication for UI dashboard.
	Dashboard() DashboardAuthentication

	// UsersStorage returns object what manipulate users in application.
	UsersStorage() UsersStorage
}

// SmtpAuthentication contains methods related to SMTP authentication to send email messages to application.
type SmtpAuthentication interface {

	// RequiresAuthentication define is need authentication for SMTP
	RequiresAuthentication() bool

	// IpsAllowList returns object what can manage IPs allowlist.
	IpsAllowList() IpsAllowList

	// ViaPasswordAuthentication returns object what can manage password for authentication.
	ViaPasswordAuthentication() ViaPasswordAuthentication

	// ViaIpAuthentication returns object what can manage IPs what can be authenticated without password.
	ViaIpAuthentication() ViaIpAuthentication
}

type DashboardAuthentication interface {

	// RequiresAuthentication define is need authentication for dashboard.
	RequiresAuthentication() bool

	// ViaPasswordAuthentication returns object what can manage password for authentication.
	ViaPasswordAuthentication() ViaPasswordAuthentication

	// ViaEmailAuthentication returns object what can manage emails what can be authenticated without password.
	ViaEmailAuthentication() ViaEmailAuthentication
}

type UsersStorage interface {

	// Exists check is username exists in storage.
	Exists(username string) bool

	// Add to auth storage.
	Add(username string) error

	// Delete from auth storage.
	Delete(username string) error

	// List from auth storage.
	List(searchQuery string, offset, limit int) ([]UserResource, int, error)
}

// IpsAllowList allow application while authentication checks IP before authenticate client and if IP not in
// allowed list then application will return unauthenticated response.
type IpsAllowList interface {

	// Enabled shows is allowlist flow enabled.
	Enabled() bool

	// Allowed check is IP in allowlist for specific user.
	Allowed(username string, ip string) bool

	// AddIp to auth allowlist storage related to specific user.
	AddIp(username string, ip string) error
	// DeleteIp from auth allowlist storage related to specific user.
	DeleteIp(username string, ip string) error
	// ClearAllIps from auth allowlist storage related to specific user.
	ClearAllIps(username string) error
}

// ViaPasswordAuthentication allow login by password.
type ViaPasswordAuthentication interface {

	// Enabled shows is authentication via email flow enabled.
	Enabled() bool

	// Authenticate check is credentials (login/password) are valid.
	Authenticate(username string, password string) bool

	// SetPassword to auth storage related to user.
	SetPassword(username string, password string) error
}

// ViaIpAuthentication allow application authenticate client by username and IP without checking password.
type ViaIpAuthentication interface {

	// Enabled shows is authentication via IP flow enabled.
	Enabled() bool

	// Authenticate check is application can bypass password and authenticate client just by username and IP.
	Authenticate(username string, ip string) bool

	// AddIp for "IP auth" to auth storage related to user.
	AddIp(username string, ip string) error
	// DeleteIp for "IP auth" from auth storage related to user.
	DeleteIp(username string, ip string) error
	// ClearAllIps for "IP auth" from auth storage related to user.
	ClearAllIps(username string) error
}

// ViaEmailAuthentication allow login by token sent to email.
type ViaEmailAuthentication interface {

	// Enabled shows is authentication via email flow enabled.
	Enabled() bool

	// SendToken to email.
	SendToken(username string, email string) error

	// Authenticate user by token sent to email.
	Authenticate(username string, email string, token string) bool

	// AddEmail for login to auth storage related to user
	AddEmail(username string, email string) error
	// DeleteEmail for login from auth storage related to user
	DeleteEmail(username string, email string) error
	// ClearAllEmails for login from auth storage related to user
	ClearAllEmails(username string) error
}

type UserResource struct {
	Username            string
	DashboardAuthEmails []string
	SmtpAuthIPs         []string
	SmtpAllowListedIPs  []string
}

type AuthenticationConfig struct {
	Smtp struct {
		IpsAllowList struct {
			Enabled bool `yaml:"enabled"`
		} `yaml:"ips_allowlist"`
		ViaIpAuthentication struct {
			Enabled bool `yaml:"enabled"`
		} `yaml:"via_ip"`
		ViaPasswordAuthentication struct {
			Enabled bool `yaml:"enabled"`
		} `yaml:"via_password"`
	} `yaml:"smtp"`
	Dashboard struct {
		ViaEmailAuthentication struct {
			Enabled bool `yaml:"enabled"`
		} `yaml:"via_email"`
		ViaPasswordAuthentication struct {
			Enabled bool `yaml:"enabled"`
		} `yaml:"via_password"`
	} `yaml:"dashboard"`
}

func CreatePasswordHash(password string) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return []byte{}, err
	}
	return bytes, nil
}
