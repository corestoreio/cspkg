package myreplicator

import (
	"context"
	"flag"
	"fmt"
	"github.com/corestoreio/csfw/storage/csdb"
	"github.com/pingcap/check"
	uuid "github.com/satori/go.uuid"
	"github.com/siddontang/go-mysql/client"
	"github.com/siddontang/go-mysql/mysql"
	"net"
	"os"
	"strconv"
	"sync"
	"testing"
	"time"
)

var testOutputLogs = flag.Bool("out", false, "output binlog event")

func TestBinLogSyncer(t *testing.T) {
	check.TestingT(t)
}

type testSyncerSuite struct {
	bls *BinlogSyncer
	con *client.Conn

	wg sync.WaitGroup

	flavor string
}

var _ = check.Suite(&testSyncerSuite{})

func (t *testSyncerSuite) SetUpSuite(c *check.C) {
}

func (t *testSyncerSuite) TearDownSuite(c *check.C) {
}

func (t *testSyncerSuite) SetUpTest(c *check.C) {
}

func (t *testSyncerSuite) TearDownTest(c *check.C) {
	defer os.RemoveAll("./testdata/var")

	if t.bls != nil {
		t.bls.Close()
		t.bls = nil
	}

	if t.con != nil {
		t.con.Close()
		t.con = nil
	}
}

func (t *testSyncerSuite) testExecute(c *check.C, query string) {
	_, err := t.con.Execute(query)
	c.Assert(err, check.IsNil)
}

func (t *testSyncerSuite) testSync(c *check.C, s *BinlogStreamer) {
	t.wg.Add(1)
	go func() {
		defer t.wg.Done()

		if s == nil {
			return
		}

		eventCount := 0
		for {
			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			e, err := s.GetEvent(ctx)
			cancel()

			if err == context.DeadlineExceeded {
				eventCount += 1
				return
			}

			c.Assert(err, check.IsNil)

			if *testOutputLogs {
				e.Dump(os.Stdout)
				os.Stdout.Sync()
			}
		}
	}()

	//use mixed format
	t.testExecute(c, "SET SESSION binlog_format = 'MIXED'")

	str := `DROP TABLE IF EXISTS test_myreplicator`
	t.testExecute(c, str)

	str = `CREATE TABLE IF NOT EXISTS test_myreplicator (
	         id BIGINT(64) UNSIGNED  NOT NULL AUTO_INCREMENT,
	         str VARCHAR(256),
	         f FLOAT,
	         d DOUBLE,
	         de DECIMAL(10,2),
	         i INT,
	         bi BIGINT,
	         e enum ("e1", "e2"),
	         b BIT(8),
	         y YEAR,
	         da DATE,
	         ts TIMESTAMP,
	         dt DATETIME,
	         tm TIME,
	         t TEXT,
	         bb BLOB,
	         se SET('a', 'b', 'c'),
	      PRIMARY KEY (id)
	       ) ENGINE=InnoDB DEFAULT CHARSET=utf8`

	t.testExecute(c, str)

	//use row format
	t.testExecute(c, "SET SESSION binlog_format = 'ROW'")

	t.testExecute(c, `INSERT INTO test_myreplicator (str, f, i, e, b, y, da, ts, dt, tm, de, t, bb, se)
		VALUES ("3", -3.14, 10, "e1", 0b0011, 1985,
		"2012-05-07", "2012-05-07 14:01:01", "2012-05-07 14:01:01",
		"14:01:01", -45363.64, "abc", "12345", "a,b")`)

	id := 100

	if t.flavor == mysql.MySQLFlavor {
		t.testExecute(c, "SET SESSION binlog_row_image = 'MINIMAL'")

		t.testExecute(c, fmt.Sprintf(`INSERT INTO test_myreplicator (id, str, f, i, bb, de) VALUES (%d, "4", -3.14, 100, "abc", -45635.64)`, id))
		t.testExecute(c, fmt.Sprintf(`UPDATE test_myreplicator SET f = -12.14, de = 555.34 WHERE id = %d`, id))
		t.testExecute(c, fmt.Sprintf(`DELETE FROM test_myreplicator WHERE id = %d`, id))
	}

	t.wg.Wait()
}

func (t *testSyncerSuite) setupTest(c *check.C, flavor string) {

	t.flavor = flavor

	var err error
	if t.con != nil {
		t.con.Close()
	}

	dsn, err := csdb.GetParsedDSN()
	if err != nil {
		c.Skip(fmt.Sprintf("Failed to get DSN from env %q with %s", csdb.EnvDSN, err))
	}

	t.con, err = client.Connect(dsn.Addr, dsn.User, dsn.Passwd, dsn.DBName)
	if err != nil {
		c.Skip(err.Error())
	}

	// _, err = t.c.Execute("CREATE DATABASE IF NOT EXISTS test")
	// c.Assert(err, check.IsNil)

	_, err = t.con.Execute("USE test")
	c.Assert(err, check.IsNil)

	if t.bls != nil {
		t.bls.Close()
	}
	host, port, _ := net.SplitHostPort(dsn.Addr)
	po, err := strconv.ParseUint(port, 10, 32)
	if err != nil {
		c.Fatal(err)
	}

	cfg := BinlogSyncerConfig{
		ServerID: 100,
		Flavor:   flavor,
		Host:     host,
		Port:     uint16(po),
		User:     dsn.User,
		Password: dsn.Passwd,
	}

	t.bls = NewBinlogSyncer(&cfg)
}

