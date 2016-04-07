# go-aws-sqs-get

usage

```
Usage:
  aws-sqs-get [OPTIONS]

Application Options:
  -r, --region= target region
  -n, --name=   target queue name
  -m, --metric= target metrics label

Help Options:
  -h, --help    Show this help message
```

```
$ aws-sqs-get -r ap-northeast-1 -n awesome-queue -m NumberOfMessagesReceived
0.98765
```
