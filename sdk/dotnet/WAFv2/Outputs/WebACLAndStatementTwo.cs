// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WAFv2.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-andstatementtwo.html
    /// </summary>
    [OutputType]
    public sealed class WebACLAndStatementTwo
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-andstatementtwo.html#cfn-wafv2-webacl-andstatementtwo-statements
        /// </summary>
        public readonly ImmutableArray<Outputs.WebACLStatementThree> Statements;

        [OutputConstructor]
        private WebACLAndStatementTwo(ImmutableArray<Outputs.WebACLStatementThree> statements)
        {
            Statements = statements;
        }
    }
}
