package service

type Cargo struct {
	ID         string `json:"id" bson:"_id"`
	Name       string `json:"name" bson:"name"`
	URI        string `json:"uri" bson:"uri"`
	MainTel    string `json:"mainTel" bson:"mainTel"`
	ManagerTel string `json:"managerTel" bson:"managerTel"`
}

func (c *Cargo) NewID(s string) {
	c.ID = s
}

func (c *Cargo) GetID() string {
	return c.ID
}

func (c *Cargo) GetName() string {
	return c.Name
}
