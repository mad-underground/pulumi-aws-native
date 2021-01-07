// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.ApiGateway.Outputs
{

    [OutputType]
    public sealed class DeploymentDeploymentCanarySettings
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-deployment-deploymentcanarysettings.html#cfn-apigateway-deployment-deploymentcanarysettings-percenttraffic
        /// </summary>
        public readonly double? PercentTraffic;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-deployment-deploymentcanarysettings.html#cfn-apigateway-deployment-deploymentcanarysettings-stagevariableoverrides
        /// </summary>
        public readonly ImmutableDictionary<string, string>? StageVariableOverrides;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-deployment-deploymentcanarysettings.html#cfn-apigateway-deployment-deploymentcanarysettings-usestagecache
        /// </summary>
        public readonly bool? UseStageCache;

        [OutputConstructor]
        private DeploymentDeploymentCanarySettings(
            double? PercentTraffic,

            ImmutableDictionary<string, string>? StageVariableOverrides,

            bool? UseStageCache)
        {
            this.PercentTraffic = PercentTraffic;
            this.StageVariableOverrides = StageVariableOverrides;
            this.UseStageCache = UseStageCache;
        }
    }
}