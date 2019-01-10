package caps

import (
	"testing"
)

func TestClear(t *testing.T) {
	c, err := FromText("cap_chown=e")
	if err != nil {
		t.Fatal("FromText: ", err)
	}

	cfv, err := c.GetFlag(CAP_CHOWN, CAP_EFFECTIVE)
	if err != nil {
		t.Fatal("GetFlag: ", err)
	}
	if cfv != CAP_SET {
		t.Errorf("CAP_CHOWN/CAP_EFFECTIVE should be set")
	}
}

func TestClearFlag(t *testing.T) {
	c, err := FromText("cap_chown=e")
	if err != nil {
		t.Fatal("FromText: ", err)
	}

	if err := c.ClearFlag(CAP_EFFECTIVE); err != nil {
		t.Fatal("ClearFlag CAP_EFFECTIVE: ", err)
	}

	cfv, err := c.GetFlag(CAP_CHOWN, CAP_EFFECTIVE)
	if err != nil {
		t.Fatal("GetFlag CAP_EFFECTIVE/CAP_CHOWN: ", err)
	}
	if cfv != CAP_CLEAR {
		t.Errorf("CAP_EFFECTIVE should not be set")
	}
}

func TestCompare(t *testing.T) {
	type cmpTest struct {
		text1    string
		text2    string
		expected error
	}

	var cmpTests = []cmpTest{
		{"", "", nil},
		{"cap_chown=e", "cap_chown=e", nil},
		{"cap_chown=e", "", ErrCapNotEqual},
		{"cap_chown=e", "cap_chown=i", ErrCapNotEqual},
	}

	for _, tt := range cmpTests {
		c1, _ := FromText(tt.text1)
		c2, _ := FromText(tt.text2)
		actual := Compare(*c1, *c2)
		if actual != tt.expected {
			t.Errorf(
				"Compare('%s', '%s'): expected %v, got %v",
				tt.text1,
				tt.text2,
				tt.expected,
				actual,
			)
		}
	}
}
