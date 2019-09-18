package ytdl

// Downloads the Content of the given Youtube-Watch-URL
// and returns it as a byte-array.
// Note:
// Downloads the Video as a webm/mp3 with the best quality it can find
func DownloadURL(url string) ([]byte, error) {
  urls, err := getDownloadURLs(url, 1400000)
  if err != nil {
    return nil, err
  }

  output := ""

  for _, value := range urls {
    downloaded, _ := downloadPart(value)

    output += downloaded
  }

  return []byte(output), nil
}
