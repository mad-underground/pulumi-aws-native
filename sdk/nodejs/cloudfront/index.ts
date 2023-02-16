// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { CachePolicyArgs } from "./cachePolicy";
export type CachePolicy = import("./cachePolicy").CachePolicy;
export const CachePolicy: typeof import("./cachePolicy").CachePolicy = null as any;
utilities.lazyLoad(exports, ["CachePolicy"], () => require("./cachePolicy"));

export { CloudFrontOriginAccessIdentityArgs } from "./cloudFrontOriginAccessIdentity";
export type CloudFrontOriginAccessIdentity = import("./cloudFrontOriginAccessIdentity").CloudFrontOriginAccessIdentity;
export const CloudFrontOriginAccessIdentity: typeof import("./cloudFrontOriginAccessIdentity").CloudFrontOriginAccessIdentity = null as any;
utilities.lazyLoad(exports, ["CloudFrontOriginAccessIdentity"], () => require("./cloudFrontOriginAccessIdentity"));

export { ContinuousDeploymentPolicyArgs } from "./continuousDeploymentPolicy";
export type ContinuousDeploymentPolicy = import("./continuousDeploymentPolicy").ContinuousDeploymentPolicy;
export const ContinuousDeploymentPolicy: typeof import("./continuousDeploymentPolicy").ContinuousDeploymentPolicy = null as any;
utilities.lazyLoad(exports, ["ContinuousDeploymentPolicy"], () => require("./continuousDeploymentPolicy"));

export { DistributionArgs } from "./distribution";
export type Distribution = import("./distribution").Distribution;
export const Distribution: typeof import("./distribution").Distribution = null as any;
utilities.lazyLoad(exports, ["Distribution"], () => require("./distribution"));

export { FunctionArgs } from "./function";
export type Function = import("./function").Function;
export const Function: typeof import("./function").Function = null as any;
utilities.lazyLoad(exports, ["Function"], () => require("./function"));

export { GetCachePolicyArgs, GetCachePolicyResult, GetCachePolicyOutputArgs } from "./getCachePolicy";
export const getCachePolicy: typeof import("./getCachePolicy").getCachePolicy = null as any;
export const getCachePolicyOutput: typeof import("./getCachePolicy").getCachePolicyOutput = null as any;
utilities.lazyLoad(exports, ["getCachePolicy","getCachePolicyOutput"], () => require("./getCachePolicy"));

export { GetCloudFrontOriginAccessIdentityArgs, GetCloudFrontOriginAccessIdentityResult, GetCloudFrontOriginAccessIdentityOutputArgs } from "./getCloudFrontOriginAccessIdentity";
export const getCloudFrontOriginAccessIdentity: typeof import("./getCloudFrontOriginAccessIdentity").getCloudFrontOriginAccessIdentity = null as any;
export const getCloudFrontOriginAccessIdentityOutput: typeof import("./getCloudFrontOriginAccessIdentity").getCloudFrontOriginAccessIdentityOutput = null as any;
utilities.lazyLoad(exports, ["getCloudFrontOriginAccessIdentity","getCloudFrontOriginAccessIdentityOutput"], () => require("./getCloudFrontOriginAccessIdentity"));

export { GetContinuousDeploymentPolicyArgs, GetContinuousDeploymentPolicyResult, GetContinuousDeploymentPolicyOutputArgs } from "./getContinuousDeploymentPolicy";
export const getContinuousDeploymentPolicy: typeof import("./getContinuousDeploymentPolicy").getContinuousDeploymentPolicy = null as any;
export const getContinuousDeploymentPolicyOutput: typeof import("./getContinuousDeploymentPolicy").getContinuousDeploymentPolicyOutput = null as any;
utilities.lazyLoad(exports, ["getContinuousDeploymentPolicy","getContinuousDeploymentPolicyOutput"], () => require("./getContinuousDeploymentPolicy"));

export { GetDistributionArgs, GetDistributionResult, GetDistributionOutputArgs } from "./getDistribution";
export const getDistribution: typeof import("./getDistribution").getDistribution = null as any;
export const getDistributionOutput: typeof import("./getDistribution").getDistributionOutput = null as any;
utilities.lazyLoad(exports, ["getDistribution","getDistributionOutput"], () => require("./getDistribution"));

export { GetFunctionArgs, GetFunctionResult, GetFunctionOutputArgs } from "./getFunction";
export const getFunction: typeof import("./getFunction").getFunction = null as any;
export const getFunctionOutput: typeof import("./getFunction").getFunctionOutput = null as any;
utilities.lazyLoad(exports, ["getFunction","getFunctionOutput"], () => require("./getFunction"));

