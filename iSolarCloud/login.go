package iSolarCloud


// var LoginType = (*Login)(nil)
//
// var _ ResourceHandler = LoginType
// // var _ api.ContentHandler = LoginType
//
// type Login struct {
// 	//GuidsAndLinks
// 	Attributes LoginAttributes
// 	ResourceType
// 	// api.ContentType
// }
//
// func (*Login) InitializeResourceType(rt *ResourceType) *ResourceType {
// 	//rt.HasVideos = true
// 	//rt.HasCaptions = true
// 	//rt.HasCountries = true
// 	//rt.HasGeoProfile = true
// 	//rt.HasTopics = true
// 	//rt.HasPlatforms = true
// 	//rt.HasImages = true
// 	//rt.HasAvailabilities = true
// 	//rt.HasLinks = true
// 	return rt
// }
//
// func (*Login) InitializeContentType(ct *api.ContentType) *api.ContentType {
// 	ct.DownloadPriority = 6
// 	ct.ResourceType = (*LoginResource)(nil)
// 	ct.CollectionType = (*LoginCollection)(nil)
// 	return ct
// }
//
// type LoginAttributes struct {
// 	CommonAttributes
// 	ResultData struct {
// 		AcceptOrderNum               int64           `json:"accept_order_num"`
// 		BackgroundColor              int64           `json:"background_color"`
// 		Countryid                    string          `json:"countryid"`
// 		Createdate                   string          `json:"createdate"`
// 		Createuserid                 string          `json:"createuserid"`
// 		CurrentOrderNum              int64           `json:"current_order_num"`
// 		DisableTime                  string          `json:"disable_time"`
// 		Email                        string          `json:"email"`
// 		Englishname                  interface{}     `json:"englishname"`
// 		ErrTimes                     string          `json:"err_times"`
// 		GcjLatitude                  interface{}     `json:"gcj_latitude"`
// 		GcjLongitude                 interface{}     `json:"gcj_longitude"`
// 		ImToken                      interface{}     `json:"im_token"`
// 		IsAfsFlag                    string          `json:"is_afs_flag"`
// 		IsAgreeGdpr                  int64           `json:"is_agree_gdpr"`
// 		IsAu                         int64           `json:"is_au"`
// 		IsCanModifyUserAccount       int64           `json:"is_can_modify_user_account"`
// 		IsDisableMap                 string          `json:"is_disable_map"`
// 		IsGdpr                       int64           `json:"is_gdpr"`
// 		IsHaveIm                     int64           `json:"is_have_im"`
// 		IsNewVersion                 int64           `json:"is_new_version"`
// 		IsOnline                     string          `json:"is_online"`
// 		IsOpenProtocol               int64           `json:"is_open_protocol"`
// 		IsReceiveNotice              int64           `json:"is_receive_notice"`
// 		IsSharePosition              int64           `json:"is_share_position"`
// 		IsUploadLocation             int64           `json:"is_upload_location"`
// 		IsUseSungrowBrand            string          `json:"is_use_sungrow_brand"`
// 		IsValidMobileEmail           int64           `json:"is_valid_mobile_email"`
// 		Isdst                        string          `json:"isdst"` // DUPLICATE
// 		IsDST                        string          `json:"isDST"` // DUPLICATE
// 		Jobs                         interface{}     `json:"jobs"`
// 		Language                     string          `json:"language"`
// 		LoginFirstDate               string          `json:"loginFirstDate"` // DUPLICATE
// 		LoginFristDate               string          `json:"loginFristDate"` // DUPLICATE
// 		LoginLastDate                string          `json:"loginLastDate"`
// 		LoginLastIP                  string          `json:"loginLastIp"`
// 		LoginTimes                   int64           `json:"loginTimes"`
// 		LoginState                   string          `json:"login_state"`
// 		Logo                         interface{}     `json:"logo"`
// 		LogoHTTPSURL                 interface{}     `json:"logo_https_url"`
// 		MapType                      string          `json:"map_type"`
// 		MinDate                      string          `json:"min_date"`
// 		MobileTel                    interface{}     `json:"mobile_tel"`
// 		OrgID                        string          `json:"org_id"`
// 		OrgName                      string          `json:"org_name"`
// 		OrgTimezone                  string          `json:"org_timezone"`
// 		PasswordIsSimple             int64           `json:"password_is_simple"`
// 		PhotoID                      interface{}     `json:"photo_id"`
// 		PhotoURL                     interface{}     `json:"photo_url"`
// 		Privileges                   LoginPrivileges `json:"privileges"`
// 		RoleID                       string          `json:"role_id"`
// 		SecondaryOrgIds              []interface{}   `json:"secondaryOrgIds"`
// 		ServerTel                    string          `json:"server_tel"`
// 		ServiceVersion               string          `json:"service_version"`
// 		Sex                          string          `json:"sex"`
// 		StyleName                    string          `json:"stylename"`
// 		TimeZone                     string          `json:"timeZone"`
// 		Timezone                     string          `json:"timezone"`
// 		TimeZoneId                   string          `json:"timezoneid"`
// 		ToggleFlag                   string          `json:"toggleflag"`
// 		Token                        string          `json:"token"`
// 		UnlockLaveMinute             int64           `json:"unlock_lave_minute"`
// 		UploadTime                   interface{}     `json:"upload_time"`
// 		UserAccount                  string          `json:"user_account"`
// 		UserAccountModifyCount       int64           `json:"user_account_modify_count"`
// 		UserAccountModifyRemainTimes int64           `json:"user_account_modify_remain_times"`
// 		UserDealerOrgCode            interface{}     `json:"user_dealer_org_code"`
// 		UserID                       string          `json:"user_id"`
// 		UserLevel                    string          `json:"user_level"`
// 		UserMasterOrgID              string          `json:"user_master_org_id"`
// 		UserMasterOrgName            string          `json:"user_master_org_name"`
// 		UserMasterOrgTimeZoneID      string          `json:"user_master_org_time_zone_id"`
// 		UserMasterOrgTimeZoneName    string          `json:"user_master_org_time_zone_name"`
// 		UserName                     string          `json:"user_name"`
// 		UserRoleIDList               []string        `json:"user_role_id_list"`
// 		UserTelNationCode            interface{}     `json:"user_tel_nation_code"`
// 		UserAuthorURL                []interface{}   `json:"userauthorURL"`
// 		UserAuthorButto              []string        `json:"userauthorbutto"`
// 		UserDescription              interface{}     `json:"userdesc"`
// 		UsePassword                  string          `json:"userpassword"`
// 		ValidFlag                    string          `json:"validflag"`
// 		Voice                        string          `json:"voice"`
// 		WelcomeText                  string          `json:"welcometext"`
// 		Wgs84Latitude                interface{}     `json:"wgs84_latitude"`
// 		Wgs84Longitude               interface{}     `json:"wgs84_longitude"`
// 		WorkTel                      interface{}     `json:"work_tel"`
// 	} `json:"result_data"`
// }
//
// type LoginPrivileges []struct {
// 	FatherID        int64       `json:"father_id"`
// 	IconURL         interface{} `json:"icon_url"`
// 	IsOpen          interface{} `json:"is_open"`
// 	IsThirdPlatform int64       `json:"is_third_platform"`
// 	MenuCode        string      `json:"menu_code"`
// 	MenuLevel       int64       `json:"menu_level"`
// 	MenuName        string      `json:"menu_name"`
// 	MenuOrder       interface{} `json:"menu_order"`
// 	MenuType        string      `json:"menu_type"`
// 	MenuURL         string      `json:"menu_url"`
// 	PrivilegeID     int64       `json:"privilege_id"`
// 	RoleID          int64       `json:"role_id"`
// 	URLTarget       string      `json:"url_target"`
// 	VueIcon         interface{} `json:"vue_icon"`
// 	VuePath         interface{} `json:"vue_path"`
// }
//
//
// type LoginResource struct {
// 	Data Login
// }
//
// type LoginCollection struct {
// 	Data []Login
// }
