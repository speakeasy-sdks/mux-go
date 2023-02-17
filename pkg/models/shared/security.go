package shared

type SchemeAccessToken struct {
	Password string `security:"name=password"`
	Username string `security:"name=username"`
}

type Security struct {
	AccessToken SchemeAccessToken `security:"scheme,type=http,subtype=basic"`
}
