package signal

type Signal struct {
	Vin              string `json:"vin"`
	Address          string `json:"address"`
	LastOnChainBlock int64  `json:"last_on_chain_block"`
	LastOnChainHash  string `json:"last_on_chain_hash"`
	LastOnChainInfo  string `json:"last_on_chain_info"`
	CreateAt         string `json:"create_at"`
	UpdateAt         string `json:"update_at"`
}
