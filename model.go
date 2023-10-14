package nfon

type method string

type response struct {
	Href   string  `json:"href"`
	Offset int     `json:"offset"`
	Total  int     `json:"total"`
	Size   int     `json:"size"`
	Links  []links `json:"links"`
	Data   []data  `json:"data"`
	Items  []struct {
		Href  string  `json:"href"`
		Links []links `json:"links"`
		Data  []data  `json:"data"`
	} `json:"items"`
}

type Response struct {
	Href   string
	Offset int
	Total  int
	Size   int
	Links  map[string]string
	Data   map[string]any
	Items  []Items
}

type Items struct {
	Href  string
	Links map[string]string
	Data  map[string]any
}

type Request struct {
	client Client
	Links  []links `json:"links,omitempty"`
	Data   []data  `json:"data,omitempty"`
}

type Error struct {
	Detail      string   `json:"detail"`
	Title       string   `json:"title"`
	DescribedBy string   `json:"described_by"`
	Errors      []Errors `json:"errors"`
}

type Errors struct {
	Message string `json:"message"`
	Path    string `json:"path"`
	Value   string `json:"value"`
}

type data struct {
	Name  string `json:"name"`
	Value any    `json:"value"`
}

type links struct {
	Rel  string `json:"rel"`
	Href string `json:"href"`
}
