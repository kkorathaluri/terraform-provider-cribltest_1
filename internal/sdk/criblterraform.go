// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package sdk

import (
	"context"
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/internal/hooks"
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/internal/utils"
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/models/shared"
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/retry"
	"net/http"
	"time"
)

// HTTPClient provides an interface for suplying the SDK with a custom HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// Bool provides a helper function to return a pointer to a bool
func Bool(b bool) *bool { return &b }

// Int provides a helper function to return a pointer to an int
func Int(i int) *int { return &i }

// Int64 provides a helper function to return a pointer to an int64
func Int64(i int64) *int64 { return &i }

// Float32 provides a helper function to return a pointer to a float32
func Float32(f float32) *float32 { return &f }

// Float64 provides a helper function to return a pointer to a float64
func Float64(f float64) *float64 { return &f }

// Pointer provides a helper function to return a pointer to a type
func Pointer[T any](v T) *T { return &v }

type sdkConfiguration struct {
	Client            HTTPClient
	Security          func(context.Context) (interface{}, error)
	ServerURL         string
	Language          string
	OpenAPIDocVersion string
	SDKVersion        string
	GenVersion        string
	UserAgent         string
	RetryConfig       *retry.Config
	Hooks             *hooks.Hooks
	Timeout           *time.Duration
}

func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
	return c.ServerURL, map[string]string{}
}

