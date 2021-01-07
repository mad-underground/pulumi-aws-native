// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.CloudFormation.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-stackset.html
    /// </summary>
    public sealed class StackSetPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-stackset.html#cfn-cloudformation-stackset-administrationrolearn
        /// </summary>
        [Input("AdministrationRoleARN")]
        public Input<string>? AdministrationRoleARN { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-stackset.html#cfn-cloudformation-stackset-autodeployment
        /// </summary>
        [Input("AutoDeployment")]
        public Input<Inputs.StackSetAutoDeploymentArgs>? AutoDeployment { get; set; }

        [Input("Capabilities")]
        private InputList<string>? _Capabilities;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-stackset.html#cfn-cloudformation-stackset-capabilities
        /// </summary>
        public InputList<string> Capabilities
        {
            get => _Capabilities ?? (_Capabilities = new InputList<string>());
            set => _Capabilities = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-stackset.html#cfn-cloudformation-stackset-description
        /// </summary>
        [Input("Description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-stackset.html#cfn-cloudformation-stackset-executionrolename
        /// </summary>
        [Input("ExecutionRoleName")]
        public Input<string>? ExecutionRoleName { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-stackset.html#cfn-cloudformation-stackset-operationpreferences
        /// </summary>
        [Input("OperationPreferences")]
        public Input<Inputs.StackSetOperationPreferencesArgs>? OperationPreferences { get; set; }

        [Input("Parameters")]
        private InputList<Inputs.StackSetParameterArgs>? _Parameters;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-stackset.html#cfn-cloudformation-stackset-parameters
        /// </summary>
        public InputList<Inputs.StackSetParameterArgs> Parameters
        {
            get => _Parameters ?? (_Parameters = new InputList<Inputs.StackSetParameterArgs>());
            set => _Parameters = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-stackset.html#cfn-cloudformation-stackset-permissionmodel
        /// </summary>
        [Input("PermissionModel", required: true)]
        public Input<string> PermissionModel { get; set; } = null!;

        [Input("StackInstancesGroup")]
        private InputList<Inputs.StackSetStackInstancesArgs>? _StackInstancesGroup;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-stackset.html#cfn-cloudformation-stackset-stackinstancesgroup
        /// </summary>
        public InputList<Inputs.StackSetStackInstancesArgs> StackInstancesGroup
        {
            get => _StackInstancesGroup ?? (_StackInstancesGroup = new InputList<Inputs.StackSetStackInstancesArgs>());
            set => _StackInstancesGroup = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-stackset.html#cfn-cloudformation-stackset-stacksetname
        /// </summary>
        [Input("StackSetName", required: true)]
        public Input<string> StackSetName { get; set; } = null!;

        [Input("Tags")]
        private InputList<Pulumi.Cloudformation.Inputs.TagArgs>? _Tags;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-stackset.html#cfn-cloudformation-stackset-tags
        /// </summary>
        public InputList<Pulumi.Cloudformation.Inputs.TagArgs> Tags
        {
            get => _Tags ?? (_Tags = new InputList<Pulumi.Cloudformation.Inputs.TagArgs>());
            set => _Tags = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-stackset.html#cfn-cloudformation-stackset-templatebody
        /// </summary>
        [Input("TemplateBody")]
        public Input<string>? TemplateBody { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-stackset.html#cfn-cloudformation-stackset-templateurl
        /// </summary>
        [Input("TemplateURL")]
        public Input<string>? TemplateURL { get; set; }

        public StackSetPropertiesArgs()
        {
        }
    }
}