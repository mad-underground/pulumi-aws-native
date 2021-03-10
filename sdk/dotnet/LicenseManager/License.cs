// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.LicenseManager
{
    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-license.html
    /// </summary>
    [AwsNativeResourceType("aws-native:LicenseManager:License")]
    public partial class License : Pulumi.CustomResource
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-license.html#cfn-licensemanager-license-beneficiary
        /// </summary>
        [Output("Beneficiary")]
        public Output<string?> Beneficiary { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-license.html#cfn-licensemanager-license-clienttoken
        /// </summary>
        [Output("ClientToken")]
        public Output<string?> ClientToken { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-license.html#cfn-licensemanager-license-consumptionconfiguration
        /// </summary>
        [Output("ConsumptionConfiguration")]
        public Output<Outputs.LicenseConsumptionConfiguration> ConsumptionConfiguration { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-license.html#cfn-licensemanager-license-entitlements
        /// </summary>
        [Output("Entitlements")]
        public Output<Outputs.LicenseEntitlementList> Entitlements { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-license.html#cfn-licensemanager-license-filters
        /// </summary>
        [Output("Filters")]
        public Output<Outputs.LicenseFilterList?> Filters { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-license.html#cfn-licensemanager-license-homeregion
        /// </summary>
        [Output("HomeRegion")]
        public Output<string> HomeRegion { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-license.html#cfn-licensemanager-license-issuer
        /// </summary>
        [Output("Issuer")]
        public Output<Outputs.LicenseIssuerData> Issuer { get; private set; } = null!;

        [Output("LicenseArn")]
        public Output<string> LicenseArn { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-license.html#cfn-licensemanager-license-licensearns
        /// </summary>
        [Output("LicenseArns")]
        public Output<Outputs.LicenseArnList?> LicenseArns { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-license.html#cfn-licensemanager-license-licensemetadata
        /// </summary>
        [Output("LicenseMetadata")]
        public Output<Outputs.LicenseMetadataList?> LicenseMetadata { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-license.html#cfn-licensemanager-license-licensename
        /// </summary>
        [Output("LicenseName")]
        public Output<string?> LicenseName { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-license.html#cfn-licensemanager-license-maxresults
        /// </summary>
        [Output("MaxResults")]
        public Output<int?> MaxResults { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-license.html#cfn-licensemanager-license-nexttoken
        /// </summary>
        [Output("NextToken")]
        public Output<string?> NextToken { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-license.html#cfn-licensemanager-license-productname
        /// </summary>
        [Output("ProductName")]
        public Output<string?> ProductName { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-license.html#cfn-licensemanager-license-productsku
        /// </summary>
        [Output("ProductSKU")]
        public Output<string?> ProductSKU { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-license.html#cfn-licensemanager-license-sourceversion
        /// </summary>
        [Output("SourceVersion")]
        public Output<string?> SourceVersion { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-license.html#cfn-licensemanager-license-status
        /// </summary>
        [Output("Status")]
        public Output<string?> Status { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-license.html#cfn-licensemanager-license-tags
        /// </summary>
        [Output("Tags")]
        public Output<Outputs.LicenseTagList?> Tags { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-license.html#cfn-licensemanager-license-validity
        /// </summary>
        [Output("Validity")]
        public Output<Outputs.LicenseValidityDateFormat> Validity { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-license.html#cfn-licensemanager-license-version
        /// </summary>
        [Output("Version")]
        public Output<string?> Version { get; private set; } = null!;


        /// <summary>
        /// Create a License resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public License(string name, LicenseArgs args, CustomResourceOptions? options = null)
            : base("aws-native:LicenseManager:License", name, args ?? new LicenseArgs(), MakeResourceOptions(options, ""))
        {
        }

        private License(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:LicenseManager:License", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing License resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static License Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new License(name, id, options);
        }
    }

    public sealed class LicenseArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-license.html#cfn-licensemanager-license-beneficiary
        /// </summary>
        [Input("Beneficiary")]
        public Input<string>? Beneficiary { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-license.html#cfn-licensemanager-license-clienttoken
        /// </summary>
        [Input("ClientToken")]
        public Input<string>? ClientToken { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-license.html#cfn-licensemanager-license-consumptionconfiguration
        /// </summary>
        [Input("ConsumptionConfiguration", required: true)]
        public Input<Inputs.LicenseConsumptionConfigurationArgs> ConsumptionConfiguration { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-license.html#cfn-licensemanager-license-entitlements
        /// </summary>
        [Input("Entitlements", required: true)]
        public Input<Inputs.LicenseEntitlementListArgs> Entitlements { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-license.html#cfn-licensemanager-license-filters
        /// </summary>
        [Input("Filters")]
        public Input<Inputs.LicenseFilterListArgs>? Filters { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-license.html#cfn-licensemanager-license-homeregion
        /// </summary>
        [Input("HomeRegion", required: true)]
        public Input<string> HomeRegion { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-license.html#cfn-licensemanager-license-issuer
        /// </summary>
        [Input("Issuer", required: true)]
        public Input<Inputs.LicenseIssuerDataArgs> Issuer { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-license.html#cfn-licensemanager-license-licensearns
        /// </summary>
        [Input("LicenseArns")]
        public Input<Inputs.LicenseArnListArgs>? LicenseArns { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-license.html#cfn-licensemanager-license-licensemetadata
        /// </summary>
        [Input("LicenseMetadata")]
        public Input<Inputs.LicenseMetadataListArgs>? LicenseMetadata { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-license.html#cfn-licensemanager-license-licensename
        /// </summary>
        [Input("LicenseName")]
        public Input<string>? LicenseName { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-license.html#cfn-licensemanager-license-maxresults
        /// </summary>
        [Input("MaxResults")]
        public Input<int>? MaxResults { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-license.html#cfn-licensemanager-license-nexttoken
        /// </summary>
        [Input("NextToken")]
        public Input<string>? NextToken { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-license.html#cfn-licensemanager-license-productname
        /// </summary>
        [Input("ProductName")]
        public Input<string>? ProductName { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-license.html#cfn-licensemanager-license-productsku
        /// </summary>
        [Input("ProductSKU")]
        public Input<string>? ProductSKU { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-license.html#cfn-licensemanager-license-sourceversion
        /// </summary>
        [Input("SourceVersion")]
        public Input<string>? SourceVersion { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-license.html#cfn-licensemanager-license-status
        /// </summary>
        [Input("Status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-license.html#cfn-licensemanager-license-tags
        /// </summary>
        [Input("Tags")]
        public Input<Inputs.LicenseTagListArgs>? Tags { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-license.html#cfn-licensemanager-license-validity
        /// </summary>
        [Input("Validity", required: true)]
        public Input<Inputs.LicenseValidityDateFormatArgs> Validity { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-license.html#cfn-licensemanager-license-version
        /// </summary>
        [Input("Version")]
        public Input<string>? Version { get; set; }

        public LicenseArgs()
        {
        }
    }
}