// CriblTerraform - Cribl API Reference: This API Reference lists available REST endpoints, along with their supported operations for accessing, creating, updating, or deleting resources. See our complementary product documentation at [docs.cribl.io](http://docs.cribl.io).
type CriblTerraform struct {
	Billing    *Billing
	V5         *V5
	Sandboxes  *Sandboxes
	Workspaces *Workspaces
	// Actions related to Projects
	Projects *Projects
	// Actions related to Subscriptions
	Subscriptions *Subscriptions
	// Actions related to Versioning
	Versioning *Versioning
	// Actions related to Git
	Git *Git
	// Actions related to data preview
	Preview *Preview
	// Actions related to samples
	Samples *Samples
	// Actions related to Pipelines
	Pipelines *Pipelines
	// Actions related to Banners
	Banners *Banners
	// Actions related to Certificates
	Certificates *Certificates
	// Actions related to Features
	Features *Features
	// Actions related to Saved Jobs
	SavedJobs *SavedJobs
	// Actions related to encryption keys
	Keys *Keys
	// Actions related to messages
	Messages *Messages
	// Actions related to Notification Targets
	NotificationTargets *NotificationTargets
	// Actions related to Notifications
	Notifications *Notifications
	// Actions related to Policies
	Policies *Policies
	// Actions related to Roles
	Roles *Roles
	// Actions related to scripts
	Scripts *Scripts
	// Actions related to Teams
	Teams *Teams
	// Actions related to users
	Users *Users
	// Actions related to Lake
	Lake *Lake
	// Actions related to DashboardCategories
	DashboardCategories *DashboardCategories
	// Actions related to Usage Groups
	UsageGroups *UsageGroups
	// Actions related to Datasets
	Datasets *Datasets
	// Actions related to Users ACL
	UsersACL *UsersACL
	// Actions related to Teams ACL
	TeamsACL *TeamsACL
	// Actions related to Appscope Configs
	AppscopeConfigs *AppscopeConfigs
	// Actions related to Grok files
	Grokfiles *Grokfiles
	// Actions related to lookups
	Lookups *Lookups
	// Actions related to parsers
	Parsers *Parsers
	// Actions related to Protobuf libraries
	Protobuflibraries *Protobuflibraries
	// Actions related to regular expressions
	Regexes *Regexes
	// Actions related to Dashboards
	Dashboards *Dashboards
	// Actions related to Macros
	Macros *Macros
	// Actions related to saved queries
	SavedQueries *SavedQueries
	// Actions related to Search
	Search *Search
	// Actions related to Database Connections
	DatabaseConnections *DatabaseConnections
	// Actions related to Event Breaker rules
	EventBreakerRules *EventBreakerRules
	// Actions related to Global Variables
	GlobalVariables *GlobalVariables
	// Actions related to HMAC functions
	HmacFunctions *HmacFunctions
	// Actions related to inputs
	Inputs *Inputs
	// Actions related to outputs
	Outputs *Outputs
	// Actions related to Parquet schemas
	Parquetschemas *Parquetschemas
	// Actions related to Profiler
	Profiler *Profiler
	// Actions related to Routes
	Routes *Routes
	// Actions related to Schemas
	Schemas *Schemas
	// Actions related to Secrets
	Secrets *Secrets
	// Actions related to Edge AppScope processes
	EdgeAppScopeProcesses *EdgeAppScopeProcesses
	// Actions enabled in Edge
	Edge *Edge
	// Actions related to EdgeEvents
	EdgeEvents *EdgeEvents
	// Actions related to Events
	Events *Events
	// Actions related to Edge Files
	EdgeFiles *EdgeFiles
	// Actions related to Edge listing
	EdgeLs *EdgeLs
	// Actions related to File
	File *File
	// Actions related to Ingest
	Ingest *Ingest
	// Actions related to FileSampler
	FileSampler *FileSampler
	// Actions related to Kube Logs
	KubeLogs *KubeLogs
	// Actions related to Kube Proxy
	KubeProxy *KubeProxy
	// Actions related to authentication. Do not use the /auth endpoints in Cribl.Cloud deployments. Instead, follow the instructions at https://docs.cribl.io/stream/api-tutorials/#criblcloud to authenticate for Cribl.Cloud.
	Auth *Auth
	// Actions related to Authorize
	Authorize *Authorize
	// Actions related to Changelog
	Changelog *Changelog
	// Actions related to system settings
	System *System
	// Actions related to ClickHouse
	ClickHouse *ClickHouse
	// Actions related to CLUI
	Clui *Clui
	// Actions related to Distributed
	Distributed *Distributed
	// Actions related to Workers
	Workers *Workers
	// Actions related to expressions
	Expressions *Expressions
	// Actions related to Conditions
	Conditions *Conditions
	// Actions related to diagnostics
	Diag *Diag
	// Actions related to REST server health
	Health *Health
	// Actions related to Jobs
	Jobs *Jobs
	// Actions related to Security
	Security *Security
	// Actions related to licenses. The <code>/licenses</code> endpoints do not apply to Cribl.Cloud deployments.
	Licenses *Licenses
	// Actions related to Logger
	Logger *Logger
	// Actions related to logging
	Logging *Logging
	// Actions related to Packs
	Packs *Packs
	// Actions related to Processes
	Processes *Processes
	// Actions related to metrics
	Metrics *Metrics
	// Actions related to UiState
	UIState *UIState
	// Actions related to Consent
	Consent *Consent
	// Actions related to Trust Policies
	TrustPolicies *TrustPolicies
	// Actions related to Edge containers
	EdgeContainers *EdgeContainers
	// Actions related to Edge processes
	EdgeProcesses *EdgeProcesses
	// Actions related to functions
	Functions *Functions
	// Actions related to Collectors
	Collectors *Collectors
	// Actions related to Executors
	Executors *Executors
	// Actions related to Groups
	Groups *Groups

	sdkConfiguration sdkConfiguration
}

type SDKOption func(*CriblTerraform)

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *CriblTerraform) {
		sdk.sdkConfiguration.Client = client
	}
}

// WithSecurity configures the SDK to use the provided security details
func WithSecurity(security shared.Security) SDKOption {
	return func(sdk *CriblTerraform) {
		sdk.sdkConfiguration.Security = utils.AsSecuritySource(security)
	}
}

// WithSecuritySource configures the SDK to invoke the Security Source function on each method call to determine authentication
func WithSecuritySource(security func(context.Context) (shared.Security, error)) SDKOption {
	return func(sdk *CriblTerraform) {
		sdk.sdkConfiguration.Security = func(ctx context.Context) (interface{}, error) {
			return security(ctx)
		}
	}
}

func WithRetryConfig(retryConfig retry.Config) SDKOption {
	return func(sdk *CriblTerraform) {
		sdk.sdkConfiguration.RetryConfig = &retryConfig
	}
}

