package latex

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"image"
	"io/ioutil"
	"net/http"
)

var Template = template.Must(template.New("latex").Parse(`
\documentclass{article}
\begin{document}
\[{{.}}\]
\pagenumbering{gobble}
\end{document}
`))

func GenerateLatexImage(latex string) (image.Image, error) {
	// Get filename
	buf := new(bytes.Buffer)
	Template.Execute(buf, latex)
	postData := Request{
		Code:   buf.String(),
		Format: "png",
	}
	b, err := json.Marshal(postData)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal postdata: %v", err)
	}
	resp, err := http.Post("http://rtex.probablyaweb.site/api/v2", "application/json", bytes.NewReader(b))
	if err != nil {
		return nil, fmt.Errorf("failed to post to latex server: %v", err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var lResp Response
	err = json.Unmarshal(body, &lResp)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %v", err)
	}
	if lResp.Status == Error {
		return nil, fmt.Errorf("failed to process input: %v", lResp.Description)
	}
	// Retrieve file
	resp, err = http.Get("http://rtex.probablyaweb.site/api/v2/" + lResp.Filename)
	if err != nil {
		return nil, fmt.Errorf("failed to get finished image: %v", err)
	}
	defer resp.Body.Close()

	img, _, _ := image.Decode(resp.Body)

	return img, nil
}
