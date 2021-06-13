# \NetworkConfigurationApi

All URIs are relative to *http://nas.local.geekspace.us/api/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NetworkConfigurationGet**](NetworkConfigurationApi.md#NetworkConfigurationGet) | **Get** /network/configuration | 
[**NetworkConfigurationPut**](NetworkConfigurationApi.md#NetworkConfigurationPut) | **Put** /network/configuration | 



## NetworkConfigurationGet

> NetworkConfigurationGet(ctx).Execute()



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
    resp, r, err := api_client.NetworkConfigurationApi.NetworkConfigurationGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkConfigurationApi.NetworkConfigurationGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkConfigurationGetRequest struct via the builder pattern


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


## NetworkConfigurationPut

> NetworkConfigurationPut(ctx).NetworkConfigurationUpdate0(networkConfigurationUpdate0).Execute()





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
    networkConfigurationUpdate0 := *openapiclient.NewNetworkConfigurationUpdate0() // NetworkConfigurationUpdate0 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkConfigurationApi.NetworkConfigurationPut(context.Background()).NetworkConfigurationUpdate0(networkConfigurationUpdate0).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkConfigurationApi.NetworkConfigurationPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNetworkConfigurationPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **networkConfigurationUpdate0** | [**NetworkConfigurationUpdate0**](NetworkConfigurationUpdate0.md) |  | 

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

