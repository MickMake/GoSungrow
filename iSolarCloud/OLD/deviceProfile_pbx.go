package iSolarCloud

import (
	"GoSungro/Only"
	"fmt"
)

func (p *SunGro) CountProfiles() error {

	for range Only.Once {
		query := p.Areas.Profile.Count()
		if query.Error != nil {
			p.Error = query.Error
			break
		}

		if p.OutputType == TypeHuman {
			_, _ = fmt.Printf("Profiles: %v\n", query.Response.Total)
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

func (p *SunGro) ListProfiles() error {

	for range Only.Once {
		query := p.Areas.Profile.List()
		if query.Error != nil {
			p.Error = query.Error
			break
		}

		if p.OutputType == TypeHuman {
			_, _ = fmt.Printf("Profiles:\n%v\n", query.Response.String())
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
			// p.Error = p.UpdateGoogleSheet("profile", data)
			break
		}
	}

	return p.Error
}

func (p *SunGro) ReadProfiles() error {

	for range Only.Once {
		query := p.Areas.Profile.Read()
		if query.Error != nil {
			p.Error = query.Error
			break
		}

		if p.OutputType == TypeHuman {
			_, _ = fmt.Printf("Profiles:\n%v\n", query.Response.String())
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
			// p.Error = p.UpdateGoogleSheet("profile", data)
			break
		}
	}

	return p.Error
}
