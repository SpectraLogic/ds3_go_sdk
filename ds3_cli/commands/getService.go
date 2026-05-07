package commands

import (
    "context"
    "fmt"
    "github.com/SpectraLogic/ds3_go_sdk/ds3"
    "github.com/SpectraLogic/ds3_go_sdk/ds3/models"
)

func getService(ctx context.Context, client *ds3.Client, args *Arguments) error {
    response, err := client.GetService(ctx, models.NewGetServiceRequest())
    if err == nil {
        for _, bucket := range response.ListAllMyBucketsResult.Buckets {
            fmt.Println(bucket.Name)
        }
    }
    return err
}


