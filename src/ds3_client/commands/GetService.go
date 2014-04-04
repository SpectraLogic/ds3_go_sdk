package commands

import (
    "fmt"
    "ds3"
    "ds3/models"
)

func getService(client *ds3.Client, args *Arguments) error {
    response, err := client.GetService(models.NewGetServiceRequest())
    if err == nil {
        //TODO: better result printing
        for _, bucket := range(response.Buckets) {
            fmt.Println(bucket.Name)
        }
    }
    return err
}


