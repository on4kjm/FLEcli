# Usage

## Overview
```
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
```
FLEcli load --help
Loads and validates a FLE type shorthand logfile

Usage:
  FLEcli load [flags]

Flags:
  -h, --help           help for load
  -i, --input string   FLE formatted input file (mandatory)
      --interpolate    Interpolates the missing time entries.

Global Flags:
      --config string   config file (default is $HOME/.FLEcli.yaml)
```


## "ADIF" command
```
FLEcli adif --help
Generates an ADIF file based on a FLE type shorthand logfile.

Usage:
  FLEcli adif [flags]

Flags:
  -h, --help            help for adif
  -i, --input string    FLE formatted input file (mandatory)
      --interpolate     Interpolates the missing time entries.
  -o, --output string   Output filename
      --overwrite       Overwrites the output file if it exisits
  -s, --sota            Generates a SOTA ready ADIF file.
  -w, --wwff            Generates a WWFF ready ADIF file.

Global Flags:
      --config string   config file (default is $HOME/.FLEcli.yaml)
```

## "CSV" command
```
FLEcli csv --help
Generates a SOTA .csv file based on a FLE type shorthand logfile.

Usage:
  FLEcli csv [flags]

Flags:
  -h, --help            help for csv
  -i, --input string    FLE formatted input file (mandatory)
      --interpolate     Interpolates the missing time entries.
  -o, --output string   Output filename
      --overwrite       Overwrites the output file if it exisits

Global Flags:
      --config string   config file (default is $HOME/.FLEcli.yaml)
```

## "VERSION" command
```
FLEcli version --help
"version" will output the current build information

Usage:
  FLEcli version [flags]

Flags:
  -d, --detailed   Prints the detailed version information
  -h, --help       help for version

Global Flags:
      --config string   config file (default is $HOME/.FLEcli.yaml)
```
The normal output looks like `FLEcli version: v0.0.1`. The detailled output gives additionaly the Git commit hash. the date and time of build and who built the release.