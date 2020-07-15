# FLEcli
Multi-platform "Fast Log Entry"(FLE) processing tool.

["Fast Log Entry"](https://www.df3cb.com/fle/) is a nifty tool, written by DF3CB. It allows quick and efficient entry of radio amateur contacts made during DXpeditions, contests or SOTA/WWFF activation. It can export the listed contacts in ADIF format ready to be uploaded to eQSL/LOTW or submitted for SOTA and WWFF. The program is only availble as a Windows executable which leaves the Mac or Linux users (like me) in the cold.

This tool is intended to process FLE formated files on other platform than Windows. The FLE file format is described on https://www.df3cb.com/fle/documentation/

But is also the alibi to learn and explore GO.

## Usage

### Installing
Running on a Mac, Linux, windows

### Commands and options

```
Usage:
  FLEcli [command]

Available Commands:
  adif        Generates an ADIF file based on a FLE type shorthand logfile.
  csv         Generates a SOTA .csv file based on a FLE type shorthand logfile.
  help        Help about any command
  load        Loads and validates a FLE type shorthand logfile

Flags:
      --config string   config file (default is $HOME/.FLEcli.yaml)
  -h, --help            help for FLEcli
  -i, --input string    FLE formatted input file (mandatory)
      --interpolate     Interpolates the missing time entries.
  -v, --version         version for FLEcli

Use "FLEcli [command] --help" for more information about a command.
````

### Example: validate the FLE syntax

To load and validate the FLE formated file (.txt):

```
./FLEcli load -i=ON4KJM@ONFF-025920200524.txt
```
This command will parse and display the QSOs in grid format. 
Fields that couldn't be succesfuly parsed are prefixed with "*". 
Parsing errors or doubts are listed at the end of the list.


### Example: generate an ADIF file

To generate an ADIF file based on `activation.txt`:

```
./FLEcli adif -i activation.txt -o output/activation.adi --overwrite --interpolate
```
The `-o` (or the long form, `--output`) specifies the path and name of the output file.
If the flag and value are omitted, the tool will generate a filename.
It is based on the input filename, just replacing the extention with `.adi`.

The `--overwrite` flag indicates that, if the output file already exsist, it should be overwritten. 

The `--interpolate` flag will interpolate the missing non-entered times based on the first and the last entered time.

### Example: generate an ADIF file for WWFF upload

To generate a WWFF ready ADIF file: 
```
./FLEcli adif -i activation.txt -o output/activation.adi --wwff --interpolate
```
The `--wwff` indicates the adif flavour to produce.


### Example: generate a SOTA csv file

Display all the options
