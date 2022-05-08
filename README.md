<h1 align="center">Go - TCP Scan :ninja:</h1>
<img align="right" src="./sage.svg" height="169"> 
Port scan is a technique used to identify the status of a port in a network, we usually use this for analisys and understand a target network.

<h3><br>  Wait... what is TCP? </br> </h3>
<p>
TCP means Transmission Control Protocol, the predominant
standard for connection-oriented, reliable communications and
the foundation of modern networking.
</p>

This application uses goRoutines to make a faster scanner, and through TCP Three-way Handshake we can discover open ports!


## How To Run

Clone the project

```bash
  git clone https://github.com/franciscofeo/go-tcpscan.git
```

Go to the project directory

```bash
  cd go-tcpscan
```


Start the application

```bash
  go run main.go <url> <workers quantity> <ports quantity>
```



## Example

After started a server locally in http://localhost:8080, we just have to write this code below to execute the scan in 100 ports with 1000 workers:

```bash
  go run main.go http://localhost:8080 1000 100
```
