package client

type Flag struct {
	Namespace   string     `json:"namespace,omitempty"`
	Name        string     `json:"name,omitempty"`
	Description string     `json:"description,omitempty"`
	Rollout     []*Rollout `json:"rollout,omitempty"`
}

type Rollout struct {
	Percentage int  `json:"percentage,omitempty"`
	Value      bool `json:"value"`
}
