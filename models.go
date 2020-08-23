package models

//Company #1 wise database
type Company struct {
	ID          string `json:"aid"`
	Type        string `json:"type"`
	CompanyID   int64  `json:"cid"`
	CompanyName string `json:"company_name,omitempty"`
	Website     string `json:"website,omitempty"`
	Status      int    `json:"status"`
}

//Access #2 login type admin,super,user,etc
type Access struct {
	ID         string `json:"aid"`
	Type       string `json:"type"`
	CompanyID  string `json:"cid"` //foreign key
	AccessID   int    `json:"access_id"`
	AccessName string `json:"access_name"`
	Status     int    `json:"status"`
}

//Login #3  all user account login table
type Login struct {
	ID         string `json:"aid"`
	Type       string `json:"type"`
	CompanyID  string `json:"cid"` //foreign key
	LoginID    int64  `json:"login_id"`
	AccessID   string `json:"access_id"` //foreign key
	AccessName string `json:"access_name"`
	UserName   string `json:"username"` //email as username
	Password   string `json:"passw"`
	CreateDate string `json:"create_date"`
	LastLogin  string `json:"last_login,omitempty"`
	Status     int    `json:"status"`
}

//Contact info for account and user
type Contact struct {
	ID        string `json:"aid"`
	Type      string `json:"type"` //table_name
	CompanyID string `json:"cid"`  //foreign key
	ContactID string `json:"contact_id"`
	//mobile,email,pager,phone,website,socialmedia
	ContactType string `json:"contact_type"`
	ContactData string `json:"contact_date"`
	//owner table info
	TableName string `json:"table_name"` //account
	PkField   string `json:"pkfield"`    //account_id
	PkValue   string `json:"pkvalue"`    //
	Status    int    `json:"status"`
}

//Address ...
type Address struct {
	AccountID   int64  `json:"account_id"` //foreign_key
	ID          string `json:"aid"`
	Type        string `json:"type"`
	CompanyID   string `json:"cid"` //foreign key
	AddressID   string `json:"address_id"`
	AddressType string `json:"address_type"` //billing,shipping
	Country     string `json:"country"`
	State       string `json:"state"`
	City        string `json:"city"`
	Address1    string `json:"address1"`
	Address2    string `json:"address2"`
	Zip         string `json:"zip"`
	Status      int    `json:"status"`
}

//Account #4 ::cid::account_id
type Account struct {
	//ParentId    string    `json:"parent,omitempty"`
	ID          string `json:"aid"`
	Type        string `json:"type"` //account
	CompanyID   string `json:"cid"`  //foreign key
	AccountID   int64  `json:"account_id"`
	AccountType string `json:"account_type"`
	AccountName string `json:"account_name"`
	LoginID     string `json:"login_id"` //foreign key

	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	DateOfBirth string `json:"dob,omitempty"`

	Mobile string `json:"mobile"` //
	Email  string `json:"email"`  //
	//ContactInfo []Contact `json:"contact_info,omitempty"`

	CreateDate string `json:"create_date"`
	UpdateDate string `json:"update_date,omitempty"`
	Status     int    `json:"status"`
}

//ActivityLog ..
type ActivityLog struct {
	ID           string `json:"aid"`
	LogID        int64  `json:"logid"`
	Type         string `json:"type"`
	CompanyID    string `json:"cid"` //foreign key
	ActivityType string `json:"activity_type"`
	//owner table info
	TableName  string `json:"table_name"`
	PkField    string `json:"pkfield"`
	PkValue    string `json:"pkvalue"`
	LogDetails string `json:"log_details"`
	IPAddress  string `json:"ip_address"`
	LoginID    string `json:"login_id"` //foreign key
	CreateDate string `json:"create_date"`
	Status     int    `json:"status"`
}

//DeviceLog tracks user device info (where they login to the system)
type DeviceLog struct {
	ID          string `json:"aid"`
	Type        string `json:"type"`
	LoginID     string `json:"login_id"` //foreign key
	DeviceType  string `json:"device_type"`
	CompanyID   string `json:"cid"` //foreign key
	Browser     string `json:"browser"`
	Os          string `json:"os"`
	Platform    string `json:"platform"`
	ScreenSize  string `json:"screen_size"`
	GeoLocation string `json:"geolocation,omitempty"`
	CreateDate  string `json:"create_date"`
	Status      int    `json:"status"`
}

//LoginSession keeps user login session for 24 hours or more
type LoginSession struct {
	ID          string `json:"aid"`
	Type        string `json:"type"`
	CompanyID   string `json:"cid"`        //foreign key
	DeviceInfo  string `json:"device_log"` //foreign key device_log::1
	SessionCode string `json:"session_code"`
	LoginID     string `json:"login_id"` //foreign key
	IPAddress   string `json:"ip_address"`
	IPCity      string `json:"city"`
	IPCountry   string `json:"country"`
	LoginTime   string `json:"login_time"`
	LogoutTime  string `json:"logout_time"`
	CreateDate  string `json:"create_date"`
	Status      int    `json:"status"`
}
