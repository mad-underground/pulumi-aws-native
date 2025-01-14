// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WaFv2.Outputs
{

    /// <summary>
    /// Default Action WebACL will take against ingress traffic when there is no matching Rule.
    /// </summary>
    [OutputType]
    public sealed class WebAclDefaultAction
    {
        public readonly Outputs.WebAclAllowAction? Allow;
        public readonly Outputs.WebAclBlockAction? Block;

        [OutputConstructor]
        private WebAclDefaultAction(
            Outputs.WebAclAllowAction? allow,

            Outputs.WebAclBlockAction? block)
        {
            Allow = allow;
            Block = block;
        }
    }
}
