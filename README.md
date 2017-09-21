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

## Why the name?
Because [@fsaravia](https://github.com/fsaravia) is smarter than me and he pointed out that this utility sorts the _body_ of a file.

## License
See [LICENSE](LICENSE).
