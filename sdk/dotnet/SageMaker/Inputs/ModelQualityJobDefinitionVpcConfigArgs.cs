// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.SageMaker.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelqualityjobdefinition-vpcconfig.html
    /// </summary>
    public sealed class ModelQualityJobDefinitionVpcConfigArgs : Pulumi.ResourceArgs
    {
        [Input("SecurityGroupIds", required: true)]
        private InputList<string>? _SecurityGroupIds;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelqualityjobdefinition-vpcconfig.html#cfn-sagemaker-modelqualityjobdefinition-vpcconfig-securitygroupids
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _SecurityGroupIds ?? (_SecurityGroupIds = new InputList<string>());
            set => _SecurityGroupIds = value;
        }

        [Input("Subnets", required: true)]
        private InputList<string>? _Subnets;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelqualityjobdefinition-vpcconfig.html#cfn-sagemaker-modelqualityjobdefinition-vpcconfig-subnets
        /// </summary>
        public InputList<string> Subnets
        {
            get => _Subnets ?? (_Subnets = new InputList<string>());
            set => _Subnets = value;
        }

        public ModelQualityJobDefinitionVpcConfigArgs()
        {
        }
    }
}
