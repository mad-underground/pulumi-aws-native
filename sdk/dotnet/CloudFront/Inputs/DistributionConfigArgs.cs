// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudFront.Inputs
{

    /// <summary>
    /// A distribution configuration.
    /// </summary>
    public sealed class DistributionConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("aliases")]
        private InputList<string>? _aliases;

        /// <summary>
        /// A complex type that contains information about CNAMEs (alternate domain names), if any, for this distribution.
        /// </summary>
        public InputList<string> Aliases
        {
            get => _aliases ?? (_aliases = new InputList<string>());
            set => _aliases = value;
        }

        [Input("cacheBehaviors")]
        private InputList<Inputs.DistributionCacheBehaviorArgs>? _cacheBehaviors;

        /// <summary>
        /// A complex type that contains zero or more ``CacheBehavior`` elements.
        /// </summary>
        public InputList<Inputs.DistributionCacheBehaviorArgs> CacheBehaviors
        {
            get => _cacheBehaviors ?? (_cacheBehaviors = new InputList<Inputs.DistributionCacheBehaviorArgs>());
            set => _cacheBehaviors = value;
        }

        [Input("cnames")]
        private InputList<string>? _cnames;
        public InputList<string> Cnames
        {
            get => _cnames ?? (_cnames = new InputList<string>());
            set => _cnames = value;
        }

        /// <summary>
        /// A comment to describe the distribution. The comment cannot be longer than 128 characters.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// The identifier of a continuous deployment policy. For more information, see ``CreateContinuousDeploymentPolicy``.
        /// </summary>
        [Input("continuousDeploymentPolicyId")]
        public Input<string>? ContinuousDeploymentPolicyId { get; set; }

        [Input("customErrorResponses")]
        private InputList<Inputs.DistributionCustomErrorResponseArgs>? _customErrorResponses;

        /// <summary>
        /// A complex type that controls the following:
        ///   +  Whether CloudFront replaces HTTP status codes in the 4xx and 5xx range with custom error messages before returning the response to the viewer.
        ///   +  How long CloudFront caches HTTP status codes in the 4xx and 5xx range.
        ///   
        ///  For more information about custom error pages, see [Customizing Error Responses](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/custom-error-pages.html) in the *Amazon CloudFront Developer Guide*.
        /// </summary>
        public InputList<Inputs.DistributionCustomErrorResponseArgs> CustomErrorResponses
        {
            get => _customErrorResponses ?? (_customErrorResponses = new InputList<Inputs.DistributionCustomErrorResponseArgs>());
            set => _customErrorResponses = value;
        }

        [Input("customOrigin")]
        public Input<Inputs.DistributionLegacyCustomOriginArgs>? CustomOrigin { get; set; }

        /// <summary>
        /// A complex type that describes the default cache behavior if you don't specify a ``CacheBehavior`` element or if files don't match any of the values of ``PathPattern`` in ``CacheBehavior`` elements. You must create exactly one default cache behavior.
        /// </summary>
        [Input("defaultCacheBehavior", required: true)]
        public Input<Inputs.DistributionDefaultCacheBehaviorArgs> DefaultCacheBehavior { get; set; } = null!;

        /// <summary>
        /// The object that you want CloudFront to request from your origin (for example, ``index.html``) when a viewer requests the root URL for your distribution (``https://www.example.com``) instead of an object in your distribution (``https://www.example.com/product-description.html``). Specifying a default root object avoids exposing the contents of your distribution.
        ///  Specify only the object name, for example, ``index.html``. Don't add a ``/`` before the object name.
        ///  If you don't want to specify a default root object when you create a distribution, include an empty ``DefaultRootObject`` element.
        ///  To delete the default root object from an existing distribution, update the distribution configuration and include an empty ``DefaultRootObject`` element.
        ///  To replace the default root object, update the distribution configuration and specify the new object.
        ///  For more information about the default root object, see [Creating a Default Root Object](https://docs.aws.amazon.com/AmazonCloudFront/latest/D
        /// </summary>
        [Input("defaultRootObject")]
        public Input<string>? DefaultRootObject { get; set; }

        /// <summary>
        /// From this field, you can enable or disable the selected distribution.
        /// </summary>
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        /// <summary>
        /// (Optional) Specify the maximum HTTP version(s) that you want viewers to use to communicate with CF. The default value for new distributions is ``http1.1``.
        ///  For viewers and CF to use HTTP/2, viewers must support TLSv1.2 or later, and must support Server Name Indication (SNI).
        ///  For viewers and CF to use HTTP/3, viewers must support TLSv1.3 and Server Name Indication (SNI). CF supports HTTP/3 connection migration to allow the viewer to switch networks without losing connection. For more information about connection migration, see [Connection Migration](https://docs.aws.amazon.com/https://www.rfc-editor.org/rfc/rfc9000.html#name-connection-migration) at RFC 9000. For more information about supported TLSv1.3 ciphers, see [Supported protocols and ciphers between viewers and CloudFront](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/secure-connections-supported-viewer-protocols-ciphers.html).
        /// </summary>
        [Input("httpVersion")]
        public Input<string>? HttpVersion { get; set; }

        /// <summary>
        /// If you want CloudFront to respond to IPv6 DNS requests with an IPv6 address for your distribution, specify ``true``. If you specify ``false``, CloudFront responds to IPv6 DNS requests with the DNS response code ``NOERROR`` and with no IP addresses. This allows viewers to submit a second request, for an IPv4 address for your distribution.
        ///  In general, you should enable IPv6 if you have users on IPv6 networks who want to access your content. However, if you're using signed URLs or signed cookies to restrict access to your content, and if you're using a custom policy that includes the ``IpAddress`` parameter to restrict the IP addresses that can access your content, don't enable IPv6. If you want to restrict access to some content by IP address and not restrict access to other content (or restrict access but not by IP address), you can create two distributions. For more information, see [Creating a Signed URL Using a Custom Policy](https://docs.aws.amazon.com/AmazonCloudFront/latest/Devel
        /// </summary>
        [Input("ipv6Enabled")]
        public Input<bool>? Ipv6Enabled { get; set; }

        /// <summary>
        /// A complex type that controls whether access logs are written for the distribution.
        ///  For more information about logging, see [Access Logs](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/AccessLogs.html) in the *Amazon CloudFront Developer Guide*.
        /// </summary>
        [Input("logging")]
        public Input<Inputs.DistributionLoggingArgs>? Logging { get; set; }

        /// <summary>
        /// A complex type that contains information about origin groups for this distribution.
        /// </summary>
        [Input("originGroups")]
        public Input<Inputs.DistributionOriginGroupsArgs>? OriginGroups { get; set; }

        [Input("origins")]
        private InputList<Inputs.DistributionOriginArgs>? _origins;

        /// <summary>
        /// A complex type that contains information about origins for this distribution.
        /// </summary>
        public InputList<Inputs.DistributionOriginArgs> Origins
        {
            get => _origins ?? (_origins = new InputList<Inputs.DistributionOriginArgs>());
            set => _origins = value;
        }

        /// <summary>
        /// The price class that corresponds with the maximum price that you want to pay for CloudFront service. If you specify ``PriceClass_All``, CloudFront responds to requests for your objects from all CloudFront edge locations.
        ///  If you specify a price class other than ``PriceClass_All``, CloudFront serves your objects from the CloudFront edge location that has the lowest latency among the edge locations in your price class. Viewers who are in or near regions that are excluded from your specified price class may encounter slower performance.
        ///  For more information about price classes, see [Choosing the Price Class for a CloudFront Distribution](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/PriceClass.html) in the *Amazon CloudFront Developer Guide*. For information about CloudFront pricing, including how price classes (such as Price Class 100) map to CloudFront regions, see [Amazon CloudFront Pricing](https://docs.aws.amazon.com/cloudfront/pricing/).
        /// </summary>
        [Input("priceClass")]
        public Input<string>? PriceClass { get; set; }

        /// <summary>
        /// A complex type that identifies ways in which you want to restrict distribution of your content.
        /// </summary>
        [Input("restrictions")]
        public Input<Inputs.DistributionRestrictionsArgs>? Restrictions { get; set; }

        [Input("s3Origin")]
        public Input<Inputs.DistributionLegacyS3OriginArgs>? S3Origin { get; set; }

        /// <summary>
        /// A Boolean that indicates whether this is a staging distribution. When this value is ``true``, this is a staging distribution. When this value is ``false``, this is not a staging distribution.
        /// </summary>
        [Input("staging")]
        public Input<bool>? Staging { get; set; }

        /// <summary>
        /// A complex type that determines the distribution's SSL/TLS configuration for communicating with viewers.
        /// </summary>
        [Input("viewerCertificate")]
        public Input<Inputs.DistributionViewerCertificateArgs>? ViewerCertificate { get; set; }

        /// <summary>
        /// A unique identifier that specifies the WAF web ACL, if any, to associate with this distribution. To specify a web ACL created using the latest version of WAF, use the ACL ARN, for example ``arn:aws:wafv2:us-east-1:123456789012:global/webacl/ExampleWebACL/473e64fd-f30b-4765-81a0-62ad96dd167a``. To specify a web ACL created using WAF Classic, use the ACL ID, for example ``473e64fd-f30b-4765-81a0-62ad96dd167a``.
        ///   WAF is a web application firewall that lets you monitor the HTTP and HTTPS requests that are forwarded to CloudFront, and lets you control access to your content. Based on conditions that you specify, such as the IP addresses that requests originate from or the values of query strings, CloudFront responds to requests either with the requested content or with an HTTP 403 status code (Forbidden). You can also configure CloudFront to return a custom error page when a request is blocked. For more information about WAF, see the [Developer Guide](https://docs.aws.amazon.com/waf/latest
        /// </summary>
        [Input("webAclId")]
        public Input<string>? WebAclId { get; set; }

        public DistributionConfigArgs()
        {
        }
        public static new DistributionConfigArgs Empty => new DistributionConfigArgs();
    }
}
