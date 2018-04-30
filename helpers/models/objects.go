package models

import (
	"github.com/SpectraLogic/ds3_go_sdk/ds3/models"
	"strings"
)

type GetObject struct {
    Name string
    Ranges []models.Range
    ChannelBuilder WriteChannelBuilder
}

type PutObject struct {
    PutObject      models.Ds3PutObject
    ChannelBuilder ReadChannelBuilder
	Metadata map[string]string
}

func (putObject *PutObject) WithMetaData(key string, values ...string) interface{} {
	if strings.HasPrefix(strings.ToLower(key), models.AMZ_META_HEADER) {
		putObject.Metadata[key] = strings.Join(values, ",")
	} else {
		putObject.Metadata[strings.ToLower(models.AMZ_META_HEADER + key)] = strings.Join(values, ",")
	}

	return putObject
}
