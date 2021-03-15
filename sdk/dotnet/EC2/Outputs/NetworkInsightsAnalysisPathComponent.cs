// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EC2.Outputs
{

    [OutputType]
    public sealed class NetworkInsightsAnalysisPathComponent
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-pathcomponent.html#cfn-ec2-networkinsightsanalysis-pathcomponent-aclrule
        /// </summary>
        public readonly Outputs.NetworkInsightsAnalysisAnalysisAclRule? AclRule;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-pathcomponent.html#cfn-ec2-networkinsightsanalysis-pathcomponent-component
        /// </summary>
        public readonly Outputs.NetworkInsightsAnalysisAnalysisComponent? Component;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-pathcomponent.html#cfn-ec2-networkinsightsanalysis-pathcomponent-destinationvpc
        /// </summary>
        public readonly Outputs.NetworkInsightsAnalysisAnalysisComponent? DestinationVpc;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-pathcomponent.html#cfn-ec2-networkinsightsanalysis-pathcomponent-inboundheader
        /// </summary>
        public readonly Outputs.NetworkInsightsAnalysisAnalysisPacketHeader? InboundHeader;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-pathcomponent.html#cfn-ec2-networkinsightsanalysis-pathcomponent-outboundheader
        /// </summary>
        public readonly Outputs.NetworkInsightsAnalysisAnalysisPacketHeader? OutboundHeader;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-pathcomponent.html#cfn-ec2-networkinsightsanalysis-pathcomponent-routetableroute
        /// </summary>
        public readonly Outputs.NetworkInsightsAnalysisAnalysisRouteTableRoute? RouteTableRoute;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-pathcomponent.html#cfn-ec2-networkinsightsanalysis-pathcomponent-securitygrouprule
        /// </summary>
        public readonly Outputs.NetworkInsightsAnalysisAnalysisSecurityGroupRule? SecurityGroupRule;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-pathcomponent.html#cfn-ec2-networkinsightsanalysis-pathcomponent-sequencenumber
        /// </summary>
        public readonly int? SequenceNumber;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-pathcomponent.html#cfn-ec2-networkinsightsanalysis-pathcomponent-sourcevpc
        /// </summary>
        public readonly Outputs.NetworkInsightsAnalysisAnalysisComponent? SourceVpc;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-pathcomponent.html#cfn-ec2-networkinsightsanalysis-pathcomponent-subnet
        /// </summary>
        public readonly Outputs.NetworkInsightsAnalysisAnalysisComponent? Subnet;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-pathcomponent.html#cfn-ec2-networkinsightsanalysis-pathcomponent-vpc
        /// </summary>
        public readonly Outputs.NetworkInsightsAnalysisAnalysisComponent? Vpc;

        [OutputConstructor]
        private NetworkInsightsAnalysisPathComponent(
            Outputs.NetworkInsightsAnalysisAnalysisAclRule? aclRule,

            Outputs.NetworkInsightsAnalysisAnalysisComponent? component,

            Outputs.NetworkInsightsAnalysisAnalysisComponent? destinationVpc,

            Outputs.NetworkInsightsAnalysisAnalysisPacketHeader? inboundHeader,

            Outputs.NetworkInsightsAnalysisAnalysisPacketHeader? outboundHeader,

            Outputs.NetworkInsightsAnalysisAnalysisRouteTableRoute? routeTableRoute,

            Outputs.NetworkInsightsAnalysisAnalysisSecurityGroupRule? securityGroupRule,

            int? sequenceNumber,

            Outputs.NetworkInsightsAnalysisAnalysisComponent? sourceVpc,

            Outputs.NetworkInsightsAnalysisAnalysisComponent? subnet,

            Outputs.NetworkInsightsAnalysisAnalysisComponent? vpc)
        {
            AclRule = aclRule;
            Component = component;
            DestinationVpc = destinationVpc;
            InboundHeader = inboundHeader;
            OutboundHeader = outboundHeader;
            RouteTableRoute = routeTableRoute;
            SecurityGroupRule = securityGroupRule;
            SequenceNumber = sequenceNumber;
            SourceVpc = sourceVpc;
            Subnet = subnet;
            Vpc = vpc;
        }
    }
}
