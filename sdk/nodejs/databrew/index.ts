// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { DatasetArgs } from "./dataset";
export type Dataset = import("./dataset").Dataset;
export const Dataset: typeof import("./dataset").Dataset = null as any;

export { GetDatasetArgs, GetDatasetResult, GetDatasetOutputArgs } from "./getDataset";
export const getDataset: typeof import("./getDataset").getDataset = null as any;
export const getDatasetOutput: typeof import("./getDataset").getDatasetOutput = null as any;

export { GetJobArgs, GetJobResult, GetJobOutputArgs } from "./getJob";
export const getJob: typeof import("./getJob").getJob = null as any;
export const getJobOutput: typeof import("./getJob").getJobOutput = null as any;

export { GetProjectArgs, GetProjectResult, GetProjectOutputArgs } from "./getProject";
export const getProject: typeof import("./getProject").getProject = null as any;
export const getProjectOutput: typeof import("./getProject").getProjectOutput = null as any;

export { GetRecipeArgs, GetRecipeResult, GetRecipeOutputArgs } from "./getRecipe";
export const getRecipe: typeof import("./getRecipe").getRecipe = null as any;
export const getRecipeOutput: typeof import("./getRecipe").getRecipeOutput = null as any;

export { GetRulesetArgs, GetRulesetResult, GetRulesetOutputArgs } from "./getRuleset";
export const getRuleset: typeof import("./getRuleset").getRuleset = null as any;
export const getRulesetOutput: typeof import("./getRuleset").getRulesetOutput = null as any;

export { GetScheduleArgs, GetScheduleResult, GetScheduleOutputArgs } from "./getSchedule";
export const getSchedule: typeof import("./getSchedule").getSchedule = null as any;
export const getScheduleOutput: typeof import("./getSchedule").getScheduleOutput = null as any;

export { JobArgs } from "./job";
export type Job = import("./job").Job;
export const Job: typeof import("./job").Job = null as any;

export { ProjectArgs } from "./project";
export type Project = import("./project").Project;
export const Project: typeof import("./project").Project = null as any;

export { RecipeArgs } from "./recipe";
export type Recipe = import("./recipe").Recipe;
export const Recipe: typeof import("./recipe").Recipe = null as any;

export { RulesetArgs } from "./ruleset";
export type Ruleset = import("./ruleset").Ruleset;
export const Ruleset: typeof import("./ruleset").Ruleset = null as any;

export { ScheduleArgs } from "./schedule";
export type Schedule = import("./schedule").Schedule;
export const Schedule: typeof import("./schedule").Schedule = null as any;

utilities.lazyLoad(exports, ["Dataset"], () => require("./dataset"));
utilities.lazyLoad(exports, ["getDataset","getDatasetOutput"], () => require("./getDataset"));
utilities.lazyLoad(exports, ["getJob","getJobOutput"], () => require("./getJob"));
utilities.lazyLoad(exports, ["getProject","getProjectOutput"], () => require("./getProject"));
utilities.lazyLoad(exports, ["getRecipe","getRecipeOutput"], () => require("./getRecipe"));
utilities.lazyLoad(exports, ["getRuleset","getRulesetOutput"], () => require("./getRuleset"));
utilities.lazyLoad(exports, ["getSchedule","getScheduleOutput"], () => require("./getSchedule"));
utilities.lazyLoad(exports, ["Job"], () => require("./job"));
utilities.lazyLoad(exports, ["Project"], () => require("./project"));
utilities.lazyLoad(exports, ["Recipe"], () => require("./recipe"));
utilities.lazyLoad(exports, ["Ruleset"], () => require("./ruleset"));
utilities.lazyLoad(exports, ["Schedule"], () => require("./schedule"));

// Export enums:
export * from "../types/enums/databrew";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:databrew:Dataset":
                return new Dataset(name, <any>undefined, { urn })
            case "aws-native:databrew:Job":
                return new Job(name, <any>undefined, { urn })
            case "aws-native:databrew:Project":
                return new Project(name, <any>undefined, { urn })
            case "aws-native:databrew:Recipe":
                return new Recipe(name, <any>undefined, { urn })
            case "aws-native:databrew:Ruleset":
                return new Ruleset(name, <any>undefined, { urn })
            case "aws-native:databrew:Schedule":
                return new Schedule(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "databrew", _module)
