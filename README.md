# FLEcli
Multi-platform "Fast Log Entry"(FLE) processing tool.

["Fast Log Entry"](https://www.df3cb.com/fle/) is a nifty tool, written by DF3CB. It allows quick and efficient entry of radio amateur contacts made during DXpeditions, contests or SOTA/WWFF activation. It can export the listed contacts in ADIF format ready to be uploaded to eQSL/LOTW or submitted for SOTA and WWFF. The program is only availble as a Windows executable which leaves the Mac or Linux users (like me) in the cold.

This tool is intended to process FLE formated files on other platform than Windows. The FLE file format is described on https://www.df3cb.com/fle/documentation/

But is also the alibi to learn and explore GO.

## Usage

### Installing
TBD: Running on a Mac, Linux, windows. 

### Commands and options

```
Usage:
  FLEcli [command]

Available Commands:
  load        Loads and validates a FLE type shorthand logfile
  adif        Generates an ADIF file based on a FLE type shorthand logfile.
  csv         Generates a SOTA .csv file based on a FLE type shorthand logfile.
  help        Help about any command

General Flags:
  -v, --version         Displays the version for FLEcli
  -i, --input string    FLE formatted input file (mandatory)
      --interpolate     Interpolates the missing time entries.
      --config string   config file (default is $HOME/.FLEcli.yaml)
  -h, --help            help for FLEcli

Use "FLEcli [command] --help" for more information about a command.
------
FLEcli adif [flags]

Flags:
  -h, --help            help for adif
  -o, --output string   Output filename
      --overwrite       Overwrites the output file if it exisits
  -s, --sota            Generates a SOTA ready ADIF file.
  -w, --wwff            Generates a WWFF ready ADIF file.
-----
FLEcli csv [flags]

Flags:
  -h, --help            help for csv
  -o, --output string   Output filename
      --overwrite       Overwrites the output file if it exisits
```

### Example: validate the FLE syntax of a file

To load and validate the FLE formated file (.txt):

```
./FLEcli load -i=myActivation.txt
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

To generate a WWFF-ready ADIF file: 
```
./FLEcli adif -i ON4KJM@ONFF-025920200524.txt --wwff --interpolate
```
The `--wwff` indicates the adif flavour to produce.
You can use the `--sota` switch to generate an ADIF file containing SOTA details.
The switch can be used together with `--wwff`

As we didn't provide an output filename, the default output, `ON4KJM@ONFF-025920200524.adi` will be used.  


### Example: generate a SOTA csv file

To generate a CSV file that can be uploaded to https://www.sotadata.org.uk/ to report SOTA activations:

```
./FLEcli csv -i activation.txt --overwrite --interpolate
```
This command will generate `sotaActivation.csv` based on the `sotaActivation.txt` FLE logfile.
If the output file exists, it will be overwritten as the `--overwrite` flag has been specified.
The `--interpolate` flag will interpolate the missing non-entered times based on the first and the last entered time.


