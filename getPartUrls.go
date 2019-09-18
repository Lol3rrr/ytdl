package ytdl

import (
  "net/url"
  "strconv"
)

func getPartUrls(data *url.URL, dataQuery *url.Values, chunkSize int) ([]string, error) {
  replaceQueryData("signature", "sig", dataQuery)

  results := make([]string, 0)
  currentBuf := -1
  totalSize, err := strconv.Atoi(dataQuery.Get("clen"))
  if err != nil {
    return nil, err
  }
  remaining := totalSize

  for {
    toDownload := chunkSize
    remaining = totalSize - currentBuf
    if remaining <= 0 {
      break
    }
    if remaining < chunkSize {
      toDownload = remaining
    }

    rangeVar := strconv.Itoa(currentBuf + 1) + "-" + strconv.Itoa(currentBuf + toDownload)

    dataQuery.Set("range", rangeVar)

    data.RawQuery = dataQuery.Encode()

    url := data.String()

    results = append(results, url)

    currentBuf += chunkSize
  }

  return results, nil
}
