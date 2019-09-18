package ytdl

import (
  "errors"

  "github.com/smallnest/goreq"
)

func downloadPart(url string) (string, error) {
  resp, data, reqErr := goreq.New().Get(url).End()

  if resp.StatusCode != 200 || reqErr != nil {
    return "", errors.New("Could not download the Video-Part")
  }

  return data, nil
}
