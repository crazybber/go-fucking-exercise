package skill

import (
	"encoding/xml"
	"fmt"
	"testing"
)

type Plant struct {
	XMLName xml.Name `xml:"plant"`
	ID      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}

func TestXmlFunction(t *testing.T) {
	coffee := &Plant{ID: 27, Name: "Coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}

	out, _ := xml.Marshal(coffee)
	t.Log(string(out))

	assertEq([]byte(`<plant id="27"><name>Coffee</name><origin>Ethiopia</origin><origin>Brazil</origin></plant>`), out)

	t.Log(xml.Header + string(out))

	assertEq(`<?xml version="1.0" encoding="UTF-8"?>`+"\n", xml.Header)

	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	t.Log(p)

	assertEq("Coffee", p.Name)
	assertEq(27, p.ID)

	tomato := &Plant{ID: 81, Name: "Tomato"}
	tomato.Origin = []string{"Mexico", "California"}

	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		Plants  []*Plant `xml:"parent>child>plant"`
	}

	nesting := &Nesting{}
	nesting.Plants = []*Plant{coffee, tomato}

	out, _ = xml.Marshal(nesting)

	assertEq(`<nesting><parent><child><plant id="27"><name>Coffee</name><origin>Ethiopia</origin><origin>Brazil</origin></plant><plant id="81"><name>Tomato</name><origin>Mexico</origin><origin>California</origin></plant></child></parent></nesting>`, string(out))

	//xmlFunction()
}
func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v", p.ID, p.Name, p.Origin)
}

func xmlFunction() {
	coffee := &Plant{ID: 27, Name: "Coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}

	out, _ := xml.MarshalIndent(coffee, " ", "  ")
	fmt.Println(string(out))

	fmt.Println(xml.Header + string(out))

	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	fmt.Println(p)

	tomato := &Plant{ID: 81, Name: "Tomato"}
	tomato.Origin = []string{"Mexico", "California"}

	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		Plants  []*Plant `xml:"parent>child>plant"`
	}

	nesting := &Nesting{}
	nesting.Plants = []*Plant{coffee, tomato}

	out, _ = xml.MarshalIndent(nesting, " ", "  ")
	fmt.Println(string(out))
}