export { GetKeyGroupArgs, GetKeyGroupResult, GetKeyGroupOutputArgs } from "./getKeyGroup";
export const getKeyGroup: typeof import("./getKeyGroup").getKeyGroup = null as any;
export const getKeyGroupOutput: typeof import("./getKeyGroup").getKeyGroupOutput = null as any;
utilities.lazyLoad(exports, ["getKeyGroup","getKeyGroupOutput"], () => require("./getKeyGroup"));

export { GetMonitoringSubscriptionArgs, GetMonitoringSubscriptionResult, GetMonitoringSubscriptionOutputArgs } from "./getMonitoringSubscription";
export const getMonitoringSubscription: typeof import("./getMonitoringSubscription").getMonitoringSubscription = null as any;
export const getMonitoringSubscriptionOutput: typeof import("./getMonitoringSubscription").getMonitoringSubscriptionOutput = null as any;
utilities.lazyLoad(exports, ["getMonitoringSubscription","getMonitoringSubscriptionOutput"], () => require("./getMonitoringSubscription"));

export { GetOriginAccessControlArgs, GetOriginAccessControlResult, GetOriginAccessControlOutputArgs } from "./getOriginAccessControl";
export const getOriginAccessControl: typeof import("./getOriginAccessControl").getOriginAccessControl = null as any;
export const getOriginAccessControlOutput: typeof import("./getOriginAccessControl").getOriginAccessControlOutput = null as any;
utilities.lazyLoad(exports, ["getOriginAccessControl","getOriginAccessControlOutput"], () => require("./getOriginAccessControl"));

export { GetOriginRequestPolicyArgs, GetOriginRequestPolicyResult, GetOriginRequestPolicyOutputArgs } from "./getOriginRequestPolicy";
export const getOriginRequestPolicy: typeof import("./getOriginRequestPolicy").getOriginRequestPolicy = null as any;
export const getOriginRequestPolicyOutput: typeof import("./getOriginRequestPolicy").getOriginRequestPolicyOutput = null as any;
utilities.lazyLoad(exports, ["getOriginRequestPolicy","getOriginRequestPolicyOutput"], () => require("./getOriginRequestPolicy"));

export { GetPublicKeyArgs, GetPublicKeyResult, GetPublicKeyOutputArgs } from "./getPublicKey";
export const getPublicKey: typeof import("./getPublicKey").getPublicKey = null as any;
export const getPublicKeyOutput: typeof import("./getPublicKey").getPublicKeyOutput = null as any;
utilities.lazyLoad(exports, ["getPublicKey","getPublicKeyOutput"], () => require("./getPublicKey"));

export { GetRealtimeLogConfigArgs, GetRealtimeLogConfigResult, GetRealtimeLogConfigOutputArgs } from "./getRealtimeLogConfig";
export const getRealtimeLogConfig: typeof import("./getRealtimeLogConfig").getRealtimeLogConfig = null as any;
export const getRealtimeLogConfigOutput: typeof import("./getRealtimeLogConfig").getRealtimeLogConfigOutput = null as any;
utilities.lazyLoad(exports, ["getRealtimeLogConfig","getRealtimeLogConfigOutput"], () => require("./getRealtimeLogConfig"));

export { GetResponseHeadersPolicyArgs, GetResponseHeadersPolicyResult, GetResponseHeadersPolicyOutputArgs } from "./getResponseHeadersPolicy";
export const getResponseHeadersPolicy: typeof import("./getResponseHeadersPolicy").getResponseHeadersPolicy = null as any;
export const getResponseHeadersPolicyOutput: typeof import("./getResponseHeadersPolicy").getResponseHeadersPolicyOutput = null as any;
utilities.lazyLoad(exports, ["getResponseHeadersPolicy","getResponseHeadersPolicyOutput"], () => require("./getResponseHeadersPolicy"));

export { GetStreamingDistributionArgs, GetStreamingDistributionResult, GetStreamingDistributionOutputArgs } from "./getStreamingDistribution";
export const getStreamingDistribution: typeof import("./getStreamingDistribution").getStreamingDistribution = null as any;
export const getStreamingDistributionOutput: typeof import("./getStreamingDistribution").getStreamingDistributionOutput = null as any;
utilities.lazyLoad(exports, ["getStreamingDistribution","getStreamingDistributionOutput"], () => require("./getStreamingDistribution"));

