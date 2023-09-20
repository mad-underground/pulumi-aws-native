// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2
{
    /// <summary>
    /// Resource Type definition for AWS::EC2::Instance
    /// </summary>
    [Obsolete(@"Instance is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")]
    [AwsNativeResourceType("aws-native:ec2:Instance")]
    public partial class Instance : global::Pulumi.CustomResource
    {
        [Output("additionalInfo")]
        public Output<string?> AdditionalInfo { get; private set; } = null!;

        [Output("affinity")]
        public Output<string?> Affinity { get; private set; } = null!;

        [Output("availabilityZone")]
        public Output<string?> AvailabilityZone { get; private set; } = null!;

        [Output("blockDeviceMappings")]
        public Output<ImmutableArray<Outputs.InstanceBlockDeviceMapping>> BlockDeviceMappings { get; private set; } = null!;

        [Output("cpuOptions")]
        public Output<Outputs.InstanceCpuOptions?> CpuOptions { get; private set; } = null!;

        [Output("creditSpecification")]
        public Output<Outputs.InstanceCreditSpecification?> CreditSpecification { get; private set; } = null!;

        [Output("disableApiTermination")]
        public Output<bool?> DisableApiTermination { get; private set; } = null!;

        [Output("ebsOptimized")]
        public Output<bool?> EbsOptimized { get; private set; } = null!;

        [Output("elasticGpuSpecifications")]
        public Output<ImmutableArray<Outputs.InstanceElasticGpuSpecification>> ElasticGpuSpecifications { get; private set; } = null!;

        [Output("elasticInferenceAccelerators")]
        public Output<ImmutableArray<Outputs.InstanceElasticInferenceAccelerator>> ElasticInferenceAccelerators { get; private set; } = null!;

        [Output("enclaveOptions")]
        public Output<Outputs.InstanceEnclaveOptions?> EnclaveOptions { get; private set; } = null!;

        [Output("hibernationOptions")]
        public Output<Outputs.InstanceHibernationOptions?> HibernationOptions { get; private set; } = null!;

        [Output("hostId")]
        public Output<string?> HostId { get; private set; } = null!;

        [Output("hostResourceGroupArn")]
        public Output<string?> HostResourceGroupArn { get; private set; } = null!;

        [Output("iamInstanceProfile")]
        public Output<string?> IamInstanceProfile { get; private set; } = null!;

        [Output("imageId")]
        public Output<string?> ImageId { get; private set; } = null!;

        [Output("instanceInitiatedShutdownBehavior")]
        public Output<string?> InstanceInitiatedShutdownBehavior { get; private set; } = null!;

        [Output("instanceType")]
        public Output<string?> InstanceType { get; private set; } = null!;

        [Output("ipv6AddressCount")]
        public Output<int?> Ipv6AddressCount { get; private set; } = null!;

        [Output("ipv6Addresses")]
        public Output<ImmutableArray<Outputs.InstanceIpv6Address>> Ipv6Addresses { get; private set; } = null!;

        [Output("kernelId")]
        public Output<string?> KernelId { get; private set; } = null!;

        [Output("keyName")]
        public Output<string?> KeyName { get; private set; } = null!;

        [Output("launchTemplate")]
        public Output<Outputs.InstanceLaunchTemplateSpecification?> LaunchTemplate { get; private set; } = null!;

        [Output("licenseSpecifications")]
        public Output<ImmutableArray<Outputs.InstanceLicenseSpecification>> LicenseSpecifications { get; private set; } = null!;

        [Output("monitoring")]
        public Output<bool?> Monitoring { get; private set; } = null!;

        [Output("networkInterfaces")]
        public Output<ImmutableArray<Outputs.InstanceNetworkInterface>> NetworkInterfaces { get; private set; } = null!;

        [Output("placementGroupName")]
        public Output<string?> PlacementGroupName { get; private set; } = null!;

        [Output("privateDnsName")]
        public Output<string> PrivateDnsName { get; private set; } = null!;

        [Output("privateDnsNameOptions")]
        public Output<Outputs.InstancePrivateDnsNameOptions?> PrivateDnsNameOptions { get; private set; } = null!;

        [Output("privateIp")]
        public Output<string> PrivateIp { get; private set; } = null!;

        [Output("privateIpAddress")]
        public Output<string?> PrivateIpAddress { get; private set; } = null!;

        [Output("propagateTagsToVolumeOnCreation")]
        public Output<bool?> PropagateTagsToVolumeOnCreation { get; private set; } = null!;

        [Output("publicDnsName")]
        public Output<string> PublicDnsName { get; private set; } = null!;

        [Output("publicIp")]
        public Output<string> PublicIp { get; private set; } = null!;

        [Output("ramdiskId")]
        public Output<string?> RamdiskId { get; private set; } = null!;

        [Output("securityGroupIds")]
        public Output<ImmutableArray<string>> SecurityGroupIds { get; private set; } = null!;

        [Output("securityGroups")]
        public Output<ImmutableArray<string>> SecurityGroups { get; private set; } = null!;

        [Output("sourceDestCheck")]
        public Output<bool?> SourceDestCheck { get; private set; } = null!;

        [Output("ssmAssociations")]
        public Output<ImmutableArray<Outputs.InstanceSsmAssociation>> SsmAssociations { get; private set; } = null!;

        [Output("subnetId")]
        public Output<string?> SubnetId { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<Outputs.InstanceTag>> Tags { get; private set; } = null!;

        [Output("tenancy")]
        public Output<string?> Tenancy { get; private set; } = null!;

        [Output("userData")]
        public Output<string?> UserData { get; private set; } = null!;

        [Output("volumes")]
        public Output<ImmutableArray<Outputs.InstanceVolume>> Volumes { get; private set; } = null!;


        /// <summary>
        /// Create a Instance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Instance(string name, InstanceArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:ec2:Instance", name, args ?? new InstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Instance(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:ec2:Instance", name, null, MakeResourceOptions(options, id))
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
                    "cpuOptions",
                    "elasticGpuSpecifications[*]",
                    "elasticInferenceAccelerators[*]",
                    "enclaveOptions",
                    "hibernationOptions",
                    "hostResourceGroupArn",
                    "imageId",
                    "ipv6AddressCount",
                    "ipv6Addresses[*]",
                    "keyName",
                    "launchTemplate",
                    "licenseSpecifications[*]",
                    "networkInterfaces[*]",
                    "placementGroupName",
                    "privateIpAddress",
                    "securityGroups[*]",
                    "subnetId",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Instance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Instance Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Instance(name, id, options);
        }
    }

    public sealed class InstanceArgs : global::Pulumi.ResourceArgs
    {
        [Input("additionalInfo")]
        public Input<string>? AdditionalInfo { get; set; }

        [Input("affinity")]
        public Input<string>? Affinity { get; set; }

        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        [Input("blockDeviceMappings")]
        private InputList<Inputs.InstanceBlockDeviceMappingArgs>? _blockDeviceMappings;
        public InputList<Inputs.InstanceBlockDeviceMappingArgs> BlockDeviceMappings
        {
            get => _blockDeviceMappings ?? (_blockDeviceMappings = new InputList<Inputs.InstanceBlockDeviceMappingArgs>());
            set => _blockDeviceMappings = value;
        }

        [Input("cpuOptions")]
        public Input<Inputs.InstanceCpuOptionsArgs>? CpuOptions { get; set; }

        [Input("creditSpecification")]
        public Input<Inputs.InstanceCreditSpecificationArgs>? CreditSpecification { get; set; }

        [Input("disableApiTermination")]
        public Input<bool>? DisableApiTermination { get; set; }

        [Input("ebsOptimized")]
        public Input<bool>? EbsOptimized { get; set; }

        [Input("elasticGpuSpecifications")]
        private InputList<Inputs.InstanceElasticGpuSpecificationArgs>? _elasticGpuSpecifications;
        public InputList<Inputs.InstanceElasticGpuSpecificationArgs> ElasticGpuSpecifications
        {
            get => _elasticGpuSpecifications ?? (_elasticGpuSpecifications = new InputList<Inputs.InstanceElasticGpuSpecificationArgs>());
            set => _elasticGpuSpecifications = value;
        }

        [Input("elasticInferenceAccelerators")]
        private InputList<Inputs.InstanceElasticInferenceAcceleratorArgs>? _elasticInferenceAccelerators;
        public InputList<Inputs.InstanceElasticInferenceAcceleratorArgs> ElasticInferenceAccelerators
        {
            get => _elasticInferenceAccelerators ?? (_elasticInferenceAccelerators = new InputList<Inputs.InstanceElasticInferenceAcceleratorArgs>());
            set => _elasticInferenceAccelerators = value;
        }

        [Input("enclaveOptions")]
        public Input<Inputs.InstanceEnclaveOptionsArgs>? EnclaveOptions { get; set; }

        [Input("hibernationOptions")]
        public Input<Inputs.InstanceHibernationOptionsArgs>? HibernationOptions { get; set; }

        [Input("hostId")]
        public Input<string>? HostId { get; set; }

        [Input("hostResourceGroupArn")]
        public Input<string>? HostResourceGroupArn { get; set; }

        [Input("iamInstanceProfile")]
        public Input<string>? IamInstanceProfile { get; set; }

        [Input("imageId")]
        public Input<string>? ImageId { get; set; }

        [Input("instanceInitiatedShutdownBehavior")]
        public Input<string>? InstanceInitiatedShutdownBehavior { get; set; }

        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        [Input("ipv6AddressCount")]
        public Input<int>? Ipv6AddressCount { get; set; }

        [Input("ipv6Addresses")]
        private InputList<Inputs.InstanceIpv6AddressArgs>? _ipv6Addresses;
        public InputList<Inputs.InstanceIpv6AddressArgs> Ipv6Addresses
        {
            get => _ipv6Addresses ?? (_ipv6Addresses = new InputList<Inputs.InstanceIpv6AddressArgs>());
            set => _ipv6Addresses = value;
        }

        [Input("kernelId")]
        public Input<string>? KernelId { get; set; }

        [Input("keyName")]
        public Input<string>? KeyName { get; set; }

        [Input("launchTemplate")]
        public Input<Inputs.InstanceLaunchTemplateSpecificationArgs>? LaunchTemplate { get; set; }

        [Input("licenseSpecifications")]
        private InputList<Inputs.InstanceLicenseSpecificationArgs>? _licenseSpecifications;
        public InputList<Inputs.InstanceLicenseSpecificationArgs> LicenseSpecifications
        {
            get => _licenseSpecifications ?? (_licenseSpecifications = new InputList<Inputs.InstanceLicenseSpecificationArgs>());
            set => _licenseSpecifications = value;
        }

        [Input("monitoring")]
        public Input<bool>? Monitoring { get; set; }

        [Input("networkInterfaces")]
        private InputList<Inputs.InstanceNetworkInterfaceArgs>? _networkInterfaces;
        public InputList<Inputs.InstanceNetworkInterfaceArgs> NetworkInterfaces
        {
            get => _networkInterfaces ?? (_networkInterfaces = new InputList<Inputs.InstanceNetworkInterfaceArgs>());
            set => _networkInterfaces = value;
        }

        [Input("placementGroupName")]
        public Input<string>? PlacementGroupName { get; set; }

        [Input("privateDnsNameOptions")]
        public Input<Inputs.InstancePrivateDnsNameOptionsArgs>? PrivateDnsNameOptions { get; set; }

        [Input("privateIpAddress")]
        public Input<string>? PrivateIpAddress { get; set; }

        [Input("propagateTagsToVolumeOnCreation")]
        public Input<bool>? PropagateTagsToVolumeOnCreation { get; set; }

        [Input("ramdiskId")]
        public Input<string>? RamdiskId { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        [Input("securityGroups")]
        private InputList<string>? _securityGroups;
        public InputList<string> SecurityGroups
        {
            get => _securityGroups ?? (_securityGroups = new InputList<string>());
            set => _securityGroups = value;
        }

        [Input("sourceDestCheck")]
        public Input<bool>? SourceDestCheck { get; set; }

        [Input("ssmAssociations")]
        private InputList<Inputs.InstanceSsmAssociationArgs>? _ssmAssociations;
        public InputList<Inputs.InstanceSsmAssociationArgs> SsmAssociations
        {
            get => _ssmAssociations ?? (_ssmAssociations = new InputList<Inputs.InstanceSsmAssociationArgs>());
            set => _ssmAssociations = value;
        }

        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        [Input("tags")]
        private InputList<Inputs.InstanceTagArgs>? _tags;
        public InputList<Inputs.InstanceTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.InstanceTagArgs>());
            set => _tags = value;
        }

        [Input("tenancy")]
        public Input<string>? Tenancy { get; set; }

        [Input("userData")]
        public Input<string>? UserData { get; set; }

        [Input("volumes")]
        private InputList<Inputs.InstanceVolumeArgs>? _volumes;
        public InputList<Inputs.InstanceVolumeArgs> Volumes
        {
            get => _volumes ?? (_volumes = new InputList<Inputs.InstanceVolumeArgs>());
            set => _volumes = value;
        }

        public InstanceArgs()
        {
        }
        public static new InstanceArgs Empty => new InstanceArgs();
    }
}