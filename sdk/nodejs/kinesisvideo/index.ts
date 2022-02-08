// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./getSignalingChannel";
export * from "./getStream";
export * from "./signalingChannel";
export * from "./stream";

// Export enums:
export * from "../types/enums/kinesisvideo";

// Import resources to register:
import { SignalingChannel } from "./signalingChannel";
import { Stream } from "./stream";

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
