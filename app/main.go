package main

import (
	"github.com/fanke15/hftoc/app/cmd"
	"github.com/fanke15/hftoc/pkg/lib/bolt"
	"github.com/fanke15/hftoc/pkg/lib/conf"
	"github.com/fanke15/hftoc/pkg/lib/log"
)

func main() {
	/**启用附加服务**/
	conf.New()
	log.New()
	bolt.New()

	/**启动应用**/
	cmd.New()
}
