// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { ExperimentArgs } from "./experiment";
export type Experiment = import("./experiment").Experiment;
export const Experiment: typeof import("./experiment").Experiment = null as any;
utilities.lazyLoad(exports, ["Experiment"], () => require("./experiment"));

export { FeatureArgs } from "./feature";
export type Feature = import("./feature").Feature;
export const Feature: typeof import("./feature").Feature = null as any;
utilities.lazyLoad(exports, ["Feature"], () => require("./feature"));

export { GetExperimentArgs, GetExperimentResult, GetExperimentOutputArgs } from "./getExperiment";
export const getExperiment: typeof import("./getExperiment").getExperiment = null as any;
export const getExperimentOutput: typeof import("./getExperiment").getExperimentOutput = null as any;
utilities.lazyLoad(exports, ["getExperiment","getExperimentOutput"], () => require("./getExperiment"));

export { GetFeatureArgs, GetFeatureResult, GetFeatureOutputArgs } from "./getFeature";
export const getFeature: typeof import("./getFeature").getFeature = null as any;
export const getFeatureOutput: typeof import("./getFeature").getFeatureOutput = null as any;
utilities.lazyLoad(exports, ["getFeature","getFeatureOutput"], () => require("./getFeature"));

export { GetLaunchArgs, GetLaunchResult, GetLaunchOutputArgs } from "./getLaunch";
export const getLaunch: typeof import("./getLaunch").getLaunch = null as any;
export const getLaunchOutput: typeof import("./getLaunch").getLaunchOutput = null as any;
utilities.lazyLoad(exports, ["getLaunch","getLaunchOutput"], () => require("./getLaunch"));

export { GetProjectArgs, GetProjectResult, GetProjectOutputArgs } from "./getProject";
export const getProject: typeof import("./getProject").getProject = null as any;
export const getProjectOutput: typeof import("./getProject").getProjectOutput = null as any;
utilities.lazyLoad(exports, ["getProject","getProjectOutput"], () => require("./getProject"));

export { GetSegmentArgs, GetSegmentResult, GetSegmentOutputArgs } from "./getSegment";
export const getSegment: typeof import("./getSegment").getSegment = null as any;
export const getSegmentOutput: typeof import("./getSegment").getSegmentOutput = null as any;
utilities.lazyLoad(exports, ["getSegment","getSegmentOutput"], () => require("./getSegment"));

export { LaunchArgs } from "./launch";
export type Launch = import("./launch").Launch;
export const Launch: typeof import("./launch").Launch = null as any;
utilities.lazyLoad(exports, ["Launch"], () => require("./launch"));

export { ProjectArgs } from "./project";
export type Project = import("./project").Project;
export const Project: typeof import("./project").Project = null as any;
utilities.lazyLoad(exports, ["Project"], () => require("./project"));

export { SegmentArgs } from "./segment";
export type Segment = import("./segment").Segment;
export const Segment: typeof import("./segment").Segment = null as any;
utilities.lazyLoad(exports, ["Segment"], () => require("./segment"));


// Export enums:
export * from "../types/enums/evidently";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:evidently:Experiment":
                return new Experiment(name, <any>undefined, { urn })
            case "aws-native:evidently:Feature":
                return new Feature(name, <any>undefined, { urn })
            case "aws-native:evidently:Launch":
                return new Launch(name, <any>undefined, { urn })
            case "aws-native:evidently:Project":
                return new Project(name, <any>undefined, { urn })
            case "aws-native:evidently:Segment":
                return new Segment(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "evidently", _module)
