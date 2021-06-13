# \SmbSharesecApi

All URIs are relative to *http://nas.local.geekspace.us/api/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SmbSharesecGet**](SmbSharesecApi.md#SmbSharesecGet) | **Get** /smb/sharesec | 
[**SmbSharesecGetaclPost**](SmbSharesecApi.md#SmbSharesecGetaclPost) | **Post** /smb/sharesec/getacl | 
[**SmbSharesecIdIdDelete**](SmbSharesecApi.md#SmbSharesecIdIdDelete) | **Delete** /smb/sharesec/id/{id} | 
[**SmbSharesecIdIdGet**](SmbSharesecApi.md#SmbSharesecIdIdGet) | **Get** /smb/sharesec/id/{id} | 
[**SmbSharesecIdIdPut**](SmbSharesecApi.md#SmbSharesecIdIdPut) | **Put** /smb/sharesec/id/{id} | 
[**SmbSharesecPost**](SmbSharesecApi.md#SmbSharesecPost) | **Post** /smb/sharesec | 
[**SmbSharesecSynchronizeAclsGet**](SmbSharesecApi.md#SmbSharesecSynchronizeAclsGet) | **Get** /smb/sharesec/synchronize_acls | 



## SmbSharesecGet

> SmbSharesecGet(ctx).Limit(limit).Offset(offset).Count(count).Sort(sort).Execute()





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
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    count := true // bool |  (optional)
    sort := "sort_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SmbSharesecApi.SmbSharesecGet(context.Background()).Limit(limit).Offset(offset).Count(count).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SmbSharesecApi.SmbSharesecGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSmbSharesecGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **count** | **bool** |  | 
 **sort** | **string** |  | 

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


## SmbSharesecGetaclPost

> SmbSharesecGetaclPost(ctx).SmbSharesecGetacl(smbSharesecGetacl).Execute()





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
    smbSharesecGetacl := *openapiclient.NewSmbSharesecGetacl() // SmbSharesecGetacl |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SmbSharesecApi.SmbSharesecGetaclPost(context.Background()).SmbSharesecGetacl(smbSharesecGetacl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SmbSharesecApi.SmbSharesecGetaclPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSmbSharesecGetaclPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **smbSharesecGetacl** | [**SmbSharesecGetacl**](SmbSharesecGetacl.md) |  | 

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


## SmbSharesecIdIdDelete

> SmbSharesecIdIdDelete(ctx, id).Execute()





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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SmbSharesecApi.SmbSharesecIdIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SmbSharesecApi.SmbSharesecIdIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSmbSharesecIdIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## SmbSharesecIdIdGet

> SmbSharesecIdIdGet(ctx, id).Limit(limit).Offset(offset).Count(count).Sort(sort).Execute()





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
    id := []interface{}{interface{}(123)} // []interface{} |  (default to [])
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    count := true // bool |  (optional)
    sort := "sort_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SmbSharesecApi.SmbSharesecIdIdGet(context.Background(), id).Limit(limit).Offset(offset).Count(count).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SmbSharesecApi.SmbSharesecIdIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**[]interface{}**](interface{}.md) |  | [default to []]

### Other Parameters

Other parameters are passed through a pointer to a apiSmbSharesecIdIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **count** | **bool** |  | 
 **sort** | **string** |  | 

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


## SmbSharesecIdIdPut

> SmbSharesecIdIdPut(ctx, id).SmbSharesecUpdate1(smbSharesecUpdate1).Execute()





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
    id := int32(56) // int32 | 
    smbSharesecUpdate1 := *openapiclient.NewSmbSharesecUpdate1() // SmbSharesecUpdate1 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SmbSharesecApi.SmbSharesecIdIdPut(context.Background(), id).SmbSharesecUpdate1(smbSharesecUpdate1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SmbSharesecApi.SmbSharesecIdIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSmbSharesecIdIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **smbSharesecUpdate1** | [**SmbSharesecUpdate1**](SmbSharesecUpdate1.md) |  | 

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


## SmbSharesecPost

> SmbSharesecPost(ctx).SmbSharesecCreate0(smbSharesecCreate0).Execute()





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
    smbSharesecCreate0 := *openapiclient.NewSmbSharesecCreate0() // SmbSharesecCreate0 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SmbSharesecApi.SmbSharesecPost(context.Background()).SmbSharesecCreate0(smbSharesecCreate0).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SmbSharesecApi.SmbSharesecPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSmbSharesecPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **smbSharesecCreate0** | [**SmbSharesecCreate0**](SmbSharesecCreate0.md) |  | 

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


## SmbSharesecSynchronizeAclsGet

> SmbSharesecSynchronizeAclsGet(ctx).Execute()





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
    resp, r, err := api_client.SmbSharesecApi.SmbSharesecSynchronizeAclsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SmbSharesecApi.SmbSharesecSynchronizeAclsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSmbSharesecSynchronizeAclsGetRequest struct via the builder pattern


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

