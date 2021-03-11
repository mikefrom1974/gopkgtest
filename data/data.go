package data

type MyData struct {
	TestString string
}

func New(s string) *MyData {
	d := new(MyData)
	d.TestString = s
	return d
}

func (m *MyData) Fetch() string {
	return m.TestString
}
