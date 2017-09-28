# `bsort` - like `sort` but keep the first line as header
I use [`sort`](http://man7.org/linux/man-pages/man1/sort.1.html) almost daily at work, but I always hated that if I want to sort something that has a header, the header gets sorted as well. Thus, `bsort` was born.

It is a very simple program: you can only use it by piping to it and only the first line is considered as the header.

## Usage
`bsort` only works by reading from `stdin` and then sub-executing the regular `sort` on all the lines but the first one. It doesn't have any parameters of its own, so any additional parameters will be sent to `sort`.

If you want to sort one file then you can do:
```sh
< path/to/file.txt bsort
```

If you want to reverse sort a file by the fourth column in a numerical fashion, you can do:
```sh
< path/to/file.txt bsort -rnk4
```

If you want to sort many files:
```sh
cat path/to/file1.txt path/to/file2.txt | bsort
```

Given that parameters are pass as they are to `sort` you could very well say _hey, I'm going to do `bsort -r file` instead of piping to it_: nothing stops you of doing that, but it will fail. You have been warned.

## Installation
If you have a running Go environment, just do

```
go install github.com/inkel/bsort
```

If you prefer instead downloading precompiled binaries, you can find them in the [releases](https://github.com/inkel/bsort/releases) section. Just download the version you're looking for, place it in your `$PATH` and make it executable. Enjoy!

## Example
Consider having the following:

```
$ cat example.txt
host        ip          heap.percent ram.percent load node.role master name
10.0.22.28  10.0.22.28            11          74 0.00 -         m      es-iot-master-4
10.0.12.201 10.0.12.201           45          73 0.00 -         m      es-iot-master-3
10.0.12.4   10.0.12.4             12          74 0.00 -         m      es-iot-master-0
10.0.32.71  10.0.32.71            69          74 1.01 -         *      es-iot-master-2
10.0.22.22  10.0.22.22            53          73 0.07 -         m      es-iot-master-1
10.0.32.29  10.0.32.29            13          73 0.00 -         m      es-iot-master-5
```

If we want to sort by `name` we can do:

```
$ < example.txt bsort -k8
host        ip          heap.percent ram.percent load node.role master name
10.0.12.4   10.0.12.4             12          74 0.00 -         m      es-iot-master-0
10.0.22.22  10.0.22.22            53          73 0.07 -         m      es-iot-master-1
10.0.32.71  10.0.32.71            69          74 1.01 -         *      es-iot-master-2
10.0.12.201 10.0.12.201           45          73 0.00 -         m      es-iot-master-3
10.0.22.28  10.0.22.28            11          74 0.00 -         m      es-iot-master-4
10.0.32.29  10.0.32.29            13          73 0.00 -         m      es-iot-master-5
```

If we wanted to do the same with the regular `sort` utility, this would be the result:

```
$ < example.txt sort -k8
10.0.12.4   10.0.12.4             12          74 0.00 -         m      es-iot-master-0
10.0.22.22  10.0.22.22            53          73 0.07 -         m      es-iot-master-1
10.0.32.71  10.0.32.71            69          74 1.01 -         *      es-iot-master-2
10.0.12.201 10.0.12.201           45          73 0.00 -         m      es-iot-master-3
10.0.22.28  10.0.22.28            11          74 0.00 -         m      es-iot-master-4
10.0.32.29  10.0.32.29            13          73 0.00 -         m      es-iot-master-5
host        ip          heap.percent ram.percent load node.role master name
```

Not cool.

## Why the name?
Because [@fsaravia](https://github.com/fsaravia) is smarter than me and he pointed out that this utility sorts the _body_ of a file.

## License
See [LICENSE](LICENSE).
