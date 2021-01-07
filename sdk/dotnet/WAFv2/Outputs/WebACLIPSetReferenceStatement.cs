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
    public sealed class WebACLIPSetReferenceStatement
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-ipsetreferencestatement.html#cfn-wafv2-webacl-ipsetreferencestatement-arn
        /// </summary>
        public readonly string Arn;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-ipsetreferencestatement.html#cfn-wafv2-webacl-ipsetreferencestatement-ipsetforwardedipconfig
        /// </summary>
        public readonly Outputs.WebACLIPSetForwardedIPConfiguration? IPSetForwardedIPConfig;

        [OutputConstructor]
        private WebACLIPSetReferenceStatement(
            string Arn,

            Outputs.WebACLIPSetForwardedIPConfiguration? IPSetForwardedIPConfig)
        {
            this.Arn = Arn;
            this.IPSetForwardedIPConfig = IPSetForwardedIPConfig;
        }
    }
}