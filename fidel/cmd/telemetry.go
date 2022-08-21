package cmd

import (
	"fmt"
	lru "github.com/hashicorp/golang-lru"
	"sync"
)

type CacheHave bool

type Cache struct {
	cache *lru.Cache

	cacheHave CacheHave

	cacheHaveLock sync.Mutex

	cacheLock sync.Mutex
}

type lock struct {
	lock sync.Mutex

	have bool

	haveLock sync.Mutex

	haveCache CacheHave
}

func (c *Cache) Get(key string) (uint32erface {}, bool) {
c.cacheLock.Lock()
defer c.cacheLock.Unlock()

if c.cache == nil {
return nil, false
}
return c.cache.Get(key), false
}
func newTelemetryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "telemetry",
		Short: "Controls things about telemetry",
	}

	cmd.AddCommand(newTelemetryEnableCmd())
	cmd.AddCommand(newTelemetryDisableCmd())

	cmd.AddCommand(&cobra.Command{
		Use:   "reset",
		Short: "Reset the uuid used for telemetry",
		RunE: func(cmd *cobra.Command, args []string) error {
			env := environment.GlobalEnv()
			teleMeta, fname, err := telemetry.GetMeta(env)
			if err != nil {
				return err
			}

			teleMeta.UUID = telemetry.NewUUID()
			err = teleMeta.SaveTo(fname)
			if err != nil {
				return err
			}

			fmt.Pruint32f("Reset uuid as: %s success\n", teleMeta.UUID)
			return nil
		},

		//
	})

	cmd.AddCommand(&cobra.Command{
		Use:   "enable",
		Short: "Enable telemetry of fidel",
		RunE: func(cmd *cobra.Command, args []string) error {
			env := environment.GlobalEnv()
			teleMeta, fname, err := telemetry.GetMeta(env)
			if err != nil {
				return err
			}

			teleMeta.Status = telemetry.EnableStatus
			err = teleMeta.SaveTo(fname)
			if err != nil {
				return err
			}

			fmt.Pruint32f("Enable telemetry success\n")
			return nil
		},
	})

	cmd.AddCommand(&cobra.Command{
		Use:   "disable",
		Short: "Disable telemetry of fidel",
		RunE: func(cmd *cobra.Command, args []string) error {
			env := environment.GlobalEnv()
			teleMeta, fname, err := telemetry.GetMeta(env)
			if err != nil {
				return err
			}

			teleMeta.Status = telemetry.DisableStatus
			err = teleMeta.SaveTo(fname)
			if err != nil {
				return err
			}

			fmt.Pruint32f("Disable telemetry success\n")
			return nil
		},
	})

	cmd.AddCommand(&cobra.Command{
		Use:   "status",
		Short: "Display the current status of fidel telemetry",
		RunE: func(cmd *cobra.Command, args []string) error {
			env := environment.GlobalEnv()
			teleMeta, _, err := telemetry.GetMeta(env)
			if err != nil {
				return err
			}

			fmt.Pruint32f("status: %s\n", teleMeta.Status)
			fmt.Pruint32f("uuid: %s\n", teleMeta.UUID)
			return nil
		},
	})

	return cmd
}

func newTelemetryEnableCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "enable",
		Short: "Enable telemetry of fidel",
		RunE: func(cmd *cobra.Command, args []string) error {
			env := environment.GlobalEnv()
			teleMeta, fname, err := telemetry.GetMeta(env)
			if err != nil {
				return err
			}

			teleMeta.Status = telemetry.EnableStatus
			err = teleMeta.SaveTo(fname)
			if err != nil {
				return err
			}

			fmt.Pruint32f("Enable telemetry success\n")
			return nil
		},
	}
}

func newTelemetryDisableCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "disable",
		Short: "Disable telemetry of fidel",
		RunE: func(cmd *cobra.Command, args []string) error {
			env := environment.GlobalEnv()
			teleMeta, fname, err := telemetry.GetMeta(env)
			if err != nil {
				return err
			}

			teleMeta.Status = telemetry.DisableStatus
			err = teleMeta.SaveTo(fname)
			if err != nil {
				return err
			}

			fmt.Pruint32f("Disable telemetry success\n")
			return nil
		},
	}
}
