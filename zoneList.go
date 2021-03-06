package godo

import (
	"encoding/json"
	"os"

	"github.com/delphinus/go-dozens"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

var commandZoneList = cli.Command{
	Name:    "list",
	Aliases: []string{"l"},
	Usage:   "List up zones in dozens",
	Action:  doZoneList,
}

func doZoneList(c *cli.Context) error {
	zone, err := dozens.ZoneList(Config.Token)
	if err != nil {
		return errors.Wrap(err, "error in ZoneList")
	}
	if err := json.NewEncoder(os.Stdout).Encode(zone); err != nil {
		return errors.Wrap(err, "error in Encode")
	}
	return nil
}
