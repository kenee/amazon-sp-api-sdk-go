/*
 * Selling Partner API for Pricing
 *
 * The Selling Partner API for Pricing helps you programmatically retrieve product pricing and offer pricing information for Amazon Marketplace products.  For more information, refer to the [Product Pricing v2022-05-01 Use Case Guide](https://developer-docs.amazon.com/sp-api/docs/product-pricing-api-v2022-05-01-use-case-guide).
 *
 * The version of the OpenAPI document: 2022-05-01
 * Generated by: https://github.com/openapitools/openapi-generator.git
 */


using Xunit;

using System;
using System.Linq;
using System.IO;
using System.Collections.Generic;
using software.amzn.spapi.Model.pricing.v2022_05_01;
using software.amzn.spapi.Client;
using System.Reflection;
using Newtonsoft.Json;

namespace software.amzn.spapi.Test.Model
{
    /// <summary>
    ///  Class for testing MoneyType
    /// </summary>
    /// <remarks>
    /// This file is automatically generated by OpenAPI Generator (https://openapi-generator.tech).
    /// Please update the test case below to test the model.
    /// </remarks>
    public class MoneyTypeTests : IDisposable
    {
        // TODO uncomment below to declare an instance variable for MoneyType
        //private MoneyType instance;

        public MoneyTypeTests()
        {
            // TODO uncomment below to create an instance of MoneyType
            //instance = new MoneyType();
        }

        public void Dispose()
        {
            // Cleanup when everything is done.
        }

        /// <summary>
        /// Test an instance of MoneyType
        /// </summary>
        [Fact]
        public void MoneyTypeInstanceTest()
        {
            // TODO uncomment below to test "IsType" MoneyType
            //Assert.IsType<MoneyType>(instance);
        }

        /// <summary>
        /// Test the property 'CurrencyCode'
        /// </summary>
        [Fact]
        public void CurrencyCodeTest()
        {
            // TODO unit test for the property 'CurrencyCode'
        }

        /// <summary>
        /// Test the property 'Amount'
        /// </summary>
        [Fact]
        public void AmountTest()
        {
            // TODO unit test for the property 'Amount'
        }
    }
}
