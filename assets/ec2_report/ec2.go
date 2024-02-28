package ec2_report

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

type EC2 struct {
	Name string
	Tags map[string]string
	State string
	LaunchTime time.Time
	StateTransitionReason string
}

func GetInstances(ctx context.Context) ([]EC2, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		fmt.Errorf("Error creating session, %w", err)
		return nil, err
	}

	svc := ec2.NewFromConfig(cfg)
	params := &ec2.DescribeInstancesInput{}

	resp, err := svc.DescribeInstances(ctx, params)
	if err != nil {
	  return nil, fmt.Errorf("Error describe EC2 instances, %w", err)
	}

  var ec2s []EC2
	for _, reservation := range resp.Reservations {
	  for _, instance := range reservation.Instances {
	    ec2 := EC2{
	      Name: *instance.InstanceId,
	      State: string(instance.State.Name),
	      LaunchTime: *instance.LaunchTime,
	      Tags: make(map[string]string),
        StateTransitionReason: *instance.StateTransitionReason,
	    }
	    fmt.Println("  State: ", string(instance.State.Name))
	    for _, tag := range instance.Tags {
	      if *tag.Key != "" {
	        ec2.Tags[*tag.Key] = *tag.Value
	      }
	    }
	    ec2s = append(ec2s, ec2)
	  }
	}

	return ec2s, nil
}
