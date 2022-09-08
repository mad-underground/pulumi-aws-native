// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { ContainerArgs } from "./container";
export type Container = import("./container").Container;
export const Container: typeof import("./container").Container = null as any;

export { GetContainerArgs, GetContainerResult, GetContainerOutputArgs } from "./getContainer";
export const getContainer: typeof import("./getContainer").getContainer = null as any;
export const getContainerOutput: typeof import("./getContainer").getContainerOutput = null as any;

utilities.lazyLoad(exports, ["Container"], () => require("./container"));
utilities.lazyLoad(exports, ["getContainer","getContainerOutput"], () => require("./getContainer"));

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:mediastore:Container":
                return new Container(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "mediastore", _module)
