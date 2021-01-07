// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.NetworkFirewall.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-firewall.html
    /// </summary>
    public sealed class FirewallPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-firewall.html#cfn-networkfirewall-firewall-deleteprotection
        /// </summary>
        [Input("DeleteProtection")]
        public Input<bool>? DeleteProtection { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-firewall.html#cfn-networkfirewall-firewall-description
        /// </summary>
        [Input("Description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-firewall.html#cfn-networkfirewall-firewall-firewallname
        /// </summary>
        [Input("FirewallName", required: true)]
        public Input<string> FirewallName { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-firewall.html#cfn-networkfirewall-firewall-firewallpolicyarn
        /// </summary>
        [Input("FirewallPolicyArn", required: true)]
        public Input<string> FirewallPolicyArn { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-firewall.html#cfn-networkfirewall-firewall-firewallpolicychangeprotection
        /// </summary>
        [Input("FirewallPolicyChangeProtection")]
        public Input<bool>? FirewallPolicyChangeProtection { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-firewall.html#cfn-networkfirewall-firewall-subnetchangeprotection
        /// </summary>
        [Input("SubnetChangeProtection")]
        public Input<bool>? SubnetChangeProtection { get; set; }

        [Input("SubnetMappings", required: true)]
        private InputList<Inputs.FirewallSubnetMappingArgs>? _SubnetMappings;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-firewall.html#cfn-networkfirewall-firewall-subnetmappings
        /// </summary>
        public InputList<Inputs.FirewallSubnetMappingArgs> SubnetMappings
        {
            get => _SubnetMappings ?? (_SubnetMappings = new InputList<Inputs.FirewallSubnetMappingArgs>());
            set => _SubnetMappings = value;
        }

        [Input("Tags")]
        private InputList<Pulumi.Cloudformation.Inputs.TagArgs>? _Tags;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-firewall.html#cfn-networkfirewall-firewall-tags
        /// </summary>
        public InputList<Pulumi.Cloudformation.Inputs.TagArgs> Tags
        {
            get => _Tags ?? (_Tags = new InputList<Pulumi.Cloudformation.Inputs.TagArgs>());
            set => _Tags = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-firewall.html#cfn-networkfirewall-firewall-vpcid
        /// </summary>
        [Input("VpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public FirewallPropertiesArgs()
        {
        }
    }
}