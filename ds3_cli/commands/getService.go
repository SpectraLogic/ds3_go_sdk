package commands

import (
    "fmt"
    "spectra/ds3_go_sdk/ds3"
    "spectra/ds3_go_sdk/ds3/models"
)

func getService(client *ds3.Client, args *Arguments) error {
    response, err := client.GetService(models.NewGetServiceRequest())
    if err == nil {
        for _, bucket := range response.ListAllMyBucketsResult.Buckets {
            fmt.Println(bucket.Name)
        }
    }
    return err
}


