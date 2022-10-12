package utils

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	ssmTypes "github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"os"
)

func StringPtr(str string) *string { return &str }

func GetAwsConfig(region string) aws.Config {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return cfg
}

func GetSsmValue(region, key string) (value string, ok bool) {
	cfg := GetAwsConfig(region)
	svc := ssm.NewFromConfig(cfg)

	res, err := svc.GetParameter(context.TODO(),
		&ssm.GetParameterInput{
			Name:           StringPtr(key),
			WithDecryption: true,
		})
	if err != nil {
		if err != nil {
			var pn *ssmTypes.ParameterNotFound
			if errors.As(err, &pn) {
				fmt.Printf("Parameter %s not Found in the SSM.\n", key)
				return value, ok
			}
		}
		fmt.Println(err)
		os.Exit(1)
	}
	value = *res.Parameter.Value
	return value, true
}
