# FLEcli
Multi-platform "Fast Log Entry"(FLE) processing tool.

![Go Build & Test](https://github.com/jmMeessen/FLEcli/workflows/Go%20Build%20&%20Test/badge.svg)

["Fast Log Entry"](https://www.df3cb.com/fle/) is a nifty tool, written by DF3CB. It allows quick and efficient entry of radio amateur contacts made during DXpeditions, contests or SOTA/WWFF activation. It can export the listed contacts in ADIF format ready to be uploaded to eQSL/LOTW or submitted for SOTA and WWFF. The program is only available as a Windows executable which leaves the Mac or Linux users (like me) in the cold.

This tool is intended to process FLE formatted files on other platform than Windows. The FLE file format is described on https://www.df3cb.com/fle/documentation/

But is also the alibi to learn and explore GO.

## Usage

### Installing
The tool is installed by downloading and unpacking the archive for your operating system.

The releases are downloaded for https://github.com/on4kjm/FLEcli/releases

The application is available in the following packaging:
* `FLEcli_v..._macOS_64-bit.tar.gz` is the MacOS version.
* `FLEcli_v..._windows_32-bit.zip` is the Windows version. Only the 32 bit is proposed as it should work also on older Windows version.
* `FLEcli_v..._Linux_64-bit.tar.gz` is the 64 bits distribution for Linux.
* `FLEcli_v..._Linux_32-bit.tar.gz` is the 32 bits distribution for Linux
* `FLEcli_v..._Linux_arm.tar.gz` is the 32 bits ARM (instruction set 6) distribution (for RaspberryPi older models for example) 
* `FLEcli_v..._Linux_arm64.tar.gz` is the 64 bits ARM (instruction set 6) distribution

Detailed installation instructions can be found here (TODO).

**Important note:** at this stage, the MacOS binary is not signed. It will fail to load on the recent OS version unless a very permissive security setting is enable (which is not advised). Therefore, it is advised to run FLEcli as a Docker application (available on DockerHub). An alternative is to compile the application. Further explanation can be found here (TODO)

### Commands and options

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
------
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
-----
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
-----
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

### Example: validate the FLE syntax of a file

To load and validate the FLE formated file (.txt):

```
./FLEcli load -i=myActivation.txt
```
This command will parse and display the QSOs in grid format. 
Fields that couldn't be successfully parsed are prefixed with "*". 
Parsing errors or doubts are listed at the end of the list.


### Example: generate an ADIF file

To generate an ADIF file based on `activation.txt`:

```
./FLEcli adif -i activation.txt -o output/activation.adi --overwrite --interpolate
```
The `-o` (or the long form, `--output`) specifies the path and name of the output file.
If the flag and value are omitted, the tool will generate a filename.
It is based on the input filename, just replacing the extension with `.adi`.

The `--overwrite` flag indicates that, if the output file already exists, it should be overwritten. 

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


