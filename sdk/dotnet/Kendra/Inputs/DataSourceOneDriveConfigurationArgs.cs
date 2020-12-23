// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.Kendra.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-onedriveconfiguration.html
    /// </summary>
    public sealed class DataSourceOneDriveConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-onedriveconfiguration.html#cfn-kendra-datasource-onedriveconfiguration-disablelocalgroups
        /// </summary>
        [Input("DisableLocalGroups")]
        public Input<bool>? DisableLocalGroups { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-onedriveconfiguration.html#cfn-kendra-datasource-onedriveconfiguration-exclusionpatterns
        /// </summary>
        [Input("ExclusionPatterns")]
        public Input<Inputs.DataSourceDataSourceInclusionsExclusionsStringsArgs>? ExclusionPatterns { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-onedriveconfiguration.html#cfn-kendra-datasource-onedriveconfiguration-fieldmappings
        /// </summary>
        [Input("FieldMappings")]
        public Input<Inputs.DataSourceDataSourceToIndexFieldMappingListArgs>? FieldMappings { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-onedriveconfiguration.html#cfn-kendra-datasource-onedriveconfiguration-inclusionpatterns
        /// </summary>
        [Input("InclusionPatterns")]
        public Input<Inputs.DataSourceDataSourceInclusionsExclusionsStringsArgs>? InclusionPatterns { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-onedriveconfiguration.html#cfn-kendra-datasource-onedriveconfiguration-onedriveusers
        /// </summary>
        [Input("OneDriveUsers", required: true)]
        public Input<Inputs.DataSourceOneDriveUsersArgs> OneDriveUsers { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-onedriveconfiguration.html#cfn-kendra-datasource-onedriveconfiguration-secretarn
        /// </summary>
        [Input("SecretArn", required: true)]
        public Input<string> SecretArn { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-onedriveconfiguration.html#cfn-kendra-datasource-onedriveconfiguration-tenantdomain
        /// </summary>
        [Input("TenantDomain", required: true)]
        public Input<string> TenantDomain { get; set; } = null!;

        public DataSourceOneDriveConfigurationArgs()
        {
        }
    }
}
