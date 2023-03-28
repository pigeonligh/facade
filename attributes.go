package facade

type IAttributes interface {
	Get(key string) (interface{}, bool)
	Render() map[string]interface{}
}

type Attributes map[string]interface{}

func (attr Attributes) Get(key string) (interface{}, bool) {
	if attr == nil {
		return nil, false
	}
	ret, found := attr[key]
	return ret, found
}

func (attr Attributes) Render() map[string]interface{} {
	return attr
}

var EmptyAttributes = Attributes{}
