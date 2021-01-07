// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.SSM.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html
    /// </summary>
    public sealed class AssociationPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-applyonlyatcroninterval
        /// </summary>
        [Input("ApplyOnlyAtCronInterval")]
        public Input<bool>? ApplyOnlyAtCronInterval { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-associationname
        /// </summary>
        [Input("AssociationName")]
        public Input<string>? AssociationName { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-automationtargetparametername
        /// </summary>
        [Input("AutomationTargetParameterName")]
        public Input<string>? AutomationTargetParameterName { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-complianceseverity
        /// </summary>
        [Input("ComplianceSeverity")]
        public Input<string>? ComplianceSeverity { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-documentversion
        /// </summary>
        [Input("DocumentVersion")]
        public Input<string>? DocumentVersion { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-instanceid
        /// </summary>
        [Input("InstanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-maxconcurrency
        /// </summary>
        [Input("MaxConcurrency")]
        public Input<string>? MaxConcurrency { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-maxerrors
        /// </summary>
        [Input("MaxErrors")]
        public Input<string>? MaxErrors { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-name
        /// </summary>
        [Input("Name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-outputlocation
        /// </summary>
        [Input("OutputLocation")]
        public Input<Inputs.AssociationInstanceAssociationOutputLocationArgs>? OutputLocation { get; set; }

        [Input("Parameters")]
        private InputMap<Inputs.AssociationParameterValuesArgs>? _Parameters;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-parameters
        /// </summary>
        public InputMap<Inputs.AssociationParameterValuesArgs> Parameters
        {
            get => _Parameters ?? (_Parameters = new InputMap<Inputs.AssociationParameterValuesArgs>());
            set => _Parameters = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-scheduleexpression
        /// </summary>
        [Input("ScheduleExpression")]
        public Input<string>? ScheduleExpression { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-synccompliance
        /// </summary>
        [Input("SyncCompliance")]
        public Input<string>? SyncCompliance { get; set; }

        [Input("Targets")]
        private InputList<Inputs.AssociationTargetArgs>? _Targets;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-targets
        /// </summary>
        public InputList<Inputs.AssociationTargetArgs> Targets
        {
            get => _Targets ?? (_Targets = new InputList<Inputs.AssociationTargetArgs>());
            set => _Targets = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html#cfn-ssm-association-waitforsuccesstimeoutseconds
        /// </summary>
        [Input("WaitForSuccessTimeoutSeconds")]
        public Input<int>? WaitForSuccessTimeoutSeconds { get; set; }

        public AssociationPropertiesArgs()
        {
        }
    }
}