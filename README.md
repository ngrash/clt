# clt - Command Line Template Engine

`clt` is a command line template engine. It renders a template to `stdout` for each line given on `stdin` by splitting those lines into values. It supports colors through ANSI escape sequences.

## Usage

```
$ clt -h
Usage: ./clt [OPTION]

  Input is read from stdin until EOF is reached.

  Available options:

  -h   print this help
  -s separator
       split input lines by separator (default " ")
  -t file
       read template from file
```

## Example

```
$ ls -lhd *
-rwxr-xr-x 1 nico nico 2.5M Jun  1 19:28 clt
drwxr-xr-x 2 nico nico 4.0K Jun  1 19:16 examples
-rw-r--r-- 1 nico nico 1.1K Jun  1 19:03 LICENSE.md
-rw-r--r-- 1 nico nico 2.6K Jun  1 19:28 main.go
-rw-r--r-- 1 nico nico 1.3K Jun  1 19:17 README.md

$ cat examples/ls.clt
- $8 is owned by $2
  size (bytes): $4

$ ls -lhd * | ./clt -t examples/ls.clt
- clt is owned by nico
  size (bytes): 2.5M

- examples is owned by nico
  size (bytes): 4.0K

- LICENSE.md is owned by nico
  size (bytes): 1.1K

- main.go is owned by nico
  size (bytes): 2.6K

- README.md is owned by nico
  size (bytes): 1.3K

```

## ANSI escape sequences

You can use ANSI escape sequences in the form of `\033[XXXm` where XXX are the usual parameters.

## License

See the [LICENSE](LICENSE.md) file for license rights and limitations (MIT).
