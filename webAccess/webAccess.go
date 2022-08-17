package webAccess

import "net/http"

// Read websites and other internet based resources

type WebAction struct {
	read int
}

type webAccess struct {
	url    string
	action WebAction
}

func (w webAccess) readWebsite() (bool, error) {
	_, err := http.Get(w.url)
	if err != nil {
		return false, err
	}
	return true, nil
}
