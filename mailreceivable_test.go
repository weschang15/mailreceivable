package mailreceivable

import "testing"

type address struct {
	email          string
	validateResult bool
	verifyMx       bool
	verifyHost     bool
}

var (
	addresses = []address{
		{email: "wes@wildearth.com", validateResult: true, verifyMx: true, verifyHost: true},
		{email: "email@example.com", validateResult: true, verifyMx: false, verifyHost: false},
		{email: "firstname.lastname@example.com", validateResult: true, verifyMx: false, verifyHost: false},
		{email: "email@subdomain.example.com", validateResult: true, verifyMx: false, verifyHost: false},
		{email: "firstname+lastname@example.com", validateResult: true, verifyMx: false, verifyHost: false},
		{email: "email@123.123.123.123", validateResult: true, verifyMx: false, verifyHost: false},
		{email: "#@%^%#$@#$@#.com", validateResult: false, verifyMx: false, verifyHost: false},
		{email: "@example.com", validateResult: false, verifyMx: false, verifyHost: false},
		{email: "Joe Smith <email@example.com>", validateResult: false, verifyMx: false, verifyHost: false},
		{email: "email.example.com", validateResult: false, verifyMx: false, verifyHost: false},
		{email: "email@example@example.com", validateResult: false, verifyMx: false, verifyHost: false},
		{email: ".email@example.com", validateResult: false, verifyMx: false, verifyHost: false},
		{email: "email.@example.com", validateResult: false, verifyMx: false, verifyHost: false},
		{email: "email..email@example.com", validateResult: false, verifyMx: false, verifyHost: false},
		{email: "あいうえお@example.com", validateResult: false, verifyMx: false, verifyHost: false},
		{email: "email@example", validateResult: false, verifyMx: false, verifyHost: false},
		{email: "email@-example.com", validateResult: false, verifyMx: false, verifyHost: false},
		{email: "email@example.web", validateResult: false, verifyMx: false, verifyHost: false},
		{email: "email@111.222.333.44444", validateResult: false, verifyMx: false, verifyHost: false},
		{email: "email@example..com", validateResult: false, verifyMx: false, verifyHost: false},
		{email: "Abc..123@example.com", validateResult: false, verifyMx: false, verifyHost: false},
	}
)

func TestValidate(t *testing.T) {
	for _, address := range addresses {
		err := Validate(address.email)
		if err != nil && address.validateResult == true {
			t.Errorf(`Validate("%s") FAILED. Unexpected error: "%v".`, address.email, err)
		}

		if err == nil && address.validateResult == false {
			t.Logf(`Validate("%s") PASSED. Expected error.`, address.email)
		}
	}
}

func TestVerifyHost(t *testing.T) {
	for _, address := range addresses {
		err := VerifyHost(address.email)
		if err != nil && address.verifyHost == true {
			t.Errorf(`VerifyHost("%s") FAILED. Unexpected error: "%v"`, address.email, err)
		}

		if err == nil && address.verifyHost == false {
			t.Logf(`VerifyHost("%s") PASSED. Expected error.`, address.email)
		}
	}
}

func TestVerifyMX(t *testing.T) {
	for _, address := range addresses {
		err := VerifyMX(address.email)
		if err != nil && address.verifyMx == true {
			t.Errorf(`VerifyMX("%s") FAILED. Unexpected error: "%v"`, address.email, err)
		}

		if err == nil && address.verifyMx == false {
			t.Logf(`VerifyMX("%s") PASSED. Expected error.`, address.email)
		}
	}
}
