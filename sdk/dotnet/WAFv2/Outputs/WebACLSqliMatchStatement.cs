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
    public sealed class WebACLSqliMatchStatement
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-sqlimatchstatement.html#cfn-wafv2-webacl-sqlimatchstatement-fieldtomatch
        /// </summary>
        public readonly Outputs.WebACLFieldToMatch FieldToMatch;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-sqlimatchstatement.html#cfn-wafv2-webacl-sqlimatchstatement-texttransformations
        /// </summary>
        public readonly ImmutableArray<Outputs.WebACLTextTransformation> TextTransformations;

        [OutputConstructor]
        private WebACLSqliMatchStatement(
            Outputs.WebACLFieldToMatch FieldToMatch,

            ImmutableArray<Outputs.WebACLTextTransformation> TextTransformations)
        {
            this.FieldToMatch = FieldToMatch;
            this.TextTransformations = TextTransformations;
        }
    }
}