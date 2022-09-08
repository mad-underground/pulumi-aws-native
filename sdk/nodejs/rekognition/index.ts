// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { CollectionArgs } from "./collection";
export type Collection = import("./collection").Collection;
export const Collection: typeof import("./collection").Collection = null as any;

export { GetCollectionArgs, GetCollectionResult, GetCollectionOutputArgs } from "./getCollection";
export const getCollection: typeof import("./getCollection").getCollection = null as any;
export const getCollectionOutput: typeof import("./getCollection").getCollectionOutput = null as any;

export { GetProjectArgs, GetProjectResult, GetProjectOutputArgs } from "./getProject";
export const getProject: typeof import("./getProject").getProject = null as any;
export const getProjectOutput: typeof import("./getProject").getProjectOutput = null as any;

export { GetStreamProcessorArgs, GetStreamProcessorResult, GetStreamProcessorOutputArgs } from "./getStreamProcessor";
export const getStreamProcessor: typeof import("./getStreamProcessor").getStreamProcessor = null as any;
export const getStreamProcessorOutput: typeof import("./getStreamProcessor").getStreamProcessorOutput = null as any;

export { ProjectArgs } from "./project";
export type Project = import("./project").Project;
export const Project: typeof import("./project").Project = null as any;

export { StreamProcessorArgs } from "./streamProcessor";
export type StreamProcessor = import("./streamProcessor").StreamProcessor;
export const StreamProcessor: typeof import("./streamProcessor").StreamProcessor = null as any;

utilities.lazyLoad(exports, ["Collection"], () => require("./collection"));
utilities.lazyLoad(exports, ["getCollection","getCollectionOutput"], () => require("./getCollection"));
utilities.lazyLoad(exports, ["getProject","getProjectOutput"], () => require("./getProject"));
utilities.lazyLoad(exports, ["getStreamProcessor","getStreamProcessorOutput"], () => require("./getStreamProcessor"));
utilities.lazyLoad(exports, ["Project"], () => require("./project"));
utilities.lazyLoad(exports, ["StreamProcessor"], () => require("./streamProcessor"));

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:rekognition:Collection":
                return new Collection(name, <any>undefined, { urn })
            case "aws-native:rekognition:Project":
                return new Project(name, <any>undefined, { urn })
            case "aws-native:rekognition:StreamProcessor":
                return new StreamProcessor(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "rekognition", _module)
