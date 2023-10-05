# ASCII Me
Simple command line tool to replace non-ASCII characters with ASCII ones in a string.

# What problem does this solve?
This tool intends to create a quick and easy way to copy text files and replace any
non-ASCII characters with ASCII substitutes.

# Installation
The most straight forward way of installing the executable is utilizing the go tool
chain. Assuming a properly configured go tool chain, you should be able to run
the following.

```bash
git clone git@github.com:EthanEFung/ascii-me.git
cd ascii-me
go install
ascii-me "“Hello, world”" # Notice the usage of curly double quotes
```
# Usage
If the above commands are run, the program returns some information about the first
non-ASCII character and prompts you to enter a ASCII character. Go ahead and insert
`"` and `enter`. At which you should be prompted yet again. Insert `"` and `enter` once
again. The program then prints to the std out `"Hello, world"`, which is exactly what
we wanted!

If there are multiple of the same non-ASCII character in the string, the program
will utilize the same ASCII character the user specified in the first encounter of the
non-ASCII character.

copy and paste the following in the terminal and run the executable replacing the
emoji with another character:
```
ascii-me "😕😕😕"
```

The executable also accepts an `input` and `output` flag.

```bash
ascii-me -input ~/path/to/text-file -output ./output-file.txt
```
The above command will slurp the input file and treat it as the string to copy.
Furthermore, because the `-output` flag was specified, a new file will be created at
the specified path, and the new string will be written to `-output`.

This tool can only consume either a string passed as the first parameter of the cli,
or consume an input file using the `-input` flag. If the `-input` flag is passed, it
will take priority over any stdin argument.

Likewise if a `-output` file path is specified, the tool will prioritize writing to
the `-output` file over printing to stdout.

# Disclaimer
This has not been tested yet, and was built to solve a very specific problem I had
at work. Use at your own discretion.
