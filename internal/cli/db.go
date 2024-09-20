package cli

import (
	"strings"

	"github.com/xpzouying/go-cmd-project-template/internal/model"

	"github.com/pkg/errors"
)

func initDB(dbConn string) error {

	dbType, dbDsn, err := parseDBConnStr(dbConn)
	if err != nil {
		return err
	}

	cliLogger.Infof("init db: type=%s, dsn=%s", dbType, dbDsn)

	if err := model.InitDB(model.DBType(dbType), dbDsn); err != nil {
		return errors.Wrap(err, "failed to init db")
	}
	return nil
}

func parseDBConnStr(dbConnStr string) (dbType, dbDsn string, err error) {

	// 使用 SplitN 将字符串分割成两部分，限制为只分割一次，确保只分离协议和后续部分
	parts := strings.SplitN(dbConnStr, "://", 2)
	if len(parts) != 2 {
		// 如果没有找到预期的分割，返回空字符串
		err = errors.New("invalid db connection string")
		return
	}

	dbType = parts[0]
	dbDsn = parts[1]
	return
}
