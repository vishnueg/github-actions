package datamodel

type Asset struct {
	Asset_id   int
	Created_by struct {
		Id               int
		Username         string
		Fullname         string
		Email            string
		Role             string
		Created_at       string
		Updated_at       string
		Is_approved      bool
		Is_staff         bool
		Is_superuser     bool
		Is_active        bool
		Last_login       string
		Groups           []interface{}
		User_permissions []interface{}
	}
	Asset_name      string
	Brand           string
	Asset_model     string
	Asset_category  string
	Description     string
	PurchaseDate    string
	DeployementDate string
	Asset_ip        string
	ConfForJetson   string
	Location        string
	Shelf           string
	Asset_position  string
	Asset_status    bool
	Serial_no       string
	Created_at      string
	Updated_at      string
}
