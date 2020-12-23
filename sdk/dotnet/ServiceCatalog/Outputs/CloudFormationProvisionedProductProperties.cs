// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.ServiceCatalog.Outputs
{

    [OutputType]
    public sealed class CloudFormationProvisionedProductProperties
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationprovisionedproduct.html#cfn-servicecatalog-cloudformationprovisionedproduct-acceptlanguage
        /// </summary>
        public readonly string? AcceptLanguage;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationprovisionedproduct.html#cfn-servicecatalog-cloudformationprovisionedproduct-notificationarns
        /// </summary>
        public readonly ImmutableArray<string> NotificationArns;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationprovisionedproduct.html#cfn-servicecatalog-cloudformationprovisionedproduct-pathid
        /// </summary>
        public readonly string? PathId;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationprovisionedproduct.html#cfn-servicecatalog-cloudformationprovisionedproduct-pathname
        /// </summary>
        public readonly string? PathName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationprovisionedproduct.html#cfn-servicecatalog-cloudformationprovisionedproduct-productid
        /// </summary>
        public readonly string? ProductId;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationprovisionedproduct.html#cfn-servicecatalog-cloudformationprovisionedproduct-productname
        /// </summary>
        public readonly string? ProductName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationprovisionedproduct.html#cfn-servicecatalog-cloudformationprovisionedproduct-provisionedproductname
        /// </summary>
        public readonly string? ProvisionedProductName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationprovisionedproduct.html#cfn-servicecatalog-cloudformationprovisionedproduct-provisioningartifactid
        /// </summary>
        public readonly string? ProvisioningArtifactId;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationprovisionedproduct.html#cfn-servicecatalog-cloudformationprovisionedproduct-provisioningartifactname
        /// </summary>
        public readonly string? ProvisioningArtifactName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationprovisionedproduct.html#cfn-servicecatalog-cloudformationprovisionedproduct-provisioningparameters
        /// </summary>
        public readonly ImmutableArray<Outputs.CloudFormationProvisionedProductProvisioningParameter> ProvisioningParameters;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationprovisionedproduct.html#cfn-servicecatalog-cloudformationprovisionedproduct-provisioningpreferences
        /// </summary>
        public readonly Outputs.CloudFormationProvisionedProductProvisioningPreferences? ProvisioningPreferences;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationprovisionedproduct.html#cfn-servicecatalog-cloudformationprovisionedproduct-tags
        /// </summary>
        public readonly ImmutableArray<Pulumi.Cloudformation.Outputs.Tag> Tags;

        [OutputConstructor]
        private CloudFormationProvisionedProductProperties(
            string? AcceptLanguage,

            ImmutableArray<string> NotificationArns,

            string? PathId,

            string? PathName,

            string? ProductId,

            string? ProductName,

            string? ProvisionedProductName,

            string? ProvisioningArtifactId,

            string? ProvisioningArtifactName,

            ImmutableArray<Outputs.CloudFormationProvisionedProductProvisioningParameter> ProvisioningParameters,

            Outputs.CloudFormationProvisionedProductProvisioningPreferences? ProvisioningPreferences,

            ImmutableArray<Pulumi.Cloudformation.Outputs.Tag> Tags)
        {
            this.AcceptLanguage = AcceptLanguage;
            this.NotificationArns = NotificationArns;
            this.PathId = PathId;
            this.PathName = PathName;
            this.ProductId = ProductId;
            this.ProductName = ProductName;
            this.ProvisionedProductName = ProvisionedProductName;
            this.ProvisioningArtifactId = ProvisioningArtifactId;
            this.ProvisioningArtifactName = ProvisioningArtifactName;
            this.ProvisioningParameters = ProvisioningParameters;
            this.ProvisioningPreferences = ProvisioningPreferences;
            this.Tags = Tags;
        }
    }
}
