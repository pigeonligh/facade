package facade

type IArchive interface {
	Name() string
	Content() IContent
	AdditionalAttributes() IAttributes

	Children() []IArchive
}

type Archive struct {
	instance IArchive

	children  []IArchive
	nameIndex map[string]int
}

func NewArchive(instance IArchive) *Archive {
	children := make([]IArchive, 0)
	nameIndex := make(map[string]int)

	for i, child := range instance.Children() {
		children = append(children, NewArchive(child))
		nameIndex[child.Name()] = i
	}

	return &Archive{
		instance:  instance,
		children:  children,
		nameIndex: nameIndex,
	}
}

func (ar *Archive) Name() string {
	return ar.instance.Name()
}

func (ar *Archive) Content() IContent {
	return ar.instance.Content()
}

func (ar *Archive) AdditionalAttributes() IAttributes {
	return ar.instance.AdditionalAttributes()
}

func (ar *Archive) Children() []IArchive {
	return ar.children
}

func (ar *Archive) Find(childname string) *Archive {
	if childname == "" {
		return ar
	}

	if index, found := ar.nameIndex[childname]; found {
		return ar.children[index].(*Archive)
	}
	return nil
}

func (ar *Archive) Walk(path ...string) *Archive {
	if len(path) == 0 {
		return ar
	}

	next := ar.Find(path[0])
	if next == nil {
		return nil
	}
	return next.Walk(path[1:]...)
}
