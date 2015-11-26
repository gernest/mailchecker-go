package mailchecker

import (
	"encoding/json"
	"errors"
	"sort"
	"strings"

	"github.com/asaskevich/govalidator"
)

var mainData = func() map[string][]string {
	d, err := getList()
	if err != nil {
		panic(err)
	}
	for k, v := range d {
		sort.Sort(sort.StringSlice(v))
		d[k] = v
	}
	return d
}()

func getList() (map[string][]string, error) {
	var rst map[string][]string
	data, err := Asset("list.json")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &rst)
	if err != nil {
		return nil, err
	}
	return rst, nil
}

func checkDisposable(emailProvider string) (string, error) {
	for _, v := range mainData {
		found := sort.SearchStrings(v, emailProvider)
		if found != len(v) {
			if v[found] == emailProvider {
				return v[found], nil
				break
			}
		}
	}
	return "", errors.New("no matches found")
}

// Valid validates email. If the email is temporary it returns false and an error explaining
// why it is not valid. For valid email it returns true, and the error is nil.
func Valid(email string) (bool, error) {
	if !govalidator.IsEmail(email) {
		return false, errors.New("Bad email address")
	}
	sep := strings.Split(email, "@")
	_, err := checkDisposable(sep[1])
	if err != nil {
		return true, nil
	}
	return false, errors.New("This is a temporary email")
}
