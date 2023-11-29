package dsl

type PrivateArgument struct {
	Url string `json:"url"`
}

type PublicArgument struct {
	Value string `json:"value"`
}

type Statements struct {
	Host string `json:"host"`
}

type ProtectionGadget struct {
	Host string `json:"host"`
}

type DSL struct {
	APIs        []ProtectionGadget `json:"gadgets"`
	Constraints []Statements       `json:"constraints"`
	Proxies     []PublicArgument   `json:"proxies"`
}
