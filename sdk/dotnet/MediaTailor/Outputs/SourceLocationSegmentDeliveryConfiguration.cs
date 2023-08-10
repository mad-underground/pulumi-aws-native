// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaTailor.Outputs
{

    /// <summary>
    /// &lt;p&gt;The segment delivery configuration settings.&lt;/p&gt;
    /// </summary>
    [OutputType]
    public sealed class SourceLocationSegmentDeliveryConfiguration
    {
        /// <summary>
        /// &lt;p&gt;The base URL of the host or path of the segment delivery server that you're using to serve segments. This is typically a content delivery network (CDN). The URL can be absolute or relative. To use an absolute URL include the protocol, such as &lt;code&gt;https://example.com/some/path&lt;/code&gt;. To use a relative URL specify the relative path, such as &lt;code&gt;/some/path*&lt;/code&gt;.&lt;/p&gt;
        /// </summary>
        public readonly string? BaseUrl;
        /// <summary>
        /// &lt;p&gt;A unique identifier used to distinguish between multiple segment delivery configurations in a source location.&lt;/p&gt;
        /// </summary>
        public readonly string? Name;

        [OutputConstructor]
        private SourceLocationSegmentDeliveryConfiguration(
            string? baseUrl,

            string? name)
        {
            BaseUrl = baseUrl;
            Name = name;
        }
    }
}
