package ytdl

import (
  "testing"
)

func TestDownloadURL(t *testing.T) {
  testURL := "https://www.youtube.com/watch?v=DyU5sbXd7SQ"

  result, err := DownloadURL(testURL)
  if err != nil {
    t.Errorf("Returned Error: '%v'", err)
  }
  if len(result) <= 0 {
    t.Errorf("Returned Data is empty")
  }
}
