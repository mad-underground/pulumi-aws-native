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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-orstatementtwo.html
    /// </summary>
    [OutputType]
    public sealed class RuleGroupOrStatementTwo
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-orstatementtwo.html#cfn-wafv2-rulegroup-orstatementtwo-statements
        /// </summary>
        public readonly ImmutableArray<Outputs.RuleGroupStatementThree> Statements;

        [OutputConstructor]
        private RuleGroupOrStatementTwo(ImmutableArray<Outputs.RuleGroupStatementThree> statements)
        {
            Statements = statements;
        }
    }
}
