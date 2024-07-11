package models

import "github.com/google/uuid"

type GObject struct {
	Name  string `json:"name,omitempty" bson:"name,omitempty"`
	Value string `json:"value,omitempty" bson:"value,omitempty"`
}

type GAuthRequest struct {
	Login    string `json:"login,omitempty" bson:"login,omitempty"`
	Password string `json:"password,omitempty" bson:"password,omitempty"`
}

type GAuthResponse struct {
	AuthId string `json:"authId"  bson:"authId,omitempty"`
	User   string `json:"user"  bson:"token,omitempty"`
}

type GAuthHeader struct {
	Token string `json:"token,omitempty" bson:"token,omitempty"`
	XAuth string `json:"x_auth,omitempty" bson:"x_auth,omitempty"`
}
type GUserFindRequest struct {
	ParentId string `json:"parentId"`
	UserId   string `json:"userId"`
	Login    string `json:"login"`
	Email    string `json:"email"`
}

type GUserFindResponse struct {
	Id              string `json:"id"`
	LastName        string `json:"lastName"`
	FirstName       string `json:"firstName"`
	Position        string `json:"position"`
	Phone           string `json:"phone"`
	ParentId        string `json:"parentId"`
	ParentName      string `json:"parentName"`
	Organization    string `json:"organization"`
	Login           string `json:"login"`
	Email           string `json:"email"`
	SessionDuration string `json:"sessionDuration"`
	Status          string `json:"status"`
	LastLogged      string `json:"lastLogged"`
}

type GUser struct {
	Id                 string                 `json:"id"`
	Name               string                 `json:"name"`
	Password           string                 `json:"password"`
	PasswordSalt       string                 `json:"passwordSalt"`
	ResetPasswordCode  string                 `json:"resetPasswordCode"`
	Email              string                 `json:"email"`
	NewEmail           string                 `json:"newEmail"`
	Phone              string                 `json:"phone"`
	Description        string                 `json:"description"`
	FirstName          string                 `json:"firstName"`
	LastName           string                 `json:"lastName"`
	Profile            string                 `json:"profile"`
	CreatedAt          string                 `json:"createdAt"`
	UpdatedAt          string                 `json:"updatedAt"`
	LastLogged         string                 `json:"lastLogged"`
	AgentGuid          string                 `json:"agentGuid"`
	IsEnabled          bool                   `json:"isEnabled"`
	IsReadOnly         bool                   `json:"isReadOnly"`
	IsConfirmed        bool                   `json:"isConfirmed"`
	ConfirmationCode   string                 `json:"confirmationCode"`
	RestoreCode        string                 `json:"restoreCode"`
	Groups             []string               `json:"groups"`
	CustomGroups       []string               `json:"customGroups"`
	MessengersSettings map[string]interface{} `json:"messengersSettings"`
	LeaderId           string                 `json:"leaderId"`
	Status             int                    `json:"status"`
	Stage              int                    `json:"stage"`
	IsDisabledMobile   bool                   `json:"isDisabledMobile"`
	Language           int                    `json:"language"`
	StatusChangeDate   string                 `json:"statusChangeDate"`
	AvatarId           string                 `json:"avatarId"`
	CuttedAvatarId     string                 `json:"cuttedAvatarId"`
	Organization       string                 `json:"organization"`
	OrganizationScope  int                    `json:"organizationScope"`
	PasswordSet        bool                   `json:"passwordSet"`
	IsDeleted          bool                   `json:"IsDeleted"`
}

type GRegionalData struct {
	Country string    `json:"country"`
	GroupId string    `json:"groupId"`
	Fields  []GObject `json:"fields"`
}

type GClient struct {
	ID               uuid.UUID       `json:"id"`
	AgentId          uuid.UUID       `json:"agentId"`
	Dealer           bool            `json:"dealer"`
	LockedDate       string          `json:"lockedDate"`
	Organization     string          `json:"organization"`
	UserSupport      uuid.UUID       `json:"userSupport"`
	UserManager      uuid.UUID       `json:"userManager"`
	OtherParametrs   string          `json:"otherParametrs"`
	District         string          `json:"district"`
	Region           string          `json:"region"`
	City             string          `json:"city"`
	Status           int             `json:"status"`
	Stage            int             `json:"stage"`
	StatusChangeDate string          `json:"statusChangeDate"`
	AccFullName      string          `json:"accFullName"`
	AccEmail         string          `json:"accEmail"`
	AccDirector      string          `json:"accDirector"`
	AccRegionalData  []GRegionalData `json:"accRegionalData"`
	Balance          int             `json:"balance"`
	BalanceSt        int             `json:"balanceSt"`
	PromisePayment   int             `json:"promisePayment"`
	Debt             int             `json:"debt"`
}

type GFeatures struct {
	VehiclesLimit         int `json:"vehicles.limit"`
	ReportsTrack          int `json:"reports.track"`
	ReportsBasic          int `json:"reports.basic"`
	BillingLightweight    int `json:"billing.lightweight"`
	ReportsGroups         int `json:"reports.groups"`
	ReportsAdditional     int `json:"reports.additional"`
	RetranslatorsLimit    int `json:"retranslators.limit"`
	RetranslatorsCapacity int `json:"retranslators.capacity"`
	Fuelcards             int `json:"fuelcards"`
	VehiclesStatuses      int `json:"vehicles.statuses"`
}

// GProfile struct for glonass
type GProfile struct {
	User                GUser     `json:"user"`
	Reason              string    `json:"reason"`
	AgentInfoType       int       `json:"agentInfoType"`
	Permissions         []string  `json:"permissions"`
	IsLimitesChecable   bool      `json:"isLimitesChecable"`
	Features            GFeatures `json:"features"`
	Notifications       string    `json:"notifications"`
	Plan                string    `json:"plan"`
	Balance             string    `json:"balance"`
	Client              GClient   `json:"client"`
	Contracts           []string  `json:"contracts"`
	AgentName           string    `json:"agentName"`
	Messages            string    `json:"messages"`
	Personalization     string    `json:"personalization"`
	IsMonitoringAllowed bool      `json:"isMonitoringAllowed"`
	Error               string    `json:"error"`
}
