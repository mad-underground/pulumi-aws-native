// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { GetPipelineArgs, GetPipelineResult, GetPipelineOutputArgs } from "./getPipeline";
export const getPipeline: typeof import("./getPipeline").getPipeline = null as any;
export const getPipelineOutput: typeof import("./getPipeline").getPipelineOutput = null as any;

export { PipelineArgs } from "./pipeline";
export type Pipeline = import("./pipeline").Pipeline;
export const Pipeline: typeof import("./pipeline").Pipeline = null as any;

utilities.lazyLoad(exports, ["getPipeline","getPipelineOutput"], () => require("./getPipeline"));
utilities.lazyLoad(exports, ["Pipeline"], () => require("./pipeline"));

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:datapipeline:Pipeline":
                return new Pipeline(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "datapipeline", _module)
