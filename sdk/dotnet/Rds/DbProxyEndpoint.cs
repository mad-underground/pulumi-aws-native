// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Rds
{
    /// <summary>
    /// Resource schema for AWS::RDS::DBProxyEndpoint.
    /// </summary>
    [AwsNativeResourceType("aws-native:rds:DbProxyEndpoint")]
    public partial class DbProxyEndpoint : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) for the DB proxy endpoint.
        /// </summary>
        [Output("dbProxyEndpointArn")]
        public Output<string> DbProxyEndpointArn { get; private set; } = null!;

        /// <summary>
        /// The identifier for the DB proxy endpoint. This name must be unique for all DB proxy endpoints owned by your AWS account in the specified AWS Region.
        /// </summary>
        [Output("dbProxyEndpointName")]
        public Output<string> DbProxyEndpointName { get; private set; } = null!;

        /// <summary>
        /// The identifier for the proxy. This name must be unique for all proxies owned by your AWS account in the specified AWS Region.
        /// </summary>
        [Output("dbProxyName")]
        public Output<string> DbProxyName { get; private set; } = null!;

        /// <summary>
        /// The endpoint that you can use to connect to the DB proxy. You include the endpoint value in the connection string for a database client application.
        /// </summary>
        [Output("endpoint")]
        public Output<string> Endpoint { get; private set; } = null!;

        /// <summary>
        /// A value that indicates whether this endpoint is the default endpoint for the associated DB proxy. Default DB proxy endpoints always have read/write capability. Other endpoints that you associate with the DB proxy can be either read/write or read-only.
        /// </summary>
        [Output("isDefault")]
        public Output<bool> IsDefault { get; private set; } = null!;

        /// <summary>
        /// An optional set of key-value pairs to associate arbitrary data of your choosing with the DB proxy endpoint.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.DbProxyEndpointTagFormat>> Tags { get; private set; } = null!;

        /// <summary>
        /// A value that indicates whether the DB proxy endpoint can be used for read/write or read-only operations.
        /// </summary>
        [Output("targetRole")]
        public Output<Pulumi.AwsNative.Rds.DbProxyEndpointTargetRole?> TargetRole { get; private set; } = null!;

        /// <summary>
        /// VPC ID to associate with the new DB proxy endpoint.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;

        /// <summary>
        /// VPC security group IDs to associate with the new DB proxy endpoint.
        /// </summary>
        [Output("vpcSecurityGroupIds")]
        public Output<ImmutableArray<string>> VpcSecurityGroupIds { get; private set; } = null!;

        /// <summary>
        /// VPC subnet IDs to associate with the new DB proxy endpoint.
        /// </summary>
        [Output("vpcSubnetIds")]
        public Output<ImmutableArray<string>> VpcSubnetIds { get; private set; } = null!;


        /// <summary>
        /// Create a DbProxyEndpoint resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DbProxyEndpoint(string name, DbProxyEndpointArgs args, CustomResourceOptions? options = null)
            : base("aws-native:rds:DbProxyEndpoint", name, args ?? new DbProxyEndpointArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DbProxyEndpoint(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:rds:DbProxyEndpoint", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "dbProxyEndpointName",
                    "dbProxyName",
                    "vpcSubnetIds[*]",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DbProxyEndpoint resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DbProxyEndpoint Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new DbProxyEndpoint(name, id, options);
        }
    }

    public sealed class DbProxyEndpointArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The identifier for the DB proxy endpoint. This name must be unique for all DB proxy endpoints owned by your AWS account in the specified AWS Region.
        /// </summary>
        [Input("dbProxyEndpointName")]
        public Input<string>? DbProxyEndpointName { get; set; }

        /// <summary>
        /// The identifier for the proxy. This name must be unique for all proxies owned by your AWS account in the specified AWS Region.
        /// </summary>
        [Input("dbProxyName", required: true)]
        public Input<string> DbProxyName { get; set; } = null!;

        [Input("tags")]
        private InputList<Inputs.DbProxyEndpointTagFormatArgs>? _tags;

        /// <summary>
        /// An optional set of key-value pairs to associate arbitrary data of your choosing with the DB proxy endpoint.
        /// </summary>
        public InputList<Inputs.DbProxyEndpointTagFormatArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.DbProxyEndpointTagFormatArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// A value that indicates whether the DB proxy endpoint can be used for read/write or read-only operations.
        /// </summary>
        [Input("targetRole")]
        public Input<Pulumi.AwsNative.Rds.DbProxyEndpointTargetRole>? TargetRole { get; set; }

        [Input("vpcSecurityGroupIds")]
        private InputList<string>? _vpcSecurityGroupIds;

        /// <summary>
        /// VPC security group IDs to associate with the new DB proxy endpoint.
        /// </summary>
        public InputList<string> VpcSecurityGroupIds
        {
            get => _vpcSecurityGroupIds ?? (_vpcSecurityGroupIds = new InputList<string>());
            set => _vpcSecurityGroupIds = value;
        }

        [Input("vpcSubnetIds", required: true)]
        private InputList<string>? _vpcSubnetIds;

        /// <summary>
        /// VPC subnet IDs to associate with the new DB proxy endpoint.
        /// </summary>
        public InputList<string> VpcSubnetIds
        {
            get => _vpcSubnetIds ?? (_vpcSubnetIds = new InputList<string>());
            set => _vpcSubnetIds = value;
        }

        public DbProxyEndpointArgs()
        {
        }
        public static new DbProxyEndpointArgs Empty => new DbProxyEndpointArgs();
    }
}