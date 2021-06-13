# \CloudsyncApi

All URIs are relative to *http://nas.local.geekspace.us/api/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudsyncCommonTaskSchemaGet**](CloudsyncApi.md#CloudsyncCommonTaskSchemaGet) | **Get** /cloudsync/common_task_schema | 
[**CloudsyncGet**](CloudsyncApi.md#CloudsyncGet) | **Get** /cloudsync | 
[**CloudsyncIdIdAbortPost**](CloudsyncApi.md#CloudsyncIdIdAbortPost) | **Post** /cloudsync/id/{id}/abort | 
[**CloudsyncIdIdDelete**](CloudsyncApi.md#CloudsyncIdIdDelete) | **Delete** /cloudsync/id/{id} | 
[**CloudsyncIdIdGet**](CloudsyncApi.md#CloudsyncIdIdGet) | **Get** /cloudsync/id/{id} | 
[**CloudsyncIdIdPut**](CloudsyncApi.md#CloudsyncIdIdPut) | **Put** /cloudsync/id/{id} | 
[**CloudsyncIdIdRestorePost**](CloudsyncApi.md#CloudsyncIdIdRestorePost) | **Post** /cloudsync/id/{id}/restore | 
[**CloudsyncIdIdSyncPost**](CloudsyncApi.md#CloudsyncIdIdSyncPost) | **Post** /cloudsync/id/{id}/sync | 
[**CloudsyncListBucketsPost**](CloudsyncApi.md#CloudsyncListBucketsPost) | **Post** /cloudsync/list_buckets | 
[**CloudsyncListDirectoryPost**](CloudsyncApi.md#CloudsyncListDirectoryPost) | **Post** /cloudsync/list_directory | 
[**CloudsyncOnedriveListDrivesPost**](CloudsyncApi.md#CloudsyncOnedriveListDrivesPost) | **Post** /cloudsync/onedrive_list_drives | 
[**CloudsyncPost**](CloudsyncApi.md#CloudsyncPost) | **Post** /cloudsync | 
[**CloudsyncProvidersGet**](CloudsyncApi.md#CloudsyncProvidersGet) | **Get** /cloudsync/providers | 
[**CloudsyncSyncOnetimePost**](CloudsyncApi.md#CloudsyncSyncOnetimePost) | **Post** /cloudsync/sync_onetime | 



## CloudsyncCommonTaskSchemaGet

> CloudsyncCommonTaskSchemaGet(ctx).Execute()



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
    resp, r, err := api_client.CloudsyncApi.CloudsyncCommonTaskSchemaGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudsyncApi.CloudsyncCommonTaskSchemaGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudsyncCommonTaskSchemaGetRequest struct via the builder pattern


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


## CloudsyncGet

> CloudsyncGet(ctx).Limit(limit).Offset(offset).Count(count).Sort(sort).Execute()





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
    resp, r, err := api_client.CloudsyncApi.CloudsyncGet(context.Background()).Limit(limit).Offset(offset).Count(count).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudsyncApi.CloudsyncGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudsyncGetRequest struct via the builder pattern


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


## CloudsyncIdIdAbortPost

> CloudsyncIdIdAbortPost(ctx, id).Body(body).Execute()





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
    body := map[string]interface{}(Object) // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudsyncApi.CloudsyncIdIdAbortPost(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudsyncApi.CloudsyncIdIdAbortPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudsyncIdIdAbortPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

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


## CloudsyncIdIdDelete

> CloudsyncIdIdDelete(ctx, id).Execute()





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
    resp, r, err := api_client.CloudsyncApi.CloudsyncIdIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudsyncApi.CloudsyncIdIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudsyncIdIdDeleteRequest struct via the builder pattern


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


## CloudsyncIdIdGet

> CloudsyncIdIdGet(ctx, id).Limit(limit).Offset(offset).Count(count).Sort(sort).Execute()





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
    resp, r, err := api_client.CloudsyncApi.CloudsyncIdIdGet(context.Background(), id).Limit(limit).Offset(offset).Count(count).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudsyncApi.CloudsyncIdIdGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudsyncIdIdGetRequest struct via the builder pattern


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


## CloudsyncIdIdPut

> CloudsyncIdIdPut(ctx, id).CloudsyncUpdate1(cloudsyncUpdate1).Execute()





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
    cloudsyncUpdate1 := *openapiclient.NewCloudsyncUpdate1() // CloudsyncUpdate1 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudsyncApi.CloudsyncIdIdPut(context.Background(), id).CloudsyncUpdate1(cloudsyncUpdate1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudsyncApi.CloudsyncIdIdPut``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudsyncIdIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudsyncUpdate1** | [**CloudsyncUpdate1**](CloudsyncUpdate1.md) |  | 

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


## CloudsyncIdIdRestorePost

> CloudsyncIdIdRestorePost(ctx, id).CloudsyncRestore1(cloudsyncRestore1).Execute()





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
    cloudsyncRestore1 := *openapiclient.NewCloudsyncRestore1() // CloudsyncRestore1 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudsyncApi.CloudsyncIdIdRestorePost(context.Background(), id).CloudsyncRestore1(cloudsyncRestore1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudsyncApi.CloudsyncIdIdRestorePost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudsyncIdIdRestorePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudsyncRestore1** | [**CloudsyncRestore1**](CloudsyncRestore1.md) |  | 

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


## CloudsyncIdIdSyncPost

> CloudsyncIdIdSyncPost(ctx, id).CloudsyncSync1(cloudsyncSync1).Execute()





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
    cloudsyncSync1 := *openapiclient.NewCloudsyncSync1() // CloudsyncSync1 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudsyncApi.CloudsyncIdIdSyncPost(context.Background(), id).CloudsyncSync1(cloudsyncSync1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudsyncApi.CloudsyncIdIdSyncPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudsyncIdIdSyncPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudsyncSync1** | [**CloudsyncSync1**](CloudsyncSync1.md) |  | 

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


## CloudsyncListBucketsPost

> CloudsyncListBucketsPost(ctx).Body(body).Execute()



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
    body := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudsyncApi.CloudsyncListBucketsPost(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudsyncApi.CloudsyncListBucketsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudsyncListBucketsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **int32** |  | 

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


## CloudsyncListDirectoryPost

> CloudsyncListDirectoryPost(ctx).CloudsyncListDirectory0(cloudsyncListDirectory0).Execute()





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
    cloudsyncListDirectory0 := *openapiclient.NewCloudsyncListDirectory0() // CloudsyncListDirectory0 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudsyncApi.CloudsyncListDirectoryPost(context.Background()).CloudsyncListDirectory0(cloudsyncListDirectory0).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudsyncApi.CloudsyncListDirectoryPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudsyncListDirectoryPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudsyncListDirectory0** | [**CloudsyncListDirectory0**](CloudsyncListDirectory0.md) |  | 

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


## CloudsyncOnedriveListDrivesPost

> CloudsyncOnedriveListDrivesPost(ctx).CloudsyncOnedriveListDrives0(cloudsyncOnedriveListDrives0).Execute()





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
    cloudsyncOnedriveListDrives0 := *openapiclient.NewCloudsyncOnedriveListDrives0() // CloudsyncOnedriveListDrives0 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudsyncApi.CloudsyncOnedriveListDrivesPost(context.Background()).CloudsyncOnedriveListDrives0(cloudsyncOnedriveListDrives0).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudsyncApi.CloudsyncOnedriveListDrivesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudsyncOnedriveListDrivesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudsyncOnedriveListDrives0** | [**CloudsyncOnedriveListDrives0**](CloudsyncOnedriveListDrives0.md) |  | 

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


## CloudsyncPost

> CloudsyncPost(ctx).CloudsyncCreate0(cloudsyncCreate0).Execute()





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
    cloudsyncCreate0 := *openapiclient.NewCloudsyncCreate0() // CloudsyncCreate0 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudsyncApi.CloudsyncPost(context.Background()).CloudsyncCreate0(cloudsyncCreate0).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudsyncApi.CloudsyncPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudsyncPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudsyncCreate0** | [**CloudsyncCreate0**](CloudsyncCreate0.md) |  | 

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


## CloudsyncProvidersGet

> CloudsyncProvidersGet(ctx).Execute()





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
    resp, r, err := api_client.CloudsyncApi.CloudsyncProvidersGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudsyncApi.CloudsyncProvidersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudsyncProvidersGetRequest struct via the builder pattern


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


## CloudsyncSyncOnetimePost

> CloudsyncSyncOnetimePost(ctx).CloudsyncSyncOnetime(cloudsyncSyncOnetime).Execute()





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
    cloudsyncSyncOnetime := *openapiclient.NewCloudsyncSyncOnetime() // CloudsyncSyncOnetime |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudsyncApi.CloudsyncSyncOnetimePost(context.Background()).CloudsyncSyncOnetime(cloudsyncSyncOnetime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudsyncApi.CloudsyncSyncOnetimePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudsyncSyncOnetimePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudsyncSyncOnetime** | [**CloudsyncSyncOnetime**](CloudsyncSyncOnetime.md) |  | 

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

