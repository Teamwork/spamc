package spamc

import (
	"fmt"
	"os"
	"testing"
)

var addr = os.Getenv("SPAMC_SA_ADDRESS") + ":783"

// Basic test to confirm that the commands return *something* until we have more
// robust test in place.
func TestCommands(t *testing.T) {
	client := New(addr, 0)
	message := "Subject: Hello, world!\r\n\r\nTest message.\r\n"

	cases := []struct {
		name string
		fun  func(...string) (*Response, error)
	}{
		{"Check", client.Check},
		{"Skip", client.Skip},
		{"Symbols", client.Symbols},
		{"Report", client.Report},
		//{"ReportIgnoreWarning", client.ReportIgnoreWarning},
		{"ReportIfSpam", client.ReportIfSpam},
		{"Process", client.Process},
		{"Header", client.Headers},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("%v", tc.name), func(t *testing.T) {
			r, err := tc.fun(message)
			if err != nil {
				t.Fatal(err)
			}
			if r == nil {
				t.Fatal("r is nil")
			}
			if r.Code != ExOK {
				t.Errorf("Code != ExOk: %v", r.Code)
			}
		})
	}
}

func TestPing(t *testing.T) {
	client := New(addr, 0)
	r, err := client.Ping()
	if err != nil {
		t.Fatal(err)
	}
	if r == nil {
		t.Fatal("r is nil")
	}
	if r.Code != ExOK {
		t.Errorf("Code != ExOk: %v", r.Code)
	}
}

func TestTell(t *testing.T) {
	client := New(addr, 0)
	message := "Subject: Hello, world!\r\n\r\nTest message.\r\n"
	r, err := client.Tell([]string{message}, &map[string]string{
		"Message-class": "spam",
		"Set":           "local",
	})
	if err != nil {
		t.Fatal(err)
	}
	if r == nil {
		t.Fatal("r is nil")
	}
	if r.Code != ExOK {
		t.Errorf("Code != ExOk: %v", r.Code)
	}

}

func TestLearn(t *testing.T) {
	client := New(addr, 0)
	message := "Subject: Hello, world!\r\n\r\nTest message.\r\n"
	r, err := client.Learn(LearnHam, message)
	if err != nil {
		t.Fatal(err)
	}
	if r == nil {
		t.Fatal("r is nil")
	}
	if r.Code != ExOK {
		t.Errorf("Code != ExOk: %v", r.Code)
	}

}