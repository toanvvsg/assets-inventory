package main

import (
  "context"
  "fmt"

  "github.com/toanvvsg/assets-inventory/assets/s3_report"
  "github.com/toanvvsg/assets-inventory/assets/ecr_report"
)


func main() {
  fmt.Println("Start Program!")
  ctx := context.TODO()
  if false {
    buckets, _ := s3_report.GetBucketsWithTags(ctx)
    fmt.Println(buckets)
  }
  x, _ := ecr_report.GetECRReposWithTags(ctx)
  fmt.Println(x)
}
