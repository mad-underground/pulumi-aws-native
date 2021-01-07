// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.Configuration.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-organizationconformancepack.html
    /// </summary>
    public sealed class OrganizationConformancePackPropertiesArgs : Pulumi.ResourceArgs
    {
        [Input("ConformancePackInputParameters")]
        private InputList<Inputs.OrganizationConformancePackConformancePackInputParameterArgs>? _ConformancePackInputParameters;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-organizationconformancepack.html#cfn-config-organizationconformancepack-conformancepackinputparameters
        /// </summary>
        public InputList<Inputs.OrganizationConformancePackConformancePackInputParameterArgs> ConformancePackInputParameters
        {
            get => _ConformancePackInputParameters ?? (_ConformancePackInputParameters = new InputList<Inputs.OrganizationConformancePackConformancePackInputParameterArgs>());
            set => _ConformancePackInputParameters = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-organizationconformancepack.html#cfn-config-organizationconformancepack-deliverys3bucket
        /// </summary>
        [Input("DeliveryS3Bucket")]
        public Input<string>? DeliveryS3Bucket { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-organizationconformancepack.html#cfn-config-organizationconformancepack-deliverys3keyprefix
        /// </summary>
        [Input("DeliveryS3KeyPrefix")]
        public Input<string>? DeliveryS3KeyPrefix { get; set; }

        [Input("ExcludedAccounts")]
        private InputList<string>? _ExcludedAccounts;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-organizationconformancepack.html#cfn-config-organizationconformancepack-excludedaccounts
        /// </summary>
        public InputList<string> ExcludedAccounts
        {
            get => _ExcludedAccounts ?? (_ExcludedAccounts = new InputList<string>());
            set => _ExcludedAccounts = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-organizationconformancepack.html#cfn-config-organizationconformancepack-organizationconformancepackname
        /// </summary>
        [Input("OrganizationConformancePackName", required: true)]
        public Input<string> OrganizationConformancePackName { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-organizationconformancepack.html#cfn-config-organizationconformancepack-templatebody
        /// </summary>
        [Input("TemplateBody")]
        public Input<string>? TemplateBody { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-organizationconformancepack.html#cfn-config-organizationconformancepack-templates3uri
        /// </summary>
        [Input("TemplateS3Uri")]
        public Input<string>? TemplateS3Uri { get; set; }

        public OrganizationConformancePackPropertiesArgs()
        {
        }
    }
}