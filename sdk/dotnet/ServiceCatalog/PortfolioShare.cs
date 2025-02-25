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
    /// Resource Type definition for AWS::ServiceCatalog::PortfolioShare
    /// </summary>
    [Obsolete(@"PortfolioShare is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")]
    [AwsNativeResourceType("aws-native:servicecatalog:PortfolioShare")]
    public partial class PortfolioShare : global::Pulumi.CustomResource
    {
        [Output("acceptLanguage")]
        public Output<string?> AcceptLanguage { get; private set; } = null!;

        [Output("accountId")]
        public Output<string> AccountId { get; private set; } = null!;

        [Output("awsId")]
        public Output<string> AwsId { get; private set; } = null!;

        [Output("portfolioId")]
        public Output<string> PortfolioId { get; private set; } = null!;

        [Output("shareTagOptions")]
        public Output<bool?> ShareTagOptions { get; private set; } = null!;


        /// <summary>
        /// Create a PortfolioShare resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PortfolioShare(string name, PortfolioShareArgs args, CustomResourceOptions? options = null)
            : base("aws-native:servicecatalog:PortfolioShare", name, args ?? new PortfolioShareArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PortfolioShare(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:servicecatalog:PortfolioShare", name, null, MakeResourceOptions(options, id))
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
                    "accountId",
                    "portfolioId",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing PortfolioShare resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PortfolioShare Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new PortfolioShare(name, id, options);
        }
    }

    public sealed class PortfolioShareArgs : global::Pulumi.ResourceArgs
    {
        [Input("acceptLanguage")]
        public Input<string>? AcceptLanguage { get; set; }

        [Input("accountId", required: true)]
        public Input<string> AccountId { get; set; } = null!;

        [Input("portfolioId", required: true)]
        public Input<string> PortfolioId { get; set; } = null!;

        [Input("shareTagOptions")]
        public Input<bool>? ShareTagOptions { get; set; }

        public PortfolioShareArgs()
        {
        }
        public static new PortfolioShareArgs Empty => new PortfolioShareArgs();
    }
}
