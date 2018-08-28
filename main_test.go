package main

import (
	"bytes"
	"reflect"
	"strings"
	"testing"
)

func TestFetchReplacements(t *testing.T) {
	testCases := []map[string]interface{}{
		map[string]interface{}{
			"data":     ``,
			"expected": []string{},
		},
		map[string]interface{}{
			"data":     `key`,
			"expected": []string{},
		},
		map[string]interface{}{
			"data":     `key=value`,
			"expected": []string{"key", "value"},
		},
		map[string]interface{}{
			"data":     `   key    =    value   `,
			"expected": []string{"key", "value"},
		},
		map[string]interface{}{
			"data": `   key    =    value   
			hello
			`,
			"expected": []string{"key", "value"},
		},
		map[string]interface{}{
			"data": `   key    =    value   
			wat = swat
			`,
			"expected": []string{"key", "value", "wat", "swat"},
		},
		map[string]interface{}{
			"data": `   key    =    value   

			j


			wat = swat


			`,
			"expected": []string{"key", "value", "wat", "swat"},
		},
	}

	for _, c := range testCases {
		d := c["data"].(string)
		ex := c["expected"].([]string)

		r := strings.NewReader(d)
		actual := fetchReplacements(r)
		if !reflect.DeepEqual(actual, ex) {
			t.Errorf("I expected that %+v and %+v are equals", actual, ex)
		}
	}
}

func TestReplace(t *testing.T) {
	testCases := []map[string]interface{}{
		map[string]interface{}{
			"src":    "db:<db_host>:<db_port>/<db_name>\napikey=<api_key>",
			"oldnew": []string{},
			"exp":    "db:<db_host>:<db_port>/<db_name>\napikey=<api_key>\n",
		},
		map[string]interface{}{
			"src":    "db:<db_host>:<db_port>/<db_name>\napikey=<api_key>",
			"oldnew": []string{"<db_host>", "host"},
			"exp":    "db:host:<db_port>/<db_name>\napikey=<api_key>\n",
		},
		map[string]interface{}{
			"src":    "",
			"oldnew": []string{"<db_host>", "host"},
			"exp":    "",
		},
		map[string]interface{}{
			"src":    "wat:<swat>",
			"oldnew": []string{"<db_host>", "host"},
			"exp":    "wat:<swat>\n",
		},
		map[string]interface{}{
			"src":    "db:<db_host>:<db_port>/<db_name>\napikey=<api_key>",
			"oldnew": []string{"<db_host>", "host", "<db_name>", "name", "<db_port>", "port", "<api_key>", "key"},
			"exp":    "db:host:port/name\napikey=key\n",
		},
	}

	for _, c := range testCases {
		src := c["src"].(string)
		oldnew := c["oldnew"].([]string)
		exp := c["exp"].(string)

		r := strings.NewReader(src)
		replacer := strings.NewReplacer(oldnew...)
		w := bytes.NewBufferString("")
		replace(r, replacer, w)

		actual := w.String()
		if actual != exp {
			t.Errorf("I expected string \"%s\" but got \"%s\"", exp, actual)
		}
	}
}
