package main


type Token struct {
	ID     int 		 	 `bson:"tokens_id"`
	AccessToken string  `bson:"accesstoken"`
	RefreshToken string `bson:"refreshtoken"`
}
