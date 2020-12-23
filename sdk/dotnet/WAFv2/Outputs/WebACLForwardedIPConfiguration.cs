// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.WAFv2.Outputs
{

    [OutputType]
    public sealed class WebACLForwardedIPConfiguration
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-forwardedipconfiguration.html#cfn-wafv2-webacl-forwardedipconfiguration-fallbackbehavior
        /// </summary>
        public readonly string FallbackBehavior;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-forwardedipconfiguration.html#cfn-wafv2-webacl-forwardedipconfiguration-headername
        /// </summary>
        public readonly string HeaderName;

        [OutputConstructor]
        private WebACLForwardedIPConfiguration(
            string FallbackBehavior,

            string HeaderName)
        {
            this.FallbackBehavior = FallbackBehavior;
            this.HeaderName = HeaderName;
        }
    }
}
