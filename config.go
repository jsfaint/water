package main

import (
	"io/ioutil"
	"strings"
)

func toLine(str string) (tables []string) {
	var sep string
	if strings.Contains(str, "\r\n") {
		sep = "\r\n"
	} else {
		sep = "\n"
	}

	for _, v := range strings.Split(str, sep) {
		//blank line
		if v == "" {
			continue
		}

		//comment
		if strings.HasPrefix(v, "#") {
			continue
		}

		tables = append(tables, v)
	}

	return tables
}

func getLineFromFile(name string) []string {
	s, err := ioutil.ReadFile(name)
	if err != nil {
		return []string{}
	}

	return toLine(string(s))
}

func loadCronFromFile(name string) []string {
	return getLineFromFile(name)
}
