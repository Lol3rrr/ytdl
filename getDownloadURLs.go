package ytdl

func getDownloadURLs(watchUrl string, chunkSize int) ([]string, error) {
  data, err := getRawDownloadURL(watchUrl)
  if err != nil {
    return nil, err
  }

  dataQuery := data.Query()

  results, err := getPartUrls(data, &dataQuery, chunkSize)
  if err != nil {
    return nil, err
  }

  return results, nil
}
