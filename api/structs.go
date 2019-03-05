package api

type hashPrms struct {
	Word string
	Salt string
}

type hashResponse struct {
	Hash    string
	Present bool
	Word    string
	Salt    string
}
