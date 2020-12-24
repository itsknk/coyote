package coyote_test

import (
	"testing"

	"github.com/itsknk/coyote"
)

func TestDownloadFunc(t *testing.T) {
	t.Run("Main", func(t *testing.T) {
		t.Parallel()
		downloader := coyote.Downloader{
			Debug: true,
		}

		t.Run("1 URL", func(t *testing.T) {
			url := []string{
				"https://images-na.ssl-images-amazon.com/images/I/81P1kr0i2RL._RI_.jpg",
			}
			fileName := []string{
				"entourage.jpg",
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
				"entourage_one.jpg",
				"entourage_two.jpg",
			}
			err := downloader.Coyote(url, fileName)

			if err != nil {
				t.Error(err)
			}
		})

		t.Run("Mismatch URL Array Length", func(t *testing.T) {
			urls := []string{
				"https://images-na.ssl-images-amazon.com/images/I/81P1kr0i2RL._RI_.jpg",
			}
			fileNames := []string{
				"entourage_one.jpg",
				"entourage_two.jpg",
			}
			err := downloader.Coyote(urls, fileNames)

			// Expecting the error not to be nil
			if err == nil {
				t.Error("Expected an error, the URL's length doesn't match the filename's length")
			}
		})

		t.Run("Mismatch Filename Array Length", func(t *testing.T) {
			urls := []string{
				"https://images-na.ssl-images-amazon.com/images/I/81P1kr0i2RL._RI_.jpg",
				"https://www.telegraph.co.uk/content/dam/film/Entourage/entouragetv-xlarge.jpg",
			}
			fileNames := []string{
				"entourage.jpg",
			}
			err := downloader.Coyote(urls, fileNames)

			// Expecting error here
			if err == nil {
				t.Error("Expected an error, the length of the URL array doesn't match the length of the filename array")
			}
		})

		t.Run("URL Doesn't Exist", func(t *testing.T) {
			url := []string{
				"http://SomeMadeUpURLForTheSakeOfTest.com",
			}
			fileName := []string{
				"entourage.jpg",
			}
			err := downloader.Coyote(url, fileName)

			if err == nil {
				t.Error("Expected a non 200 status code error")
			}
		})
	})
}
