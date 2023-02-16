// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { GetScalingPlanArgs, GetScalingPlanResult, GetScalingPlanOutputArgs } from "./getScalingPlan";
export const getScalingPlan: typeof import("./getScalingPlan").getScalingPlan = null as any;
export const getScalingPlanOutput: typeof import("./getScalingPlan").getScalingPlanOutput = null as any;
utilities.lazyLoad(exports, ["getScalingPlan","getScalingPlanOutput"], () => require("./getScalingPlan"));

export { ScalingPlanArgs } from "./scalingPlan";
export type ScalingPlan = import("./scalingPlan").ScalingPlan;
export const ScalingPlan: typeof import("./scalingPlan").ScalingPlan = null as any;
utilities.lazyLoad(exports, ["ScalingPlan"], () => require("./scalingPlan"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:autoscalingplans:ScalingPlan":
                return new ScalingPlan(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "autoscalingplans", _module)
