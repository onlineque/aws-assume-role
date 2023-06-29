package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"log"
)

func getConfig() (aws.Config, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return cfg, err
	}
	return cfg, nil
}

func assumeRole(roleArn string) (aws.Config, error) {
	cfg, err := getConfig()
	if err != nil {
		return cfg, err
	}

	stsSvc := sts.NewFromConfig(cfg)
	creds := stscreds.NewAssumeRoleProvider(stsSvc, roleArn)

	cfg.Credentials = aws.NewCredentialsCache(creds)
	return cfg, nil
}

func main() {
	const roleArn string = "arn:aws:iam::012345678912:role/roleName"

	cfg, err := assumeRole(roleArn)
	if err != nil {
		log.Fatal(err)
	}

	stsSvc := sts.NewFromConfig(cfg)
	result, err := stsSvc.GetCallerIdentity(context.TODO(), &sts.GetCallerIdentityInput{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Assumed role: %s\n", *result.Arn)
}
