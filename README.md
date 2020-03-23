# Prop - Command-Line / sed like properties editor


## Goal

The Goal for this project is to have a fast, easy to use and reliable way to check, edit and add values to a properties file


## Building

- clone the repository and run

```sh
$ make
```

### Building for Other SOs 

If you run make on your platform, a native binary will be generated for your platform, however, it is possible to build for another platforms, let's say you are on a Linux Machine, and want to build a Windows Executable, you can run:

```sh
$ make windows
```

and a ```prop.exe``` file will be generated

Options are:
- linux
- windows
- macos