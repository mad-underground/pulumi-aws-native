// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { GetSignalingChannelArgs, GetSignalingChannelResult, GetSignalingChannelOutputArgs } from "./getSignalingChannel";
export const getSignalingChannel: typeof import("./getSignalingChannel").getSignalingChannel = null as any;
export const getSignalingChannelOutput: typeof import("./getSignalingChannel").getSignalingChannelOutput = null as any;

export { GetStreamArgs, GetStreamResult, GetStreamOutputArgs } from "./getStream";
export const getStream: typeof import("./getStream").getStream = null as any;
export const getStreamOutput: typeof import("./getStream").getStreamOutput = null as any;

export { SignalingChannelArgs } from "./signalingChannel";
export type SignalingChannel = import("./signalingChannel").SignalingChannel;
export const SignalingChannel: typeof import("./signalingChannel").SignalingChannel = null as any;

export { StreamArgs } from "./stream";
export type Stream = import("./stream").Stream;
export const Stream: typeof import("./stream").Stream = null as any;

utilities.lazyLoad(exports, ["getSignalingChannel","getSignalingChannelOutput"], () => require("./getSignalingChannel"));
utilities.lazyLoad(exports, ["getStream","getStreamOutput"], () => require("./getStream"));
utilities.lazyLoad(exports, ["SignalingChannel"], () => require("./signalingChannel"));
utilities.lazyLoad(exports, ["Stream"], () => require("./stream"));

// Export enums:
export * from "../types/enums/kinesisvideo";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:kinesisvideo:SignalingChannel":
                return new SignalingChannel(name, <any>undefined, { urn })
            case "aws-native:kinesisvideo:Stream":
                return new Stream(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "kinesisvideo", _module)
