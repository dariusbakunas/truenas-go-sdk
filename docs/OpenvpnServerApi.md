# \OpenvpnServerApi

All URIs are relative to *http://nas.local.geekspace.us/api/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OpenvpnServerAuthenticationAlgorithmChoicesGet**](OpenvpnServerApi.md#OpenvpnServerAuthenticationAlgorithmChoicesGet) | **Get** /openvpn/server/authentication_algorithm_choices | 
[**OpenvpnServerCipherChoicesGet**](OpenvpnServerApi.md#OpenvpnServerCipherChoicesGet) | **Get** /openvpn/server/cipher_choices | 
[**OpenvpnServerClientConfigurationGenerationPost**](OpenvpnServerApi.md#OpenvpnServerClientConfigurationGenerationPost) | **Post** /openvpn/server/client_configuration_generation | 
[**OpenvpnServerGet**](OpenvpnServerApi.md#OpenvpnServerGet) | **Get** /openvpn/server | 
[**OpenvpnServerPut**](OpenvpnServerApi.md#OpenvpnServerPut) | **Put** /openvpn/server | 
[**OpenvpnServerRenewStaticKeyGet**](OpenvpnServerApi.md#OpenvpnServerRenewStaticKeyGet) | **Get** /openvpn/server/renew_static_key | 



## OpenvpnServerAuthenticationAlgorithmChoicesGet

> OpenvpnServerAuthenticationAlgorithmChoicesGet(ctx).Execute()





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
    resp, r, err := api_client.OpenvpnServerApi.OpenvpnServerAuthenticationAlgorithmChoicesGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenvpnServerApi.OpenvpnServerAuthenticationAlgorithmChoicesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOpenvpnServerAuthenticationAlgorithmChoicesGetRequest struct via the builder pattern


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


## OpenvpnServerCipherChoicesGet

> OpenvpnServerCipherChoicesGet(ctx).Execute()





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
    resp, r, err := api_client.OpenvpnServerApi.OpenvpnServerCipherChoicesGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenvpnServerApi.OpenvpnServerCipherChoicesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOpenvpnServerCipherChoicesGetRequest struct via the builder pattern


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


## OpenvpnServerClientConfigurationGenerationPost

> OpenvpnServerClientConfigurationGenerationPost(ctx).OpenvpnServerClientConfigurationGeneration(openvpnServerClientConfigurationGeneration).Execute()





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
    openvpnServerClientConfigurationGeneration := *openapiclient.NewOpenvpnServerClientConfigurationGeneration() // OpenvpnServerClientConfigurationGeneration |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OpenvpnServerApi.OpenvpnServerClientConfigurationGenerationPost(context.Background()).OpenvpnServerClientConfigurationGeneration(openvpnServerClientConfigurationGeneration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenvpnServerApi.OpenvpnServerClientConfigurationGenerationPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpenvpnServerClientConfigurationGenerationPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **openvpnServerClientConfigurationGeneration** | [**OpenvpnServerClientConfigurationGeneration**](OpenvpnServerClientConfigurationGeneration.md) |  | 

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


## OpenvpnServerGet

> OpenvpnServerGet(ctx).Execute()



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
    resp, r, err := api_client.OpenvpnServerApi.OpenvpnServerGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenvpnServerApi.OpenvpnServerGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOpenvpnServerGetRequest struct via the builder pattern


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


## OpenvpnServerPut

> OpenvpnServerPut(ctx).OpenvpnServerUpdate0(openvpnServerUpdate0).Execute()





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
    openvpnServerUpdate0 := *openapiclient.NewOpenvpnServerUpdate0() // OpenvpnServerUpdate0 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OpenvpnServerApi.OpenvpnServerPut(context.Background()).OpenvpnServerUpdate0(openvpnServerUpdate0).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenvpnServerApi.OpenvpnServerPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpenvpnServerPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **openvpnServerUpdate0** | [**OpenvpnServerUpdate0**](OpenvpnServerUpdate0.md) |  | 

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


## OpenvpnServerRenewStaticKeyGet

> OpenvpnServerRenewStaticKeyGet(ctx).Execute()





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
    resp, r, err := api_client.OpenvpnServerApi.OpenvpnServerRenewStaticKeyGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenvpnServerApi.OpenvpnServerRenewStaticKeyGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOpenvpnServerRenewStaticKeyGetRequest struct via the builder pattern


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

