// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./destination";
export * from "./deviceProfile";
export * from "./serviceProfile";
export * from "./taskDefinition";
export * from "./wirelessDevice";
export * from "./wirelessGateway";

// Import resources to register:
import { Destination } from "./destination";
import { DeviceProfile } from "./deviceProfile";
import { ServiceProfile } from "./serviceProfile";
import { TaskDefinition } from "./taskDefinition";
import { WirelessDevice } from "./wirelessDevice";
import { WirelessGateway } from "./wirelessGateway";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:iotwireless:Destination":
                return new Destination(name, <any>undefined, { urn })
            case "aws-native:iotwireless:DeviceProfile":
                return new DeviceProfile(name, <any>undefined, { urn })
            case "aws-native:iotwireless:ServiceProfile":
                return new ServiceProfile(name, <any>undefined, { urn })
            case "aws-native:iotwireless:TaskDefinition":
                return new TaskDefinition(name, <any>undefined, { urn })
            case "aws-native:iotwireless:WirelessDevice":
                return new WirelessDevice(name, <any>undefined, { urn })
            case "aws-native:iotwireless:WirelessGateway":
                return new WirelessGateway(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "iotwireless", _module)