// WithTimeout Optional request timeout applied to each operation
func WithTimeout(timeout time.Duration) SDKOption {
	return func(sdk *CriblTerraform) {
		sdk.sdkConfiguration.Timeout = &timeout
	}
}

// New creates a new instance of the SDK with the provided serverURL and options
func New(serverURL string, opts ...SDKOption) *CriblTerraform {
	sdk := &CriblTerraform{
		sdkConfiguration: sdkConfiguration{
			Language:          "go",
			OpenAPIDocVersion: "4.12.0-alpha.1742471106527-a8a53ddb",
			SDKVersion:        "0.0.7",
			GenVersion:        "2.563.0",
			UserAgent:         "speakeasy-sdk/terraform 0.0.7 2.563.0 4.12.0-alpha.1742471106527-a8a53ddb github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk",
			ServerURL:         serverURL,
			Hooks:             hooks.New(),
		},
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk.sdkConfiguration.Client == nil {
		sdk.sdkConfiguration.Client = &http.Client{Timeout: 60 * time.Second}
	}

	currentServerURL := serverURL
	serverURL, sdk.sdkConfiguration.Client = sdk.sdkConfiguration.Hooks.SDKInit(currentServerURL, sdk.sdkConfiguration.Client)
	if serverURL != currentServerURL {
		sdk.sdkConfiguration.ServerURL = serverURL
	}

	sdk.Billing = newBilling(sdk.sdkConfiguration)

	sdk.V5 = newV5(sdk.sdkConfiguration)

	sdk.Sandboxes = newSandboxes(sdk.sdkConfiguration)

	sdk.Workspaces = newWorkspaces(sdk.sdkConfiguration)

	sdk.Projects = newProjects(sdk.sdkConfiguration)

	sdk.Subscriptions = newSubscriptions(sdk.sdkConfiguration)

	sdk.Versioning = newVersioning(sdk.sdkConfiguration)

	sdk.Git = newGit(sdk.sdkConfiguration)

	sdk.Preview = newPreview(sdk.sdkConfiguration)

	sdk.Samples = newSamples(sdk.sdkConfiguration)

	sdk.Pipelines = newPipelines(sdk.sdkConfiguration)

	sdk.Banners = newBanners(sdk.sdkConfiguration)

	sdk.Certificates = newCertificates(sdk.sdkConfiguration)

	sdk.Features = newFeatures(sdk.sdkConfiguration)

	sdk.SavedJobs = newSavedJobs(sdk.sdkConfiguration)

	sdk.Keys = newKeys(sdk.sdkConfiguration)

	sdk.Messages = newMessages(sdk.sdkConfiguration)

	sdk.NotificationTargets = newNotificationTargets(sdk.sdkConfiguration)

	sdk.Notifications = newNotifications(sdk.sdkConfiguration)

	sdk.Policies = newPolicies(sdk.sdkConfiguration)

	sdk.Roles = newRoles(sdk.sdkConfiguration)

	sdk.Scripts = newScripts(sdk.sdkConfiguration)

	sdk.Teams = newTeams(sdk.sdkConfiguration)

	sdk.Users = newUsers(sdk.sdkConfiguration)

	sdk.Lake = newLake(sdk.sdkConfiguration)

	sdk.DashboardCategories = newDashboardCategories(sdk.sdkConfiguration)

	sdk.UsageGroups = newUsageGroups(sdk.sdkConfiguration)

	sdk.Datasets = newDatasets(sdk.sdkConfiguration)

	sdk.UsersACL = newUsersACL(sdk.sdkConfiguration)

	sdk.TeamsACL = newTeamsACL(sdk.sdkConfiguration)

	sdk.AppscopeConfigs = newAppscopeConfigs(sdk.sdkConfiguration)

	sdk.Grokfiles = newGrokfiles(sdk.sdkConfiguration)

	sdk.Lookups = newLookups(sdk.sdkConfiguration)

	sdk.Parsers = newParsers(sdk.sdkConfiguration)

	sdk.Protobuflibraries = newProtobuflibraries(sdk.sdkConfiguration)

	sdk.Regexes = newRegexes(sdk.sdkConfiguration)

	sdk.Dashboards = newDashboards(sdk.sdkConfiguration)

	sdk.Macros = newMacros(sdk.sdkConfiguration)

	sdk.SavedQueries = newSavedQueries(sdk.sdkConfiguration)

	sdk.Search = newSearch(sdk.sdkConfiguration)

	sdk.DatabaseConnections = newDatabaseConnections(sdk.sdkConfiguration)

	sdk.EventBreakerRules = newEventBreakerRules(sdk.sdkConfiguration)

	sdk.GlobalVariables = newGlobalVariables(sdk.sdkConfiguration)

	sdk.HmacFunctions = newHmacFunctions(sdk.sdkConfiguration)

	sdk.Inputs = newInputs(sdk.sdkConfiguration)

	sdk.Outputs = newOutputs(sdk.sdkConfiguration)

	sdk.Parquetschemas = newParquetschemas(sdk.sdkConfiguration)

	sdk.Profiler = newProfiler(sdk.sdkConfiguration)

	sdk.Routes = newRoutes(sdk.sdkConfiguration)

	sdk.Schemas = newSchemas(sdk.sdkConfiguration)

	sdk.Secrets = newSecrets(sdk.sdkConfiguration)

	sdk.EdgeAppScopeProcesses = newEdgeAppScopeProcesses(sdk.sdkConfiguration)

	sdk.Edge = newEdge(sdk.sdkConfiguration)

	sdk.EdgeEvents = newEdgeEvents(sdk.sdkConfiguration)

	sdk.Events = newEvents(sdk.sdkConfiguration)

	sdk.EdgeFiles = newEdgeFiles(sdk.sdkConfiguration)

	sdk.EdgeLs = newEdgeLs(sdk.sdkConfiguration)

	sdk.File = newFile(sdk.sdkConfiguration)

	sdk.Ingest = newIngest(sdk.sdkConfiguration)

	sdk.FileSampler = newFileSampler(sdk.sdkConfiguration)

	sdk.KubeLogs = newKubeLogs(sdk.sdkConfiguration)

	sdk.KubeProxy = newKubeProxy(sdk.sdkConfiguration)

	sdk.Auth = newAuth(sdk.sdkConfiguration)

	sdk.Authorize = newAuthorize(sdk.sdkConfiguration)

	sdk.Changelog = newChangelog(sdk.sdkConfiguration)

	sdk.System = newSystem(sdk.sdkConfiguration)

	sdk.ClickHouse = newClickHouse(sdk.sdkConfiguration)

	sdk.Clui = newClui(sdk.sdkConfiguration)

	sdk.Distributed = newDistributed(sdk.sdkConfiguration)

	sdk.Workers = newWorkers(sdk.sdkConfiguration)

	sdk.Expressions = newExpressions(sdk.sdkConfiguration)

	sdk.Conditions = newConditions(sdk.sdkConfiguration)

	sdk.Diag = newDiag(sdk.sdkConfiguration)

	sdk.Health = newHealth(sdk.sdkConfiguration)

	sdk.Jobs = newJobs(sdk.sdkConfiguration)

	sdk.Security = newSecurity(sdk.sdkConfiguration)

	sdk.Licenses = newLicenses(sdk.sdkConfiguration)

	sdk.Logger = newLogger(sdk.sdkConfiguration)

	sdk.Logging = newLogging(sdk.sdkConfiguration)

	sdk.Packs = newPacks(sdk.sdkConfiguration)

	sdk.Processes = newProcesses(sdk.sdkConfiguration)

	sdk.Metrics = newMetrics(sdk.sdkConfiguration)

	sdk.UIState = newUIState(sdk.sdkConfiguration)

	sdk.Consent = newConsent(sdk.sdkConfiguration)

	sdk.TrustPolicies = newTrustPolicies(sdk.sdkConfiguration)

	sdk.EdgeContainers = newEdgeContainers(sdk.sdkConfiguration)

	sdk.EdgeProcesses = newEdgeProcesses(sdk.sdkConfiguration)

	sdk.Functions = newFunctions(sdk.sdkConfiguration)

	sdk.Collectors = newCollectors(sdk.sdkConfiguration)

	sdk.Executors = newExecutors(sdk.sdkConfiguration)

	sdk.Groups = newGroups(sdk.sdkConfiguration)

	return sdk
}
