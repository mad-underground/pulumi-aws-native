// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.Greengrass.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinitionversion-secretsmanagersecretresourcedata.html
    /// </summary>
    public sealed class ResourceDefinitionVersionSecretsManagerSecretResourceDataArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinitionversion-secretsmanagersecretresourcedata.html#cfn-greengrass-resourcedefinitionversion-secretsmanagersecretresourcedata-arn
        /// </summary>
        [Input("ARN", required: true)]
        public Input<string> ARN { get; set; } = null!;

        [Input("AdditionalStagingLabelsToDownload")]
        private InputList<string>? _AdditionalStagingLabelsToDownload;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinitionversion-secretsmanagersecretresourcedata.html#cfn-greengrass-resourcedefinitionversion-secretsmanagersecretresourcedata-additionalstaginglabelstodownload
        /// </summary>
        public InputList<string> AdditionalStagingLabelsToDownload
        {
            get => _AdditionalStagingLabelsToDownload ?? (_AdditionalStagingLabelsToDownload = new InputList<string>());
            set => _AdditionalStagingLabelsToDownload = value;
        }

        public ResourceDefinitionVersionSecretsManagerSecretResourceDataArgs()
        {
        }
    }
}