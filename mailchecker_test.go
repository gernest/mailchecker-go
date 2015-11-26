package mailchecker

import "testing"

func TestGetList(t *testing.T) {
	_, err := getList()
	if err != nil {
		t.Error(err)
	}
}

func TestCheckDisposable(t *testing.T) {
	sampleFound := []string{
		"20minutemail.com", "2prong.com", "maileater.com", "mailexpire.com",
	}
	for _, v := range sampleFound {
		_, err := checkDisposable(v)
		if err != nil {
			t.Error(err)
		}
	}
}

func TestValid(t *testing.T) {
	sample := []struct {
		email string
		pass  bool
	}{
		{"gernest@20minutemail.com", false},
		{"gernest@live.com", true},
	}
	for _, v := range sample {
		ok, err := Valid(v.email)
		if ok != v.pass {
			t.Errorf("expected %v  for %sgot %v %v", v.pass, v.email, ok, err)
		}
	}
}
