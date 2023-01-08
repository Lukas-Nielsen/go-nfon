package nfon

import (
	"fmt"
)

type PhoneExtension struct {
	*Config
}

func (c *Config) NewPhoneExtension() *PhoneExtension {
	return &PhoneExtension{
		Config: c,
	}
}

func NewPhoneExtensionOption(name string, extension int) *Option {
	var temp Option
	return temp.
		SetData(DISPLAY_NAME, name).
		SetData(EXTENSION_NUMBER, string(rune(extension))).
		SetData(ABANDON_OTHER_SOFTPHONES, true).
		SetData(NCONTROL_ENABLED, true).
		SetData(CALL_WAITING_INDICATION, true).
		SetData(ACCESS_CENTRAL_PHONE_BOOK, true)
}

func (p *PhoneExtension) DeleteEntry(id string) bool {
	if status, err := p.NewRequest().
		send(DELETE, "/api/customers/"+p.sysid+"/targets/phone-extensions/"+id, nil); status == DELETE_SUCCESS {
		return true
	} else {
		err.log()
		return false
	}
}

func (p *PhoneExtension) Get(offset int, pagesize int) apiResponse {
	query := "?_offset=" + fmt.Sprint(offset) + "&_pagesize=" + fmt.Sprint(pagesize)
	var data apiResponse

	if status, err := p.NewRequest().
		send(GET, "/api/customers/"+p.sysid+"/targets/phone-extensions"+query, &data); status == GET_SUCCESS {
		return data
	} else {
		err.log()
		return apiResponse{}
	}
}

func (p *PhoneExtension) GetEntry(id string) apiResponse {
	var data apiResponse

	if status, err := p.NewRequest().
		send(GET, "/api/customers/"+p.sysid+"/targets/phone-extensions/"+id, &data); status == GET_SUCCESS {
		return data
	} else {
		err.log()
		return apiResponse{}
	}
}

func (p *PhoneExtension) Post(input *Option) apiResponse {
	var data apiResponse

	req := p.NewRequest()

	for name, value := range input.data {
		name := name
		value := value
		req = req.AddData(string(name), value)
	}

	for rel, href := range input.link {
		rel := rel
		href := href
		req = req.AddLink(string(rel), href)
	}

	if status, err := req.
		send(POST, "/api/customers/"+p.sysid+"/targets/phone-extensions", &data); status == POST_SUCCESS {
		return data
	} else {
		err.log()
		return apiResponse{}
	}
}

func (p *PhoneExtension) PutEntry(id string, input *Option) bool {
	var data apiResponse

	req := p.NewRequest()

	for name, value := range input.data {
		name := name
		value := value
		req = req.AddData(string(name), value)
	}

	for rel, href := range input.link {
		rel := rel
		href := href
		req = req.AddLink(string(rel), href)
	}

	if status, err := req.
		send(POST, "/api/customers/"+p.sysid+"/targets/phone-extensions/"+id, &data); status == PUT_SUCCESS {
		return true
	} else {
		err.log()
		return false
	}
}
