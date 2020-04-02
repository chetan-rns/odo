package pipelines

import (
	"fmt"
	"strings"

	"github.com/openshift/odo/pkg/odo/genericclioptions"
	"github.com/openshift/odo/pkg/pipelines"
	"github.com/spf13/cobra"

	ktemplates "k8s.io/kubernetes/pkg/kubectl/util/templates"
)

const (
	// InitRecommendedCommandName the recommended command name
	InitRecommendedCommandName = "init"
)

var (
	initExample = ktemplates.Examples(`
	# Initialise OpenShift pipelines in a cluster
	%[1]s 
	`)

	initLongDesc  = ktemplates.LongDesc(`Initialise GitOps CI/CD Pipelines`)
	initShortDesc = `Initialise pipelines`
)

// InitParameters encapsulates the parameters for the odo pipelines init command.
type InitParameters struct {
	gitOpsRepo               string // repo to store Gitops resources e.g. org/repo
	gitOpsWebhookSecret      string // used to create Github's shared webhook secret for gitops repo
	output                   string // path to add Gitops resources
	prefix                   string // used to generate the environments in a shared cluster
	skipChecks               bool   // skip Tekton installation checks
	appGitRepo               string
	appWebhookSecret         string
	appImageRepo             string
	internalRegistryHostname string
	envName                  string
	dockerConfigJSON         string
	// generic context options common to all commands
	*genericclioptions.Context
}

// NewInitParameters bootstraps a InitParameters instance.
func NewInitParameters() *InitParameters {
	return &InitParameters{}
}

// Complete completes InitParameters after they've been created.
//
// If the prefix provided doesn't have a "-" then one is added, this makes the
// generated environment names nicer to read.
func (io *InitParameters) Complete(name string, cmd *cobra.Command, args []string) error {
	if io.prefix != "" && !strings.HasSuffix(io.prefix, "-") {
		io.prefix = io.prefix + "-"
	}
	return nil
}

// Validate validates the parameters of the InitParameters.
func (io *InitParameters) Validate() error {
	// TODO: this won't work with GitLab as the repo can have more path elements.
	if len(strings.Split(io.gitOpsRepo, "/")) != 2 {
		return fmt.Errorf("repo must be org/repo: %s", io.gitOpsRepo)
	}

	if io.appGitRepo != "" && len(strings.Split(io.appGitRepo, "/")) != 2 {
		return fmt.Errorf("repo must be org/repo: %s", io.appGitRepo)
	}

	if io.appGitRepo != "" {
		return checkAppParameters(io)
	} else {
		if !containsEmpty(io.appWebhookSecret, io.appImageRepo, io.envName) {
			return missingParameterError("app-git-repo")
		}
	}
	return nil
}

func containsEmpty(params ...string) bool {
	for _, param := range params {
		if param != "" {
			return false
		}
	}
	return true
}

func checkAppParameters(o *InitParameters) error {
	if o.appWebhookSecret == "" {
		return missingParameterError("app-webhook-secret")
	}
	if o.appImageRepo == "" {
		return missingParameterError("app-image-repo")
	}
	if o.envName == "" {
		return missingParameterError("env-name")
	}
	return nil
}

func missingParameterError(param string) error {
	return fmt.Errorf("Flag:%s required to initialize application", param)
}

// Run runs the project bootstrap command.
func (io *InitParameters) Run() error {
	options := pipelines.InitParameters{
		GitOpsWebhookSecret:      io.gitOpsWebhookSecret,
		GitOpsRepo:               io.gitOpsRepo,
		Output:                   io.output,
		Prefix:                   io.prefix,
		SkipChecks:               io.skipChecks,
		AppGitRepo:               io.appGitRepo,
		AppWebhookSecret:         io.appWebhookSecret,
		AppImageRepo:             io.appImageRepo,
		InternalRegistryHostname: io.internalRegistryHostname,
		EnvName:                  io.envName,
		DockerConfigJSON:         io.dockerConfigJSON,
	}
	return pipelines.Init(&options)
}

// NewCmdInit creates the project init command.
func NewCmdInit(name, fullName string) *cobra.Command {
	o := NewInitParameters()

	initCmd := &cobra.Command{
		Use:     name,
		Short:   initShortDesc,
		Long:    initLongDesc,
		Example: fmt.Sprintf(initExample, fullName),
		Run: func(cmd *cobra.Command, args []string) {
			genericclioptions.GenericRun(o, cmd, args)
		},
	}

	initCmd.Flags().StringVar(&o.gitOpsRepo, "gitops-repo", "", "CI/CD pipelines configuration Git repository in this form <username>/<repository>")
	initCmd.MarkFlagRequired("gitops-repo")
	initCmd.Flags().StringVar(&o.gitOpsWebhookSecret, "gitops-webhook-secret", "", "provide the GitHub webhook secret for gitops repository")
	initCmd.MarkFlagRequired("gitops-webhook-secret")
	initCmd.Flags().StringVar(&o.output, "output", ".", "folder path to add Gitops resources")
	initCmd.Flags().StringVarP(&o.prefix, "prefix", "p", "", "add a prefix to the environment names")
	initCmd.Flags().BoolVarP(&o.skipChecks, "skip-checks", "s", false, "skip Tekton installation checks")

	initCmd.Flags().StringVar(&o.appGitRepo, "app-git-repo", "", "CI/CD pipelines configuration Git repository in this form <username>/<repository>")
	initCmd.Flags().StringVar(&o.appWebhookSecret, "app-webhook-secret", "", "Provide the webhook secret of the app git repository")
	initCmd.Flags().StringVar(&o.appImageRepo, "app-image-repo", "", "Image repository name in form <username>/<repository>")
	initCmd.Flags().StringVar(&o.envName, "env-name", "", "Add the name of the environment(namespace) to which the pipelines should be bootstrapped")
	initCmd.Flags().StringVar(&o.dockerConfigJSON, "dockercfgjson", "", "Add the docker auth.json file path")
	return initCmd
}
