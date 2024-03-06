// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DocDb
{
    /// <summary>
    /// Resource Type definition for AWS::DocDB::DBInstance
    /// </summary>
    [Obsolete(@"DbInstance is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")]
    [AwsNativeResourceType("aws-native:docdb:DbInstance")]
    public partial class DbInstance : global::Pulumi.CustomResource
    {
        [Output("autoMinorVersionUpgrade")]
        public Output<bool?> AutoMinorVersionUpgrade { get; private set; } = null!;

        [Output("availabilityZone")]
        public Output<string?> AvailabilityZone { get; private set; } = null!;

        [Output("awsId")]
        public Output<string> AwsId { get; private set; } = null!;

        [Output("caCertificateIdentifier")]
        public Output<string?> CaCertificateIdentifier { get; private set; } = null!;

        [Output("certificateRotationRestart")]
        public Output<bool?> CertificateRotationRestart { get; private set; } = null!;

        [Output("dbClusterIdentifier")]
        public Output<string> DbClusterIdentifier { get; private set; } = null!;

        [Output("dbInstanceClass")]
        public Output<string> DbInstanceClass { get; private set; } = null!;

        [Output("dbInstanceIdentifier")]
        public Output<string?> DbInstanceIdentifier { get; private set; } = null!;

        [Output("enablePerformanceInsights")]
        public Output<bool?> EnablePerformanceInsights { get; private set; } = null!;

        [Output("endpoint")]
        public Output<string> Endpoint { get; private set; } = null!;

        [Output("port")]
        public Output<string> Port { get; private set; } = null!;

        [Output("preferredMaintenanceWindow")]
        public Output<string?> PreferredMaintenanceWindow { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a DbInstance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DbInstance(string name, DbInstanceArgs args, CustomResourceOptions? options = null)
            : base("aws-native:docdb:DbInstance", name, args ?? new DbInstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DbInstance(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:docdb:DbInstance", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "availabilityZone",
                    "dbClusterIdentifier",
                    "dbInstanceIdentifier",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DbInstance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DbInstance Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new DbInstance(name, id, options);
        }
    }

    public sealed class DbInstanceArgs : global::Pulumi.ResourceArgs
    {
        [Input("autoMinorVersionUpgrade")]
        public Input<bool>? AutoMinorVersionUpgrade { get; set; }

        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        [Input("caCertificateIdentifier")]
        public Input<string>? CaCertificateIdentifier { get; set; }

        [Input("certificateRotationRestart")]
        public Input<bool>? CertificateRotationRestart { get; set; }

        [Input("dbClusterIdentifier", required: true)]
        public Input<string> DbClusterIdentifier { get; set; } = null!;

        [Input("dbInstanceClass", required: true)]
        public Input<string> DbInstanceClass { get; set; } = null!;

        [Input("dbInstanceIdentifier")]
        public Input<string>? DbInstanceIdentifier { get; set; }

        [Input("enablePerformanceInsights")]
        public Input<bool>? EnablePerformanceInsights { get; set; }

        [Input("preferredMaintenanceWindow")]
        public Input<string>? PreferredMaintenanceWindow { get; set; }

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        public DbInstanceArgs()
        {
        }
        public static new DbInstanceArgs Empty => new DbInstanceArgs();
    }
}
