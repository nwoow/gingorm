package models

//City Model
type City struct {
	CityName string
	StateID  int
	CityID   int
}

//AngelClient ...
type AngelClient struct {
	ClientType     string
	Clients        int
	OfflineClients int
}

//UpdateBookInput ...
type UpdateBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

//Datacode ...
type Datacode struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}
