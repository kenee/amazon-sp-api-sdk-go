/*
 * Selling Partner API for Direct Fulfillment Orders
 *
 * The Selling Partner API for Direct Fulfillment Orders provides programmatic access to a direct fulfillment vendor's order data.
 *
 * The version of the OpenAPI document: 2021-12-28
 * Generated by: https://github.com/openapitools/openapi-generator.git
 */


using Xunit;

using System;
using System.Linq;
using System.IO;
using System.Collections.Generic;
using software.amzn.spapi.Model.vendor.df.orders.v2021_12_28;
using software.amzn.spapi.Client;
using System.Reflection;
using Newtonsoft.Json;

namespace software.amzn.spapi.Test.Model
{
    /// <summary>
    ///  Class for testing OrderItemAcknowledgement
    /// </summary>
    /// <remarks>
    /// This file is automatically generated by OpenAPI Generator (https://openapi-generator.tech).
    /// Please update the test case below to test the model.
    /// </remarks>
    public class OrderItemAcknowledgementTests : IDisposable
    {
        // TODO uncomment below to declare an instance variable for OrderItemAcknowledgement
        //private OrderItemAcknowledgement instance;

        public OrderItemAcknowledgementTests()
        {
            // TODO uncomment below to create an instance of OrderItemAcknowledgement
            //instance = new OrderItemAcknowledgement();
        }

        public void Dispose()
        {
            // Cleanup when everything is done.
        }

        /// <summary>
        /// Test an instance of OrderItemAcknowledgement
        /// </summary>
        [Fact]
        public void OrderItemAcknowledgementInstanceTest()
        {
            // TODO uncomment below to test "IsType" OrderItemAcknowledgement
            //Assert.IsType<OrderItemAcknowledgement>(instance);
        }

        /// <summary>
        /// Test the property 'ItemSequenceNumber'
        /// </summary>
        [Fact]
        public void ItemSequenceNumberTest()
        {
            // TODO unit test for the property 'ItemSequenceNumber'
        }

        /// <summary>
        /// Test the property 'BuyerProductIdentifier'
        /// </summary>
        [Fact]
        public void BuyerProductIdentifierTest()
        {
            // TODO unit test for the property 'BuyerProductIdentifier'
        }

        /// <summary>
        /// Test the property 'VendorProductIdentifier'
        /// </summary>
        [Fact]
        public void VendorProductIdentifierTest()
        {
            // TODO unit test for the property 'VendorProductIdentifier'
        }

        /// <summary>
        /// Test the property 'AcknowledgedQuantity'
        /// </summary>
        [Fact]
        public void AcknowledgedQuantityTest()
        {
            // TODO unit test for the property 'AcknowledgedQuantity'
        }
    }
}
