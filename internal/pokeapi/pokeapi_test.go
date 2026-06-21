package pokeapi

import (
	"fmt"
	"testing"
)

//func TestGetLocation(t *testing.T) {
//	cases := []struct {
//		input1   string
//		input2   int
//		expected mapData
//	}{
//		{
//			input1:		"https://pokeapi.co/api/v2/location-area/",
//			input2:		1,
//			expected:	mapData{},
//		},
//		// add more cases here
//	}
//
//	for _, c := range cases {
//		_, err := GetLocation(c.input1, c.input2)
//		if err != nil {
//			t.Errorf(err.Error())
//		}
//	}
//}

func TestGetFromURL(t *testing.T) {
	cases := []struct {
		input string
	}{
		{ // Test anything is replied with other than an error
			input: "https://pokeapi.co/api/v2/location-area/1/",
		},
		// add more cases here
	}

	for _, c := range cases {
		_, err := getFromURL(c.input)
		if err != nil {
			t.Error(err)
		}
	}
}

func TestGetLocation(t *testing.T) {
	cases := []struct {
		url      string
		expected string
	}{
		{ // Test anything is replied with other than an error
			url:      "https://pokeapi.co/api/v2/location-area/1/",
			expected: "canalave-city-area",
		},
		// add more cases here
	}

	for _, c := range cases {
		actual, err := GetLocation(c.url)
		if err != nil {
			t.Error(err)
		}
		fmt.Print(actual)
		if string(actual) != c.expected {
			t.Errorf("Incorrect location returned\nExpected: %s\nActual: %s", actual, c.expected)
		}
	}
}
