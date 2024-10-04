# ListDeviceStat

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListDeviceStat**](ListDeviceStat.md#ListDeviceStat) | **Get** /v3/devices/stats | Retrieve email statistics by device type.



## ListDeviceStat

> []ListClientStat200ResponseInner ListDeviceStat(ctx, StartDateoptional)

Retrieve email statistics by device type.

**This endpoint allows you to retrieve your email statistics segmented by the device type.**  **We only store up to 7 days of email activity in our database.** By default, 500 items will be returned per request via the Advanced Stats API endpoints.  ## Available Device Types | **Device** | **Description** | **Example** | |---|---|---| | Desktop | Email software on desktop computer. | I.E., Outlook, Sparrow, or Apple Mail. | | Webmail | A web-based email client. | I.E., Yahoo, Google, AOL, or Outlook.com. | | Phone | A smart phone. | iPhone, Android, Blackberry, etc. | Tablet | A tablet computer. | iPad, android based tablet, etc. | | Other | An unrecognized device. |  Advanced Stats provide a more in-depth view of your email statistics and the actions taken by your recipients. You can segment these statistics by geographic location, device type, client type, browser, and mailbox provider. For more information about statistics, please see our [Statistics Overview](https://sendgrid.com/docs/ui/analytics-and-reporting/stats-overview/).

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListDeviceStatParams struct


Name | Type | Description
------------- | ------------- | -------------
**Onbehalfof** | **string** | The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.
**Limit** | **int32** | The number of results to return.
**Offset** | **int32** | The point in the list to begin retrieving results.
**AggregatedBy** | [**AggregatedBy3**](AggregatedBy3AggregatedBy3.md) | How to group the statistics. Must be either \"day\", \"week\", or \"month\".
**EndDate** | **string** | The end date of the statistics to retrieve. Defaults to today. Must follow format YYYY-MM-DD.

### Return type

[**[]ListClientStat200ResponseInner**](ListClientStat200ResponseInner.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

