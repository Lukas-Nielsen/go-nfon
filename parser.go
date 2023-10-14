package nfon

func dataToMap(data []data) map[string]any {
	result := make(map[string]any)
	for _, entry := range data {
		result[entry.Name] = entry.Value
	}
	return result
}

func linksToMap(data []links) map[string]string {
	result := make(map[string]string)
	for _, entry := range data {
		result[entry.Rel] = entry.Href
	}
	return result
}

func (r response) parse(c *Response) {
	c.Href = r.Href
	c.Offset = r.Offset
	c.Size = r.Size
	c.Total = r.Total

	c.Links = linksToMap(r.Links)
	c.Data = dataToMap(r.Data)

	var t []Items
	for _, e := range r.Items {
		e := e
		t = append(t, Items{
			Href:  e.Href,
			Links: linksToMap(e.Links),
			Data:  dataToMap(e.Data),
		})
	}
	c.Items = t

}
