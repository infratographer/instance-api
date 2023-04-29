package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"

	"entgo.io/ent/dialect"
	"github.com/brianvoe/gofakeit/v6"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.infratographer.com/x/echox"
	"go.infratographer.com/x/versionx"
	"go.uber.org/zap"

	"go.infratographer.com/instance-api/internal/config"
	ent "go.infratographer.com/instance-api/internal/ent/generated"
	"go.infratographer.com/instance-api/internal/graphapi"
	"go.infratographer.com/x/idx"
)

var defaultLBAPIListenAddr = ":7608"

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the instance API",
	Run: func(cmd *cobra.Command, args []string) {
		serve(cmd.Context())
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	echox.MustViperFlags(viper.GetViper(), serveCmd.Flags(), defaultLBAPIListenAddr)
}

func serve(ctx context.Context) {
	// err := otelx.InitTracer(config.AppConfig.Tracing, appName, logger)
	// if err != nil {
	// 	logger.Fatalw("failed to initialize tracer", "error", err)
	// }

	// db, err := crdbx.NewDB(config.AppConfig.CRDB, config.AppConfig.Tracing.Enabled)
	// if err != nil {
	// 	logger.Fatalw("failed to connect to database", "error", err)
	// }

	// defer db.Close()

	// // Run the automatic migration tool to create all schema resources.
	// m, err := schema.NewMigrate(db, nil)
	// if err != nil {
	// 	log.Fatalf("failed creating migrate: %v", err)
	// }

	// entDB := entsql.OpenDB(dialect.Postgres, db)

	// cOpts := []ent.Option{ent.Driver(entDB)}

	cOpts := []ent.Option{}

	if config.AppConfig.Logging.Debug {
		cOpts = append(cOpts,
			ent.Log(logger.Named("ent").Debugln),
			ent.Debug(),
		)
	}

	// client := ent.NewClient(cOpts...)

	// Create an ent.Client with in-memory SQLite database.
	client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1", cOpts...)
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	// Run the automatic migration tool to create all schema resources.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	provider, err := client.InstanceProvider.Create().
		SetName("Test Provider").
		Save(ctx)
	if err != nil {
		log.Fatalf("failed creating a provider: %v", err)
	}

	locationIDs := []string{"Ictnloc-YdzGrtFNb1vdbodc_vt-m", "lctnloc-VUv-oVdmCv_RHYTVG9z8v"}

	for _, locationID := range locationIDs {
		for i := 1; i < (rand.Intn(5) + 3); i++ {
			instance, err := client.Instance.Create().
				SetName(gofakeit.DomainName()).
				SetInstanceProvider(provider).
				SetTenantID(idx.MustNewID("tnntten")).
				SetLocationID(idx.PrefixedID(locationID)).
				Save(ctx)
			if err != nil {
				log.Fatalf("failed creating an instance: %v", err)
			}

			_, err = client.InstanceMetadata.Create().
				SetNamespace("seed.data").
				SetData(json.RawMessage(
					[]byte(fmt.Sprintf(
						`{"company": "%s", "contacts": ["%s", "%s"], "accessEnabled": %t}`,
						gofakeit.Company(),
						gofakeit.Name(),
						gofakeit.Name(),
						gofakeit.Bool(),
					)))).
				SetInstance(instance).
				Save(ctx)
			if err != nil {
				log.Fatalf("failed creating a instance metadata: %v", err)
			}
		}
	}

	srv, err := echox.NewServer(
		logger.Desugar(),
		echox.Config{
			Listen:              viper.GetString("server.listen"),
			ShutdownGracePeriod: viper.GetDuration("server.shutdown-grace-period"),
		},
		versionx.BuildDetails(),
	)
	if err != nil {
		logger.Fatal("failed to create server", zap.Error(err))
	}

	handler := graphapi.NewHandler(client)

	srv.AddHandler(handler)

	if err := srv.Run(); err != nil {
		logger.Fatal("failed to run server", zap.Error(err))
	}

}
