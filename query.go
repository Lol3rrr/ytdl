package ytdl

import (
  "net/url"
)

func replaceQueryData(inputName, outputName string, query *url.Values) {
  tmpSignature := query.Get(inputName)
	tmpSig := query.Get(outputName)
	if tmpSig == "" && tmpSignature != "" {
		query.Del(inputName)
		query.Set(outputName, tmpSignature)
	}
}
