// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ServiceCatalog
{
    /// <summary>
    /// Resource Type definition for AWS::ServiceCatalog::PortfolioPrincipalAssociation
    /// </summary>
    [Obsolete(@"PortfolioPrincipalAssociation is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")]
    [AwsNativeResourceType("aws-native:servicecatalog:PortfolioPrincipalAssociation")]
    public partial class PortfolioPrincipalAssociation : global::Pulumi.CustomResource
    {
        [Output("acceptLanguage")]
        public Output<string?> AcceptLanguage { get; private set; } = null!;

        [Output("awsId")]
        public Output<string> AwsId { get; private set; } = null!;

        [Output("portfolioId")]
        public Output<string> PortfolioId { get; private set; } = null!;

        [Output("principalArn")]
        public Output<string> PrincipalArn { get; private set; } = null!;

        [Output("principalType")]
        public Output<string> PrincipalType { get; private set; } = null!;


        /// <summary>
        /// Create a PortfolioPrincipalAssociation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PortfolioPrincipalAssociation(string name, PortfolioPrincipalAssociationArgs args, CustomResourceOptions? options = null)
            : base("aws-native:servicecatalog:PortfolioPrincipalAssociation", name, args ?? new PortfolioPrincipalAssociationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PortfolioPrincipalAssociation(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:servicecatalog:PortfolioPrincipalAssociation", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "acceptLanguage",
                    "portfolioId",
                    "principalArn",
                    "principalType",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing PortfolioPrincipalAssociation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PortfolioPrincipalAssociation Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new PortfolioPrincipalAssociation(name, id, options);
        }
    }

    public sealed class PortfolioPrincipalAssociationArgs : global::Pulumi.ResourceArgs
    {
        [Input("acceptLanguage")]
        public Input<string>? AcceptLanguage { get; set; }

        [Input("portfolioId", required: true)]
        public Input<string> PortfolioId { get; set; } = null!;

        [Input("principalArn", required: true)]
        public Input<string> PrincipalArn { get; set; } = null!;

        [Input("principalType", required: true)]
        public Input<string> PrincipalType { get; set; } = null!;

        public PortfolioPrincipalAssociationArgs()
        {
        }
        public static new PortfolioPrincipalAssociationArgs Empty => new PortfolioPrincipalAssociationArgs();
    }
}
