package db

import (
	"context"
	"fmt"
	"sync"

	sqle "github.com/dolthub/go-mysql-server"
	"github.com/dolthub/go-mysql-server/memory"
	"github.com/dolthub/go-mysql-server/server"
	"github.com/dolthub/go-mysql-server/sql"
	"github.com/dolthub/go-mysql-server/sql/mysql_db"
	"github.com/dolthub/go-mysql-server/sql/types"
	"github.com/dolthub/vitess/go/mysql"
)

var (
	dbName    = "utTestGenesis"
	tableName = "mytable"
	dbAddress = "localhost"
	dbPort    = 3306
)

func StartMysqlServer() {

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		db := memory.NewDatabase(dbName)
		db.BaseDatabase.EnablePrimaryKeyIndexes()

		pro := memory.NewDBProvider(db)
		session := memory.NewSession(sql.NewBaseSession(), pro)
		ctx := sql.NewContext(context.Background(), sql.WithSession(session))

		tableName = "k8s_configs"

		table := memory.NewTable(db, tableName, sql.NewPrimaryKeySchema(sql.Schema{
			{Name: "created_time", Type: types.Datetime, Nullable: true, Source: tableName, PrimaryKey: false},
			{Name: "updated_time", Type: types.Datetime, Nullable: true, Source: tableName, PrimaryKey: false},
			{Name: "id", Type: types.Text, Nullable: false, Source: tableName, PrimaryKey: true},
			{Name: "owner_id", Type: types.Text, Nullable: true, Source: tableName, PrimaryKey: false},
			{Name: "owner_name", Type: types.Text, Nullable: true, Source: tableName, PrimaryKey: false},
			{Name: "owner_type", Type: types.Text, Nullable: true, Source: tableName, PrimaryKey: false},
			{Name: "cluster_name", Type: types.Text, Nullable: true, Source: tableName, PrimaryKey: false},
			{Name: "tenant_id", Type: types.Text, Nullable: true, Source: tableName, PrimaryKey: false},
			{Name: "tenant_name", Type: types.Text, Nullable: true, Source: tableName, PrimaryKey: false},
			{Name: "workspace", Type: types.Text, Nullable: true, Source: tableName, PrimaryKey: false},
			{Name: "region", Type: types.Text, Nullable: true, Source: tableName, PrimaryKey: false},
			{Name: "available_zone", Type: types.Text, Nullable: true, Source: tableName, PrimaryKey: false},
			{Name: "data_center", Type: types.Text, Nullable: true, Source: tableName, PrimaryKey: false},
			{Name: "config_info", Type: types.Text, Nullable: true, Source: tableName, PrimaryKey: false},
			{Name: "typ", Type: types.Text, Nullable: true, Source: tableName, PrimaryKey: false},
			{Name: "nebula_cluster_name", Type: types.Text, Nullable: true, Source: tableName, PrimaryKey: false},
			{Name: "tenant_name", Type: types.Text, Nullable: true, Source: tableName, PrimaryKey: false},
		}), db.GetForeignKeyCollection())

		db.AddTable(tableName, table)

		engine := sqle.NewDefault(pro)
		//session := memory.NewSession(sql.NewBaseSession(), pro)
		//ctx := sql.NewContext(context.Background(), sql.WithSession(session))
		ctx.SetCurrentDatabase(dbName)

		config := server.Config{
			Protocol: "tcp",
			Address:  fmt.Sprintf("%s:%d", dbAddress, dbPort),
		}

		sb := func(ctx context.Context, conn *mysql.Conn, addr string) (sql.Session, error) {
			host := ""
			user := ""
			mysqlConnectionUser, ok := conn.UserData.(mysql_db.MysqlConnectionUser)
			if ok {
				host = mysqlConnectionUser.Host
				user = mysqlConnectionUser.User
			}

			client := sql.Client{Address: host, User: user, Capabilities: conn.Capabilities}
			baseSession := sql.NewBaseSessionWithClientServer(addr, client, conn.ConnectionID)
			return memory.NewSession(baseSession, pro), nil
		}
		s, err := server.NewServer(config, engine, sb, nil)
		if err != nil {
			panic(err)
		}

		//ctx := sql.NewEmptyContext()
		/*engine := sqle.NewDefault(
		memory.NewDBProvider(
			createTestDatabase(ctx),
		))*/

		// This variable may be found in the "users_example.go" file. Please refer to that file for a walkthrough on how to
		// set up the "mysql" database to allow user creation and user checking when establishing connections. This is set
		// to false for this example, but feel free to play around with it and see how it works.
		/*if enableUsers {
			if err := enableUserAccounts(ctx, engine); err != nil {
				panic(err)
			}
		}*/

		/*config := server.Config{
			Protocol: "tcp",
			Address:  fmt.Sprintf("%s:%d", dbAddress, dbPort),
		}
		//s, err := server.NewDefaultServer(config, engine)

		s, err := server.NewServer(config, engine, func(ctx context.Context, conn *mysql.Conn, addr string) (sql.Session, error) {
			builder, err := server.DefaultSessionBuilder(ctx, conn, addr)
			if err != nil {
				return nil, err
			}
			return memory.NewSession(builder.(*sql.BaseSession), pro), nil
		}, nil)*/

		if err = s.Start(); err != nil {
			panic(err)
		}
	}()
	fmt.Println("Start Mysql mock server success.")
	wg.Wait()
}
