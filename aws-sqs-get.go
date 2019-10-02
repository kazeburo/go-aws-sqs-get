package main

import (
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/jessevdk/go-flags"
)

type options struct {
	Region string `short:"r" long:"region" arg:"String" required:"true" description:"target region"`
	Name   string `short:"n" long:"name" arg:"String" required:"true" description:"target queue name"`
	Metric string `short:"m" long:"metric" arg:"String" requried:"true" description:"target metrics label"`
}

func main() {
	os.Exit(_main())
}

func _main() (st int) {
	st = 1

	opts := options{}
	psr := flags.NewParser(&opts, flags.Default)
	_, err := psr.Parse()
	if err != nil {
		return
	}

	var home = os.Getenv("HOME")
	_, err = os.Stat("/root/.aws/credentials")
	if home == "" && err == nil {
		os.Setenv("HOME", "/root")
	}

	maxRetry := 3
	svc := cloudwatch.New(session.New(), &aws.Config{Region: aws.String(opts.Region), MaxRetries: &maxRetry})
	now := time.Now()
	prev := now.Add(time.Duration(630) * time.Second * -1) // 10 min (to fetch at least 1 data-point)

	params := &cloudwatch.GetMetricStatisticsInput{
		StartTime:  aws.Time(prev),
		EndTime:    aws.Time(now),
		MetricName: aws.String(opts.Metric),
		Namespace:  aws.String("AWS/SQS"),
		Period:     aws.Int64(60),
		Statistics: []*string{
			aws.String("Average"),
		},
		Dimensions: []*cloudwatch.Dimension{
			{ // Required
				Name:  aws.String("QueueName"), // Required
				Value: aws.String(opts.Name),   // Required
			},
			// More values...
		},
	}
	resp, err := svc.GetMetricStatistics(params)
	if err != nil {
		fmt.Fprintf(os.Stderr, "svc.GetMetricStatistics failed: %s\n", err)
		return
	}

	datapoints := resp.Datapoints
	if len(datapoints) == 0 {
		fmt.Fprintf(os.Stderr, "svc.GetMetricStatistics failed: %s\n", "fetched no datapoints")
		return
	}
	// fmt.Println(resp)
	latest := *datapoints[0].Timestamp
	var latestVal float64
	for _, dp := range datapoints {
		if dp.Timestamp.After(latest) {
			continue
		}
		latest = *dp.Timestamp
		latestVal = *dp.Average
	}
	st = 0
	fmt.Fprintf(os.Stdout, "%f\n", latestVal)
	return
}