export { KeyGroupArgs } from "./keyGroup";
export type KeyGroup = import("./keyGroup").KeyGroup;
export const KeyGroup: typeof import("./keyGroup").KeyGroup = null as any;
utilities.lazyLoad(exports, ["KeyGroup"], () => require("./keyGroup"));

export { MonitoringSubscriptionArgs } from "./monitoringSubscription";
export type MonitoringSubscription = import("./monitoringSubscription").MonitoringSubscription;
export const MonitoringSubscription: typeof import("./monitoringSubscription").MonitoringSubscription = null as any;
utilities.lazyLoad(exports, ["MonitoringSubscription"], () => require("./monitoringSubscription"));

export { OriginAccessControlArgs } from "./originAccessControl";
export type OriginAccessControl = import("./originAccessControl").OriginAccessControl;
export const OriginAccessControl: typeof import("./originAccessControl").OriginAccessControl = null as any;
utilities.lazyLoad(exports, ["OriginAccessControl"], () => require("./originAccessControl"));

export { OriginRequestPolicyArgs } from "./originRequestPolicy";
export type OriginRequestPolicy = import("./originRequestPolicy").OriginRequestPolicy;
export const OriginRequestPolicy: typeof import("./originRequestPolicy").OriginRequestPolicy = null as any;
utilities.lazyLoad(exports, ["OriginRequestPolicy"], () => require("./originRequestPolicy"));

export { PublicKeyArgs } from "./publicKey";
export type PublicKey = import("./publicKey").PublicKey;
export const PublicKey: typeof import("./publicKey").PublicKey = null as any;
utilities.lazyLoad(exports, ["PublicKey"], () => require("./publicKey"));

export { RealtimeLogConfigArgs } from "./realtimeLogConfig";
export type RealtimeLogConfig = import("./realtimeLogConfig").RealtimeLogConfig;
export const RealtimeLogConfig: typeof import("./realtimeLogConfig").RealtimeLogConfig = null as any;
utilities.lazyLoad(exports, ["RealtimeLogConfig"], () => require("./realtimeLogConfig"));

export { ResponseHeadersPolicyArgs } from "./responseHeadersPolicy";
export type ResponseHeadersPolicy = import("./responseHeadersPolicy").ResponseHeadersPolicy;
export const ResponseHeadersPolicy: typeof import("./responseHeadersPolicy").ResponseHeadersPolicy = null as any;
utilities.lazyLoad(exports, ["ResponseHeadersPolicy"], () => require("./responseHeadersPolicy"));

export { StreamingDistributionArgs } from "./streamingDistribution";
export type StreamingDistribution = import("./streamingDistribution").StreamingDistribution;
export const StreamingDistribution: typeof import("./streamingDistribution").StreamingDistribution = null as any;
utilities.lazyLoad(exports, ["StreamingDistribution"], () => require("./streamingDistribution"));


// Export enums:
export * from "../types/enums/cloudfront";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:cloudfront:CachePolicy":
                return new CachePolicy(name, <any>undefined, { urn })
            case "aws-native:cloudfront:CloudFrontOriginAccessIdentity":
                return new CloudFrontOriginAccessIdentity(name, <any>undefined, { urn })
            case "aws-native:cloudfront:ContinuousDeploymentPolicy":
                return new ContinuousDeploymentPolicy(name, <any>undefined, { urn })
            case "aws-native:cloudfront:Distribution":
                return new Distribution(name, <any>undefined, { urn })
            case "aws-native:cloudfront:Function":
                return new Function(name, <any>undefined, { urn })
            case "aws-native:cloudfront:KeyGroup":
                return new KeyGroup(name, <any>undefined, { urn })
            case "aws-native:cloudfront:MonitoringSubscription":
                return new MonitoringSubscription(name, <any>undefined, { urn })
            case "aws-native:cloudfront:OriginAccessControl":
                return new OriginAccessControl(name, <any>undefined, { urn })
            case "aws-native:cloudfront:OriginRequestPolicy":
                return new OriginRequestPolicy(name, <any>undefined, { urn })
            case "aws-native:cloudfront:PublicKey":
                return new PublicKey(name, <any>undefined, { urn })
            case "aws-native:cloudfront:RealtimeLogConfig":
                return new RealtimeLogConfig(name, <any>undefined, { urn })
            case "aws-native:cloudfront:ResponseHeadersPolicy":
                return new ResponseHeadersPolicy(name, <any>undefined, { urn })
            case "aws-native:cloudfront:StreamingDistribution":
                return new StreamingDistribution(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "cloudfront", _module)
