// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { EnabledControlArgs } from "./enabledControl";
export type EnabledControl = import("./enabledControl").EnabledControl;
export const EnabledControl: typeof import("./enabledControl").EnabledControl = null as any;
utilities.lazyLoad(exports, ["EnabledControl"], () => require("./enabledControl"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:controltower:EnabledControl":
                return new EnabledControl(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "controltower", _module)
