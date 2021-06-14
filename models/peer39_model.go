package model


type  Cat_score struct{

	category string `json:"cat"`
	score    string `json:"score"`

}

type  Peer_models struct{

     values []Cat_score `json:"response"`

}
