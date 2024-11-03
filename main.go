package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"strings"
	"time"
)

const (
	// USAGE
	USAGE = `Usage: watch [OPTIONS]... command`
	// HELP
	HELP = `A command line utility to execute commands on file updates.

Usage: watch [OPTIONS]... command

  -f    comma seperated lists of files to listen for updates.
  -h    list help option.
  -i    interval between checks. default 100ms.
  -s    turn off update messages.`
	// START
	START = `
 __          __   _       _     _____
 \ \        / /  | |     | |   |  __ \
  \ \  /\  / /_ _| |_ ___| |__ | |  | | ___   __ _
   \ \/  \/ / _  | __/ __|  _ \| |  | |/ _ \ / _  |
    \  /\  / (_| | || (__| | | | |__| | (_) | (_| |
     \/  \/ \__,_|\__\___|_| |_|_____/ \___/ \__, |
                                              __/ |
                                             |___/

          -                                -
           --                            --
           =-=-                        -=-=
           ---==                      ==---
           -==-=-                    -=-==-
           -==--=--                --=--==-
            ++=====--            --=====++
            +++==*==--------------==*==+++
            -++=-=*--=-=------=-=-=*=-=++-
             ++=-====================-=++
             ++--=========**=========--++
             ++---=*===*======*===*=---++
             -+=-=*--=**==++==**=--*=-=+-
              --===-==+-==**==-+==-===--
              --==---=+-=-==-=-+=---==--
              -=*==----=------=----==*=-
             -=====-------==-------=====-
              ---=====--======--=====---
              ----====--======--====----
                ---==--========--==---
                 --==--+**++**+--==-- 
                ---==--+++==+++--==---
                --=======*==*=======--
                =-======------======-=
               -=-=++++=------=++++=-=-
               =====++++======++++=====
           *#*--==-=-====-==-====-=-==--*#*
            ==------==-++++++++-==------==
            -=-==**=----=++++=----=**==-=-
              ----=----------------=---- 
                 ----===-=**=-===----
                     ----====----
`
)

// Program entry point
func main() {
	fFlag := flag.String("f", "", "files to watch. accepts multiple files or directory, each seperated by a comma.")
	hFlag := flag.Bool("h", false, "help")
	sFlag := flag.Bool("s", false, "silence update messages")
	iFlag := flag.Int("i", 100, "check intervals")

	flag.Parse() // Parse all defined flags

	if *hFlag {
		fmt.Println(HELP)
		os.Exit(0)
	}

	cmdArg := flag.Args()

	if len(cmdArg) == 0 {
		fmt.Println(fmt.Sprintf("No command specified.\n%s", USAGE))
		os.Exit(2)
	}

	var files []string

	if *fFlag != "" {
		files = strings.Split(*fFlag, ",")
	} else {
		// Get current working directory if files are left empty
		exPath, err := os.Executable()

		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		files = []string{path.Dir(exPath)}
	}

	if len(files) > 10 {
		fmt.Println("Too many files to watch (maximum is 10)")
		os.Exit(2)
	}

	fileStats := make(map[string]*FileStat)

	for _, file := range files {
		stat, err := checkFile(file)
		if err != nil {
			fmt.Printf("Error checking file %s: %v\n", file, err)
			os.Exit(2)
		}
		fileStats[file] = stat
	}

	ticker := time.NewTicker(time.Duration(*iFlag) * time.Millisecond)
	defer ticker.Stop()
	logger := NewLogger(*sFlag)

  logger.LogStartup(files, cmdArg[0], cmdArg[1:], time.Duration(*iFlag))
	logger.LogCommandStart(cmdArg[0], cmdArg[1:])
	out, err := execute(cmdArg)
	logger.LogCommandOutput(err == nil, string(out))
	logger.LogCommandEnd(err == nil, time.Since(time.Now()))

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	// Listen for ticker updates
	// Rerun command if timestamp changes
	for range ticker.C {
		var changed = false

		for _, file := range files {
			curStat, err := checkFile(file)

			if err != nil {
				fmt.Printf("Error checking file %s: %v\n", file, err)
				continue
			}

			prevStat := fileStats[file]

			// Check if both file and if the the file has been modified at any point
			if curStat.exists != prevStat.exists ||
				(curStat.exists && !curStat.modified.Equal(prevStat.modified)) {

				logger.LogFileChange(file, curStat.modified)

				changed = true
				fileStats[file] = curStat
			}
		}

		if changed {
			startTime := time.Now()
			logger.LogCommandStart(cmdArg[0], cmdArg[1:])

			out, err := execute(cmdArg)

			logger.LogCommandOutput(err == nil, string(out))
			logger.LogCommandEnd(err == nil, time.Since(startTime))

			if err != nil {
				continue
			}
		}
	}
}
