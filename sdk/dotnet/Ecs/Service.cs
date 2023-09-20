// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ecs
{
    /// <summary>
    /// Resource Type definition for AWS::ECS::Service
    /// </summary>
    [AwsNativeResourceType("aws-native:ecs:Service")]
    public partial class Service : global::Pulumi.CustomResource
    {
        [Output("capacityProviderStrategy")]
        public Output<ImmutableArray<Outputs.ServiceCapacityProviderStrategyItem>> CapacityProviderStrategy { get; private set; } = null!;

        [Output("cluster")]
        public Output<string?> Cluster { get; private set; } = null!;

        [Output("deploymentConfiguration")]
        public Output<Outputs.ServiceDeploymentConfiguration?> DeploymentConfiguration { get; private set; } = null!;

        [Output("deploymentController")]
        public Output<Outputs.ServiceDeploymentController?> DeploymentController { get; private set; } = null!;

        [Output("desiredCount")]
        public Output<int?> DesiredCount { get; private set; } = null!;

        [Output("enableEcsManagedTags")]
        public Output<bool?> EnableEcsManagedTags { get; private set; } = null!;

        [Output("enableExecuteCommand")]
        public Output<bool?> EnableExecuteCommand { get; private set; } = null!;

        [Output("healthCheckGracePeriodSeconds")]
        public Output<int?> HealthCheckGracePeriodSeconds { get; private set; } = null!;

        [Output("launchType")]
        public Output<Pulumi.AwsNative.Ecs.ServiceLaunchType?> LaunchType { get; private set; } = null!;

        [Output("loadBalancers")]
        public Output<ImmutableArray<Outputs.ServiceLoadBalancer>> LoadBalancers { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("networkConfiguration")]
        public Output<Outputs.ServiceNetworkConfiguration?> NetworkConfiguration { get; private set; } = null!;

        [Output("placementConstraints")]
        public Output<ImmutableArray<Outputs.ServicePlacementConstraint>> PlacementConstraints { get; private set; } = null!;

        [Output("placementStrategies")]
        public Output<ImmutableArray<Outputs.ServicePlacementStrategy>> PlacementStrategies { get; private set; } = null!;

        [Output("platformVersion")]
        public Output<string?> PlatformVersion { get; private set; } = null!;

        [Output("propagateTags")]
        public Output<Pulumi.AwsNative.Ecs.ServicePropagateTags?> PropagateTags { get; private set; } = null!;

        [Output("role")]
        public Output<string?> Role { get; private set; } = null!;

        [Output("schedulingStrategy")]
        public Output<Pulumi.AwsNative.Ecs.ServiceSchedulingStrategy?> SchedulingStrategy { get; private set; } = null!;

        [Output("serviceArn")]
        public Output<string> ServiceArn { get; private set; } = null!;

        [Output("serviceConnectConfiguration")]
        public Output<Outputs.ServiceConnectConfiguration?> ServiceConnectConfiguration { get; private set; } = null!;

        [Output("serviceName")]
        public Output<string?> ServiceName { get; private set; } = null!;

        [Output("serviceRegistries")]
        public Output<ImmutableArray<Outputs.ServiceRegistry>> ServiceRegistries { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<Outputs.ServiceTag>> Tags { get; private set; } = null!;

        [Output("taskDefinition")]
        public Output<string?> TaskDefinition { get; private set; } = null!;


        /// <summary>
        /// Create a Service resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Service(string name, ServiceArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:ecs:Service", name, args ?? new ServiceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Service(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:ecs:Service", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "cluster",
                    "deploymentController",
                    "launchType",
                    "role",
                    "schedulingStrategy",
                    "serviceName",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Service resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Service Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Service(name, id, options);
        }
    }

    public sealed class ServiceArgs : global::Pulumi.ResourceArgs
    {
        [Input("capacityProviderStrategy")]
        private InputList<Inputs.ServiceCapacityProviderStrategyItemArgs>? _capacityProviderStrategy;
        public InputList<Inputs.ServiceCapacityProviderStrategyItemArgs> CapacityProviderStrategy
        {
            get => _capacityProviderStrategy ?? (_capacityProviderStrategy = new InputList<Inputs.ServiceCapacityProviderStrategyItemArgs>());
            set => _capacityProviderStrategy = value;
        }

        [Input("cluster")]
        public Input<string>? Cluster { get; set; }

        [Input("deploymentConfiguration")]
        public Input<Inputs.ServiceDeploymentConfigurationArgs>? DeploymentConfiguration { get; set; }

        [Input("deploymentController")]
        public Input<Inputs.ServiceDeploymentControllerArgs>? DeploymentController { get; set; }

        [Input("desiredCount")]
        public Input<int>? DesiredCount { get; set; }

        [Input("enableEcsManagedTags")]
        public Input<bool>? EnableEcsManagedTags { get; set; }

        [Input("enableExecuteCommand")]
        public Input<bool>? EnableExecuteCommand { get; set; }

        [Input("healthCheckGracePeriodSeconds")]
        public Input<int>? HealthCheckGracePeriodSeconds { get; set; }

        [Input("launchType")]
        public Input<Pulumi.AwsNative.Ecs.ServiceLaunchType>? LaunchType { get; set; }

        [Input("loadBalancers")]
        private InputList<Inputs.ServiceLoadBalancerArgs>? _loadBalancers;
        public InputList<Inputs.ServiceLoadBalancerArgs> LoadBalancers
        {
            get => _loadBalancers ?? (_loadBalancers = new InputList<Inputs.ServiceLoadBalancerArgs>());
            set => _loadBalancers = value;
        }

        [Input("networkConfiguration")]
        public Input<Inputs.ServiceNetworkConfigurationArgs>? NetworkConfiguration { get; set; }

        [Input("placementConstraints")]
        private InputList<Inputs.ServicePlacementConstraintArgs>? _placementConstraints;
        public InputList<Inputs.ServicePlacementConstraintArgs> PlacementConstraints
        {
            get => _placementConstraints ?? (_placementConstraints = new InputList<Inputs.ServicePlacementConstraintArgs>());
            set => _placementConstraints = value;
        }

        [Input("placementStrategies")]
        private InputList<Inputs.ServicePlacementStrategyArgs>? _placementStrategies;
        public InputList<Inputs.ServicePlacementStrategyArgs> PlacementStrategies
        {
            get => _placementStrategies ?? (_placementStrategies = new InputList<Inputs.ServicePlacementStrategyArgs>());
            set => _placementStrategies = value;
        }

        [Input("platformVersion")]
        public Input<string>? PlatformVersion { get; set; }

        [Input("propagateTags")]
        public Input<Pulumi.AwsNative.Ecs.ServicePropagateTags>? PropagateTags { get; set; }

        [Input("role")]
        public Input<string>? Role { get; set; }

        [Input("schedulingStrategy")]
        public Input<Pulumi.AwsNative.Ecs.ServiceSchedulingStrategy>? SchedulingStrategy { get; set; }

        [Input("serviceConnectConfiguration")]
        public Input<Inputs.ServiceConnectConfigurationArgs>? ServiceConnectConfiguration { get; set; }

        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        [Input("serviceRegistries")]
        private InputList<Inputs.ServiceRegistryArgs>? _serviceRegistries;
        public InputList<Inputs.ServiceRegistryArgs> ServiceRegistries
        {
            get => _serviceRegistries ?? (_serviceRegistries = new InputList<Inputs.ServiceRegistryArgs>());
            set => _serviceRegistries = value;
        }

        [Input("tags")]
        private InputList<Inputs.ServiceTagArgs>? _tags;
        public InputList<Inputs.ServiceTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.ServiceTagArgs>());
            set => _tags = value;
        }

        [Input("taskDefinition")]
        public Input<string>? TaskDefinition { get; set; }

        public ServiceArgs()
        {
        }
        public static new ServiceArgs Empty => new ServiceArgs();
    }
}