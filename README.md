# FLEcli
Multi-platform "Fast Log Entry"(FLE) processing tool.

![Go Build & Test](https://github.com/jmMeessen/FLEcli/workflows/Go%20Build%20&%20Test/badge.svg)
[![codecov](https://codecov.io/gh/on4kjm/FLEcli/branch/master/graph/badge.svg)](https://codecov.io/gh/on4kjm/FLEcli)

["Fast Log Entry"](https://www.df3cb.com/fle/) is a nifty tool, written by DF3CB. It allows quick and efficient entry of radio amateur contacts made during DXpeditions, contests or SOTA/WWFF/POTA activation. It can export the listed contacts in ADIF format ready to be uploaded to eQSL/LOTW or submitted for SOTA, POTA, and WWFF. The original program is only available as a Windows executable which leaves the Mac or Linux users (like me) in the cold.

The original FLE format was extented with release V1.4 to support POTA log files.

This tool is intended to process FLE formatted files on other platform than Windows. The FLE input file format is described on https://www.df3cb.com/fle/documentation/

------

If you like what I did and want to keep me going, you can "buy me a coffee" via the link below. 

[!["Buy Me A Coffee"](https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png)](https://www.buymeacoffee.com/on4kjm) 

------


## Installing
The tool is installed by downloading and unpacking the archive for your operating system.

A version is available for Windows, MacOS, Linux and ARM (RaspberryPi). 
Most of them are available as 32 or 64 bit applications.
A docker version is also available.

The releases can be downloaded from https://github.com/on4kjm/FLEcli/releases

For MacOS users, `homebrew` is the easiest installation method.

* add the Homebrew tap with `brew tap on4kjm/tap`. 
* install the application with `brew install flecli`.


Detailed installation instructions can be found [at this page](doc/install.md).



## Usage

Detailed explanations can be found [on this page](doc/usage.md).

Most common use cases are described hereafter as examples.

### Example: validate the FLE syntax of a file

To load and validate the FLE formatted file (`myActivation.txt`):

```
./FLEcli load myActivation.txt
```
This command will parse and display the QSOs in grid format. 
Fields that couldn't be successfully parsed are prefixed with "*". 
Parsing errors or doubts are listed at the end of the list.


### Example: generate an ADIF file

To generate an ADIF file based on `activation.txt`:

```
./FLEcli adif -i --overwrite activation.txt output/activation.adi
```
The second element after the input specifies the path and name of the output file.
If is omitted, the tool will generate a filename.
It is based on the input filename, just replacing the extension with `.adi`.

The `-o` or `--overwrite` flag indicates that, if the output file already exists, it should be overwritten. 

The `-i` or `--interpolate` flag will interpolate the missing non-entered times based on the first and the last entered time.

### Example: generate an ADIF file for WWFF upload

To generate a WWFF-ready ADIF file: 
```
./FLEcli adif -i --wwff ON4KJM@ONFF-025920200524.txt 
```
The `--wwff` (or `-w` in the short form) indicates the adif flavour to produce.
You can also use the `--sota` switch to generate an ADIF file containing SOTA details.
The switch can be used together with `--wwff`

As we didn't provide an output filename, the default output, `ON4KJM@ONFF-025920200524.adi` will be used.  

### Example: generate an ADIF file for POTA upload

To generate a POTA-ready ADIF file: 
```
./FLEcli adif -i --pota ON4KJM@ON-000120211224.txt 
```
The `--pota` (or `-p` in the short form) indicates the adif flavour to produce.

As we didn't provide an output filename, the default output, `ON4KJM@ON-000120211224.adi` will be used.  

A sample POTA FLE input file can be found in [the sample directory](test/data/sample_pota.txt).


### Example: generate a SOTA csv file

To generate a CSV file that can be uploaded to https://www.sotadata.org.uk/ to report SOTA activations:

```
./FLEcli csv --interpolate --overwrite sotaActivation.txt
```
or the short form
```
./FLEcli csv -i -o sotaActivation.txt
```
This command will generate `sotaActivation.csv` based on the `sotaActivation.txt` FLE logfile.
If the output file exists, it will be overwritten as the `--overwrite` flag has been specified.
The `--interpolate` flag will interpolate the missing non-entered times based on the first and the last entered time.


