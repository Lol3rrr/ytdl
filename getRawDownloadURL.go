package ytdl

import (
  "net/url"

  "github.com/rylio/ytdl"
)

func getRawDownloadURL(watchUrl string) (*url.URL, error) {
  video, err := ytdl.GetVideoInfo(watchUrl)
	if err != nil {
		return nil, err
	}

	format, err := getFormatBestExtension(video, "webm")
	if err != nil {
		return nil, err
	}

	data, err := video.GetDownloadURL(format)
	if err != nil {
		return nil, err
	}

  return data, nil
}