func (t *testSyncerSuite) testPositionSync(c *check.C) {
	//get current master binlog file and position
	r, err := t.con.Execute("SHOW MASTER STATUS")
	c.Assert(err, check.IsNil)
	binFile, _ := r.GetString(0, 0)
	binPos, _ := r.GetInt(0, 1)

	s, err := t.bls.StartSync(csdb.MasterStatus{File: binFile, Position: uint(binPos)})
	c.Assert(err, check.IsNil)

	// Test re-sync.
	time.Sleep(100 * time.Millisecond)
	t.bls.con.SetReadDeadline(time.Now().Add(time.Millisecond))
	time.Sleep(100 * time.Millisecond)

	t.testSync(c, s)
}

func (t *testSyncerSuite) TestMysqlPositionSync(c *check.C) {
	t.setupTest(c, mysql.MySQLFlavor)
	t.testPositionSync(c)
}

func (t *testSyncerSuite) TestMysqlGTIDSync(c *check.C) {
	t.setupTest(c, mysql.MySQLFlavor)

	r, err := t.con.Execute("SELECT @@gtid_mode")
	c.Assert(err, check.IsNil)
	modeOn, _ := r.GetString(0, 0)
	if modeOn != "ON" {
		c.Skip("GTID mode is not ON")
	}

	r, err = t.con.Execute("SHOW GLOBAL VARIABLES LIKE 'SERVER_UUID'")
	c.Assert(err, check.IsNil)

	var masterUuid uuid.UUID
	if s, _ := r.GetString(0, 1); len(s) > 0 && s != "NONE" {
		masterUuid, err = uuid.FromString(s)
		c.Assert(err, check.IsNil)
	}

	set, _ := mysql.ParseMysqlGTIDSet(fmt.Sprintf("%s:%d-%d", masterUuid.String(), 1, 2))

	s, err := t.bls.StartSyncGTID(set)
	c.Assert(err, check.IsNil)

	t.testSync(c, s)
}

func (t *testSyncerSuite) TestMariadbPositionSync(c *check.C) {
	c.Skip("todo MariaDB Tests")

	t.setupTest(c, mysql.MariaDBFlavor)

	t.testPositionSync(c)
}

func (t *testSyncerSuite) TestMariadbGTIDSync(c *check.C) {
	c.Skip("todo MariaDB Tests")

	t.setupTest(c, mysql.MariaDBFlavor)

	// get current master gtid binlog pos
	r, err := t.con.Execute("SELECT @@gtid_binlog_pos")
	c.Assert(err, check.IsNil)

	str, _ := r.GetString(0, 0)
	set, _ := mysql.ParseMariadbGTIDSet(str)

	s, err := t.bls.StartSyncGTID(set)
	c.Assert(err, check.IsNil)

	t.testSync(c, s)
}

func (t *testSyncerSuite) TestMysqlSemiPositionSync(c *check.C) {
	t.setupTest(c, mysql.MySQLFlavor)

	t.bls.cfg.SemiSyncEnabled = true

	t.testPositionSync(c)
}

func (t *testSyncerSuite) TestMysqlBinlogCodec(c *check.C) {
	t.setupTest(c, mysql.MySQLFlavor)

	t.testExecute(c, "RESET MASTER")

	var wg sync.WaitGroup
	wg.Add(1)
	defer wg.Wait()

	go func() {
		defer wg.Done()

		t.testSync(c, nil)

		t.testExecute(c, "FLUSH LOGS")

		t.testSync(c, nil)
	}()

	if err := os.RemoveAll("./testdata/var"); err != nil {
		c.Error(err)
	}

	err := t.bls.StartBackup("./testdata/var", csdb.MasterStatus{Position: uint(0)}, 2*time.Second)
	if err != nil {
		c.Fatalf("%+v", err)
	}
	c.Assert(err, check.IsNil)

	p := NewBinlogParser()

	f := func(e *BinlogEvent) error {
		if *testOutputLogs {
			e.Dump(os.Stdout)
			os.Stdout.Sync()
		}
		return nil
	}

	err = p.ParseFile("./var/mysql.000001", 0, f)
	c.Assert(err, check.IsNil)

	err = p.ParseFile("./var/mysql.000002", 0, f)
	c.Assert(err, check.IsNil)
}