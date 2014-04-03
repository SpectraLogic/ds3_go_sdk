package commands

import (
    "fmt"
    "ds3"
)

func getService(client *ds3.Client) error {
    getServiceResponse, err := client.GetService(models.NewGetServiceRequest())
    if err == nil {
        //TODO: better result printing
        fmt.Println(getServiceResponse)
    }
    return err
}


