package md

type JBOV struct {
	Cname          string `json:"cname"`
	Uniqid         string `json:"uniqid"`
	LastMountPoint string `json:"last-mount-point"`
	Volumes        map[string]*Volume `json:"volumes""`
	Rules          []Rule `json:"rules,omitempty"`
	Deleted        map[string]*Deleted `json:"deleted,omitempty"`
}

type Volume struct {
	Uniqid         string `json:"uniqid"`
	LastMountPoint string `json:"last-mount-point"`
	Deprecated     bool `json:"deprecated,omitempty""`
}

type Rule struct {
	Pattern        string `json:"pattern"`
	AtLeastACopyIn string `json:"at-least-a-copy-in,omitempty"`
	Ncopies        int `json:"ncopies,omitempty"`
}

type Deleted struct {
	Ts      int `json:"ts"`
	Pending []string `json:"pending"`
}
