package youtrack

import (
	"io"
	"net/url"
	"net/http"
	"encoding/json"
	"bytes"
)

type Repository interface {
	
}

type RepositoryHelper struct {
	urlValues url.Values
}
func (this *RepositoryHelper) GetJson (model interface{}) io.Reader {
	s, _ := json.Marshal(model)
	b := bytes.NewBuffer(s)
	return b
}
func (this *RepositoryHelper) Load(res *http.Response, model interface{}) {
	json.NewDecoder(res.Body).Decode(model)
}