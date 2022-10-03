package client

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"steampipe-plugin-aci/container"
)

type ServiceManager struct {
	MOURL  string
	client *Client
}

func NewServiceManager(moURL string, client *Client) *ServiceManager {

	sm := &ServiceManager{
		MOURL:  moURL,
		client: client,
	}
	return sm
}

func (sm *ServiceManager) Get(dn string) (*container.Container, error) {
	finalURL := fmt.Sprintf("%s/%s.json", sm.MOURL, dn)
	req, err := sm.client.MakeRestRequest("GET", finalURL, nil, true)

	if err != nil {
		return nil, err
	}

	obj, _, err := sm.client.Do(req)
	if err != nil {
		return nil, err
	}

	if obj == nil {
		return nil, errors.New("Empty response body")
	}
	log.Printf("[DEBUG] Exit from GET %s", finalURL)
	return obj, CheckForErrors(obj, "GET", sm.client.skipLoggingPayload)
}

func createJsonPayload(payload map[string]string) (*container.Container, error) {
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
			}
		}
	}`, payload["classname"]))

	return container.ParseJSON(containerJSON)
}

func StripQuotes(word string) string {
	if strings.HasPrefix(word, "\"") && strings.HasSuffix(word, "\"") {
		return strings.TrimSuffix(strings.TrimPrefix(word, "\""), "\"")
	}
	return word
}

func G(cont *container.Container, key string) string {
	return StripQuotes(cont.S(key).String())
}

// CheckForErrors parses the response and checks of there is an error attribute in the response
func CheckForErrors(cont *container.Container, method string, skipLoggingPayload bool) error {
	number, err := strconv.Atoi(G(cont, "totalCount"))
	if err != nil {
		if !skipLoggingPayload {
			log.Printf("[DEBUG] Exit from errors, Unable to parse error count from response %v", cont)
		} else {
			log.Printf("[DEBUG] Exit from errors %s", err.Error())
		}
		return err
	}
	imdata := cont.S("imdata").Index(0)
	if number > 0 {

		if imdata.Exists("error") {
			errorCode := StripQuotes(imdata.Path("error.attributes.code").String())
			// Ignore errors of type "Cannot create object"
			if errorCode == "103" {
				if !skipLoggingPayload {
					log.Printf("[DEBUG] Exit from error 103 %v", cont)
				}
				return nil
			} else if method == "DELETE" && (errorCode == "1" || errorCode == "107") { // Ignore errors of type "Cannot delete object"
				if !skipLoggingPayload {
					log.Printf("[DEBUG] Exit from error 1 or 107 %v", cont)
				}
				return nil
			} else {
				if StripQuotes(imdata.Path("error.attributes.text").String()) == "" && errorCode == "403" {
					if !skipLoggingPayload {
						log.Printf("[DEBUG] Exit from authentication error 403 %v", cont)
					}
					return errors.New("Unable to authenticate. Please check your credentials")
				}
				if !skipLoggingPayload {
					log.Printf("[DEBUG] Exit from errors %v", cont)
				}

				return errors.New(StripQuotes(imdata.Path("error.attributes.text").String()))
			}
		}

	}

	if imdata.String() == "{}" && method == "GET" {
		if !skipLoggingPayload {
			log.Printf("[DEBUG] Exit from error (Empty response) %v", cont)
		}

		return errors.New("Error retrieving Object: Object may not exists")
	}
	if !skipLoggingPayload {
		log.Printf("[DEBUG] Exit from errors %v", cont)
	}
	return nil
}

func (sm *ServiceManager) GetViaURL(url string) (*container.Container, error) {
	req, err := sm.client.MakeRestRequest("GET", url, nil, true)

	if err != nil {
		return nil, err
	}

	obj, _, err := sm.client.Do(req)
	if !sm.client.skipLoggingPayload {
		log.Printf("Getvia url %+v", obj)
	}
	if err != nil {
		return nil, err
	}

	if obj == nil {
		return nil, errors.New("Empty response body")
	}
	return obj, CheckForErrors(obj, "GET", sm.client.skipLoggingPayload)

}
