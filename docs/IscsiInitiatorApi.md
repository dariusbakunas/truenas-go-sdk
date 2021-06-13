# \IscsiInitiatorApi

All URIs are relative to *http://nas.local.geekspace.us/api/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IscsiInitiatorGet**](IscsiInitiatorApi.md#IscsiInitiatorGet) | **Get** /iscsi/initiator | 
[**IscsiInitiatorIdIdDelete**](IscsiInitiatorApi.md#IscsiInitiatorIdIdDelete) | **Delete** /iscsi/initiator/id/{id} | 
[**IscsiInitiatorIdIdGet**](IscsiInitiatorApi.md#IscsiInitiatorIdIdGet) | **Get** /iscsi/initiator/id/{id} | 
[**IscsiInitiatorIdIdPut**](IscsiInitiatorApi.md#IscsiInitiatorIdIdPut) | **Put** /iscsi/initiator/id/{id} | 
[**IscsiInitiatorPost**](IscsiInitiatorApi.md#IscsiInitiatorPost) | **Post** /iscsi/initiator | 



## IscsiInitiatorGet

> IscsiInitiatorGet(ctx).Limit(limit).Offset(offset).Count(count).Sort(sort).Execute()



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
    resp, r, err := api_client.IscsiInitiatorApi.IscsiInitiatorGet(context.Background()).Limit(limit).Offset(offset).Count(count).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IscsiInitiatorApi.IscsiInitiatorGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIscsiInitiatorGetRequest struct via the builder pattern


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


## IscsiInitiatorIdIdDelete

> IscsiInitiatorIdIdDelete(ctx, id).Execute()





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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IscsiInitiatorApi.IscsiInitiatorIdIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IscsiInitiatorApi.IscsiInitiatorIdIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiIscsiInitiatorIdIdDeleteRequest struct via the builder pattern


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


## IscsiInitiatorIdIdGet

> IscsiInitiatorIdIdGet(ctx, id).Limit(limit).Offset(offset).Count(count).Sort(sort).Execute()



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
    resp, r, err := api_client.IscsiInitiatorApi.IscsiInitiatorIdIdGet(context.Background(), id).Limit(limit).Offset(offset).Count(count).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IscsiInitiatorApi.IscsiInitiatorIdIdGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiIscsiInitiatorIdIdGetRequest struct via the builder pattern


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


## IscsiInitiatorIdIdPut

> IscsiInitiatorIdIdPut(ctx, id).IscsiInitiatorUpdate1(iscsiInitiatorUpdate1).Execute()





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
    iscsiInitiatorUpdate1 := *openapiclient.NewIscsiInitiatorUpdate1() // IscsiInitiatorUpdate1 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IscsiInitiatorApi.IscsiInitiatorIdIdPut(context.Background(), id).IscsiInitiatorUpdate1(iscsiInitiatorUpdate1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IscsiInitiatorApi.IscsiInitiatorIdIdPut``: %v\n", err)
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

Other parameters are passed through a pointer to a apiIscsiInitiatorIdIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iscsiInitiatorUpdate1** | [**IscsiInitiatorUpdate1**](IscsiInitiatorUpdate1.md) |  | 

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


## IscsiInitiatorPost

> IscsiInitiatorPost(ctx).IscsiInitiatorCreate0(iscsiInitiatorCreate0).Execute()





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
    iscsiInitiatorCreate0 := *openapiclient.NewIscsiInitiatorCreate0() // IscsiInitiatorCreate0 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IscsiInitiatorApi.IscsiInitiatorPost(context.Background()).IscsiInitiatorCreate0(iscsiInitiatorCreate0).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IscsiInitiatorApi.IscsiInitiatorPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIscsiInitiatorPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iscsiInitiatorCreate0** | [**IscsiInitiatorCreate0**](IscsiInitiatorCreate0.md) |  | 

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

