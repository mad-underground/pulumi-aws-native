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
    /// Size Constraint statement.
    /// </summary>
    [OutputType]
    public sealed class WebAclSizeConstraintStatement
    {
        public readonly Pulumi.AwsNative.WaFv2.WebAclSizeConstraintStatementComparisonOperator ComparisonOperator;
        public readonly Outputs.WebAclFieldToMatch FieldToMatch;
        public readonly double Size;
        public readonly ImmutableArray<Outputs.WebAclTextTransformation> TextTransformations;

        [OutputConstructor]
        private WebAclSizeConstraintStatement(
            Pulumi.AwsNative.WaFv2.WebAclSizeConstraintStatementComparisonOperator comparisonOperator,

            Outputs.WebAclFieldToMatch fieldToMatch,

            double size,

            ImmutableArray<Outputs.WebAclTextTransformation> textTransformations)
        {
            ComparisonOperator = comparisonOperator;
            FieldToMatch = fieldToMatch;
            Size = size;
            TextTransformations = textTransformations;
        }
    }
}