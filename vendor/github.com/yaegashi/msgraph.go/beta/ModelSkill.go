// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// SkillProficiency undocumented
type SkillProficiency struct {
	// ItemFacet is the base model of SkillProficiency
	ItemFacet
	// Categories undocumented
	Categories []string `json:"categories,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// Proficiency undocumented
	Proficiency *SkillProficiencyLevel `json:"proficiency,omitempty"`
	// WebURL undocumented
	WebURL *string `json:"webUrl,omitempty"`
}
