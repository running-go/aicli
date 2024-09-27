package cmd

import (
	"fmt"
	"github.com/morikuni/aec"
	"runtime"
)

// printLogo prints an ASCII logo, which was generated with figlet
func printLogo() {
	figletColoured := aec.YellowF.Apply(figletStr)
	if runtime.GOOS == "windows" {
		figletColoured = aec.YellowF.Apply(figletStr)
	}
	fmt.Printf(figletColoured)
}

const figletStr = `  ___                   _____           ____
    mm      mmmmmm      mmmm   mm         mmmmmm  
   ####     ""##""    ##""""#  ##         ""##""  
   ####       ##     ##"       ##           ##    
  ##  ##      ##     ##        ##           ##    
  ######      ##     ##m       ##           ##    
 m##  ##m   mm##mm    ##mmmm#  ##mmmmmm   mm##mm  
 ""    ""   """"""      """"   """"""""   """"""
`
