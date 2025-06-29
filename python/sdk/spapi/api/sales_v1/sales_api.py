# coding: utf-8

"""
    Selling Partner API for Sales

    The Selling Partner API for Sales provides APIs related to sales performance.

    The version of the OpenAPI document: v1
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


from __future__ import absolute_import

import re  # noqa: F401

# python 2 and python 3 compatibility library
import six

from spapi.api_client import ApiClient


class SalesApi(object):
    """NOTE: This class is auto generated by the openapi generator.

    Do not edit the class manually.
    """

    api_models_module = "spapi.models.sales_v1"

    def __init__(self, api_client=None):
        if api_client is None:
            api_client = ApiClient()
        self.api_client = api_client
        self.classFileName = 'sales_api'

    def get_order_metrics(self, marketplace_ids, interval, granularity, **kwargs):  # noqa: E501
        """get_order_metrics  # noqa: E501

        Returns aggregated order metrics for given interval, broken down by granularity, for given buyer type.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | .5 | 15 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).  # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True
        >>> thread = api.get_order_metrics(marketplace_ids, interval, granularity, async_req=True)
        >>> result = thread.get()

        :param async_req bool
        :param List[str] marketplace_ids: A marketplace identifier. This specifies the marketplace in which the order was placed. Only one marketplace can be specified.  For example, ATVPDKIKX0DER indicates the US marketplace. (required)
        :param str interval: A time interval used for selecting order metrics. This takes the form of two dates separated by two hyphens (first date is inclusive; second date is exclusive). Dates are in ISO8601 format and must represent absolute time (either Z notation or offset notation). Example: 2018-09-01T00:00:00-07:00--2018-09-04T00:00:00-07:00 requests order metrics for Sept 1st, 2nd and 3rd in the -07:00 zone. (required)
        :param str granularity: The granularity of the grouping of order metrics, based on a unit of time. Specifying granularity=Hour results in a successful request only if the interval specified is less than or equal to 30 days from now. For all other granularities, the interval specified must be less or equal to 2 years from now. Specifying granularity=Total results in order metrics that are aggregated over the entire interval that you specify. If the interval start and end date don’t align with the specified granularity, the head and tail end of the response interval will contain partial data. Example: Day to get a daily breakdown of the request interval, where the day boundary is defined by the granularityTimeZone. (required)
        :param str granularity_time_zone: An IANA-compatible time zone for determining the day boundary. Required when specifying a granularity value greater than Hour. The granularityTimeZone value must align with the offset of the specified interval value. For example, if the interval value uses Z notation, then granularityTimeZone must be UTC. If the interval value uses an offset, then granularityTimeZone must be an IANA-compatible time zone that matches the offset. Example: US/Pacific to compute day boundaries, accounting for daylight time savings, for US/Pacific zone.
        :param str buyer_type: Filters the results by the buyer type that you specify, B2B (business to business) or B2C (business to customer). Example: B2B, if you want the response to include order metrics for only B2B buyers.
        :param str fulfillment_network: Filters the results by the fulfillment network that you specify, MFN (merchant fulfillment network) or AFN (Amazon fulfillment network). Do not include this filter if you want the response to include order metrics for all fulfillment networks. Example: AFN, if you want the response to include order metrics for only Amazon fulfillment network.
        :param str first_day_of_week: Specifies the day that the week starts on when granularity=Week, either Monday or Sunday. Default: Monday. Example: Sunday, if you want the week to start on a Sunday.
        :param str asin: Filters the results by the ASIN that you specify. Specifying both ASIN and SKU returns an error. Do not include this filter if you want the response to include order metrics for all ASINs. Example: B0792R1RSN, if you want the response to include order metrics for only ASIN B0792R1RSN.
        :param str sku: Filters the results by the SKU that you specify. Specifying both ASIN and SKU returns an error. Do not include this filter if you want the response to include order metrics for all SKUs. Example: TestSKU, if you want the response to include order metrics for only SKU TestSKU.
        :param str amazon_program: Filters the results by the Amazon program that you specify. Do not include this filter if you want the response to include order metrics for all programs. **Example:** `AmazonHaul` returns order metrics for the Amazon Haul program only.
        :return: GetOrderMetricsResponse
                 If the method is called asynchronously,
                 returns the request thread.
        """
        kwargs['_return_http_data_only'] = True
        if kwargs.get('async_req'):
            return self.get_order_metrics_with_http_info(marketplace_ids, interval, granularity, **kwargs)  # noqa: E501
        else:
            (data) = self.get_order_metrics_with_http_info(marketplace_ids, interval, granularity, **kwargs)  # noqa: E501
            return data

    def get_order_metrics_with_http_info(self, marketplace_ids, interval, granularity, **kwargs):  # noqa: E501
        """get_order_metrics  # noqa: E501

        Returns aggregated order metrics for given interval, broken down by granularity, for given buyer type.  **Usage Plan:**  | Rate (requests per second) | Burst | | ---- | ---- | | .5 | 15 |  The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).  # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True
        >>> thread = api.get_order_metrics_with_http_info(marketplace_ids, interval, granularity, async_req=True)
        >>> result = thread.get()

        :param async_req bool
        :param List[str] marketplace_ids: A marketplace identifier. This specifies the marketplace in which the order was placed. Only one marketplace can be specified.  For example, ATVPDKIKX0DER indicates the US marketplace. (required)
        :param str interval: A time interval used for selecting order metrics. This takes the form of two dates separated by two hyphens (first date is inclusive; second date is exclusive). Dates are in ISO8601 format and must represent absolute time (either Z notation or offset notation). Example: 2018-09-01T00:00:00-07:00--2018-09-04T00:00:00-07:00 requests order metrics for Sept 1st, 2nd and 3rd in the -07:00 zone. (required)
        :param str granularity: The granularity of the grouping of order metrics, based on a unit of time. Specifying granularity=Hour results in a successful request only if the interval specified is less than or equal to 30 days from now. For all other granularities, the interval specified must be less or equal to 2 years from now. Specifying granularity=Total results in order metrics that are aggregated over the entire interval that you specify. If the interval start and end date don’t align with the specified granularity, the head and tail end of the response interval will contain partial data. Example: Day to get a daily breakdown of the request interval, where the day boundary is defined by the granularityTimeZone. (required)
        :param str granularity_time_zone: An IANA-compatible time zone for determining the day boundary. Required when specifying a granularity value greater than Hour. The granularityTimeZone value must align with the offset of the specified interval value. For example, if the interval value uses Z notation, then granularityTimeZone must be UTC. If the interval value uses an offset, then granularityTimeZone must be an IANA-compatible time zone that matches the offset. Example: US/Pacific to compute day boundaries, accounting for daylight time savings, for US/Pacific zone.
        :param str buyer_type: Filters the results by the buyer type that you specify, B2B (business to business) or B2C (business to customer). Example: B2B, if you want the response to include order metrics for only B2B buyers.
        :param str fulfillment_network: Filters the results by the fulfillment network that you specify, MFN (merchant fulfillment network) or AFN (Amazon fulfillment network). Do not include this filter if you want the response to include order metrics for all fulfillment networks. Example: AFN, if you want the response to include order metrics for only Amazon fulfillment network.
        :param str first_day_of_week: Specifies the day that the week starts on when granularity=Week, either Monday or Sunday. Default: Monday. Example: Sunday, if you want the week to start on a Sunday.
        :param str asin: Filters the results by the ASIN that you specify. Specifying both ASIN and SKU returns an error. Do not include this filter if you want the response to include order metrics for all ASINs. Example: B0792R1RSN, if you want the response to include order metrics for only ASIN B0792R1RSN.
        :param str sku: Filters the results by the SKU that you specify. Specifying both ASIN and SKU returns an error. Do not include this filter if you want the response to include order metrics for all SKUs. Example: TestSKU, if you want the response to include order metrics for only SKU TestSKU.
        :param str amazon_program: Filters the results by the Amazon program that you specify. Do not include this filter if you want the response to include order metrics for all programs. **Example:** `AmazonHaul` returns order metrics for the Amazon Haul program only.
        :return: GetOrderMetricsResponse
                 If the method is called asynchronously,
                 returns the request thread.
        """

        all_params = ['marketplace_ids', 'interval', 'granularity', 'granularity_time_zone', 'buyer_type', 'fulfillment_network', 'first_day_of_week', 'asin', 'sku', 'amazon_program']  # noqa: E501
        all_params.append('async_req')
        all_params.append('_return_http_data_only')
        all_params.append('_preload_content')
        all_params.append('_request_timeout')

        params = locals()
        for key, val in six.iteritems(params['kwargs']):
            if key not in all_params:
                raise TypeError(
                    "Got an unexpected keyword argument '%s'"
                    " to method get_order_metrics" % key
                )
            params[key] = val
        del params['kwargs']
        # verify the required parameter 'marketplace_ids' is set
        if self.api_client.client_side_validation and ('marketplace_ids' not in params or
                                                       params['marketplace_ids'] is None):  # noqa: E501
            raise ValueError("Missing the required parameter `marketplace_ids` when calling `get_order_metrics`")  # noqa: E501
        # verify the required parameter 'interval' is set
        if self.api_client.client_side_validation and ('interval' not in params or
                                                       params['interval'] is None):  # noqa: E501
            raise ValueError("Missing the required parameter `interval` when calling `get_order_metrics`")  # noqa: E501
        # verify the required parameter 'granularity' is set
        if self.api_client.client_side_validation and ('granularity' not in params or
                                                       params['granularity'] is None):  # noqa: E501
            raise ValueError("Missing the required parameter `granularity` when calling `get_order_metrics`")  # noqa: E501

        collection_formats = {}

        path_params = {}

        query_params = []
        if 'marketplace_ids' in params:
            query_params.append(('marketplaceIds', params['marketplace_ids']))  # noqa: E501
            collection_formats['marketplaceIds'] = 'csv'  # noqa: E501
        if 'interval' in params:
            query_params.append(('interval', params['interval']))  # noqa: E501
        if 'granularity_time_zone' in params:
            query_params.append(('granularityTimeZone', params['granularity_time_zone']))  # noqa: E501
        if 'granularity' in params:
            query_params.append(('granularity', params['granularity']))  # noqa: E501
        if 'buyer_type' in params:
            query_params.append(('buyerType', params['buyer_type']))  # noqa: E501
        if 'fulfillment_network' in params:
            query_params.append(('fulfillmentNetwork', params['fulfillment_network']))  # noqa: E501
        if 'first_day_of_week' in params:
            query_params.append(('firstDayOfWeek', params['first_day_of_week']))  # noqa: E501
        if 'asin' in params:
            query_params.append(('asin', params['asin']))  # noqa: E501
        if 'sku' in params:
            query_params.append(('sku', params['sku']))  # noqa: E501
        if 'amazon_program' in params:
            query_params.append(('amazonProgram', params['amazon_program']))  # noqa: E501

        header_params = {}

        form_params = []
        local_var_files = {}

        body_params = None
        # HTTP header `Accept`
        header_params['Accept'] = self.api_client.select_header_accept(
            ['application/json', 'payload'])  # noqa: E501

        # Authentication setting
        auth_settings = []  # noqa: E501

        return self.api_client.call_api(
            '/sales/v1/orderMetrics', 'GET',
            path_params,
            query_params,
            header_params,
            body=body_params,
            post_params=form_params,
            files=local_var_files,
            response_type='GetOrderMetricsResponse',  # noqa: E501
            auth_settings=auth_settings,
            async_req=params.get('async_req'),
            _return_http_data_only=params.get('_return_http_data_only'),
            _preload_content=params.get('_preload_content', True),
            _request_timeout=params.get('_request_timeout'),
            collection_formats=collection_formats, api_models_module=self.api_models_module)
