# Usage

## Overview

```text
A Command Line "Fast Log Entry" (FLE) processor

Usage:
  FLEcli [command]

Available Commands:
  adif        Generates an ADIF file based on a FLE type shorthand logfile.
  csv         Generates a SOTA .csv file based on a FLE type shorthand logfile.
  help        Help about any command
  load        Loads and validates a FLE type shorthand logfile
  version     "version" will output the current build information

Flags:
      --config string   config file (default is $HOME/.FLEcli.yaml)
  -h, --help            help for FLEcli

Use "FLEcli [command] --help" for more information about a command.
```

## "LOAD" command

```text
Loads and validates a FLE type shorthand logfile

Usage:
  FLEcli load [flags] inputFile

Flags:
  -h, --help          help for load
  -i, --interpolate   Interpolates the missing time entries.

Global Flags:
      --config string   config file (default is $HOME/.FLEcli.yaml)
```

## "ADIF" command

```text
Generates an ADIF file based on a FLE type shorthand logfile.

Usage:
  FLEcli adif [flags] inputFile [outputFile]

Flags:
  -h, --help          help for adif
  -i, --interpolate   Interpolates the missing time entries.
  -o, --overwrite     Overwrites the output file if it exisits
  -p, --pota          Generates a POTA ready ADIF file.
  -s, --sota          Generates a SOTA ready ADIF file.
  -w, --wwff          Generates a WWFF ready ADIF file.

Global Flags:
      --config string   config file (default is $HOME/.FLEcli.yaml)
```

## "CSV" command

```text
Generates a SOTA .csv file based on a FLE type shorthand logfile.

Usage:
  FLEcli csv [flags] inputFile [outputFile]

Flags:
  -h, --help          help for csv
  -i, --interpolate   Interpolates the missing time entries.
  -o, --overwrite     Overwrites the output file if it exisits

Global Flags:
      --config string   config file (default is $HOME/.FLEcli.yaml)
```

## "VERSION" command

```text
"version" will output the current build information

Usage:
  FLEcli version [flags]

Flags:
  -d, --detailed   Prints the detailed version information
  -h, --help       help for version

Global Flags:
      --config string   config file (default is $HOME/.FLEcli.yaml)
```

The normal output looks like `FLEcli version: v0.1.2`. The detailled output gives additionaly the Git commit hash. the date and time of build and who built the release.
