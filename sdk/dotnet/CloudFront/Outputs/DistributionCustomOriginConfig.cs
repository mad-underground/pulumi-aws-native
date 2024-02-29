// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudFront.Outputs
{

    /// <summary>
    /// A custom origin. A custom origin is any origin that is *not* an Amazon S3 bucket, with one exception. An Amazon S3 bucket that is [configured with static website hosting](https://docs.aws.amazon.com/AmazonS3/latest/dev/WebsiteHosting.html) *is* a custom origin.
    /// </summary>
    [OutputType]
    public sealed class DistributionCustomOriginConfig
    {
        /// <summary>
        /// The HTTP port that CloudFront uses to connect to the origin. Specify the HTTP port that the origin listens on.
        /// </summary>
        public readonly int? HttpPort;
        /// <summary>
        /// The HTTPS port that CloudFront uses to connect to the origin. Specify the HTTPS port that the origin listens on.
        /// </summary>
        public readonly int? HttpsPort;
        /// <summary>
        /// Specifies how long, in seconds, CloudFront persists its connection to the origin. The minimum timeout is 1 second, the maximum is 60 seconds, and the default (if you don't specify otherwise) is 5 seconds.
        ///  For more information, see [Origin Keep-alive Timeout](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-web-values-specify.html#DownloadDistValuesOriginKeepaliveTimeout) in the *Amazon CloudFront Developer Guide*.
        /// </summary>
        public readonly int? OriginKeepaliveTimeout;
        /// <summary>
        /// Specifies the protocol (HTTP or HTTPS) that CloudFront uses to connect to the origin. Valid values are:
        ///   +   ``http-only`` – CloudFront always uses HTTP to connect to the origin.
        ///   +   ``match-viewer`` – CloudFront connects to the origin using the same protocol that the viewer used to connect to CloudFront.
        ///   +   ``https-only`` – CloudFront always uses HTTPS to connect to the origin.
        /// </summary>
        public readonly string OriginProtocolPolicy;
        /// <summary>
        /// Specifies how long, in seconds, CloudFront waits for a response from the origin. This is also known as the *origin response timeout*. The minimum timeout is 1 second, the maximum is 60 seconds, and the default (if you don't specify otherwise) is 30 seconds.
        ///  For more information, see [Origin Response Timeout](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-web-values-specify.html#DownloadDistValuesOriginResponseTimeout) in the *Amazon CloudFront Developer Guide*.
        /// </summary>
        public readonly int? OriginReadTimeout;
        /// <summary>
        /// Specifies the minimum SSL/TLS protocol that CloudFront uses when connecting to your origin over HTTPS. Valid values include ``SSLv3``, ``TLSv1``, ``TLSv1.1``, and ``TLSv1.2``.
        ///  For more information, see [Minimum Origin SSL Protocol](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-web-values-specify.html#DownloadDistValuesOriginSSLProtocols) in the *Amazon CloudFront Developer Guide*.
        /// </summary>
        public readonly ImmutableArray<string> OriginSslProtocols;

        [OutputConstructor]
        private DistributionCustomOriginConfig(
            int? httpPort,

            int? httpsPort,

            int? originKeepaliveTimeout,

            string originProtocolPolicy,

            int? originReadTimeout,

            ImmutableArray<string> originSslProtocols)
        {
            HttpPort = httpPort;
            HttpsPort = httpsPort;
            OriginKeepaliveTimeout = originKeepaliveTimeout;
            OriginProtocolPolicy = originProtocolPolicy;
            OriginReadTimeout = originReadTimeout;
            OriginSslProtocols = originSslProtocols;
        }
    }
}
