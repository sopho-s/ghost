# What is ghost

Ghost is a software made to hide files that would normally be flagged as suspicious so that they can be transferred to other devices without being flagged

# How to use ghost

## Linux

```shell
./ghost <OPTIONS> -in <INPUT FILE>
```

# Options

- -r: Revive, turns the ghosted file back to its original form (spooky)
- -f: Foil, makes a foil ghost, these ghosts have wrappers that do not require ghost to unwrap, they will simply unwrap themselves upon execution, but will need a second execution to then run the unwrapped executable
- -ct: Crypt type, decides the method that ghost will use to hide the file

# To add

- Directly load exes into memory
- Exe wrapper so that ghost is not needed on the other system, exe can simply be run and will be automatically revived
- Exe splice and scramble
