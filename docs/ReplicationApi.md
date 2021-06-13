# \ReplicationApi

All URIs are relative to *http://nas.local.geekspace.us/api/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReplicationCountEligibleManualSnapshotsPost**](ReplicationApi.md#ReplicationCountEligibleManualSnapshotsPost) | **Post** /replication/count_eligible_manual_snapshots | 
[**ReplicationCreateDatasetPost**](ReplicationApi.md#ReplicationCreateDatasetPost) | **Post** /replication/create_dataset | 
[**ReplicationGet**](ReplicationApi.md#ReplicationGet) | **Get** /replication | 
[**ReplicationIdIdDelete**](ReplicationApi.md#ReplicationIdIdDelete) | **Delete** /replication/id/{id} | 
[**ReplicationIdIdGet**](ReplicationApi.md#ReplicationIdIdGet) | **Get** /replication/id/{id} | 
[**ReplicationIdIdPut**](ReplicationApi.md#ReplicationIdIdPut) | **Put** /replication/id/{id} | 
[**ReplicationIdIdRestorePost**](ReplicationApi.md#ReplicationIdIdRestorePost) | **Post** /replication/id/{id}/restore | 
[**ReplicationIdIdRunPost**](ReplicationApi.md#ReplicationIdIdRunPost) | **Post** /replication/id/{id}/run | 
[**ReplicationListDatasetsPost**](ReplicationApi.md#ReplicationListDatasetsPost) | **Post** /replication/list_datasets | 
[**ReplicationListNamingSchemasGet**](ReplicationApi.md#ReplicationListNamingSchemasGet) | **Get** /replication/list_naming_schemas | 
[**ReplicationPost**](ReplicationApi.md#ReplicationPost) | **Post** /replication | 
[**ReplicationTargetUnmatchedSnapshotsPost**](ReplicationApi.md#ReplicationTargetUnmatchedSnapshotsPost) | **Post** /replication/target_unmatched_snapshots | 



## ReplicationCountEligibleManualSnapshotsPost

> ReplicationCountEligibleManualSnapshotsPost(ctx).ReplicationCountEligibleManualSnapshots(replicationCountEligibleManualSnapshots).Execute()





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
    replicationCountEligibleManualSnapshots := *openapiclient.NewReplicationCountEligibleManualSnapshots() // ReplicationCountEligibleManualSnapshots |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReplicationApi.ReplicationCountEligibleManualSnapshotsPost(context.Background()).ReplicationCountEligibleManualSnapshots(replicationCountEligibleManualSnapshots).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReplicationApi.ReplicationCountEligibleManualSnapshotsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReplicationCountEligibleManualSnapshotsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **replicationCountEligibleManualSnapshots** | [**ReplicationCountEligibleManualSnapshots**](ReplicationCountEligibleManualSnapshots.md) |  | 

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


## ReplicationCreateDatasetPost

> ReplicationCreateDatasetPost(ctx).ReplicationCreateDataset(replicationCreateDataset).Execute()





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
    replicationCreateDataset := *openapiclient.NewReplicationCreateDataset() // ReplicationCreateDataset |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReplicationApi.ReplicationCreateDatasetPost(context.Background()).ReplicationCreateDataset(replicationCreateDataset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReplicationApi.ReplicationCreateDatasetPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReplicationCreateDatasetPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **replicationCreateDataset** | [**ReplicationCreateDataset**](ReplicationCreateDataset.md) |  | 

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


## ReplicationGet

> ReplicationGet(ctx).Limit(limit).Offset(offset).Count(count).Sort(sort).Execute()



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
    resp, r, err := api_client.ReplicationApi.ReplicationGet(context.Background()).Limit(limit).Offset(offset).Count(count).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReplicationApi.ReplicationGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReplicationGetRequest struct via the builder pattern


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


## ReplicationIdIdDelete

> ReplicationIdIdDelete(ctx, id).Execute()





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
    resp, r, err := api_client.ReplicationApi.ReplicationIdIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReplicationApi.ReplicationIdIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiReplicationIdIdDeleteRequest struct via the builder pattern


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


## ReplicationIdIdGet

> ReplicationIdIdGet(ctx, id).Limit(limit).Offset(offset).Count(count).Sort(sort).Execute()



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
    resp, r, err := api_client.ReplicationApi.ReplicationIdIdGet(context.Background(), id).Limit(limit).Offset(offset).Count(count).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReplicationApi.ReplicationIdIdGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiReplicationIdIdGetRequest struct via the builder pattern


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


## ReplicationIdIdPut

> ReplicationIdIdPut(ctx, id).ReplicationUpdate1(replicationUpdate1).Execute()





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
    replicationUpdate1 := *openapiclient.NewReplicationUpdate1() // ReplicationUpdate1 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReplicationApi.ReplicationIdIdPut(context.Background(), id).ReplicationUpdate1(replicationUpdate1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReplicationApi.ReplicationIdIdPut``: %v\n", err)
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

Other parameters are passed through a pointer to a apiReplicationIdIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **replicationUpdate1** | [**ReplicationUpdate1**](ReplicationUpdate1.md) |  | 

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


## ReplicationIdIdRestorePost

> ReplicationIdIdRestorePost(ctx, id).ReplicationRestore1(replicationRestore1).Execute()





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
    replicationRestore1 := *openapiclient.NewReplicationRestore1() // ReplicationRestore1 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReplicationApi.ReplicationIdIdRestorePost(context.Background(), id).ReplicationRestore1(replicationRestore1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReplicationApi.ReplicationIdIdRestorePost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiReplicationIdIdRestorePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **replicationRestore1** | [**ReplicationRestore1**](ReplicationRestore1.md) |  | 

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


## ReplicationIdIdRunPost

> ReplicationIdIdRunPost(ctx, id).Body(body).Execute()





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
    resp, r, err := api_client.ReplicationApi.ReplicationIdIdRunPost(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReplicationApi.ReplicationIdIdRunPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiReplicationIdIdRunPostRequest struct via the builder pattern


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


## ReplicationListDatasetsPost

> ReplicationListDatasetsPost(ctx).ReplicationListDatasets(replicationListDatasets).Execute()





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
    replicationListDatasets := *openapiclient.NewReplicationListDatasets() // ReplicationListDatasets |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReplicationApi.ReplicationListDatasetsPost(context.Background()).ReplicationListDatasets(replicationListDatasets).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReplicationApi.ReplicationListDatasetsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReplicationListDatasetsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **replicationListDatasets** | [**ReplicationListDatasets**](ReplicationListDatasets.md) |  | 

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


## ReplicationListNamingSchemasGet

> ReplicationListNamingSchemasGet(ctx).Execute()





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
    resp, r, err := api_client.ReplicationApi.ReplicationListNamingSchemasGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReplicationApi.ReplicationListNamingSchemasGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiReplicationListNamingSchemasGetRequest struct via the builder pattern


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


## ReplicationPost

> ReplicationPost(ctx).ReplicationCreate0(replicationCreate0).Execute()





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
    replicationCreate0 := *openapiclient.NewReplicationCreate0() // ReplicationCreate0 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReplicationApi.ReplicationPost(context.Background()).ReplicationCreate0(replicationCreate0).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReplicationApi.ReplicationPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReplicationPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **replicationCreate0** | [**ReplicationCreate0**](ReplicationCreate0.md) |  | 

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


## ReplicationTargetUnmatchedSnapshotsPost

> ReplicationTargetUnmatchedSnapshotsPost(ctx).ReplicationTargetUnmatchedSnapshots(replicationTargetUnmatchedSnapshots).Execute()





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
    replicationTargetUnmatchedSnapshots := *openapiclient.NewReplicationTargetUnmatchedSnapshots() // ReplicationTargetUnmatchedSnapshots |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReplicationApi.ReplicationTargetUnmatchedSnapshotsPost(context.Background()).ReplicationTargetUnmatchedSnapshots(replicationTargetUnmatchedSnapshots).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReplicationApi.ReplicationTargetUnmatchedSnapshotsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReplicationTargetUnmatchedSnapshotsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **replicationTargetUnmatchedSnapshots** | [**ReplicationTargetUnmatchedSnapshots**](ReplicationTargetUnmatchedSnapshots.md) |  | 

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

