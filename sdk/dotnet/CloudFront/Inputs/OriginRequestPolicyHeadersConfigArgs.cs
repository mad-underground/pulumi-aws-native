// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.CloudFront.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-originrequestpolicy-headersconfig.html
    /// </summary>
    public sealed class OriginRequestPolicyHeadersConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-originrequestpolicy-headersconfig.html#cfn-cloudfront-originrequestpolicy-headersconfig-headerbehavior
        /// </summary>
        [Input("HeaderBehavior", required: true)]
        public Input<string> HeaderBehavior { get; set; } = null!;

        [Input("Headers")]
        private InputList<string>? _Headers;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-originrequestpolicy-headersconfig.html#cfn-cloudfront-originrequestpolicy-headersconfig-headers
        /// </summary>
        public InputList<string> Headers
        {
            get => _Headers ?? (_Headers = new InputList<string>());
            set => _Headers = value;
        }

        public OriginRequestPolicyHeadersConfigArgs()
        {
        }
    }
}
