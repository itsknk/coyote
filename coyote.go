package coyote

import (
	"fmt"
	"bytes"
	"errors"
	"os"
	"io"
	"github.com/schollz/progressbar"
	"github.com/valyala/fasthttp"
)

type downloader struct {
	Client fasthttp.Client
	Debug bool
}

func (s *downloader) Coyote(url[]string, filename [] string) error {
	l := len(url)
	if l != len(filename) {
		return errors.New("Length won't match")
	}
	downch := make(chan []byte, l)
	errorch := make(chan error, l)

	for i, url := range url {
		go func(url string, filename string) {
			result, err := s.download(url, filename)
			if err != nil {
				errorch <- err
				downch <- nil
				return
			}
			downch <- result
			errorch <- err
		}(url, filename[i])
	}
	var errorString string
	for i:=0; i<l; i++ {
		if err := <-errorch; err != nil {
			errorString = errorString + " " + err.Error()
		}
	}
	var err error
	if errorString != "" {
		err = errors.New(errorString)
	}
	return err
}

func (s *downloader) download(url string, filename string) ([]byte, error) {
	if s.Debug == true {
		defer func() {
			fmt.Printf("Download Finished URL: %s, FILE: %s \n", url, filename)
		}()
	}
	statuscode, body, err := s.Client.Get(nil, url)
	if err != nil {
		return nil, err
	}
	if statuscode != 200 {
		return nil, errors.New("Not OK")
	}
	var out io.Writer;
	fn, err := os.Create(filename)
	if err != nil {
		return nil, err
	}
	pb := progressbar.NewOptions(
		int(len(body)),
		progressbar.OptionSetWidth(int(len(body))),
	)
	out = io.MultiWriter(fn, pb)
	var data bytes.Buffer
	rd := bytes.NewReader(body)
	_,err = io.Copy(out, rd)
	print("\n")
	if err != nil{
		return nil, err
	}
	return data.Bytes(), nil
}