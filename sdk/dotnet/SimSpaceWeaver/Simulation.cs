// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SimSpaceWeaver
{
    /// <summary>
    /// AWS::SimSpaceWeaver::Simulation resource creates an AWS Simulation.
    /// </summary>
    [AwsNativeResourceType("aws-native:simspaceweaver:Simulation")]
    public partial class Simulation : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Json object with all simulation details
        /// </summary>
        [Output("describePayload")]
        public Output<string> DescribePayload { get; private set; } = null!;

        /// <summary>
        /// The name of the simulation.
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        /// <summary>
        /// Role ARN.
        /// </summary>
        [Output("roleArn")]
        public Output<string?> RoleArn { get; private set; } = null!;

        [Output("schemaS3Location")]
        public Output<Outputs.SimulationS3Location?> SchemaS3Location { get; private set; } = null!;


        /// <summary>
        /// Create a Simulation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Simulation(string name, SimulationArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:simspaceweaver:Simulation", name, args ?? new SimulationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Simulation(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:simspaceweaver:Simulation", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Simulation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Simulation Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Simulation(name, id, options);
        }
    }

    public sealed class SimulationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the simulation.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Role ARN.
        /// </summary>
        [Input("roleArn")]
        public Input<string>? RoleArn { get; set; }

        [Input("schemaS3Location")]
        public Input<Inputs.SimulationS3LocationArgs>? SchemaS3Location { get; set; }

        public SimulationArgs()
        {
        }
        public static new SimulationArgs Empty => new SimulationArgs();
    }
}
