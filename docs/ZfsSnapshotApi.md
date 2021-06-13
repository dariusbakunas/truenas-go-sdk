# \ZfsSnapshotApi

All URIs are relative to *http://nas.local.geekspace.us/api/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ZfsSnapshotClonePost**](ZfsSnapshotApi.md#ZfsSnapshotClonePost) | **Post** /zfs/snapshot/clone | 
[**ZfsSnapshotGet**](ZfsSnapshotApi.md#ZfsSnapshotGet) | **Get** /zfs/snapshot | 
[**ZfsSnapshotIdIdDelete**](ZfsSnapshotApi.md#ZfsSnapshotIdIdDelete) | **Delete** /zfs/snapshot/id/{id} | 
[**ZfsSnapshotIdIdGet**](ZfsSnapshotApi.md#ZfsSnapshotIdIdGet) | **Get** /zfs/snapshot/id/{id} | 
[**ZfsSnapshotPost**](ZfsSnapshotApi.md#ZfsSnapshotPost) | **Post** /zfs/snapshot | 
[**ZfsSnapshotRemovePost**](ZfsSnapshotApi.md#ZfsSnapshotRemovePost) | **Post** /zfs/snapshot/remove | 
[**ZfsSnapshotRollbackPost**](ZfsSnapshotApi.md#ZfsSnapshotRollbackPost) | **Post** /zfs/snapshot/rollback | 



## ZfsSnapshotClonePost

> ZfsSnapshotClonePost(ctx).ZfsSnapshotClone0(zfsSnapshotClone0).Execute()





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
    zfsSnapshotClone0 := *openapiclient.NewZfsSnapshotClone0() // ZfsSnapshotClone0 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ZfsSnapshotApi.ZfsSnapshotClonePost(context.Background()).ZfsSnapshotClone0(zfsSnapshotClone0).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZfsSnapshotApi.ZfsSnapshotClonePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiZfsSnapshotClonePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zfsSnapshotClone0** | [**ZfsSnapshotClone0**](ZfsSnapshotClone0.md) |  | 

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


## ZfsSnapshotGet

> ZfsSnapshotGet(ctx).Limit(limit).Offset(offset).Count(count).Sort(sort).Execute()





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
    resp, r, err := api_client.ZfsSnapshotApi.ZfsSnapshotGet(context.Background()).Limit(limit).Offset(offset).Count(count).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZfsSnapshotApi.ZfsSnapshotGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiZfsSnapshotGetRequest struct via the builder pattern


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


## ZfsSnapshotIdIdDelete

> ZfsSnapshotIdIdDelete(ctx, id).ZfsSnapshotDelete1(zfsSnapshotDelete1).Execute()





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
    zfsSnapshotDelete1 := *openapiclient.NewZfsSnapshotDelete1() // ZfsSnapshotDelete1 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ZfsSnapshotApi.ZfsSnapshotIdIdDelete(context.Background(), id).ZfsSnapshotDelete1(zfsSnapshotDelete1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZfsSnapshotApi.ZfsSnapshotIdIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiZfsSnapshotIdIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **zfsSnapshotDelete1** | [**ZfsSnapshotDelete1**](ZfsSnapshotDelete1.md) |  | 

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


## ZfsSnapshotIdIdGet

> ZfsSnapshotIdIdGet(ctx, id).Limit(limit).Offset(offset).Count(count).Sort(sort).Execute()





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
    resp, r, err := api_client.ZfsSnapshotApi.ZfsSnapshotIdIdGet(context.Background(), id).Limit(limit).Offset(offset).Count(count).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZfsSnapshotApi.ZfsSnapshotIdIdGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiZfsSnapshotIdIdGetRequest struct via the builder pattern


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


## ZfsSnapshotPost

> ZfsSnapshotPost(ctx).ZfsSnapshotCreate0(zfsSnapshotCreate0).Execute()





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
    zfsSnapshotCreate0 := *openapiclient.NewZfsSnapshotCreate0() // ZfsSnapshotCreate0 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ZfsSnapshotApi.ZfsSnapshotPost(context.Background()).ZfsSnapshotCreate0(zfsSnapshotCreate0).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZfsSnapshotApi.ZfsSnapshotPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiZfsSnapshotPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zfsSnapshotCreate0** | [**ZfsSnapshotCreate0**](ZfsSnapshotCreate0.md) |  | 

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


## ZfsSnapshotRemovePost

> ZfsSnapshotRemovePost(ctx).ZfsSnapshotRemove0(zfsSnapshotRemove0).Execute()





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
    zfsSnapshotRemove0 := *openapiclient.NewZfsSnapshotRemove0() // ZfsSnapshotRemove0 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ZfsSnapshotApi.ZfsSnapshotRemovePost(context.Background()).ZfsSnapshotRemove0(zfsSnapshotRemove0).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZfsSnapshotApi.ZfsSnapshotRemovePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiZfsSnapshotRemovePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zfsSnapshotRemove0** | [**ZfsSnapshotRemove0**](ZfsSnapshotRemove0.md) |  | 

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


## ZfsSnapshotRollbackPost

> ZfsSnapshotRollbackPost(ctx).ZfsSnapshotRollback(zfsSnapshotRollback).Execute()





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
    zfsSnapshotRollback := *openapiclient.NewZfsSnapshotRollback() // ZfsSnapshotRollback |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ZfsSnapshotApi.ZfsSnapshotRollbackPost(context.Background()).ZfsSnapshotRollback(zfsSnapshotRollback).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZfsSnapshotApi.ZfsSnapshotRollbackPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiZfsSnapshotRollbackPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zfsSnapshotRollback** | [**ZfsSnapshotRollback**](ZfsSnapshotRollback.md) |  | 

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

