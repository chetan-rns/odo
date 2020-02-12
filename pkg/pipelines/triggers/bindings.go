package triggers

import (
	pipelinev2 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1alpha2"
	triggersv1 "github.com/tektoncd/triggers/pkg/apis/triggers/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GetTriggerBindings returns a slice of trigger bindings
func GetTriggerBindings() []triggersv1.TriggerBinding {
	triggerBindings := []triggersv1.TriggerBinding{}
	triggerBindings = append(
		triggerBindings,
		createDevCDDeployBinding(),
		createDevCIBuildBinding(),
		createStageCDDeployBinding(),
		createStageCIDryRunBinding(),
	)
	return triggerBindings
}

func createDevCDDeployBinding() triggersv1.TriggerBinding {
	return triggersv1.TriggerBinding{
		TypeMeta:   createTypeMeta(),
		ObjectMeta: createObjectMeta("dev-cd-deploy-from-master-binding"),
		Spec: triggersv1.TriggerBindingSpec{
			Params: []pipelinev2.Param{
				createBindingParam("gitref", "$(body.head_commit.id)"),
				createBindingParam("gitrepositoryurl", "$(body.repository.clone_url)"),
			},
		},
	}
}

func createDevCIBuildBinding() triggersv1.TriggerBinding {
	return triggersv1.TriggerBinding{
		TypeMeta:   createTypeMeta(),
		ObjectMeta: createObjectMeta("dev-ci-build-from-pr-binding"),
		Spec: triggersv1.TriggerBindingSpec{
			Params: []pipelinev2.Param{
				createBindingParam("gitref", "$(body.pull_request.head.ref)"),
				createBindingParam("gitsha", "$(body.pull_request.head.sha)"),
				createBindingParam("gitrepositoryurl", "$(body.repository.clone_url)"),
				createBindingParam("fullname", "$(body.repository.full_name)"),
			},
		},
	}
}

func createStageCDDeployBinding() triggersv1.TriggerBinding {
	return triggersv1.TriggerBinding{
		TypeMeta:   createTypeMeta(),
		ObjectMeta: createObjectMeta("stage-cd-deploy-from-push-binding"),
		Spec: triggersv1.TriggerBindingSpec{
			Params: []pipelinev2.Param{
				createBindingParam("gitref", "$(body.ref)"),
				createBindingParam("gitsha", "$(body.commits.0.id)"),
				createBindingParam("gitrepositoryurl", "$(body.repository.clone_url)"),
			},
		},
	}
}

func createStageCIDryRunBinding() triggersv1.TriggerBinding {
	return triggersv1.TriggerBinding{
		TypeMeta:   createTypeMeta(),
		ObjectMeta: createObjectMeta("stage-ci-dryrun-from-pr-binding"),
		Spec: triggersv1.TriggerBindingSpec{
			Params: []pipelinev2.Param{
				createBindingParam("gitref", "$(body.pull_request.head.ref)"),
				createBindingParam("gitrepositoryurl", "$(body.repository.clone_url)"),
			},
		},
	}
}

func createTypeMeta() v1.TypeMeta {
	return v1.TypeMeta{
		Kind:       "TriggerBinding",
		APIVersion: "tekton.dev/v1alpha1",
	}
}

func createObjectMeta(name string) v1.ObjectMeta {
	return v1.ObjectMeta{
		Name: name,
	}
}

func createBindingParam(name string, value string) pipelinev2.Param {
	return pipelinev2.Param{
		Name: name,
		Value: pipelinev2.ArrayOrString{
			StringVal: value,
		},
	}
}
