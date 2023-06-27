// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IAM
{
    /// <summary>
    /// Resource Type definition for AWS::IAM::ServiceLinkedRole
    /// </summary>
    [AwsNativeResourceType("aws-native:iam:ServiceLinkedRole")]
    public partial class ServiceLinkedRole : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The service principal for the AWS service to which this role is attached.
        /// </summary>
        [Output("aWSServiceName")]
        public Output<string?> AWSServiceName { get; private set; } = null!;

        /// <summary>
        /// A string that you provide, which is combined with the service-provided prefix to form the complete role name.
        /// </summary>
        [Output("customSuffix")]
        public Output<string?> CustomSuffix { get; private set; } = null!;

        /// <summary>
        /// The description of the role.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The name of the role.
        /// </summary>
        [Output("roleName")]
        public Output<string> RoleName { get; private set; } = null!;


        /// <summary>
        /// Create a ServiceLinkedRole resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServiceLinkedRole(string name, ServiceLinkedRoleArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:iam:ServiceLinkedRole", name, args ?? new ServiceLinkedRoleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServiceLinkedRole(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:iam:ServiceLinkedRole", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ServiceLinkedRole resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServiceLinkedRole Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ServiceLinkedRole(name, id, options);
        }
    }

    public sealed class ServiceLinkedRoleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The service principal for the AWS service to which this role is attached.
        /// </summary>
        [Input("aWSServiceName")]
        public Input<string>? AWSServiceName { get; set; }

        /// <summary>
        /// A string that you provide, which is combined with the service-provided prefix to form the complete role name.
        /// </summary>
        [Input("customSuffix")]
        public Input<string>? CustomSuffix { get; set; }

        /// <summary>
        /// The description of the role.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        public ServiceLinkedRoleArgs()
        {
        }
        public static new ServiceLinkedRoleArgs Empty => new ServiceLinkedRoleArgs();
    }
}
