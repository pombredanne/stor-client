# storclient
--
    import "github.com/avast/stor-client/client"

Package storclient to download samples from stor service

### SYNOPSIS

    client := storclient.New(storageUrl, storclient.StorClientOpts{})

    client.Start()

    for _, sha := range shaList {
    	client.Download(sha)
    }

    downloadStatus := client.Wait()

## Usage

```go
const (
	DefaultMax           = 4
	DefaultTimeout       = 30 * time.Second
	DefaultRetryAttempts = 10
	DefaultRetryDelay    = 1e5 * time.Microsecond
	DefaultS3Template    = "{{.FirstShaByte}}/{{.SecondShaByte}}/{{.ThirdShaByte}}/{{.Sha}}"
)
```

#### type DownPool

```go
type DownPool struct {
}
```


#### type DownStat

```go
type DownStat struct {
	Size     int64
	Duration time.Duration
	Status   DownloadStatus
}
```


#### type DownloadStatus

```go
type DownloadStatus int
```


```go
const (
	// DOWN_FAIL - downlad fail (default)
	DOWN_FAIL DownloadStatus = iota
	// DOWN_SKIP - downlad skipped because file is downlad
	DOWN_SKIP
	// DOWN_OK - downlad ok
	DOWN_OK
)
```

#### type StorClient

```go
type StorClient struct {
	StorClientOpts
}
```


#### func  New

```go
func New(storUrl url.URL, downloadDir string, opts StorClientOpts) (*StorClient, error)
```
Create new instance of stor client

#### func (*StorClient) Download

```go
func (client *StorClient) Download(sha hashutil.Hash)
```
add sha to douwnload queue

#### func (*StorClient) Start

```go
func (client *StorClient) Start()
```
start stor downloading process

#### func (*StorClient) Wait

```go
func (client *StorClient) Wait() TotalStat
```
wait to all downloads return download stats

#### type StorClientOpts

```go
type StorClientOpts struct {
	//	max size of download pool
	Max int
	//	write to devnull instead of file
	Devnull bool
	//	connection timeout
	//
	//	-1 means no limit (no timeout)
	Timeout time.Duration
	// exponential retry - start delay time
	// default is 10e5 microseconds
	RetryDelay time.Duration
	// count of tries of retry
	// default is 10
	RetryAttempts uint
	// downladed file suffix
	// e.g. .dat => SHA.dat file
	// default ("") means without suffix
	Suffix string
	// name of file will be upper case (not applied to extension)
	UpperCase bool
	// host to s3 endpoint with bucket e.g. https://bucket.s3.eu-central-1.amazonaws.com, if is s3url set, first will be use S3, then fallback to stor
	S3URL *url.URL
	// template to S3 path
	S3Template string
}
```

StorClientOpts is base struct

#### type TotalStat

```go
type TotalStat struct {
	Size     int64
	Duration time.Duration
	// Count of downloaded files
	Count int
	// Count of skipped files
	Skip int
}
```

Size and Duration is duplicate, becuse embedding not works, because

    https://stackoverflow.com/questions/41686692/embedding-structs-in-golang-gives-error-unknown-field

#### func (TotalStat) Print

```go
func (total TotalStat) Print(startTime time.Time)
```
format and log total stats

#### func (TotalStat) Status

```go
func (total TotalStat) Status() bool
```
Status return true if all files are downloaded
