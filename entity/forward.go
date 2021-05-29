package entity

type Rule struct {
}

type DataObject struct {
	Forward Forward `yaml: "forward"`
}

type Forward struct {
	Name    string `yaml:"name,omitempty"`
	Value   string `yaml:"value"`
	Content string `yaml:"content,omitempty"`
	Rules   []Rule `yaml:"rules,omitempty"`
	Squad   string `yaml:"squad,omitempty"`
}
