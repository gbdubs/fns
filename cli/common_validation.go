package cli

import "github.com/gbdubs/fns/util"

func validateInteractiveAndDryrunNotBothSet(errs *util.ErrList, i bool, d bool) {
	if i && d {
		errs.Appendf("Both the --dryrun and the --interactive flags are set, which doesn't make sense. Dryrun means no work will be done, just printed out. Interactive will similarly print out the set of work to be done, and give you a prompt of whether to proceed. Please pick one of these options and proceed.")
	}
}
