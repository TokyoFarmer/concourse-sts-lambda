package handler

import (
	"strings"
	"text/template"
)

// Configuration passed to the Lambda, pointing to an S3 object with the team configuration.
type Configuration struct {
	Bucket string `json:"bucket"`
	Key    string `json:"key"`
}

// Team represents the configuration for a single team.
type Team struct {
	Name     string     `json:"name"`
	Accounts []*Account `json:"accounts"`
	// TODO make enum "parameter-store" | "secrets-manager"
	ParameterType string `json:"parameterType"`
}

// Account represents the configuration for an assumable role.
type Account struct {
	Name     string `json:"name"`
	RoleArn  string `json:"roleArn"`
	Duration int64  `json:"duration"`
}

// SecretPath represents the path used to write secrets into Secrets manager.
type SecretPath struct {
	Team     string
	Account  string
	Template string
}

// ParameterStorePath represents the path used to write secrets into Secrets manager.
type ParameterStorePath struct {
	Team     string
	Account  string
	Template string
}

// NewSecretPath ...
func NewSecretPath(team, account, template string) *SecretPath {
	return &SecretPath{
		Team:     team,
		Account:  account,
		Template: template,
	}
}

// NewParameterStorePath ...
func NewParameterStorePath(team, account, template string) *ParameterStorePath {
	return &ParameterStorePath{
		Team:     team,
		Account:  account,
		Template: template,
	}
}

func (p *SecretPath) String() (string, error) {
	t, err := template.New("path").Option("missingkey=error").Parse(p.Template)
	if err != nil {
		return "", err
	}
	var s strings.Builder
	if err = t.Execute(&s, p); err != nil {
		return "", err
	}
	return s.String(), nil
}

func (p *ParameterStorePath) String() (string, error) {
	t, err := template.New("path").Option("missingkey=error").Parse(p.Template)
	if err != nil {
		return "", err
	}
	var s strings.Builder
	if err = t.Execute(&s, p); err != nil {
		return "", err
	}
	return s.String(), nil
}
