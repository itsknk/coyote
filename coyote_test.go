package coyote

import (
	"testing"
)

func FileDownloadTest(t *testing.T) {
	t.Run("Main", func(t *testing.T) {
		t.Parallel()
		downloader := Downloader{
			Debug: true,
		}

		t.Run("1 URL", func(t *testing.T) {
			url := []string{
				"https://images-na.ssl-images-amazon.com/images/I/81P1kr0i2RL._RI_.jpg",
			}
			fileName := []string{
				"something.png",
			}
			err := downloader.Coyote(url, fileName)

			if err != nil {
				t.Error(err)
			}
		})

		t.Run("2 URLs", func(t *testing.T) {
			url := []string{
				"https://images-na.ssl-images-amazon.com/images/I/81P1kr0i2RL._RI_.jpg",
				"https://www.telegraph.co.uk/content/dam/film/Entourage/entouragetv-xlarge.jpg",
			}
			fileName := []string{
				"something.png",
				"nothing.png",
			}
			err := downloader.Coyote(url, fileName)

			if err != nil {
				t.Error(err)
			}
		})

		t.Run("Mismatch URL Array Length", func(t *testing.T) {
			urls := []string{
				"https://www.telegraph.co.uk/content/dam/film/Entourage/entouragetv-xlarge.jpg",
			}
			fileNames := []string{
				"something.png",
				"nothing.png",
			}
			err := downloader.Coyote(urls, fileNames)

			if err == nil {
				t.Error("Expected Error: the URL's length doesn't match the filename's length")
			}
		})

		t.Run("Mismatch Filename Array Length", func(t *testing.T) {
			urls := []string{
				"https://images-na.ssl-images-amazon.com/images/I/81P1kr0i2RL._RI_.jpg",
				"https://www.telegraph.co.uk/content/dam/film/Entourage/entouragetv-xlarge.jpg",
			}
			fileNames := []string{
				"something.png",
			}
			err := downloader.Coyote(urls, fileNames)

			// Expecting error here
			if err == nil {
				t.Error("Expected Error: the length of the URL array doesn't match the length of the filename array")
			}
		})

		t.Run("URL Doesn't Exist", func(t *testing.T) {
			url := []string{
				"http://somethingelseforthesakeoftestingthispkg.com",
			}
			fileName := []string{
				"something.png",
			}
			err := downloader.Coyote(url, fileName)

			if err == nil {
				t.Error("Not OK")
			}
		})
	})
}