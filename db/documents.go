package documents

type ProductDocument struct {
	Id string `bson:"_id,omitempty"`
	Name string
  Phone string
}
