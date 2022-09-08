// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { AssessmentArgs } from "./assessment";
export type Assessment = import("./assessment").Assessment;
export const Assessment: typeof import("./assessment").Assessment = null as any;

export { GetAssessmentArgs, GetAssessmentResult, GetAssessmentOutputArgs } from "./getAssessment";
export const getAssessment: typeof import("./getAssessment").getAssessment = null as any;
export const getAssessmentOutput: typeof import("./getAssessment").getAssessmentOutput = null as any;

utilities.lazyLoad(exports, ["Assessment"], () => require("./assessment"));
utilities.lazyLoad(exports, ["getAssessment","getAssessmentOutput"], () => require("./getAssessment"));

// Export enums:
export * from "../types/enums/auditmanager";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:auditmanager:Assessment":
                return new Assessment(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "auditmanager", _module)
