package iSolarCloud

import (
	"GoSungro/Only"
	"fmt"
)

func (p *SunGro) CountModels() error {

	for range Only.Once {
		query := p.Areas.Model.Count()
		if query.Error != nil {
			p.Error = query.Error
			break
		}

		if p.OutputType == TypeHuman {
			_, _ = fmt.Printf("Models: %v\n", query.Response)
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

func (p *SunGro) ListModels() error {

	for range Only.Once {
		query := p.Areas.Model.List()
		if query.Error != nil {
			p.Error = query.Error
			break
		}

		if p.OutputType == TypeHuman {
			_, _ = fmt.Printf("Models:\n%v\n", query.Response.String())
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
			// p.Error = p.UpdateGoogleSheet("model", data)
			break
		}
	}

	return p.Error
}

func (p *SunGro) ReadModels() error {

	for range Only.Once {
		query := p.Areas.Model.Read()
		if query.Error != nil {
			p.Error = query.Error
			break
		}

		switch p.OutputType {
			case TypeHuman:
				_, _ = fmt.Printf("Models:\n%v\n", query.Response.String())
				break

			case TypeJson:
				//_, _ = fmt.Fprintf(os.Stderr, "# Domains ")
				//_, _ = fmt.Printf("%s", query.Response.JsonString())
				p.OutputString = query.Response.JsonString()
				break

			case TypeGoogle:
				p.OutputArray = query.Response.ToArray()

				// data := query.Response.ToArray()
				// p.Error = p.UpdateGoogleSheet("model", data)
				break
		}
	}

	return p.Error
}
