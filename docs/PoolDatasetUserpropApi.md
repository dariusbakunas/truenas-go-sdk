# \PoolDatasetUserpropApi

All URIs are relative to *http://nas.local.geekspace.us/api/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PoolDatasetUserpropGet**](PoolDatasetUserpropApi.md#PoolDatasetUserpropGet) | **Get** /pool/dataset/userprop | 
[**PoolDatasetUserpropIdIdDelete**](PoolDatasetUserpropApi.md#PoolDatasetUserpropIdIdDelete) | **Delete** /pool/dataset/userprop/id/{id} | 
[**PoolDatasetUserpropIdIdGet**](PoolDatasetUserpropApi.md#PoolDatasetUserpropIdIdGet) | **Get** /pool/dataset/userprop/id/{id} | 
[**PoolDatasetUserpropIdIdPut**](PoolDatasetUserpropApi.md#PoolDatasetUserpropIdIdPut) | **Put** /pool/dataset/userprop/id/{id} | 
[**PoolDatasetUserpropPost**](PoolDatasetUserpropApi.md#PoolDatasetUserpropPost) | **Post** /pool/dataset/userprop | 



## PoolDatasetUserpropGet

> PoolDatasetUserpropGet(ctx).Limit(limit).Offset(offset).Count(count).Sort(sort).Execute()





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
    resp, r, err := api_client.PoolDatasetUserpropApi.PoolDatasetUserpropGet(context.Background()).Limit(limit).Offset(offset).Count(count).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoolDatasetUserpropApi.PoolDatasetUserpropGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPoolDatasetUserpropGetRequest struct via the builder pattern


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


## PoolDatasetUserpropIdIdDelete

> PoolDatasetUserpropIdIdDelete(ctx, id).PoolDatasetUserpropDelete1(poolDatasetUserpropDelete1).Execute()





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
    poolDatasetUserpropDelete1 := *openapiclient.NewPoolDatasetUserpropDelete1() // PoolDatasetUserpropDelete1 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoolDatasetUserpropApi.PoolDatasetUserpropIdIdDelete(context.Background(), id).PoolDatasetUserpropDelete1(poolDatasetUserpropDelete1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoolDatasetUserpropApi.PoolDatasetUserpropIdIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPoolDatasetUserpropIdIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **poolDatasetUserpropDelete1** | [**PoolDatasetUserpropDelete1**](PoolDatasetUserpropDelete1.md) |  | 

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


## PoolDatasetUserpropIdIdGet

> PoolDatasetUserpropIdIdGet(ctx, id).Limit(limit).Offset(offset).Count(count).Sort(sort).Execute()





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
    resp, r, err := api_client.PoolDatasetUserpropApi.PoolDatasetUserpropIdIdGet(context.Background(), id).Limit(limit).Offset(offset).Count(count).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoolDatasetUserpropApi.PoolDatasetUserpropIdIdGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPoolDatasetUserpropIdIdGetRequest struct via the builder pattern


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


## PoolDatasetUserpropIdIdPut

> PoolDatasetUserpropIdIdPut(ctx, id).PoolDatasetUserpropUpdate1(poolDatasetUserpropUpdate1).Execute()





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
    poolDatasetUserpropUpdate1 := *openapiclient.NewPoolDatasetUserpropUpdate1() // PoolDatasetUserpropUpdate1 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoolDatasetUserpropApi.PoolDatasetUserpropIdIdPut(context.Background(), id).PoolDatasetUserpropUpdate1(poolDatasetUserpropUpdate1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoolDatasetUserpropApi.PoolDatasetUserpropIdIdPut``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPoolDatasetUserpropIdIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **poolDatasetUserpropUpdate1** | [**PoolDatasetUserpropUpdate1**](PoolDatasetUserpropUpdate1.md) |  | 

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


## PoolDatasetUserpropPost

> PoolDatasetUserpropPost(ctx).PoolDatasetUserpropCreate0(poolDatasetUserpropCreate0).Execute()





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
    poolDatasetUserpropCreate0 := *openapiclient.NewPoolDatasetUserpropCreate0() // PoolDatasetUserpropCreate0 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoolDatasetUserpropApi.PoolDatasetUserpropPost(context.Background()).PoolDatasetUserpropCreate0(poolDatasetUserpropCreate0).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoolDatasetUserpropApi.PoolDatasetUserpropPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPoolDatasetUserpropPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **poolDatasetUserpropCreate0** | [**PoolDatasetUserpropCreate0**](PoolDatasetUserpropCreate0.md) |  | 

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

