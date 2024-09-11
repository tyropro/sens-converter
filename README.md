# go-sens-converter

This is a tool to convert sensitivities between:
 - Different games and
 - Different DPIs

## Usage

Run the program in a terminal with the `-h` flag.

The output should look like: 
```
$ sens-converter -h
Usage of sens-converter.exe:
  -b float
        Original DPI to convert from (default 800)
  -e    Prints the eDPI instead of the sensitivity
  -f string
        Game to convert from (default "undefined")
  -n float
        New DPI to convert to (default 800)
  -s float
        The sensitivity to convert (default 1)
  -t string
        Game to convert to (default "undefined")
```
