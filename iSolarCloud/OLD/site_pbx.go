package iSolarCloud

import (
	"GoSungro/Only"
	"fmt"
)

func (p *SunGro) CountSites(domain string) error {

	for range Only.Once {
		domain = p.VerifyDomain(domain)

		query := p.Areas.Site.Count(domain)
		if query.Error != nil {
			p.Error = query.Error
			break
		}

		if p.OutputType == TypeHuman {
			_, _ = fmt.Printf("Sites: %v\n", query.Response.Total)
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

func (p *SunGro) ListSites(domain string) error {

	for range Only.Once {
		domain = p.VerifyDomain(domain)

		query := p.Areas.Site.List(domain)
		if query.Error != nil {
			p.Error = query.Error
			break
		}

		if p.OutputType == TypeHuman {
			_, _ = fmt.Printf("Sites:\n%v\n", query.Response.String())
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
			// p.Error = p.UpdateGoogleSheet("site", data)
			break
		}
	}

	return p.Error
}

func (p *SunGro) ReadSites(domain string) error {

	for range Only.Once {
		domain = p.VerifyDomain(domain)

		query := p.Areas.Site.Read(domain)
		if query.Error != nil {
			p.Error = query.Error
			break
		}

		if p.OutputType == TypeHuman {
			_, _ = fmt.Printf("Sites:\n%v\n", query.Response.String())
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
			// p.Error = p.UpdateGoogleSheet("site", data)
			break
		}
	}

	return p.Error
}
