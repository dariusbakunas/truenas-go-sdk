# \StatsApi

All URIs are relative to *http://nas.local.geekspace.us/api/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StatsGetDataPost**](StatsApi.md#StatsGetDataPost) | **Post** /stats/get_data | 
[**StatsGetDatasetInfoPost**](StatsApi.md#StatsGetDatasetInfoPost) | **Post** /stats/get_dataset_info | 
[**StatsGetSourcesGet**](StatsApi.md#StatsGetSourcesGet) | **Get** /stats/get_sources | 



## StatsGetDataPost

> StatsGetDataPost(ctx).StatsGetData(statsGetData).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    statsGetData := *openapiclient.NewStatsGetData() // StatsGetData |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StatsApi.StatsGetDataPost(context.Background()).StatsGetData(statsGetData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatsApi.StatsGetDataPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStatsGetDataPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **statsGetData** | [**StatsGetData**](StatsGetData.md) |  | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StatsGetDatasetInfoPost

> StatsGetDatasetInfoPost(ctx).StatsGetDatasetInfo(statsGetDatasetInfo).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    statsGetDatasetInfo := *openapiclient.NewStatsGetDatasetInfo() // StatsGetDatasetInfo |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StatsApi.StatsGetDatasetInfoPost(context.Background()).StatsGetDatasetInfo(statsGetDatasetInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatsApi.StatsGetDatasetInfoPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStatsGetDatasetInfoPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **statsGetDatasetInfo** | [**StatsGetDatasetInfo**](StatsGetDatasetInfo.md) |  | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StatsGetSourcesGet

> StatsGetSourcesGet(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StatsApi.StatsGetSourcesGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatsApi.StatsGetSourcesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiStatsGetSourcesGetRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

