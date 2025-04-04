// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type Git struct {
	Commit       *string  `json:"commit,omitempty"`
	LocalChanges *float64 `json:"localChanges,omitempty"`
	Log          []Commit `json:"log,omitempty"`
}

func (o *Git) GetCommit() *string {
	if o == nil {
		return nil
	}
	return o.Commit
}

func (o *Git) GetLocalChanges() *float64 {
	if o == nil {
		return nil
	}
	return o.LocalChanges
}

func (o *Git) GetLog() []Commit {
	if o == nil {
		return nil
	}
	return o.Log
}

type ConfigGroup struct {
	Cloud                   *ConfigGroupCloud `json:"cloud,omitempty"`
	ConfigVersion           string            `json:"configVersion"`
	DeployingWorkerCount    *float64          `json:"deployingWorkerCount,omitempty"`
	Description             *string           `json:"description,omitempty"`
	EstimatedIngestRate     *float64          `json:"estimatedIngestRate,omitempty"`
	Git                     *Git              `json:"git,omitempty"`
	ID                      string            `json:"id"`
	IncompatibleWorkerCount *float64          `json:"incompatibleWorkerCount,omitempty"`
	Inherits                *string           `json:"inherits,omitempty"`
	IsFleet                 *bool             `json:"isFleet,omitempty"`
	IsSearch                *bool             `json:"isSearch,omitempty"`
	Name                    *string           `json:"name,omitempty"`
	OnPrem                  *bool             `json:"onPrem,omitempty"`
	Provisioned             *bool             `json:"provisioned,omitempty"`
	Streamtags              []string          `json:"streamtags,omitempty"`
	Tags                    *string           `json:"tags,omitempty"`
	UpgradeVersion          *string           `json:"upgradeVersion,omitempty"`
	WorkerCount             *float64          `json:"workerCount,omitempty"`
	WorkerRemoteAccess      *bool             `json:"workerRemoteAccess,omitempty"`
}

func (o *ConfigGroup) GetCloud() *ConfigGroupCloud {
	if o == nil {
		return nil
	}
	return o.Cloud
}

func (o *ConfigGroup) GetConfigVersion() string {
	if o == nil {
		return ""
	}
	return o.ConfigVersion
}

func (o *ConfigGroup) GetDeployingWorkerCount() *float64 {
	if o == nil {
		return nil
	}
	return o.DeployingWorkerCount
}

func (o *ConfigGroup) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *ConfigGroup) GetEstimatedIngestRate() *float64 {
	if o == nil {
		return nil
	}
	return o.EstimatedIngestRate
}

func (o *ConfigGroup) GetGit() *Git {
	if o == nil {
		return nil
	}
	return o.Git
}

func (o *ConfigGroup) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ConfigGroup) GetIncompatibleWorkerCount() *float64 {
	if o == nil {
		return nil
	}
	return o.IncompatibleWorkerCount
}

func (o *ConfigGroup) GetInherits() *string {
	if o == nil {
		return nil
	}
	return o.Inherits
}

func (o *ConfigGroup) GetIsFleet() *bool {
	if o == nil {
		return nil
	}
	return o.IsFleet
}

func (o *ConfigGroup) GetIsSearch() *bool {
	if o == nil {
		return nil
	}
	return o.IsSearch
}

func (o *ConfigGroup) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ConfigGroup) GetOnPrem() *bool {
	if o == nil {
		return nil
	}
	return o.OnPrem
}

func (o *ConfigGroup) GetProvisioned() *bool {
	if o == nil {
		return nil
	}
	return o.Provisioned
}

func (o *ConfigGroup) GetStreamtags() []string {
	if o == nil {
		return nil
	}
	return o.Streamtags
}

func (o *ConfigGroup) GetTags() *string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *ConfigGroup) GetUpgradeVersion() *string {
	if o == nil {
		return nil
	}
	return o.UpgradeVersion
}

func (o *ConfigGroup) GetWorkerCount() *float64 {
	if o == nil {
		return nil
	}
	return o.WorkerCount
}

func (o *ConfigGroup) GetWorkerRemoteAccess() *bool {
	if o == nil {
		return nil
	}
	return o.WorkerRemoteAccess
}
