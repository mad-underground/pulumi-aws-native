// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Transfer
{
    /// <summary>
    /// Resource Type definition for AWS::Transfer::Server
    /// </summary>
    [Obsolete(@"Server is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")]
    [AwsNativeResourceType("aws-native:transfer:Server")]
    public partial class Server : global::Pulumi.CustomResource
    {
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        [Output("certificate")]
        public Output<string?> Certificate { get; private set; } = null!;

        [Output("domain")]
        public Output<string?> Domain { get; private set; } = null!;

        [Output("endpointDetails")]
        public Output<Outputs.ServerEndpointDetails?> EndpointDetails { get; private set; } = null!;

        [Output("endpointType")]
        public Output<string?> EndpointType { get; private set; } = null!;

        [Output("identityProviderDetails")]
        public Output<Outputs.ServerIdentityProviderDetails?> IdentityProviderDetails { get; private set; } = null!;

        [Output("identityProviderType")]
        public Output<string?> IdentityProviderType { get; private set; } = null!;

        [Output("loggingRole")]
        public Output<string?> LoggingRole { get; private set; } = null!;

        [Output("postAuthenticationLoginBanner")]
        public Output<string?> PostAuthenticationLoginBanner { get; private set; } = null!;

        [Output("preAuthenticationLoginBanner")]
        public Output<string?> PreAuthenticationLoginBanner { get; private set; } = null!;

        [Output("protocolDetails")]
        public Output<Outputs.ServerProtocolDetails?> ProtocolDetails { get; private set; } = null!;

        [Output("protocols")]
        public Output<ImmutableArray<Outputs.ServerProtocol>> Protocols { get; private set; } = null!;

        [Output("s3StorageOptions")]
        public Output<Outputs.ServerS3StorageOptions?> S3StorageOptions { get; private set; } = null!;

        [Output("securityPolicyName")]
        public Output<string?> SecurityPolicyName { get; private set; } = null!;

        [Output("serverId")]
        public Output<string> ServerId { get; private set; } = null!;

        [Output("structuredLogDestinations")]
        public Output<ImmutableArray<Outputs.ServerStructuredLogDestination>> StructuredLogDestinations { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;

        [Output("workflowDetails")]
        public Output<Outputs.ServerWorkflowDetails?> WorkflowDetails { get; private set; } = null!;


        /// <summary>
        /// Create a Server resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Server(string name, ServerArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:transfer:Server", name, args ?? new ServerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Server(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:transfer:Server", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "domain",
                    "identityProviderType",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Server resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Server Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Server(name, id, options);
        }
    }

    public sealed class ServerArgs : global::Pulumi.ResourceArgs
    {
        [Input("certificate")]
        public Input<string>? Certificate { get; set; }

        [Input("domain")]
        public Input<string>? Domain { get; set; }

        [Input("endpointDetails")]
        public Input<Inputs.ServerEndpointDetailsArgs>? EndpointDetails { get; set; }

        [Input("endpointType")]
        public Input<string>? EndpointType { get; set; }

        [Input("identityProviderDetails")]
        public Input<Inputs.ServerIdentityProviderDetailsArgs>? IdentityProviderDetails { get; set; }

        [Input("identityProviderType")]
        public Input<string>? IdentityProviderType { get; set; }

        [Input("loggingRole")]
        public Input<string>? LoggingRole { get; set; }

        [Input("postAuthenticationLoginBanner")]
        public Input<string>? PostAuthenticationLoginBanner { get; set; }

        [Input("preAuthenticationLoginBanner")]
        public Input<string>? PreAuthenticationLoginBanner { get; set; }

        [Input("protocolDetails")]
        public Input<Inputs.ServerProtocolDetailsArgs>? ProtocolDetails { get; set; }

        [Input("protocols")]
        private InputList<Inputs.ServerProtocolArgs>? _protocols;
        public InputList<Inputs.ServerProtocolArgs> Protocols
        {
            get => _protocols ?? (_protocols = new InputList<Inputs.ServerProtocolArgs>());
            set => _protocols = value;
        }

        [Input("s3StorageOptions")]
        public Input<Inputs.ServerS3StorageOptionsArgs>? S3StorageOptions { get; set; }

        [Input("securityPolicyName")]
        public Input<string>? SecurityPolicyName { get; set; }

        [Input("structuredLogDestinations")]
        private InputList<Inputs.ServerStructuredLogDestinationArgs>? _structuredLogDestinations;
        public InputList<Inputs.ServerStructuredLogDestinationArgs> StructuredLogDestinations
        {
            get => _structuredLogDestinations ?? (_structuredLogDestinations = new InputList<Inputs.ServerStructuredLogDestinationArgs>());
            set => _structuredLogDestinations = value;
        }

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        [Input("workflowDetails")]
        public Input<Inputs.ServerWorkflowDetailsArgs>? WorkflowDetails { get; set; }

        public ServerArgs()
        {
        }
        public static new ServerArgs Empty => new ServerArgs();
    }
}
