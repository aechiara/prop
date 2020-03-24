# Prop - Command-Line / sed like properties editor


## Goal

The Goal for this project is to have a fast, easy to use and reliable way to check, edit and add values to a properties file

## Binaries

The binaries can be downloaded form the URLs bellow

### Version 0.8.1
- [Linux](https://storage.googleapis.com/cloudbackup-apps/prop-linux-0.8.1.zip)
- [MacOS](https://storage.googleapis.com/cloudbackup-apps/prop-macos-0.8.1.zip)
- [Windows](https://storage.googleapis.com/cloudbackup-apps/prop-windows-0.8.1.zip)


## Building

clone the repository and run

```sh
$ make
```

### Building for Other OSs 

If you run make on your platform, a native binary will be generated for your platform, however, it is possible to build for another platforms, let's say you are on a Linux Machine, and want to build a Windows Executable, you can run:

```sh
$ make windows
```

and a ```prop.exe``` file will be generated

Options are:
- linux
- windows
- macos