package cmd

import (
	"context"
	"encoding/json"
	"log"

	"entgo.io/ent/dialect"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.infratographer.com/x/echox"
	"go.infratographer.com/x/versionx"
	"go.uber.org/zap"

	_ "github.com/mattn/go-sqlite3"

	"go.infratographer.com/instance-api/internal/config"
	"go.infratographer.com/instance-api/internal/ent"
	"go.infratographer.com/instance-api/internal/graphapi"
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

	cOpts := []ent.Option{}

	if config.AppConfig.Logging.Debug {
		cOpts = append(cOpts,
			ent.Log(logger.Named("ent").Debugln),
			ent.Debug(),
		)
	}

	// Create an ent.Client with in-memory SQLite database.
	client, err := ent.Open(dialect.SQLite, "file:instance.db?cache=shared&_fk=1", cOpts...)

	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	// Run the automatic migration tool to create all schema resources.
	if err := client.Schema.Create(
		ctx,
	); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	provider, err := client.InstanceProvider.Create().
		SetName("Test Provider").
		Save(ctx)
	if err != nil {
		log.Fatalf("failed creating a provider: %v", err)
	}

	instance, err := client.Instance.Create().
		SetName("Test Instance").
		SetInstanceProvider(provider).
		SetTenantID(uuid.New()).
		Save(ctx)
	if err != nil {
		log.Fatalf("failed creating a provider: %v", err)
	}

	_, err = client.InstanceMetadata.Create().
		SetNamespace("instance-api.namespace").
		SetData(json.RawMessage([]byte(`{"name": "Nicole", "nested": {"field": "here"}}`))).
		SetInstance(instance).
		Save(ctx)
	if err != nil {
		log.Fatalf("failed creating a instance-api.namespace metadata: %v", err)
	}

	srv := echox.NewServer(
		logger.Desugar(),
		echox.Config{
			Listen:              viper.GetString("server.listen"),
			ShutdownGracePeriod: viper.GetDuration("server.shutdown-grace-period"),
		},
		versionx.BuildDetails(),
	)

	handler := graphapi.NewHandler(client)

	srv.AddHandler(handler)

	if err := srv.Run(); err != nil {
		logger.Fatal("failed to run server", zap.Error(err))
	}

}
