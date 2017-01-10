# Tracing Microservices

![traces](https://raw.githubusercontent.com/goadesign/examples/master/xray/trace.png)

This example contains two dependent services: the upstream service `fetcher`
fetches HTTP documents given their URLs and stores them in a downstream service,
the `archiver`. The requests sent by the fetcher to retrieve the external
requests as well as the ones sent to the archiver service to store the results
are traced using AWS X-Ray.

Running the example requires an active [AWS X-Ray](https://aws.amazon.com/xray/)
account.

## Setup

### goa and goagen

First let's make sure goa is properly installed and updated:
```
go get -u github.com/goadesign/goa
```

### AWS X-Ray Daemon

AWS X-Ray uses a
[daemon](http://docs.aws.amazon.com/xray/latest/devguide/xray-daemon.html) to
forward the tracing data to the AWS service. The service sends the traces to the
daemon over UDP which aggrates the information and forwards it to the AWS X-Ray
service. The page linked above includes multiple download links for various
platforms. The Linux binary can be download
[here](https://s3.amazonaws.com/aws-xray-assets.us-east-1/xray-daemon/aws-xray-daemon-linux-1.x.zip).

Since the daemon sends data to AWS X-Ray it needs access to the required
credentials. Save the credentials in the file `~/.aws/credentials` using the
format:
```
[default]
aws_access_key_id = AKIAIOSFODNN7EXAMPLE
aws_secret_access_key = wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY
```
Next you need to edit the daemon configuration to set the AWS region where you
want the traces to be reported. The `LocalMode` setting also needs to be changed
to `true` otherwise the daemon assumes it's running on an EC2 instance and
attempts to contact the EC2 metadata server. The default configuration file is
part of the zip downloaded above, its name is `cfg.yaml`. Open the file and
change:
```yaml
  # Uncommment Region to send segments to AWS X-Ray service that in a specific region
  # Region: ""

  #...

# Don’t check for EC2 instance metadata.
LocalMode: false
```
to:

```yaml
  # Uncommment Region to send segments to AWS X-Ray service that in a specific region
  Region: "us-west-2" # YOUR REGION HERE

  #...

# Don’t check for EC2 instance metadata.
LocalMode: true
```

As a convenience the example repo provides a `cfg.yaml` with the updated content.
Next start the daemon:
```
./xray
2017-01-07T22:57:04-08:00 [Info] Initializing AWS X-Ray daemon 1.0.1
2017-01-07T22:57:04-08:00 [Info] Using memory limit of 799 MB
2017-01-07T22:57:04-08:00 [Info] 5113 segment buffers allocated
```

### Code Generation And Build

The files [main.go](main.go) of both the `fetcher` and `archiver` service
contain a `go:generate` comment which invokes `goagen` to generate the necessary
files.
```
cd fetcher
go generate
go build
cd ../archiver
go generate
go build
```

## Running the Example

Simply start the services:
```
cd archiver
./archiver
cd ../fetcher
./fetcher
```

The fetcher makes health check requests to the archiver to make sure it's
running on start. So make sure to start the archiver first (there's a grace
period of 1 minute to allow for concurrent launches in environments such as
docker-compose)

## Compile and Run the Fetcher Client

Compile and run the fetcher service client in a different terminal. The command
below fetches [https://goa.design](https://goa.design).
```
cd fetcher/tool/fetcher-cli
go build
./fetcher-cli fetch fetcher /fetch/https://goa.design
2017/01/07 23:17:56 [INFO] started id=uedsR5bZ GET=http://localhost:8081/fetch/https://goa.designr-cli
2017/01/07 23:17:56 [INFO] completed id=uedsR5bZ status=200 time=782.791415ms
{"archive_href":"/archive/1","status":200}
```

Now we can retrieve the archived response from the archiver service:
```
cd ../../../archiver/tool/archiver-cli
go build
./archiver-cli read archiver /archive/1
2017/01/07 23:21:50 [INFO] started id=z4BMNcC+ GET=http://localhost:8080/archive/1
2017/01/07 23:21:50 [INFO] completed id=z4BMNcC+ status=200 time=1.070221ms
{"body":"\n\u003c!DOCTYPE html\u003e\n\n\n\n\n\u003chtml class=\"not-ie\" lang=\"en\"\u003e\n    \n\n    \u003chead\u003e\n\t\u003cmeta name=\"generator\" content=\"Hugo 0.17\" /\u003e\n
...
```
Next head over to the AWS X-Ray console and check out the graphs.
