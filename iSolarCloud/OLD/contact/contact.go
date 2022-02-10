package contact

import (
	"GoSungro/Only"
	"errors"
	"net/url"
)

func (t *Contact) SetBaseUrl(u string) error {
	for range Only.Once {
		if u == "" {
			t.Error = errors.New("url is empty")
			break
		}

		t.Web.Url, t.Error = url.Parse(u)
		if t.Error != nil {
			break
		}

		t.count.Url = t.Web.Url
		t.create.Url = t.Web.Url
		t.delete.Url = t.Web.Url
		t.list.Url = t.Web.Url
		t.read.Url = t.Web.Url
		t.update.Url = t.Web.Url
	}

	return t.Error
}

func (t *Contact) Count(domain string, user string) Count {
	for range Only.Once {
		t.count.Request = CountRequest{
			Domain: domain,
			User:   user,
		}

		t.count.Error = t.Web.Get(t.count, t.count.Request, &t.count.Response)
		if t.count.Error != nil {
			break
		}
	}

	return t.count
}

func (t *Contact) Create(domain string, user string, firstName string, lastName string) (Create, error) {
	for range Only.Once {
		t.create.Request = CreateRequest{
			Domain:    domain,
			User:      user,
			FirstName: firstName,
			LastName:  lastName,
		}

		t.create.Error = t.Web.Get(t.create, t.create.Request, &t.create.Response)
		if t.create.Error != nil {
			break
		}
	}

	return t.create, t.Error
}

func (t *Contact) Delete(domain string, user string) (Delete, error) {
	for range Only.Once {
		t.delete.Request = DeleteRequest{
			Domain: domain,
			User:   user,
		}

		t.delete.Error = t.Web.Get(t.delete, t.delete.Request, &t.delete.Response)
		if t.delete.Error != nil {
			break
		}
	}

	return t.delete, t.Error
}

func (t *Contact) List(domain string, user string) List {
	for range Only.Once {
		t.list.Request = ListRequest{
			Domain: domain,
			User:   user,
		}

		t.list.Error = t.Web.Get(t.list, t.list.Request, &t.list.Response)
		if t.list.Error != nil {
			break
		}
	}

	return t.list
}

func (t *Contact) Read(domain string, user string) Read {
	for range Only.Once {
		t.read.Request = ReadRequest{
			Domain: domain,
			User:   user,
		}

		t.read.Error = t.Web.Get(t.read, t.read.Request, &t.read.Response)
		if t.read.Error != nil {
			break
		}
	}

	return t.read
}

func (t *Contact) Update(domain string, user string, firstName string, lastName string) (Update, error) {
	for range Only.Once {
		t.update.Request = UpdateRequest{
			Domain:    domain,
			User:      user,
			FirstName: firstName,
			LastName:  lastName,
		}

		t.update.Error = t.Web.Get(t.update, t.update.Request, &t.update.Response)
		if t.update.Error != nil {
			break
		}
	}

	return t.update, t.Error
}
