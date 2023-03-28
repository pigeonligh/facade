package facade

type IContent interface {
	Title() string
	Render() map[string]interface{}
}

type DefaultContent struct {
	title   string
	content string
}

func (c DefaultContent) Title() string {
	return c.title
}

func (c DefaultContent) Render() map[string]interface{} {
	ret := make(map[string]interface{})
	ret["title"] = c.title
	ret["content"] = c.content
	return ret
}

func NewEmptyContent(title string) *DefaultContent {
	return &DefaultContent{
		title: title,
	}
}

func NewDefaultContent(title, content string) *DefaultContent {
	return &DefaultContent{
		title:   title,
		content: content,
	}
}
