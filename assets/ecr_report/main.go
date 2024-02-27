package ecr_report

import (
	"context"
	"fmt"
	// "sync"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
)

type Repository struct {
	Name string
	Tags map[string]string
}

func GetECRReposWithTags(ctx context.Context) ([]Repository, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		fmt.Println("Error creating session,", err)
		return nil, err
	}

	client := ecr.NewFromConfig(cfg)

	input := &ecr.DescribeRepositoriesInput{}

	resp, err := client.DescribeRepositories(context.TODO(), input)
	if err != nil {
		fmt.Println("Error describing repositories,", err)
		return nil, err
	}


	repoWithTags := make([]Repository, 0, len(resp.Repositories))
	for _, repo := range resp.Repositories {
		repoObj := Repository{Name: *repo.RepositoryName}
		repoObj.Tags, _ = getTags(ctx, client, repo.RepositoryArn)
		repoWithTags = append(repoWithTags, repoObj)
	}

	return repoWithTags, nil 
}

func getTags(ctx context.Context, svc *ecr.Client, repoArn *string) (map[string]string, error) {
		tagInput := &ecr.ListTagsForResourceInput{
			ResourceArn: repoArn,
		}

		tagResp, _ := svc.ListTagsForResource(ctx, tagInput)

	  tags := make(map[string]string)
		for _, tag := range tagResp.Tags {
	  	tags[*tag.Key] = *tag.Value
		}

	  return tags, nil
}
