// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./assessmentTarget";
export * from "./assessmentTemplate";
export * from "./resourceGroup";

// Import resources to register:
import { AssessmentTarget } from "./assessmentTarget";
import { AssessmentTemplate } from "./assessmentTemplate";
import { ResourceGroup } from "./resourceGroup";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "cloudformation:Inspector:AssessmentTarget":
                return new AssessmentTarget(name, <any>undefined, { urn })
            case "cloudformation:Inspector:AssessmentTemplate":
                return new AssessmentTemplate(name, <any>undefined, { urn })
            case "cloudformation:Inspector:ResourceGroup":
                return new ResourceGroup(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("cloudformation", "Inspector", _module)
