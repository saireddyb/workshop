package model

type Movies struct {
	Id int `bson:"_id,omitempty"`
	Name string 
	Rating int 
	Year int 
	ImgUrl string 
	Category string 
	Watched bool 
}
