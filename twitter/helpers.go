package twitter

import (
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func httpRequest(url, method string, headers map[string][]string) (*string, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Errorf("Error creating request:\n %v", err)
		return nil, err
	}
	req.Header = headers

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Errorf("Error making request:\n %v", err)
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Errorf("Error reading res.Body:\n %v", err)
		return nil, err
	}
	s := string(body)
	log.WithFields(log.Fields{
		"body": s,
	}).Debug()
	return &s, nil
}
