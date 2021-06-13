# \PoolResilverApi

All URIs are relative to *http://nas.local.geekspace.us/api/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PoolResilverGet**](PoolResilverApi.md#PoolResilverGet) | **Get** /pool/resilver | 
[**PoolResilverPut**](PoolResilverApi.md#PoolResilverPut) | **Put** /pool/resilver | 



## PoolResilverGet

> PoolResilverGet(ctx).Execute()



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
    resp, r, err := api_client.PoolResilverApi.PoolResilverGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoolResilverApi.PoolResilverGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPoolResilverGetRequest struct via the builder pattern


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


## PoolResilverPut

> PoolResilverPut(ctx).PoolResilverUpdate0(poolResilverUpdate0).Execute()





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
    poolResilverUpdate0 := *openapiclient.NewPoolResilverUpdate0() // PoolResilverUpdate0 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoolResilverApi.PoolResilverPut(context.Background()).PoolResilverUpdate0(poolResilverUpdate0).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoolResilverApi.PoolResilverPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPoolResilverPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **poolResilverUpdate0** | [**PoolResilverUpdate0**](PoolResilverUpdate0.md) |  | 

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

