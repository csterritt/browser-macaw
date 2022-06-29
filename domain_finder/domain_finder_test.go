package domain_finder

import "testing"

func TestEmptyUrl(t *testing.T) {
	domain := DomainFromUrl("")
	expected := ""
	if domain != expected {
		t.Errorf("Expected '%s', got '%s'",
			expected, domain)
	}
}

func TestSimpleHttpUrl(t *testing.T) {
	domain := DomainFromUrl("http://domain.com")
	expected := "domain.com"
	if domain != expected {
		t.Errorf("Expected '%s', got '%s'",
			expected, domain)
	}
}

func TestSimpleHttpsUrl(t *testing.T) {
	domain := DomainFromUrl("https://wails.io")
	expected := "wails.io"
	if domain != expected {
		t.Errorf("Expected '%s', got '%s'",
			expected, domain)
	}
}

func TestDomainWithSubDomain(t *testing.T) {
	domain := DomainFromUrl("https://help.wails.io")
	expected := "wails.io"
	if domain != expected {
		t.Errorf("Expected '%s', got '%s'",
			expected, domain)
	}
}

func TestDomainWithDeepSubDomain(t *testing.T) {
	domain := DomainFromUrl("https://very.detailed.help.wails.io")
	expected := "wails.io"
	if domain != expected {
		t.Errorf("Expected '%s', got '%s'",
			expected, domain)
	}
}

func TestDomainWithShortTwoLetterDomain(t *testing.T) {
	domain := DomainFromUrl("https://fb.me")
	expected := "fb.me"
	if domain != expected {
		t.Errorf("Expected '%s', got '%s'",
			expected, domain)
	}
}

func TestDomainInForeignCountry(t *testing.T) {
	domain := DomainFromUrl("https://some.company.co.uk")
	expected := "company.co.uk"
	if domain != expected {
		t.Errorf("Expected '%s', got '%s'",
			expected, domain)
	}
}
