# \SystemAdvancedApi

All URIs are relative to *http://nas.local.geekspace.us/api/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SystemAdvancedGet**](SystemAdvancedApi.md#SystemAdvancedGet) | **Get** /system/advanced | 
[**SystemAdvancedPut**](SystemAdvancedApi.md#SystemAdvancedPut) | **Put** /system/advanced | 
[**SystemAdvancedSedGlobalPasswordGet**](SystemAdvancedApi.md#SystemAdvancedSedGlobalPasswordGet) | **Get** /system/advanced/sed_global_password | 
[**SystemAdvancedSerialPortChoicesGet**](SystemAdvancedApi.md#SystemAdvancedSerialPortChoicesGet) | **Get** /system/advanced/serial_port_choices | 



## SystemAdvancedGet

> SystemAdvancedGet(ctx).Execute()



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
    resp, r, err := api_client.SystemAdvancedApi.SystemAdvancedGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemAdvancedApi.SystemAdvancedGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSystemAdvancedGetRequest struct via the builder pattern


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


## SystemAdvancedPut

> SystemAdvancedPut(ctx).SystemAdvancedUpdate0(systemAdvancedUpdate0).Execute()





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
    systemAdvancedUpdate0 := *openapiclient.NewSystemAdvancedUpdate0() // SystemAdvancedUpdate0 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SystemAdvancedApi.SystemAdvancedPut(context.Background()).SystemAdvancedUpdate0(systemAdvancedUpdate0).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemAdvancedApi.SystemAdvancedPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSystemAdvancedPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **systemAdvancedUpdate0** | [**SystemAdvancedUpdate0**](SystemAdvancedUpdate0.md) |  | 

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


## SystemAdvancedSedGlobalPasswordGet

> SystemAdvancedSedGlobalPasswordGet(ctx).Execute()





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
    resp, r, err := api_client.SystemAdvancedApi.SystemAdvancedSedGlobalPasswordGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemAdvancedApi.SystemAdvancedSedGlobalPasswordGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSystemAdvancedSedGlobalPasswordGetRequest struct via the builder pattern


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


## SystemAdvancedSerialPortChoicesGet

> SystemAdvancedSerialPortChoicesGet(ctx).Execute()





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
    resp, r, err := api_client.SystemAdvancedApi.SystemAdvancedSerialPortChoicesGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemAdvancedApi.SystemAdvancedSerialPortChoicesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSystemAdvancedSerialPortChoicesGetRequest struct via the builder pattern


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

