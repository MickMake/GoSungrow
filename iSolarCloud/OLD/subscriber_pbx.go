package iSolarCloud

import (
	"GoSungro/Only"
	"fmt"
)

func (p *SunGro) CountUsers(domain string) error {

	for range Only.Once {
		domain = p.VerifyDomain(domain)

		query := p.Areas.Subscriber.Count(domain)
		if query.Error != nil {
			p.Error = query.Error
			break
		}

		if p.OutputType == TypeHuman {
			_, _ = fmt.Printf("Users: %v\n", query.Response.Total)
			break
		}
		if p.OutputType == TypeJson {
			//_, _ = fmt.Fprintf(os.Stderr, "# Domains ")
			//_, _ = fmt.Printf("%s", query.Response.JsonString())
			p.OutputString = query.Response.JsonString()
			break
		}
	}

	return p.Error
}

func (p *SunGro) ListUsers(domain string) error {

	for range Only.Once {
		domain = p.VerifyDomain(domain)

		query := p.Areas.Subscriber.List(domain)
		if query.Error != nil {
			p.Error = query.Error
			break
		}

		if p.OutputType == TypeHuman {
			_, _ = fmt.Printf("Users:\n%v\n", query.Response.String())
			break
		}
		if p.OutputType == TypeJson {
			//_, _ = fmt.Fprintf(os.Stderr, "# Domains ")
			//_, _ = fmt.Printf("%s", query.Response.JsonString())
			p.OutputString = query.Response.JsonString()
			break
		}
		if p.OutputType == TypeGoogle {
			p.OutputArray = query.Response.ToArray()

			// data := query.Response.ToArray()
			// p.Error = p.UpdateGoogleSheet("user", data)
			break
		}
	}

	return p.Error
}

func (p *SunGro) ReadUsers(domain string, user string) error {

	for range Only.Once {
		domain = p.VerifyDomain(domain)
		user = p.VerifyUser(user)

		query := p.Areas.Subscriber.Read(domain, user)
		if query.Error != nil {
			p.Error = query.Error
			break
		}

		if p.OutputType == TypeHuman {
			_, _ = fmt.Printf("Users:\n%v\n", query.Response.String())
			break
		}
		if p.OutputType == TypeJson {
			//_, _ = fmt.Fprintf(os.Stderr, "# Domains ")
			//_, _ = fmt.Printf("%s", query.Response.JsonString())
			p.OutputString = query.Response.JsonString()
			break
		}
		if p.OutputType == TypeGoogle {
			p.OutputArray = query.Response.ToArray()

			// data := query.Response.ToArray()
			// p.Error = p.UpdateGoogleSheet("user", data)
			break
		}
	}

	return p.Error
}
