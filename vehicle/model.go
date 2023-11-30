package vehicle

type Vehicle struct {
	Vin         string `redis:"vin" json:"vin"`
	PublicKey   string `redis:"public_key" json:"public_key"`
	PrivateKey  string `redis:"private_key" json:"-"`
	MobilePhone string `redis:"mobile_phone" json:"mobile_phone"`
}
