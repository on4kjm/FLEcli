# Installation

## Downloading

The application can be downloaded from the latest release page, found at https://github.com/on4kjm/FLEcli/releases.

The application is available in the following packaging:
* `FLEcli_v..._macOS_64-bit.tar.gz` is the MacOS version.
* `FLEcli_v..._windows_32-bit.zip` is the Windows version. Only the 32 bit is proposed as it should work also on older Windows version.
* `FLEcli_v..._Linux_64-bit.tar.gz` is the 64 bits distribution for Linux.
* `FLEcli_v..._Linux_32-bit.tar.gz` is the 32 bits distribution for Linux
* `FLEcli_v..._Linux_arm.tar.gz` is the 32 bits ARM (instruction set 6) distribution (for RaspberryPi older models for example) 
* `FLEcli_v..._Linux_arm64.tar.gz` is the 64 bits ARM (instruction set 6) distribution

After downloading the appropriate archive, you can validate its sha256 and compare it with the value in the `checksum.txt` file.

## Compiling
The application can also be compiled from sources if Golang is available. How to proceed is out of the scope of this notice.

## Installing

Unpack it in a convenient location. Depending on your Operating System, you can make it available from any location: 
* Adding the directory to the path for Windows
* Creating a (permanent) alias on Linux or MacOS
* Moving the executable to a directory that is in the path, such as `/usr/local/bin` for Linux or MacOs.

### HomeBrew
Currently there is no Homebrew installation available.

### Docker
The application is also available as a docker container. 
It can be started with: `docker run --rm -i --user $(id -u):$(id -g) -v $(pwd):/FLEcli_data on4kjm/flecli <FLEcli command>`. If no command is specified, help is displayed.
To avoid typing the whole command, an alias is recomended as described in the [docker specific instructions](../docker/README.md).

**Important note:** when specifying the path of a file (input or output), it must be relative to the directory the container was started in.

### MacOS

**WARNING:** at this stage, the MacOS binary is not signed. It will fail to load on the recent OS version unless a very permissive security setting is enable (which is not advised). 

Please find [here a detailed procedure](install_mac.md) to use FLEcli on a Mac despite the missing signature.