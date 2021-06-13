# \ConfigApi

All URIs are relative to *http://nas.local.geekspace.us/api/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConfigResetPost**](ConfigApi.md#ConfigResetPost) | **Post** /config/reset | 



## ConfigResetPost

> ConfigResetPost(ctx).ConfigReset0(configReset0).Execute()





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
    configReset0 := *openapiclient.NewConfigReset0() // ConfigReset0 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigApi.ConfigResetPost(context.Background()).ConfigReset0(configReset0).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.ConfigResetPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConfigResetPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **configReset0** | [**ConfigReset0**](ConfigReset0.md) |  | 

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

