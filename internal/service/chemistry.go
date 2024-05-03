package service

type Chemistry struct {
	ID         string `json:"id" bson:"_id"`
	Name       string `json:"name" bson:"name"`
	SellValue  int    `json:"sellValue" bson:"sellValue"`
	ProbeValue int    `json:"probeValue" bson:"probeValue"`
	ProbeCount int    `json:"probeCount" bson:"probeCount"`
}

func (c *Chemistry) NewID(s string) {
	c.ID = s
}

func (c *Chemistry) GetID() string {
	return c.ID
}

func (c *Chemistry) GetName() string {
	return c.Name
}
