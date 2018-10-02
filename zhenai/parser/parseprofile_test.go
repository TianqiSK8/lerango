package parser

import (
	"crawler/model"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("parseprofile_test_data.html")
	if err != nil {
		panic(err)
	}


	result := ParseCity(contents)
	fmt.Println(len(result.Items))
	//if len(result.Items) != 1 {
	//	t.Errorf("Items should contail 1 element; but was %v", result.Items)
	//}

	profile := result.Items[0].(model.Profile)

	expected := model.Profile{
		Age:		"",
		Heigth:		"",
		Weigth:		"",
		Income:     "3001-5000",
		Gender:		"女",
		Name:		"asdad",
		Xinzuo:		"111",
		Occupation: "",
		Marriage:	"",
		House:		"",
		Hokou:		"",
		Car:		"",
		Education:	"",
	}

	if profile != expected {
		t.Errorf("expected %v: but was %v", expected, profile)
	}
}
