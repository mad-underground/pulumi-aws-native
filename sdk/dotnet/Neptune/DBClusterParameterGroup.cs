// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Neptune
{
    /// <summary>
    /// Resource Type definition for AWS::Neptune::DBClusterParameterGroup
    /// </summary>
    [Obsolete(@"DBClusterParameterGroup is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")]
    [AwsNativeResourceType("aws-native:neptune:DBClusterParameterGroup")]
    public partial class DBClusterParameterGroup : global::Pulumi.CustomResource
    {
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        [Output("family")]
        public Output<string> Family { get; private set; } = null!;

        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        [Output("parameters")]
        public Output<object> Parameters { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<Outputs.DBClusterParameterGroupTag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a DBClusterParameterGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DBClusterParameterGroup(string name, DBClusterParameterGroupArgs args, CustomResourceOptions? options = null)
            : base("aws-native:neptune:DBClusterParameterGroup", name, args ?? new DBClusterParameterGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DBClusterParameterGroup(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:neptune:DBClusterParameterGroup", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing DBClusterParameterGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DBClusterParameterGroup Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new DBClusterParameterGroup(name, id, options);
        }
    }

    public sealed class DBClusterParameterGroupArgs : global::Pulumi.ResourceArgs
    {
        [Input("description", required: true)]
        public Input<string> Description { get; set; } = null!;

        [Input("family", required: true)]
        public Input<string> Family { get; set; } = null!;

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("parameters", required: true)]
        public Input<object> Parameters { get; set; } = null!;

        [Input("tags")]
        private InputList<Inputs.DBClusterParameterGroupTagArgs>? _tags;
        public InputList<Inputs.DBClusterParameterGroupTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.DBClusterParameterGroupTagArgs>());
            set => _tags = value;
        }

        public DBClusterParameterGroupArgs()
        {
        }
        public static new DBClusterParameterGroupArgs Empty => new DBClusterParameterGroupArgs();
    }
}
