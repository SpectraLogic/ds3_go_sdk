package commands

import (
    "fmt"
    "ds3"
)

const maxInt = int(^uint(0) >> 1)
const defaultMaxKeys = 1000

func getMinInt(a, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}

func getBucket(client *ds3.Client, args *Arguments) error {
    objects, err := getBucketObjects(client, args)
    if err != nil {
        return err
    }
    for _, obj := range objects {
        fmt.Println(obj.Key)
    }
    return nil
}

