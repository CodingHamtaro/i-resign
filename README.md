# iResign

The right tool for your graceful resignation.

This tool will help you to generate a common yet effective resignation letter for your growing career.


## Usage

Linux:
```bash
$ ./dist/linux[your_system_arch]/resign
```

Windows:

Just double-click the .exe file located under "dist/win[your_system_arch]/" directory.

## Compiling

Install first the dependencies by using the command below:
```bash
$ glide install
```

After installing the dependencies, compile the tool using the Go command line:
```bash
$ GOOS=windows GOARCH=386 go build -o dist/win32/resign.exe *.go
$ GOOS=windows GOARCH=amd64 go build -o dist/win64/resign.exe *.go
$ GOOS=linux GOARCH=386 go build -o dist/linux32/resign *.go
$ GOOS=linux GOARCH=amd64 go build -o dist/linux64/resign *.go
```

MIT License

Copyright (c) 2017 XHamtaro

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.