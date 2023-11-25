# Dormite

Shutdown your computer from any device on your local network.

# uses:

Execute the program and go to
    
    http://[YOUR IP]:3004/?t=[MINUTS TO SLEEP]

if you want to cancel the shutdown schedule just go to
    http://[YOUR IP]:3004/?t=c

# Configuration

To prevent the remote execute the subnet ip is harcoded on the main.go 

    if strings.Contains(r.RemoteAddr, "192.168")

change that line with your subnet ip.