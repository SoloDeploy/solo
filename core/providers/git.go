package providers

import (
	"context"
	"fmt"
	"os/exec"
	"time"

	"github.com/SoloDeploy/solo/core/configuration"
	"github.com/SoloDeploy/solo/core/output"
	pb "github.com/SoloDeploy/solo/grpc/git"
	"google.golang.org/grpc"
)

type Repository struct {
	name string
}

// GitProvider defines the abstraction wrapper for all functions to interact with
// the Git Provider implementation as configured in the .solo/config.yml file.
type GitProvider struct {
	port       int
	client     pb.GitProviderClient
	connection *grpc.ClientConn
	cmd        *exec.Cmd
}

func (provider *GitProvider) Close() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	_, err := provider.client.Close(ctx, &pb.CloseRequest{})
	if err != nil {
		output.FPrintlnWarn("Couldn't close Git Provider process:\n%v", err)
	}
	provider.connection.Close()
}

func NewGitProvider(configuration *configuration.Configuration) (provider *GitProvider, err error) {
	port := 5000 // TODO: randomise this
	provider = &GitProvider{
		port: port,
	}

	output.FPrintlnLog("Starting the Git Provider at port %v", port)
	output.FPrintlnLog("Provider Options:\n%v", configuration.Providers.Git.Options)
	gitProviderPath, err := GetProviderPath("git")
	if err != nil {
		return nil, err
	}

	provider.cmd = exec.Command(gitProviderPath, "start", "-p", fmt.Sprint(port))
	// TODO: process these streams to indicate where the logs are coming from
	provider.cmd.Stdout = NewStreamWriter(output.FPrintLog, "[GitProvider]")
	provider.cmd.Stderr = NewStreamWriter(output.FPrintLog, "[GitProvider]")
	err = provider.cmd.Start()
	if err != nil {
		return
	}

	output.PrintlnLog("Git Provider started")

	// Set up a connection to the server.
	url := fmt.Sprint("localhost:", provider.port)
	output.FPrintlnLog("Opening connection to %v", url)
	conn, err := grpc.Dial(url, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	provider.connection = conn
	output.PrintlnLog("Creating gRPC client")
	provider.client = pb.NewGitProviderClient(conn)
	return
}

// GetAllRepositories returns a list of all the Repository objects in the associated
// Git Provider.
func (provider *GitProvider) GetAllRepositories() (repositories []Repository, err error) {
	repositories = []Repository{}
	return
}

// GetRepositoryNames returns a list of all the Git repository names.
func (provider *GitProvider) GetRepositoryNames() (names []string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	r, err := provider.client.GetRepositoryNames(ctx, &pb.GetRepositoryNamesRequest{})
	if err != nil {
		return
	}
	names = r.GetNames()
	return
}

// RepositoryExists checks if the repository exists.
func (provider *GitProvider) RepositoryExists(repositoryName string) (exists bool, err error) {
	return
}

// CreateRepository creates a new repository with the specified name.
func (provider *GitProvider) CreateRepository(repositoryName string) error {
	return nil
}

// CloneRepository clones a repository with the specified name into the specified local folder path.
func (provider *GitProvider) CloneRepository(repositoryName string, path string) error {
	return nil
}

// FetchRemotes performs a Git fetch on a local Git directory.
func (provider *GitProvider) FetchRemotes(path string) error {
	return nil
}

// CheckoutBranch checks out an existing branch for a local Git directory.
func (provider *GitProvider) CheckoutBranch(path string, branchName string) error {
	return nil
}

// CheckoutNewBranch checks out a new branch for a local Git directory.
func (provider *GitProvider) CheckoutNewBranch(path string, branchName string) error {
	return nil
}

// CreatePullRequest creates a new pull request
func (provider *GitProvider) CreatePullRequest(repositoryName string, title string, sourceBranchName string, targetBranchName string) (pullRequestId string, err error) {
	return
}

// MergePullRequest merges an existing pull request
func (provider *GitProvider) MergePullRequest(repositoryName string, pullRequestId string) error {
	return nil
}

// CancelPullRequest cancels an existing pull request
func (provider *GitProvider) CancelPullRequest(repositoryName string, pullRequestId string) error {
	return nil
}
