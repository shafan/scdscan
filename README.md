# scdscan
Tools for pentesting Source Code Disclosure

Source managed by GIT and SVN

The first tool [Finder](#Finder) allows to check that a domain list has an exposed source code.

The second tool [Dumper](#Dumper) (todo) will allow to retrieve an exposed source code

## Tools

### Finder
Find domains having git or svn repos exposed publically.

#### Usage
```
Usage:
  scdscan find URL [flags]

Flags:
  -h, --help   help for find
```

You can pass either a url in parameter 
```
$ scdscan find http://domain.com
```
or use the command with the pipe
```
$ cat list.txt | scdscan find
```

The script will output discovered domains to stdout.

#### How does it work?
- GIT : Checks if the .git/HEAD file contains refs/heads.
- SVN : Checks if the .svn/wc.db or .svn/entries file exist.
- The search for multiple urls is performed asynchronously, which saves a lot of time.

### Dumper
TODO

## TODO
- [] Test the code
- [] Dump repository

## Why a new git repository tool?
Two reasons led me to develop this tool.

I am a developer and I like to discover new languages. I wanted to discover [Go](https://golang.org). To learn at best, you need theory and practice. 
The theoretical part is reading online resources like Maximilien Andile's book: [Practical Go Lessons](https://www.practical-go-lessons.com). The practical part, well it's precisely this tool that I develop above all to practice developing in Go.

I also have a penchant for pentesting. In this context, I have to use several tools. Several tools are used to check if git repositories are exposed on URLS. One of my favorite is [GitTools](https://github.com/internetwache/GitTools). But this one is written in Python and I didn't find any in Go. So I found it interesting to write one in Go.
But to be different from this tool, I brought a feature I use a lot: the possibility to use the command in a pipeline.
I also extended the search to the svn repository.

If you are experienced in Go (or not), feel free to look at the code and give me feedback, that's also how I could progress.

## License

All tools are licensed using the MIT license. See [LICENSE.md](LICENSE.md)
