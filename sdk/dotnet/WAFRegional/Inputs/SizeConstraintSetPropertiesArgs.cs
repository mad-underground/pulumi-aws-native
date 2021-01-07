// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.WAFRegional.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-sizeconstraintset.html
    /// </summary>
    public sealed class SizeConstraintSetPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-sizeconstraintset.html#cfn-wafregional-sizeconstraintset-name
        /// </summary>
        [Input("Name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("SizeConstraints")]
        private InputList<Inputs.SizeConstraintSetSizeConstraintArgs>? _SizeConstraints;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-sizeconstraintset.html#cfn-wafregional-sizeconstraintset-sizeconstraints
        /// </summary>
        public InputList<Inputs.SizeConstraintSetSizeConstraintArgs> SizeConstraints
        {
            get => _SizeConstraints ?? (_SizeConstraints = new InputList<Inputs.SizeConstraintSetSizeConstraintArgs>());
            set => _SizeConstraints = value;
        }

        public SizeConstraintSetPropertiesArgs()
        {
        }
    }
}