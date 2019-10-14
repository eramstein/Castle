Bearlibtermnial Go bindings for Windows:
https://bitbucket.org/cfyzium/bearlibterminal/issues/31/go-bindings-require-compiling-from-c

I did get the Go bindings to work. Here is what I had to do on my Windows 7 machine:

Create folder src\bearlibterminal in Go workspace directory
Copy BearLibTerminal.go into folder created in step 1
Copy BearLibTerminal.h into folder created in step 1
Install TDM-GCC-64 version 5.1.0-2
Add TDM-GCC-64\bin to PATH
Copy BearLibTerminal.dll and BearLibTerminal.lib to TDM-GCC-64\lib
Only after I did all these things was I able to import "bearlibterminal" and use the bindings. Note that BearLibTerminal.dll and BearLibTerminal.lib must also be present in the same folder as the compiled executable or Go file to be run.