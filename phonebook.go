package nfon

import "fmt"

type Phonebook struct {
	*Config
}

func (c *Config) NewPhonebook() *Phonebook {
	return &Phonebook{
		Config: c,
	}
}

func (p *Phonebook) Delete() bool {
	if status, err := p.NewRequest().
		send(DELETE, "/api/customers/"+p.sysid+"/phone-books", nil); status == DELETE_SUCCESS {
		return true
	} else {
		err.log()
		return false
	}
}

func (p *Phonebook) DeleteEntry(id string) bool {
	if status, err := p.NewRequest().
		send(DELETE, "/api/customers/"+p.sysid+"/phone-books/"+id, nil); status == DELETE_SUCCESS {
		return true
	} else {
		err.log()
		return false
	}
}

func (p *Phonebook) Get(offset int, pagesize int) apiResponse {
	query := "?_offset=" + fmt.Sprint(offset) + "&_pagesize=" + fmt.Sprint(pagesize)
	var data apiResponse

	if status, err := p.NewRequest().
		send(GET, "/api/customers/"+p.sysid+"/phone-books"+query, &data); status == GET_SUCCESS {
		return data
	} else {
		err.log()
		return apiResponse{}
	}
}

func (p *Phonebook) GetEntry(id string) apiResponse {
	var data apiResponse

	if status, err := p.NewRequest().
		send(GET, "/api/customers/"+p.sysid+"/phone-books/"+id, &data); status == GET_SUCCESS {
		return data
	} else {
		err.log()
		return apiResponse{}
	}
}

func (p *Phonebook) Post(name string, number string) apiResponse {
	var data apiResponse

	if status, err := p.NewRequest().
		AddData("displayName", name).
		AddData("displayNumber", number).
		send(POST, "/api/customers/"+p.sysid+"/phone-books", &data); status == POST_SUCCESS {
		return data
	} else {
		err.log()
		return apiResponse{}
	}
}

func (p *Phonebook) PutEntry(id string, name string, number string) bool {
	var data apiResponse

	if status, err := p.NewRequest().
		AddData("displayName", name).
		AddData("displayNumber", number).
		send(PUT, "/api/customers/"+p.sysid+"/phone-books/"+id, &data); status == PUT_SUCCESS {
		return true
	} else {
		err.log()
		return false
	}
}

func (p *Phonebook) GetEntryVisibilities(id string) apiResponse {
	var data apiResponse

	if status, err := p.NewRequest().
		send(GET, "/api/customers/"+p.sysid+"/phone-books/"+id+"/visibilities", &data); status == GET_SUCCESS {
		return data
	} else {
		err.log()
		return apiResponse{}
	}
}

func (p *Phonebook) GetEntryVisibilitiesExtension(id string, extension string) apiResponse {
	var data apiResponse

	if status, err := p.NewRequest().
		send(GET, "/api/customers/"+p.sysid+"/phone-books/"+id+"/visibilities/"+extension, &data); status == GET_SUCCESS {
		return data
	} else {
		err.log()
		return apiResponse{}
	}
}

func (p *Phonebook) PostEntryVisibilities(id string, extension string) apiResponse {
	var data apiResponse

	if status, err := p.NewRequest().
		AddLink("phoneExtension", "/api/customers/"+p.sysid+"/targets/phone-extensions/"+extension).
		send(POST, "/api/customers/"+p.sysid+"/phone-books/"+id+"/visibilities/", &data); status == POST_SUCCESS {
		return data
	} else {
		err.log()
		return apiResponse{}
	}
}

func (p *Phonebook) DeleteEntryVisibilitiesExtension(id string, extension string) apiResponse {
	var data apiResponse

	if status, err := p.NewRequest().
		send(DELETE, "/api/customers/"+p.sysid+"/phone-books/"+id+"/visibilities/"+extension, &data); status == DELETE_SUCCESS {
		return data
	} else {
		err.log()
		return apiResponse{}
	}
}

func (p *Phonebook) GetEntryVisibilitiesAvailableExtensions(id string, extension string) apiResponse {
	var data apiResponse

	if status, err := p.NewRequest().
		send(GET, "/api/customers/"+p.sysid+"/phone-books/"+id+"/visibilities/available-phone-extensions", &data); status == GET_SUCCESS {
		return data
	} else {
		err.log()
		return apiResponse{}
	}
}
