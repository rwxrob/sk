package main

import (
	"fmt"

	"github.com/rwxrob/cmdtab"
)

// extra semicolons for safety depending on how users add to their bashrc
const shell = `
export _sk=$(which sk);

sk () {
  case "$1" in 
		psessions|plog|pinvoice|reserved|genid|bashrc|roll|email|income) $_sk $* ;;
		*) if [[ $1 == d ]]; then
		     shift;
			 fi
			 mdir=$($_sk d $*);
       if [[ $? == 0 ]]; then 
         echo $mdir;
         cd "$mdir";
       fi
       ;; 
  esac;
};

export -f sk;
complete -C sk sk;
`

func init() {
	_cmd := cmdtab.New("bashrc")
	_cmd.Summary = "print the recommended lines to add to a bashrc"
	_cmd.Usage = ""
	_cmd.Description = `
            Prints the lines to add to a bashrc (or simply eval) providing the following:

            * Adds the *sk* function
            * Adds tab completion *sk*
      `
	_cmd.Method = func(args []string) error {
		fmt.Println(shell)
		return nil
	}
}
