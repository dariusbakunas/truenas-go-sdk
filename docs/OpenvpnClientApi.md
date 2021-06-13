# \OpenvpnClientApi

All URIs are relative to *http://nas.local.geekspace.us/api/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OpenvpnClientAuthenticationAlgorithmChoicesGet**](OpenvpnClientApi.md#OpenvpnClientAuthenticationAlgorithmChoicesGet) | **Get** /openvpn/client/authentication_algorithm_choices | 
[**OpenvpnClientCipherChoicesGet**](OpenvpnClientApi.md#OpenvpnClientCipherChoicesGet) | **Get** /openvpn/client/cipher_choices | 
[**OpenvpnClientGet**](OpenvpnClientApi.md#OpenvpnClientGet) | **Get** /openvpn/client | 
[**OpenvpnClientPut**](OpenvpnClientApi.md#OpenvpnClientPut) | **Put** /openvpn/client | 



## OpenvpnClientAuthenticationAlgorithmChoicesGet

> OpenvpnClientAuthenticationAlgorithmChoicesGet(ctx).Execute()





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
    resp, r, err := api_client.OpenvpnClientApi.OpenvpnClientAuthenticationAlgorithmChoicesGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenvpnClientApi.OpenvpnClientAuthenticationAlgorithmChoicesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOpenvpnClientAuthenticationAlgorithmChoicesGetRequest struct via the builder pattern


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


## OpenvpnClientCipherChoicesGet

> OpenvpnClientCipherChoicesGet(ctx).Execute()





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
    resp, r, err := api_client.OpenvpnClientApi.OpenvpnClientCipherChoicesGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenvpnClientApi.OpenvpnClientCipherChoicesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOpenvpnClientCipherChoicesGetRequest struct via the builder pattern


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


## OpenvpnClientGet

> OpenvpnClientGet(ctx).Execute()



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
    resp, r, err := api_client.OpenvpnClientApi.OpenvpnClientGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenvpnClientApi.OpenvpnClientGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOpenvpnClientGetRequest struct via the builder pattern


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


## OpenvpnClientPut

> OpenvpnClientPut(ctx).OpenvpnClientUpdate0(openvpnClientUpdate0).Execute()





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
    openvpnClientUpdate0 := *openapiclient.NewOpenvpnClientUpdate0() // OpenvpnClientUpdate0 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OpenvpnClientApi.OpenvpnClientPut(context.Background()).OpenvpnClientUpdate0(openvpnClientUpdate0).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenvpnClientApi.OpenvpnClientPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpenvpnClientPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **openvpnClientUpdate0** | [**OpenvpnClientUpdate0**](OpenvpnClientUpdate0.md) |  | 

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

