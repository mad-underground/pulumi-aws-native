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
    /// Resource Type definition for AWS::Route53::RecordSet
    /// </summary>
    [Obsolete(@"RecordSet is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")]
    [AwsNativeResourceType("aws-native:route53:RecordSet")]
    public partial class RecordSet : global::Pulumi.CustomResource
    {
        [Output("aliasTarget")]
        public Output<Outputs.RecordSetAliasTarget?> AliasTarget { get; private set; } = null!;

        [Output("awsId")]
        public Output<string> AwsId { get; private set; } = null!;

        [Output("cidrRoutingConfig")]
        public Output<Outputs.RecordSetCidrRoutingConfig?> CidrRoutingConfig { get; private set; } = null!;

        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        [Output("failover")]
        public Output<string?> Failover { get; private set; } = null!;

        [Output("geoLocation")]
        public Output<Outputs.RecordSetGeoLocation?> GeoLocation { get; private set; } = null!;

        [Output("geoProximityLocation")]
        public Output<Outputs.RecordSetGeoProximityLocation?> GeoProximityLocation { get; private set; } = null!;

        [Output("healthCheckId")]
        public Output<string?> HealthCheckId { get; private set; } = null!;

        [Output("hostedZoneId")]
        public Output<string?> HostedZoneId { get; private set; } = null!;

        [Output("hostedZoneName")]
        public Output<string?> HostedZoneName { get; private set; } = null!;

        [Output("multiValueAnswer")]
        public Output<bool?> MultiValueAnswer { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("region")]
        public Output<string?> Region { get; private set; } = null!;

        [Output("resourceRecords")]
        public Output<ImmutableArray<string>> ResourceRecords { get; private set; } = null!;

        [Output("setIdentifier")]
        public Output<string?> SetIdentifier { get; private set; } = null!;

        [Output("ttl")]
        public Output<string?> Ttl { get; private set; } = null!;

        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        [Output("weight")]
        public Output<int?> Weight { get; private set; } = null!;


        /// <summary>
        /// Create a RecordSet resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RecordSet(string name, RecordSetArgs args, CustomResourceOptions? options = null)
            : base("aws-native:route53:RecordSet", name, args ?? new RecordSetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RecordSet(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:route53:RecordSet", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "hostedZoneId",
                    "hostedZoneName",
                    "name",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing RecordSet resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RecordSet Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new RecordSet(name, id, options);
        }
    }

    public sealed class RecordSetArgs : global::Pulumi.ResourceArgs
    {
        [Input("aliasTarget")]
        public Input<Inputs.RecordSetAliasTargetArgs>? AliasTarget { get; set; }

        [Input("cidrRoutingConfig")]
        public Input<Inputs.RecordSetCidrRoutingConfigArgs>? CidrRoutingConfig { get; set; }

        [Input("comment")]
        public Input<string>? Comment { get; set; }

        [Input("failover")]
        public Input<string>? Failover { get; set; }

        [Input("geoLocation")]
        public Input<Inputs.RecordSetGeoLocationArgs>? GeoLocation { get; set; }

        [Input("geoProximityLocation")]
        public Input<Inputs.RecordSetGeoProximityLocationArgs>? GeoProximityLocation { get; set; }

        [Input("healthCheckId")]
        public Input<string>? HealthCheckId { get; set; }

        [Input("hostedZoneId")]
        public Input<string>? HostedZoneId { get; set; }

        [Input("hostedZoneName")]
        public Input<string>? HostedZoneName { get; set; }

        [Input("multiValueAnswer")]
        public Input<bool>? MultiValueAnswer { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("resourceRecords")]
        private InputList<string>? _resourceRecords;
        public InputList<string> ResourceRecords
        {
            get => _resourceRecords ?? (_resourceRecords = new InputList<string>());
            set => _resourceRecords = value;
        }

        [Input("setIdentifier")]
        public Input<string>? SetIdentifier { get; set; }

        [Input("ttl")]
        public Input<string>? Ttl { get; set; }

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        [Input("weight")]
        public Input<int>? Weight { get; set; }

        public RecordSetArgs()
        {
        }
        public static new RecordSetArgs Empty => new RecordSetArgs();
    }
}
