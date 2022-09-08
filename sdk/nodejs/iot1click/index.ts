// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { DeviceArgs } from "./device";
export type Device = import("./device").Device;
export const Device: typeof import("./device").Device = null as any;

export { GetDeviceArgs, GetDeviceResult, GetDeviceOutputArgs } from "./getDevice";
export const getDevice: typeof import("./getDevice").getDevice = null as any;
export const getDeviceOutput: typeof import("./getDevice").getDeviceOutput = null as any;

export { GetPlacementArgs, GetPlacementResult, GetPlacementOutputArgs } from "./getPlacement";
export const getPlacement: typeof import("./getPlacement").getPlacement = null as any;
export const getPlacementOutput: typeof import("./getPlacement").getPlacementOutput = null as any;

export { GetProjectArgs, GetProjectResult, GetProjectOutputArgs } from "./getProject";
export const getProject: typeof import("./getProject").getProject = null as any;
export const getProjectOutput: typeof import("./getProject").getProjectOutput = null as any;

export { PlacementArgs } from "./placement";
export type Placement = import("./placement").Placement;
export const Placement: typeof import("./placement").Placement = null as any;

export { ProjectArgs } from "./project";
export type Project = import("./project").Project;
export const Project: typeof import("./project").Project = null as any;

utilities.lazyLoad(exports, ["Device"], () => require("./device"));
utilities.lazyLoad(exports, ["getDevice","getDeviceOutput"], () => require("./getDevice"));
utilities.lazyLoad(exports, ["getPlacement","getPlacementOutput"], () => require("./getPlacement"));
utilities.lazyLoad(exports, ["getProject","getProjectOutput"], () => require("./getProject"));
utilities.lazyLoad(exports, ["Placement"], () => require("./placement"));
utilities.lazyLoad(exports, ["Project"], () => require("./project"));

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:iot1click:Device":
                return new Device(name, <any>undefined, { urn })
            case "aws-native:iot1click:Placement":
                return new Placement(name, <any>undefined, { urn })
            case "aws-native:iot1click:Project":
                return new Project(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "iot1click", _module)
