package v_3_0_0

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const masterTemplateURIFmt = "https://raw.githubusercontent.com/giantswarm/k8scloudconfig-templates/%s/master_template.tmpl"

func GetMasterTemplate(version string) (string, error) {
	uri := fmt.Sprintf(masterTemplateURIFmt, version)

	resp, err := http.Get(uri)
	if err != nil {
		// handle error
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		return "", err
	}

	return string(body), nil
}
