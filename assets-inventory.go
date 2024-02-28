package main

import (
  "context"
  "fmt"

  "github.com/toanvvsg/assets-inventory/assets/s3_report"
  "github.com/toanvvsg/assets-inventory/assets/ecr_report"
  "github.com/toanvvsg/assets-inventory/assets/ec2_report"
)


func main() {
  fmt.Println("Start Program!")
  ctx := context.TODO()
  if false {
    buckets, _ := s3_report.GetBuckets(ctx)
    fmt.Println(buckets)
    ecrs, _ := ecr_report.GetRepos(ctx)
    fmt.Println(ecrs)
  }
  ec2s, _ := ec2_report.GetInstances(ctx)
  for _, ec2 := range ec2s {
    fmt.Println(ec2.Name)
    fmt.Println(" - ", ec2.State)
    if ec2.StateTransitionReason != "" {
      fmt.Println(" - ", ec2.StateTransitionReason)
    }
  }
}
