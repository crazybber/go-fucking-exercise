package skill

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func TestJasonFunction(t *testing.T) {

	bolB, _ := json.Marshal(true)
	assertEq("true", string(bolB))

	intB, _ := json.Marshal(1)
	assertEq("1", string(intB))

	fltB, _ := json.Marshal(2.34)
	assertEq([]byte("2.34"), fltB)

	strB, _ := json.Marshal("gopher")
	assertEq([]byte(`"gopher"`), strB)

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD) //jason -->list
	assertEq([]byte(`["apple","peach","pear"]`), slcB)

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD) //Json -->object
	assertEq([]byte(`{"apple":5,"lettuce":7}`), mapB)

	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	assertEq([]byte(`{"Page":1,"Fruits":["apple","peach","pear"]}`), res1B)

	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	assertEq(`{"page":1,"fruits":["apple","peach","pear"]}`, string(res2B))

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	num := dat["num"].(float64)
	assertEq(6.13, num)

	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	assertEq("a", str1)
	assertEq("b", strs[1].(string))

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	assertEq(response2{Page: 1, Fruits: []string{"apple", "peach"}}, res)
	assertEq("apple", res.Fruits[0])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	t.Logf("show to console:")
	enc.Encode(d)

}

func jasonFunction() {

	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	num := dat["num"].(float64)
	fmt.Println(num)

	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
