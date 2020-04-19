package latex

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
)

var LatexTemplate = template.Must(template.New("latex").Parse(`
\documentclass{article}
\begin{document}
{{.}}
\pagenumbering{gobble}
\end{document}
`))

func GeneratePNGFromLatex(latex string) (io.Reader, error) {
	// Get filename
	buf := new(bytes.Buffer)
	LatexTemplate.Execute(buf, latex)
	postData := LatexRequest{
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
	var lResp LatexResponse
	err = json.Unmarshal(body, &lResp)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %v", err)
	}
	if lResp.Status == Error {
		return nil, fmt.Errorf("failed to process input: %v", lResp.Description)
	}
	// Retrieve file
	fmt.Println(lResp.Filename)
	resp, err = http.Get("http://rtex.probablyaweb.site/api/v2/" + lResp.Filename)
	if err != nil {
		return nil, fmt.Errorf("failed to get finished image: %v", err)
	}
	defer resp.Body.Close()
	body, _ = ioutil.ReadAll(resp.Body)
	return bytes.NewReader(body), nil
}
