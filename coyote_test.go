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
				"http://pluspng.com/img-png/google-logo-png-open-2000.png",
			}
			fileName := []string{
				"google.png",
			}
			err := downloader.Coyote(url, fileName)

			if err != nil {
				t.Error(err)
			}
		})

		t.Run("2 URLs", func(t *testing.T) {
			url := []string{
				"http://pluspng.com/img-png/google-logo-png-open-2000.png",
				"https://img.icons8.com/cotton/512/000000/chrome.png",
			}
			fileName := []string{
				"google.png",
				"chrome.png",
			}
			err := downloader.Coyote(url, fileName)

			if err != nil {
				t.Error(err)
			}
		})

		t.Run("Mismatch URL Array Length", func(t *testing.T) {
			urls := []string{
				"http://pluspng.com/img-png/google-logo-png-open-2000.png",
			}
			fileNames := []string{
				"google.png",
				"chrome.png",
			}
			err := downloader.Coyote(urls, fileNames)

			// Expecting the error not to be nil
			if err == nil {
				t.Error("Expected an error, the URL's length doesn't match the filename's length")
			}
		})

		t.Run("Mismatch Filename Array Length", func(t *testing.T) {
			urls := []string{
				"http://pluspng.com/img-png/google-logo-png-open-2000.png",
				"http://pluspng.com/img-png/google-logo-png-open-2000.png",
			}
			fileNames := []string{
				"google.png",
			}
			err := downloader.Coyote(urls, fileNames)

			// Expecting error here
			if err == nil {
				t.Error("Expected an error, the length of the URL array doesn't match the length of the filename array")
			}
		})

		t.Run("URL Doesn't Exist", func(t *testing.T) {
			url := []string{
				"http://randomImageWebsiteOrSomething.com",
			}
			fileName := []string{
				"google.png",
			}
			err := downloader.Coyote(url, fileName)

			if err == nil {
				t.Error("Expected a non 200 status code error")
			}
		})
	})
}