// Package filterflags implements command line flags to set up a filter
package filterflags

import (
	"github.com/ncw/rclone/fs/config/flags"
	"github.com/ncw/rclone/fs/filter"
	"github.com/ncw/rclone/fs/rc"
	"github.com/spf13/pflag"
)

// Options set by command line flags
var (
	Opt = filter.DefaultOpt
)

// Reload the filters from the flags
func Reload() (err error) {
	filter.Active, err = filter.NewFilter(&Opt)
	return err
}

// AddFlags adds the non filing system specific flags to the command
func AddFlags(flagSet *pflag.FlagSet) {
	rc.AddOptionReload("filter", &Opt, Reload)
	flags.BoolVarP(flagSet, &Opt.DeleteExcluded, "delete-excluded", "", false, "Delete files on dest excluded from sync")
	flags.StringArrayVarP(flagSet, &Opt.FilterRule, "filter", "f", nil, "Add a file-filtering rule")
	flags.StringArrayVarP(flagSet, &Opt.FilterFrom, "filter-from", "", nil, "Read filtering patterns from a file")
	flags.StringArrayVarP(flagSet, &Opt.ExcludeRule, "exclude", "", nil, "Exclude files matching pattern")
	flags.StringArrayVarP(flagSet, &Opt.ExcludeFrom, "exclude-from", "", nil, "Read exclude patterns from file")
	flags.StringVarP(flagSet, &Opt.ExcludeFile, "exclude-if-present", "", "", "Exclude directories if filename is present")
	flags.StringArrayVarP(flagSet, &Opt.IncludeRule, "include", "", nil, "Include files matching pattern")
	flags.StringArrayVarP(flagSet, &Opt.IncludeFrom, "include-from", "", nil, "Read include patterns from file")
	flags.StringArrayVarP(flagSet, &Opt.FilesFrom, "files-from", "", nil, "Read list of source-file names from file")
	flags.FVarP(flagSet, &Opt.MinAge, "min-age", "", "Only transfer files older than this in s or suffix ms|s|m|h|d|w|M|y")
	flags.FVarP(flagSet, &Opt.MaxAge, "max-age", "", "Only transfer files younger than this in s or suffix ms|s|m|h|d|w|M|y")
	flags.FVarP(flagSet, &Opt.MinSize, "min-size", "", "Only transfer files bigger than this in k or suffix b|k|M|G")
	flags.FVarP(flagSet, &Opt.MaxSize, "max-size", "", "Only transfer files smaller than this in k or suffix b|k|M|G")
	flags.BoolVarP(flagSet, &Opt.IgnoreCase, "ignore-case", "", false, "Ignore case in filters (case insensitive)")
	//cvsExclude     = BoolP("cvs-exclude", "C", false, "Exclude files in the same way CVS does")
}
