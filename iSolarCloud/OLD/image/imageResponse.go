package image

import (
	"GoSungro/Only"
	"GoSungro/iSolarCloud/web"
	"fmt"
)

func (t *CountResponse) String() string {
	var str string
	for range Only.Once {
		str = fmt.Sprintf("%d", t.Total)
	}

	return str
}
func (t *CountResponse) JsonString() string {
	var str string
	for range Only.Once {
		j, err := web.PrintAsJson(t)
		if err != nil {
			break
		}
		str = fmt.Sprintf(j)
	}

	return str
}

func (t *CreateResponse) String() string {
	var str string
	for range Only.Once {
		str = fmt.Sprintf("%v\n", t)
	}

	return str
}
func (t *CreateResponse) JsonString() string {
	var str string
	for range Only.Once {
		j, err := web.PrintAsJson(t)
		if err != nil {
			break
		}

		str = fmt.Sprintf(j)
	}

	return str
}

func (t *DeleteResponse) String() string {
	var str string
	for range Only.Once {
		str = fmt.Sprintf("%v\n", t)
	}

	return str
}
func (t *DeleteResponse) JsonString() string {
	var str string
	for range Only.Once {
		j, err := web.PrintAsJson(t)
		if err != nil {
			break
		}

		str = fmt.Sprintf(j)
	}

	return str
}

func (t *ListResponse) String() string {
	var str string
	for range Only.Once {
		for i, k := range *t {
			if i == 0 {
				str += fmt.Sprintf("%s\n", web.PrintHeader(k))
			}
			str += fmt.Sprintf("%v\n", k.String())
		}
	}

	return str
}
func (t *ListResponse) JsonString() string {
	var str string
	for range Only.Once {
		j, err := web.PrintAsJson(t)
		if err != nil {
			break
		}

		str = fmt.Sprintf(j)
	}

	return str
}
func (t *ListResponse) ToArray() [][]interface{} {
	var ret [][]interface{}
	for i, k := range *t {
		if i == 0 {
			ret = append(ret, web.HeaderAsArray(k))
		}
		ret = append(ret, web.AsArray(k))
	}
	return ret
}

func (t *ListResponseSingle) String() string {
	return web.PrintValue(*t)
}
func (t *ListResponseSingle) JsonString() string {
	var str string
	for range Only.Once {
		j, err := web.PrintAsJson(t)
		if err != nil {
			break
		}

		str = fmt.Sprintf(j)
	}

	return str
}

func (t *ReadResponse) String() string {
	var str string
	for range Only.Once {
		str += web.PrintHeader(*t)
		str += web.PrintValue(*t)
	}

	return str
}
func (t *ReadResponse) JsonString() string {
	var str string
	for range Only.Once {
		j, err := web.PrintAsJson(t)
		if err != nil {
			break
		}

		str = fmt.Sprintf(j)
	}

	return str
}
func (t *ReadResponse) ToArray() [][]interface{} {
	var ret [][]interface{}
	for i, k := range *t {
		if i == 0 {
			ret = append(ret, web.HeaderAsArray(k))
		}
		ret = append(ret, web.AsArray(k))
	}
	return ret
}

func (t *ReadResponseSingle) String() string {
	return web.PrintValue(*t)
}
func (t *ReadResponseSingle) JsonString() string {
	var str string
	for range Only.Once {
		j, err := web.PrintAsJson(t)
		if err != nil {
			break
		}

		str = fmt.Sprintf(j)
	}

	return str
}

func (t *UpdateResponse) String() string {
	var str string
	for range Only.Once {
		str = fmt.Sprintf("%v\n", t)
	}

	return str
}
func (t *UpdateResponse) JsonString() string {
	var str string
	for range Only.Once {
		j, err := web.PrintAsJson(t)
		if err != nil {
			break
		}

		str = fmt.Sprintf(j)
	}

	return str
}
