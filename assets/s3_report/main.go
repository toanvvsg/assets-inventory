package s3_report

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// S3Bucket struct remains the same
type S3Bucket struct {
	Name string
	Tags map[string]string
}

// getS3BucketsWithTags function using AWS SDK v2
func GetBuckets(ctx context.Context) ([]S3Bucket, error) {
	// Load default AWS configuration
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("error loading AWS configuration: %w", err)
	}

	// Create an S3 service client
	svc := s3.NewFromConfig(cfg)

	// List S3 buckets
	params := &s3.ListBucketsInput{}
	output, err := svc.ListBuckets(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("error listing S3 buckets: %w", err)
	}

	bucketsWithTags := make([]S3Bucket, 0, len(output.Buckets))
	for _, bucket := range output.Buckets {
		bucketObj := S3Bucket{Name: *bucket.Name}
		bucketObj.Tags, err = getTags(ctx, svc, *bucket.Name)
		if err != nil {
			fmt.Errorf("error getting tags for bucket %s: %w\n", bucket.Name, err)
			continue
		}
		bucketsWithTags = append(bucketsWithTags, bucketObj)
	}

	return bucketsWithTags, nil
}

// getBucketTags function using AWS SDK v2
func getTags(ctx context.Context, svc *s3.Client, bucketName string) (map[string]string, error) {
	tagsInput := &s3.GetBucketTaggingInput{
		Bucket: aws.String(bucketName),
	}
	tagsOutput, err := svc.GetBucketTagging(ctx, tagsInput)
	if err != nil {
		return nil, fmt.Errorf("error getting tags for bucket %s: %w", bucketName, err)
	}

	tags := make(map[string]string)
	if tagsOutput.TagSet != nil {
		for _, tag := range tagsOutput.TagSet {
			tags[*tag.Key] = *tag.Value
		}
	}

	return tags, nil
}
