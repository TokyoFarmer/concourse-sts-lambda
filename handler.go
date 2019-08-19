package handler

import (
	"github.com/sirupsen/logrus"
)

// New lambda handler with the provided settings.
func New(manager *Manager, secretTemplate string, logger *logrus.Logger) func(Team) error {
	return func(team Team) error {
		// team, err := manager.ReadConfig(config.Bucket, config.Key)
		// if err != nil {
		// 	return fmt.Errorf("failed to read configuration: %s", err)
		// }
		// Loop through teams and assume roles/write credentials for
		// all accounts controlled by the team.
		for _, account := range team.Accounts {
			log := logger.WithFields(logrus.Fields{
				"team":     team.Name,
				"account":  account.Name,
				"role":     account.RoleArn,
				"duration": account.Duration,
			})

			path := ""

			creds, err := manager.AssumeRole(account.RoleArn, team.Name, account.Duration)
			if err != nil {
				log.Warnf("failed to assume role: %s", err)
				continue
			}

			if team.ParameterType == "secrets-manager" {
				path, err = NewSecretPath(team.Name, account.Name, secretTemplate).String()
				if err != nil {
					log.Warnf("failed to parse secret path: %s", err)
					continue
				}
				if err := manager.WriteCredentialsSecretsManager(creds, path); err != nil {
					log.Warnf("failed to write credentials: %s", err)
					continue
				}
			} else {
				path, err = NewParameterStorePath(team.Name, account.Name, secretTemplate).String()
				if err != nil {
					log.Warnf("failed to parse secret path: %s", err)
					continue
				}
				if err := manager.WriteCredentialsParameterStore(creds, path); err != nil {
					log.Warnf("failed to write credentials: %s", err)
					continue
				}
			}

		}
		return nil
	}
}
