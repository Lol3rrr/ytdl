package ytdl

import (
  "errors"
  "net/http"
  "io/ioutil"
)

func downloadPart(url string) (string, error) {
  resp, err := http.Get(url)
  if err != nil {
    return "", err
  }
  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return "", err
  }

  if resp.StatusCode != 200 {
    return "", errors.New("Could not download the Video-Part")
  }

  return string(body), nil
}
