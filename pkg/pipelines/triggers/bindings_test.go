package triggers

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	pipelinev2 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1alpha2"
	triggersv1 "github.com/tektoncd/triggers/pkg/apis/triggers/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestCreateDevCDDeployBinding(t *testing.T) {
	validDevCDBinding := triggersv1.TriggerBinding{
		TypeMeta: v1.TypeMeta{
			Kind:       "TriggerBinding",
			APIVersion: "tekton.dev/v1alpha1",
		},
		ObjectMeta: v1.ObjectMeta{
			Name: "dev-cd-deploy-from-master-binding",
		},
		Spec: triggersv1.TriggerBindingSpec{
			Params: []pipelinev2.Param{
				pipelinev2.Param{
					Name: "gitref",
					Value: pipelinev2.ArrayOrString{
						StringVal: "$(body.head_commit.id)",
					},
				},
				pipelinev2.Param{
					Name: "gitrepositoryurl",
					Value: pipelinev2.ArrayOrString{
						StringVal: "$(body.repository.clone_url)",
					},
				},
			},
		},
	}
	binding := createDevCDDeployBinding()
	if diff := cmp.Diff(validDevCDBinding, binding); diff != "" {
		t.Errorf("createDevCDDeployBinding() failed:\n%s", diff)
	}
}

func TestCreateDevCIBuildBinding(t *testing.T) {
	validDevCIBinding := triggersv1.TriggerBinding{
		TypeMeta: v1.TypeMeta{
			Kind:       "TriggerBinding",
			APIVersion: "tekton.dev/v1alpha1",
		},
		ObjectMeta: v1.ObjectMeta{
			Name: "dev-ci-build-from-pr-binding",
		},
		Spec: triggersv1.TriggerBindingSpec{
			Params: []pipelinev2.Param{
				pipelinev2.Param{
					Name: "gitref",
					Value: pipelinev2.ArrayOrString{
						StringVal: "$(body.pull_request.head.ref)",
					},
				},
				pipelinev2.Param{
					Name: "gitsha",
					Value: pipelinev2.ArrayOrString{
						StringVal: "$(body.pull_request.head.sha)",
					},
				},
				pipelinev2.Param{
					Name: "gitrepositoryurl",
					Value: pipelinev2.ArrayOrString{
						StringVal: "$(body.repository.clone_url)",
					},
				},
				pipelinev2.Param{
					Name: "fullname",
					Value: pipelinev2.ArrayOrString{
						StringVal: "$(body.repository.full_name)",
					},
				},
			},
		},
	}
	binding := createDevCIBuildBinding()
	if diff := cmp.Diff(validDevCIBinding, binding); diff != "" {
		t.Errorf("createDevCIBuildBinding() failed:\n%s", diff)
	}
}

func TestCreateStageCDDeployBinding(t *testing.T) {
	validStageCDDeployBinding := triggersv1.TriggerBinding{
		TypeMeta: v1.TypeMeta{
			Kind:       "TriggerBinding",
			APIVersion: "tekton.dev/v1alpha1",
		},
		ObjectMeta: v1.ObjectMeta{
			Name: "stage-cd-deploy-from-push-binding",
		},
		Spec: triggersv1.TriggerBindingSpec{
			Params: []pipelinev2.Param{
				pipelinev2.Param{
					Name: "gitref",
					Value: pipelinev2.ArrayOrString{
						StringVal: "$(body.ref)",
					},
				},
				pipelinev2.Param{
					Name: "gitsha",
					Value: pipelinev2.ArrayOrString{
						StringVal: "$(body.commits.0.id)",
					},
				},
				pipelinev2.Param{
					Name: "gitrepositoryurl",
					Value: pipelinev2.ArrayOrString{
						StringVal: "$(body.repository.clone_url)",
					},
				},
			},
		},
	}
	binding := createStageCDDeployBinding()
	if diff := cmp.Diff(validStageCDDeployBinding, binding); diff != "" {
		t.Errorf("createDevCIBuildBinding() failed:\n%s", diff)
	}
}

func TestCreateStageCIDryRunBinding(t *testing.T) {
	validStageCIDryRunBinding := triggersv1.TriggerBinding{
		TypeMeta: v1.TypeMeta{
			Kind:       "TriggerBinding",
			APIVersion: "tekton.dev/v1alpha1",
		},
		ObjectMeta: v1.ObjectMeta{
			Name: "stage-ci-dryrun-from-pr-binding",
		},
		Spec: triggersv1.TriggerBindingSpec{
			Params: []pipelinev2.Param{
				pipelinev2.Param{
					Name: "gitref",
					Value: pipelinev2.ArrayOrString{
						StringVal: "$(body.pull_request.head.ref)",
					},
				},
				pipelinev2.Param{
					Name: "gitrepositoryurl",
					Value: pipelinev2.ArrayOrString{
						StringVal: "$(body.repository.clone_url)",
					},
				},
			},
		},
	}
	binding := createStageCIDryRunBinding()
	if diff := cmp.Diff(validStageCIDryRunBinding, binding); diff != "" {
		t.Errorf("createStageCIDryRunBinding() failed:\n%s", diff)
	}
}
