// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { AssetArgs } from "./asset";
export type Asset = import("./asset").Asset;
export const Asset: typeof import("./asset").Asset = null as any;
utilities.lazyLoad(exports, ["Asset"], () => require("./asset"));

export { ChannelArgs } from "./channel";
export type Channel = import("./channel").Channel;
export const Channel: typeof import("./channel").Channel = null as any;
utilities.lazyLoad(exports, ["Channel"], () => require("./channel"));

export { GetAssetArgs, GetAssetResult, GetAssetOutputArgs } from "./getAsset";
export const getAsset: typeof import("./getAsset").getAsset = null as any;
export const getAssetOutput: typeof import("./getAsset").getAssetOutput = null as any;
utilities.lazyLoad(exports, ["getAsset","getAssetOutput"], () => require("./getAsset"));

export { GetChannelArgs, GetChannelResult, GetChannelOutputArgs } from "./getChannel";
export const getChannel: typeof import("./getChannel").getChannel = null as any;
export const getChannelOutput: typeof import("./getChannel").getChannelOutput = null as any;
utilities.lazyLoad(exports, ["getChannel","getChannelOutput"], () => require("./getChannel"));

export { GetOriginEndpointArgs, GetOriginEndpointResult, GetOriginEndpointOutputArgs } from "./getOriginEndpoint";
export const getOriginEndpoint: typeof import("./getOriginEndpoint").getOriginEndpoint = null as any;
export const getOriginEndpointOutput: typeof import("./getOriginEndpoint").getOriginEndpointOutput = null as any;
utilities.lazyLoad(exports, ["getOriginEndpoint","getOriginEndpointOutput"], () => require("./getOriginEndpoint"));

export { GetPackagingConfigurationArgs, GetPackagingConfigurationResult, GetPackagingConfigurationOutputArgs } from "./getPackagingConfiguration";
export const getPackagingConfiguration: typeof import("./getPackagingConfiguration").getPackagingConfiguration = null as any;
export const getPackagingConfigurationOutput: typeof import("./getPackagingConfiguration").getPackagingConfigurationOutput = null as any;
utilities.lazyLoad(exports, ["getPackagingConfiguration","getPackagingConfigurationOutput"], () => require("./getPackagingConfiguration"));

export { GetPackagingGroupArgs, GetPackagingGroupResult, GetPackagingGroupOutputArgs } from "./getPackagingGroup";
export const getPackagingGroup: typeof import("./getPackagingGroup").getPackagingGroup = null as any;
export const getPackagingGroupOutput: typeof import("./getPackagingGroup").getPackagingGroupOutput = null as any;
utilities.lazyLoad(exports, ["getPackagingGroup","getPackagingGroupOutput"], () => require("./getPackagingGroup"));

export { OriginEndpointArgs } from "./originEndpoint";
export type OriginEndpoint = import("./originEndpoint").OriginEndpoint;
export const OriginEndpoint: typeof import("./originEndpoint").OriginEndpoint = null as any;
utilities.lazyLoad(exports, ["OriginEndpoint"], () => require("./originEndpoint"));

export { PackagingConfigurationArgs } from "./packagingConfiguration";
export type PackagingConfiguration = import("./packagingConfiguration").PackagingConfiguration;
export const PackagingConfiguration: typeof import("./packagingConfiguration").PackagingConfiguration = null as any;
utilities.lazyLoad(exports, ["PackagingConfiguration"], () => require("./packagingConfiguration"));

export { PackagingGroupArgs } from "./packagingGroup";
export type PackagingGroup = import("./packagingGroup").PackagingGroup;
export const PackagingGroup: typeof import("./packagingGroup").PackagingGroup = null as any;
utilities.lazyLoad(exports, ["PackagingGroup"], () => require("./packagingGroup"));


// Export enums:
export * from "../types/enums/mediapackage";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:mediapackage:Asset":
                return new Asset(name, <any>undefined, { urn })
            case "aws-native:mediapackage:Channel":
                return new Channel(name, <any>undefined, { urn })
            case "aws-native:mediapackage:OriginEndpoint":
                return new OriginEndpoint(name, <any>undefined, { urn })
            case "aws-native:mediapackage:PackagingConfiguration":
                return new PackagingConfiguration(name, <any>undefined, { urn })
            case "aws-native:mediapackage:PackagingGroup":
                return new PackagingGroup(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "mediapackage", _module)
