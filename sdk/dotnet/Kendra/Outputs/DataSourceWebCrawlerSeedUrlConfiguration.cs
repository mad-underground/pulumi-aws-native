// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Kendra.Outputs
{

    [OutputType]
    public sealed class DataSourceWebCrawlerSeedUrlConfiguration
    {
        public readonly ImmutableArray<string> SeedUrls;
        public readonly Pulumi.AwsNative.Kendra.DataSourceWebCrawlerSeedUrlConfigurationWebCrawlerMode? WebCrawlerMode;

        [OutputConstructor]
        private DataSourceWebCrawlerSeedUrlConfiguration(
            ImmutableArray<string> seedUrls,

            Pulumi.AwsNative.Kendra.DataSourceWebCrawlerSeedUrlConfigurationWebCrawlerMode? webCrawlerMode)
        {
            SeedUrls = seedUrls;
            WebCrawlerMode = webCrawlerMode;
        }
    }
}
