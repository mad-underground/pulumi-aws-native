// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Route53
{
    /// <summary>
    /// Resource schema for AWS::Route53::CidrCollection.
    /// </summary>
    [AwsNativeResourceType("aws-native:route53:CidrCollection")]
    public partial class CidrCollection : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Amazon resource name (ARN) to uniquely identify the AWS resource.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// UUID of the CIDR collection.
        /// </summary>
        [Output("awsId")]
        public Output<string> AwsId { get; private set; } = null!;

        /// <summary>
        /// A complex type that contains information about the list of CIDR locations.
        /// </summary>
        [Output("locations")]
        public Output<ImmutableArray<Outputs.CidrCollectionLocation>> Locations { get; private set; } = null!;

        /// <summary>
        /// A unique name for the CIDR collection.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a CidrCollection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CidrCollection(string name, CidrCollectionArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:route53:CidrCollection", name, args ?? new CidrCollectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CidrCollection(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:route53:CidrCollection", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "name",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing CidrCollection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CidrCollection Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new CidrCollection(name, id, options);
        }
    }

    public sealed class CidrCollectionArgs : global::Pulumi.ResourceArgs
    {
        [Input("locations")]
        private InputList<Inputs.CidrCollectionLocationArgs>? _locations;

        /// <summary>
        /// A complex type that contains information about the list of CIDR locations.
        /// </summary>
        public InputList<Inputs.CidrCollectionLocationArgs> Locations
        {
            get => _locations ?? (_locations = new InputList<Inputs.CidrCollectionLocationArgs>());
            set => _locations = value;
        }

        /// <summary>
        /// A unique name for the CIDR collection.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public CidrCollectionArgs()
        {
        }
        public static new CidrCollectionArgs Empty => new CidrCollectionArgs();
    }
}
