package ytdl

import (
  "errors"

  "github.com/rylio/ytdl"
)

func getFormatBestExtension(videoInfo *ytdl.VideoInfo, extension string) (ytdl.Format, error) {
  itags := videoInfo.Formats

  set := false
  format := ytdl.Format{}

	for _, value := range itags {
    if value.Extension == extension && value.AudioBitrate > format.AudioBitrate {
      format = value
      set = true
    }
	}

  var returnErr error
  if !set {
    returnErr = errors.New("Didnt find a fitting Value")
  }

  return format, returnErr
}
