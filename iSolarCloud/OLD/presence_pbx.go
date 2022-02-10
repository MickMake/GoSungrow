package iSolarCloud

import (
	"GoSungro/Only"
	"fmt"
)

func (p *SunGro) CountPresence(domain string) error {

	for range Only.Once {
		domain = p.VerifyDomain(domain)

		query := p.Areas.Presence.Count(domain)
		if query.Error != nil {
			p.Error = query.Error
			break
		}

		if p.OutputType == TypeHuman {
			_, _ = fmt.Printf("Presence: %v\n", query.Response.Total)
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

func (p *SunGro) ListPresence(domain string) error {

	for range Only.Once {
		domain = p.VerifyDomain(domain)

		query := p.Areas.Presence.List(domain)
		if query.Error != nil {
			p.Error = query.Error
			break
		}

		if p.OutputType == TypeHuman {
			_, _ = fmt.Printf("Presence:\n%v\n", query.Response.String())
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
			// p.Error = p.UpdateGoogleSheet("presence", data)
			break
		}
	}

	return p.Error
}

func (p *SunGro) ReadPresence(domain string) error {

	for range Only.Once {
		domain = p.VerifyDomain(domain)

		query := p.Areas.Presence.Read(domain)
		if query.Error != nil {
			p.Error = query.Error
			break
		}

		if p.OutputType == TypeHuman {
			_, _ = fmt.Printf("Presence:\n%v\n", query.Response.String())
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
			// p.Error = p.UpdateGoogleSheet("presence", data)
			break
		}
	}

	return p.Error
}
