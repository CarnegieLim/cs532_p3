package db

type Record struct {
	Review string `json:"review"`
	Sentiment string `json:"sentiment"`
	CreateAt string `json:"create_at"`
}

type Result struct {
	CreateAt string `json:"create_at"`
	BinData []byte `bson:"bin_data"`
}