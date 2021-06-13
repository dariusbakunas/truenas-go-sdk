# \ActivedirectoryApi

All URIs are relative to *http://nas.local.geekspace.us/api/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivedirectoryChangeTrustAccountPwGet**](ActivedirectoryApi.md#ActivedirectoryChangeTrustAccountPwGet) | **Get** /activedirectory/change_trust_account_pw | 
[**ActivedirectoryDomainInfoGet**](ActivedirectoryApi.md#ActivedirectoryDomainInfoGet) | **Get** /activedirectory/domain_info | 
[**ActivedirectoryGet**](ActivedirectoryApi.md#ActivedirectoryGet) | **Get** /activedirectory | 
[**ActivedirectoryGetSpnListGet**](ActivedirectoryApi.md#ActivedirectoryGetSpnListGet) | **Get** /activedirectory/get_spn_list | 
[**ActivedirectoryGetStateGet**](ActivedirectoryApi.md#ActivedirectoryGetStateGet) | **Get** /activedirectory/get_state | 
[**ActivedirectoryLeavePost**](ActivedirectoryApi.md#ActivedirectoryLeavePost) | **Post** /activedirectory/leave | 
[**ActivedirectoryNssInfoChoicesGet**](ActivedirectoryApi.md#ActivedirectoryNssInfoChoicesGet) | **Get** /activedirectory/nss_info_choices | 
[**ActivedirectoryPut**](ActivedirectoryApi.md#ActivedirectoryPut) | **Put** /activedirectory | 
[**ActivedirectoryStartedGet**](ActivedirectoryApi.md#ActivedirectoryStartedGet) | **Get** /activedirectory/started | 



## ActivedirectoryChangeTrustAccountPwGet

> ActivedirectoryChangeTrustAccountPwGet(ctx).Execute()





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
    resp, r, err := api_client.ActivedirectoryApi.ActivedirectoryChangeTrustAccountPwGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivedirectoryApi.ActivedirectoryChangeTrustAccountPwGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiActivedirectoryChangeTrustAccountPwGetRequest struct via the builder pattern


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


## ActivedirectoryDomainInfoGet

> ActivedirectoryDomainInfoGet(ctx).Execute()





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
    resp, r, err := api_client.ActivedirectoryApi.ActivedirectoryDomainInfoGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivedirectoryApi.ActivedirectoryDomainInfoGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiActivedirectoryDomainInfoGetRequest struct via the builder pattern


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


## ActivedirectoryGet

> ActivedirectoryGet(ctx).Execute()



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
    resp, r, err := api_client.ActivedirectoryApi.ActivedirectoryGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivedirectoryApi.ActivedirectoryGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiActivedirectoryGetRequest struct via the builder pattern


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


## ActivedirectoryGetSpnListGet

> ActivedirectoryGetSpnListGet(ctx).Execute()





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
    resp, r, err := api_client.ActivedirectoryApi.ActivedirectoryGetSpnListGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivedirectoryApi.ActivedirectoryGetSpnListGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiActivedirectoryGetSpnListGetRequest struct via the builder pattern


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


## ActivedirectoryGetStateGet

> ActivedirectoryGetStateGet(ctx).Execute()





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
    resp, r, err := api_client.ActivedirectoryApi.ActivedirectoryGetStateGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivedirectoryApi.ActivedirectoryGetStateGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiActivedirectoryGetStateGetRequest struct via the builder pattern


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


## ActivedirectoryLeavePost

> ActivedirectoryLeavePost(ctx).ActivedirectoryLeave0(activedirectoryLeave0).Execute()





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
    activedirectoryLeave0 := *openapiclient.NewActivedirectoryLeave0() // ActivedirectoryLeave0 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ActivedirectoryApi.ActivedirectoryLeavePost(context.Background()).ActivedirectoryLeave0(activedirectoryLeave0).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivedirectoryApi.ActivedirectoryLeavePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiActivedirectoryLeavePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **activedirectoryLeave0** | [**ActivedirectoryLeave0**](ActivedirectoryLeave0.md) |  | 

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


## ActivedirectoryNssInfoChoicesGet

> ActivedirectoryNssInfoChoicesGet(ctx).Execute()





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
    resp, r, err := api_client.ActivedirectoryApi.ActivedirectoryNssInfoChoicesGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivedirectoryApi.ActivedirectoryNssInfoChoicesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiActivedirectoryNssInfoChoicesGetRequest struct via the builder pattern


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


## ActivedirectoryPut

> ActivedirectoryPut(ctx).ActivedirectoryUpdate0(activedirectoryUpdate0).Execute()





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
    activedirectoryUpdate0 := *openapiclient.NewActivedirectoryUpdate0() // ActivedirectoryUpdate0 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ActivedirectoryApi.ActivedirectoryPut(context.Background()).ActivedirectoryUpdate0(activedirectoryUpdate0).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivedirectoryApi.ActivedirectoryPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiActivedirectoryPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **activedirectoryUpdate0** | [**ActivedirectoryUpdate0**](ActivedirectoryUpdate0.md) |  | 

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


## ActivedirectoryStartedGet

> ActivedirectoryStartedGet(ctx).Execute()





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
    resp, r, err := api_client.ActivedirectoryApi.ActivedirectoryStartedGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivedirectoryApi.ActivedirectoryStartedGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiActivedirectoryStartedGetRequest struct via the builder pattern


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

