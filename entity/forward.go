package entity

type Rule struct {
	HasAttachment bool `yaml:"has_attachment,omitempty"`
}

type DataObject struct {
	Forward Forward `yaml: "forward"`
}

type Forward struct {
	Name    string `yaml:"name,omitempty"`
	Value   string `yaml: "value,omitempty"`
	Content string `yaml: "content,omitempty"`
	Rules   []Rule `yaml: "rule,omitempty"`
	Assign  string `yaml: "assign,omitempty"`
}
