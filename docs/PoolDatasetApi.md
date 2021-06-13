# \PoolDatasetApi

All URIs are relative to *http://nas.local.geekspace.us/api/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PoolDatasetChangeKeyPost**](PoolDatasetApi.md#PoolDatasetChangeKeyPost) | **Post** /pool/dataset/change_key | 
[**PoolDatasetCompressionChoicesGet**](PoolDatasetApi.md#PoolDatasetCompressionChoicesGet) | **Get** /pool/dataset/compression_choices | 
[**PoolDatasetEncryptionAlgorithmChoicesGet**](PoolDatasetApi.md#PoolDatasetEncryptionAlgorithmChoicesGet) | **Get** /pool/dataset/encryption_algorithm_choices | 
[**PoolDatasetEncryptionSummaryPost**](PoolDatasetApi.md#PoolDatasetEncryptionSummaryPost) | **Post** /pool/dataset/encryption_summary | 
[**PoolDatasetExportKeyPost**](PoolDatasetApi.md#PoolDatasetExportKeyPost) | **Post** /pool/dataset/export_key | 
[**PoolDatasetGet**](PoolDatasetApi.md#PoolDatasetGet) | **Get** /pool/dataset | 
[**PoolDatasetIdIdAttachmentsPost**](PoolDatasetApi.md#PoolDatasetIdIdAttachmentsPost) | **Post** /pool/dataset/id/{id}/attachments | 
[**PoolDatasetIdIdDelete**](PoolDatasetApi.md#PoolDatasetIdIdDelete) | **Delete** /pool/dataset/id/{id} | 
[**PoolDatasetIdIdGet**](PoolDatasetApi.md#PoolDatasetIdIdGet) | **Get** /pool/dataset/id/{id} | 
[**PoolDatasetIdIdGetQuotaPost**](PoolDatasetApi.md#PoolDatasetIdIdGetQuotaPost) | **Post** /pool/dataset/id/{id}/get_quota | 
[**PoolDatasetIdIdPermissionPost**](PoolDatasetApi.md#PoolDatasetIdIdPermissionPost) | **Post** /pool/dataset/id/{id}/permission | 
[**PoolDatasetIdIdProcessesPost**](PoolDatasetApi.md#PoolDatasetIdIdProcessesPost) | **Post** /pool/dataset/id/{id}/processes | 
[**PoolDatasetIdIdPromotePost**](PoolDatasetApi.md#PoolDatasetIdIdPromotePost) | **Post** /pool/dataset/id/{id}/promote | 
[**PoolDatasetIdIdPut**](PoolDatasetApi.md#PoolDatasetIdIdPut) | **Put** /pool/dataset/id/{id} | 
[**PoolDatasetIdIdSetQuotaPost**](PoolDatasetApi.md#PoolDatasetIdIdSetQuotaPost) | **Post** /pool/dataset/id/{id}/set_quota | 
[**PoolDatasetInheritParentEncryptionPropertiesPost**](PoolDatasetApi.md#PoolDatasetInheritParentEncryptionPropertiesPost) | **Post** /pool/dataset/inherit_parent_encryption_properties | 
[**PoolDatasetLockPost**](PoolDatasetApi.md#PoolDatasetLockPost) | **Post** /pool/dataset/lock | 
[**PoolDatasetPost**](PoolDatasetApi.md#PoolDatasetPost) | **Post** /pool/dataset | 
[**PoolDatasetRecommendedZvolBlocksizePost**](PoolDatasetApi.md#PoolDatasetRecommendedZvolBlocksizePost) | **Post** /pool/dataset/recommended_zvol_blocksize | 
[**PoolDatasetUnlockPost**](PoolDatasetApi.md#PoolDatasetUnlockPost) | **Post** /pool/dataset/unlock | 



## PoolDatasetChangeKeyPost

> PoolDatasetChangeKeyPost(ctx).PoolDatasetChangeKey(poolDatasetChangeKey).Execute()





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
    poolDatasetChangeKey := *openapiclient.NewPoolDatasetChangeKey() // PoolDatasetChangeKey |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoolDatasetApi.PoolDatasetChangeKeyPost(context.Background()).PoolDatasetChangeKey(poolDatasetChangeKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoolDatasetApi.PoolDatasetChangeKeyPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPoolDatasetChangeKeyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **poolDatasetChangeKey** | [**PoolDatasetChangeKey**](PoolDatasetChangeKey.md) |  | 

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


## PoolDatasetCompressionChoicesGet

> PoolDatasetCompressionChoicesGet(ctx).Execute()





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
    resp, r, err := api_client.PoolDatasetApi.PoolDatasetCompressionChoicesGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoolDatasetApi.PoolDatasetCompressionChoicesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPoolDatasetCompressionChoicesGetRequest struct via the builder pattern


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


## PoolDatasetEncryptionAlgorithmChoicesGet

> PoolDatasetEncryptionAlgorithmChoicesGet(ctx).Execute()





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
    resp, r, err := api_client.PoolDatasetApi.PoolDatasetEncryptionAlgorithmChoicesGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoolDatasetApi.PoolDatasetEncryptionAlgorithmChoicesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPoolDatasetEncryptionAlgorithmChoicesGetRequest struct via the builder pattern


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


## PoolDatasetEncryptionSummaryPost

> PoolDatasetEncryptionSummaryPost(ctx).PoolDatasetEncryptionSummary(poolDatasetEncryptionSummary).Execute()





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
    poolDatasetEncryptionSummary := *openapiclient.NewPoolDatasetEncryptionSummary() // PoolDatasetEncryptionSummary |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoolDatasetApi.PoolDatasetEncryptionSummaryPost(context.Background()).PoolDatasetEncryptionSummary(poolDatasetEncryptionSummary).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoolDatasetApi.PoolDatasetEncryptionSummaryPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPoolDatasetEncryptionSummaryPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **poolDatasetEncryptionSummary** | [**PoolDatasetEncryptionSummary**](PoolDatasetEncryptionSummary.md) |  | 

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


## PoolDatasetExportKeyPost

> PoolDatasetExportKeyPost(ctx).PoolDatasetExportKey(poolDatasetExportKey).Execute()





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
    poolDatasetExportKey := *openapiclient.NewPoolDatasetExportKey() // PoolDatasetExportKey |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoolDatasetApi.PoolDatasetExportKeyPost(context.Background()).PoolDatasetExportKey(poolDatasetExportKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoolDatasetApi.PoolDatasetExportKeyPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPoolDatasetExportKeyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **poolDatasetExportKey** | [**PoolDatasetExportKey**](PoolDatasetExportKey.md) |  | 

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


## PoolDatasetGet

> PoolDatasetGet(ctx).Limit(limit).Offset(offset).Count(count).Sort(sort).Execute()





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
    resp, r, err := api_client.PoolDatasetApi.PoolDatasetGet(context.Background()).Limit(limit).Offset(offset).Count(count).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoolDatasetApi.PoolDatasetGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPoolDatasetGetRequest struct via the builder pattern


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


## PoolDatasetIdIdAttachmentsPost

> PoolDatasetIdIdAttachmentsPost(ctx, id).Body(body).Execute()





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
    body := map[string]interface{}(Object) // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoolDatasetApi.PoolDatasetIdIdAttachmentsPost(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoolDatasetApi.PoolDatasetIdIdAttachmentsPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPoolDatasetIdIdAttachmentsPostRequest struct via the builder pattern


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


## PoolDatasetIdIdDelete

> PoolDatasetIdIdDelete(ctx, id).PoolDatasetDelete1(poolDatasetDelete1).Execute()





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
    poolDatasetDelete1 := *openapiclient.NewPoolDatasetDelete1() // PoolDatasetDelete1 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoolDatasetApi.PoolDatasetIdIdDelete(context.Background(), id).PoolDatasetDelete1(poolDatasetDelete1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoolDatasetApi.PoolDatasetIdIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPoolDatasetIdIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **poolDatasetDelete1** | [**PoolDatasetDelete1**](PoolDatasetDelete1.md) |  | 

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


## PoolDatasetIdIdGet

> PoolDatasetIdIdGet(ctx, id).Limit(limit).Offset(offset).Count(count).Sort(sort).Execute()





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
    resp, r, err := api_client.PoolDatasetApi.PoolDatasetIdIdGet(context.Background(), id).Limit(limit).Offset(offset).Count(count).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoolDatasetApi.PoolDatasetIdIdGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPoolDatasetIdIdGetRequest struct via the builder pattern


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


## PoolDatasetIdIdGetQuotaPost

> PoolDatasetIdIdGetQuotaPost(ctx, id).PoolDatasetGetQuota(poolDatasetGetQuota).Execute()





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
    poolDatasetGetQuota := *openapiclient.NewPoolDatasetGetQuota() // PoolDatasetGetQuota |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoolDatasetApi.PoolDatasetIdIdGetQuotaPost(context.Background(), id).PoolDatasetGetQuota(poolDatasetGetQuota).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoolDatasetApi.PoolDatasetIdIdGetQuotaPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPoolDatasetIdIdGetQuotaPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **poolDatasetGetQuota** | [**PoolDatasetGetQuota**](PoolDatasetGetQuota.md) |  | 

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


## PoolDatasetIdIdPermissionPost

> PoolDatasetIdIdPermissionPost(ctx, id).PoolDatasetPermission1(poolDatasetPermission1).Execute()





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
    poolDatasetPermission1 := *openapiclient.NewPoolDatasetPermission1() // PoolDatasetPermission1 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoolDatasetApi.PoolDatasetIdIdPermissionPost(context.Background(), id).PoolDatasetPermission1(poolDatasetPermission1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoolDatasetApi.PoolDatasetIdIdPermissionPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPoolDatasetIdIdPermissionPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **poolDatasetPermission1** | [**PoolDatasetPermission1**](PoolDatasetPermission1.md) |  | 

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


## PoolDatasetIdIdProcessesPost

> PoolDatasetIdIdProcessesPost(ctx, id).Body(body).Execute()





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
    body := map[string]interface{}(Object) // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoolDatasetApi.PoolDatasetIdIdProcessesPost(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoolDatasetApi.PoolDatasetIdIdProcessesPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPoolDatasetIdIdProcessesPostRequest struct via the builder pattern


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


## PoolDatasetIdIdPromotePost

> PoolDatasetIdIdPromotePost(ctx, id).Body(body).Execute()





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
    body := map[string]interface{}(Object) // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoolDatasetApi.PoolDatasetIdIdPromotePost(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoolDatasetApi.PoolDatasetIdIdPromotePost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPoolDatasetIdIdPromotePostRequest struct via the builder pattern


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


## PoolDatasetIdIdPut

> PoolDatasetIdIdPut(ctx, id).PoolDatasetUpdate1(poolDatasetUpdate1).Execute()





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
    poolDatasetUpdate1 := *openapiclient.NewPoolDatasetUpdate1() // PoolDatasetUpdate1 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoolDatasetApi.PoolDatasetIdIdPut(context.Background(), id).PoolDatasetUpdate1(poolDatasetUpdate1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoolDatasetApi.PoolDatasetIdIdPut``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPoolDatasetIdIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **poolDatasetUpdate1** | [**PoolDatasetUpdate1**](PoolDatasetUpdate1.md) |  | 

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


## PoolDatasetIdIdSetQuotaPost

> PoolDatasetIdIdSetQuotaPost(ctx, id).RequestBody(requestBody).Execute()





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
    requestBody := []map[string]interface{}{map[string]interface{}(123)} // []map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoolDatasetApi.PoolDatasetIdIdSetQuotaPost(context.Background(), id).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoolDatasetApi.PoolDatasetIdIdSetQuotaPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPoolDatasetIdIdSetQuotaPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **[]map[string]interface{}** |  | 

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


## PoolDatasetInheritParentEncryptionPropertiesPost

> PoolDatasetInheritParentEncryptionPropertiesPost(ctx).Body(body).Execute()





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
    body := "body_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoolDatasetApi.PoolDatasetInheritParentEncryptionPropertiesPost(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoolDatasetApi.PoolDatasetInheritParentEncryptionPropertiesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPoolDatasetInheritParentEncryptionPropertiesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** |  | 

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


## PoolDatasetLockPost

> PoolDatasetLockPost(ctx).PoolDatasetLock(poolDatasetLock).Execute()





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
    poolDatasetLock := *openapiclient.NewPoolDatasetLock() // PoolDatasetLock |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoolDatasetApi.PoolDatasetLockPost(context.Background()).PoolDatasetLock(poolDatasetLock).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoolDatasetApi.PoolDatasetLockPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPoolDatasetLockPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **poolDatasetLock** | [**PoolDatasetLock**](PoolDatasetLock.md) |  | 

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


## PoolDatasetPost

> PoolDatasetPost(ctx).PoolDatasetCreate0(poolDatasetCreate0).Execute()





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
    poolDatasetCreate0 := *openapiclient.NewPoolDatasetCreate0() // PoolDatasetCreate0 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoolDatasetApi.PoolDatasetPost(context.Background()).PoolDatasetCreate0(poolDatasetCreate0).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoolDatasetApi.PoolDatasetPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPoolDatasetPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **poolDatasetCreate0** | [**PoolDatasetCreate0**](PoolDatasetCreate0.md) |  | 

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


## PoolDatasetRecommendedZvolBlocksizePost

> PoolDatasetRecommendedZvolBlocksizePost(ctx).Body(body).Execute()





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
    body := "body_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoolDatasetApi.PoolDatasetRecommendedZvolBlocksizePost(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoolDatasetApi.PoolDatasetRecommendedZvolBlocksizePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPoolDatasetRecommendedZvolBlocksizePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** |  | 

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


## PoolDatasetUnlockPost

> PoolDatasetUnlockPost(ctx).PoolDatasetUnlock(poolDatasetUnlock).Execute()





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
    poolDatasetUnlock := *openapiclient.NewPoolDatasetUnlock() // PoolDatasetUnlock |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoolDatasetApi.PoolDatasetUnlockPost(context.Background()).PoolDatasetUnlock(poolDatasetUnlock).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoolDatasetApi.PoolDatasetUnlockPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPoolDatasetUnlockPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **poolDatasetUnlock** | [**PoolDatasetUnlock**](PoolDatasetUnlock.md) |  | 

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

