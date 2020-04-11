package util

import (
	"fmt"
	"github.com/jonas747/dca"
	"io"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestDCA(t *testing.T) {
	fP, _ := filepath.Abs("assets/test.mp3")

	encSes, err := dca.EncodeFile(fP, dca.StdEncodeOptions)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(encSes)
	fmt.Println(encSes.Stats())

	defer encSes.Cleanup()
	output, _ := os.Create("assets/test.dca")

	log.Println(encSes.FFMPEGMessages())

	io.Copy(output, encSes)
}
