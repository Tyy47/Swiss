package network

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"

	"swiss/utils"
)

type FileReturn struct {
	name string
	file *os.File
	err  string
}

var outputFile = FileReturn{
	name: "swiss_net_output",
	err:  "File already exists",
}

var (
	endpoint = utils.CheckArguments(utils.Arguments, 3, 3)
	port     = utils.CheckArguments(utils.Arguments, 4, 4)
)

func networkCrashError(err error, data string) {
	if err != nil {
		utils.Error("Unable to gather " + data + " from " + endpoint + ".")
		os.Exit(1)
	}
}

// Connects to an IP Address or domain via TCP and returns an HTTP status code that states the result of the connection
func Connection() {
	// Connects to the domain or IP address
	conn, err := net.Dial("tcp", endpoint)
	if err != nil {
		utils.Crash(err)
	}
	
	// Prints the status code neatly
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		utils.Crash(err)
	}
	
	// Closes the connection when everything is complete
	defer conn.Close()
	// Prints the result of the http status code.
	fmt.Println(status)
}

// Connects to an IP Address or Domain through a port, prints a statement regarding the status of the port.
func GetPortStatus() {
	// Grabs the port from the package arguments and checks if it is valid.
	port, err := strconv.Atoi(port)
	if err != nil || port < 1 || port > 65535 {
		utils.Error("Port " + strconv.Itoa(port) + " is CLOSED on " + endpoint + ".")
		utils.Reason("Port exceeds or is under port range 0-65535")
		return
	}
	
	// Concats the address into a full string with the port
	address := endpoint + ":" + strconv.Itoa(port)
	// Timeout configured for the port.
	timeout := 3 * time.Second
	
	// Connects to the address with the newly given port
	conn, err := net.DialTimeout("tcp", address, timeout)

	// Handles an error if the connection fails
	if err != nil {
		utils.Error("Port " + strconv.Itoa(port) + " is CLOSED on " + endpoint + ".")
		utils.Reason(err.Error())
		return
	}
	
	// Closes the connection when the connection attempt is completed
	defer conn.Close()

	// Success message if the port is open on the address.
	utils.Success("Port " + strconv.Itoa(port) + " is OPEN on " + endpoint + ".")
}

// Takes an endpoint as a string and prints the IPv4 and v6 address of the domain.
func GetAddresses() {
	// Connects to the domain via a given address
	conn, err := net.LookupIP(endpoint)
	// Crashes the program if there is a network connection related error.
	networkCrashError(err, "IP addresses") 
	
	// Prints all IP addresses that are tied to that address to the console.
	for rows := range conn {
		fmt.Println(conn[rows])
	}
}

// Connects to an address and writes collected addresses to a network output file.
func writeAddresses() {
	// Connects to the endpoint
	conn, err := net.LookupIP(endpoint)
	// Checks if the function crashes
	networkCrashError(err, "IP Addresses")
	// Open a new writer to write address output.
	write := bufio.NewWriter(outputFile.file)
	
	// Writes IP Addresses title in output file and crashes if unable to write
	if _, err := write.WriteString("\nIP Addresses: \n"); err != nil {
		utils.CrashCheck(err)
	}
	
	// Collects all available IP addresses and converts them to strings.
	for rows := range conn {
		val := conn[rows].String()
	
		// Writes all available addresses to output file and crashes if unable to write.
		if _, err := write.WriteString(val + "\n"); err != nil {
			utils.CrashCheck(err)
		}
	}
	
	// Closes the writer once the function is complete
	write.Flush()
}

// Collects nameservers from a given address and prints it to the console.
func GetNameServer() {
	// Connects to the endpoint to retrieve nameservers.
	conn, err := net.LookupNS(endpoint)
	// Crashes the program if unable to connect
	networkCrashError(err, "Name servers")
	
	// Prints all available nameservers to the console.
	utils.Success("\nNameservers for " + endpoint + ".")
	for rows := range conn {
		val := conn[rows].Host + "\n"
		fmt.Println(val)
	}
}

// Collects namesevers from a given address and writes them to a network output file.
func writeNameServer() {
	// Connects to the endpoint to retrieve nameservers.
	conn, err := net.LookupNS(endpoint)
	// Crashes the program if unable to connect.
	networkCrashError(err, "Name servers")

	// Opens up the file in a new "writer"
	write := bufio.NewWriter(outputFile.file)
	// Writes all nameservers to output file, crashes the program if unable to write.
	if _, err := write.WriteString("\nName Servers: \n"); err != nil {
		utils.CrashCheck(err)
	}
	for rows := range conn {
		val := conn[rows].Host + "\n"

		_, err := write.WriteString(val)
		if err != nil {
			utils.Crash(err)
		}
	}
	
	// Closes the writer once the function is complete
	write.Flush()
}

// Connects to a given address and retrieves known CNAME records.
func GetCNameRecords() {
	// Connects to the endpoint to retreive the records
	conn, err := net.LookupCNAME(endpoint)
	// Crashes if unable to retrieve records
	networkCrashError(err, "CNAME records")
	
	// Prints the known CNAME records to the console.
	utils.Success("CNAME Records for " + endpoint + ".")
	fmt.Println(conn)
}

// Connects to a given address, retrieves known CNAME records, and writes them to a network output file.
func writeCNameRecords() {
	// Connects to the endpoint to get the CNAME records
	conn, err := net.LookupCNAME(endpoint)
	// Crashes the program if unable to connect
	networkCrashError(err, "CNAME records")
	
	// Opens up a writer to write in the file
	write := bufio.NewWriter(outputFile.file)
	// Writes the CNAME records title in file.
	if _, err := write.WriteString("\nCNAME Records: \n"); err != nil {
		utils.CrashCheck(err)
	}
	
	// Writes CNAME records into network output file.
	_, errr := write.WriteString(conn)
	if errr != nil {
		utils.CrashCheck(err)
	}
	
	// Closes the file once function is complete
	write.Flush()
}

// Connects to an endpoint, retrieves TXT Records, and prints them to the console.
func GetTXTRecords() {
	// Connects to the endpoint to retrieve TXT records
	conn, err := net.LookupTXT(endpoint)
	// Crashes the program if unable to connect
	networkCrashError(err, "TXT records")
	
	// Prints all known TXT records to the console.
	utils.Success("TXT Records for " + endpoint + ".")
	for rows := range conn {
		fmt.Println(conn[rows])
	}
}

// Connects to a given endpoint, retrieves TXT records, and writes them to a network output file.
func writeTXTRecords() {
	// Connects to the endpoint to retrieve the records
	conn, err := net.LookupTXT(endpoint)
	// Crashes if unable to connect
	networkCrashError(err, "TXT records")
	
	// Opens a new writer to write output file
	write := bufio.NewWriter(outputFile.file)
	// Writes TXT records title in output file
	if _, err := write.WriteString("\n\nTXT Records: \n"); err != nil {
		utils.CrashCheck(err)
	}
	
	// Writes all TXT records to network output file
	for rows := range conn {
		_, err := write.WriteString(conn[rows] + "\n")
		utils.CrashCheck(err)
	}
	
	// Closes the writer when the function is complete
	write.Flush()
}

// Connects to a given endpoint, collects known MX records, and prints them to the console.
func GetMXRecords() {
	// Connects to the endpoint to retrieve MX records
	conn, err := net.LookupMX(endpoint)
	// Crashes the program if unable to connect.
	networkCrashError(err, "MX records")
	
	// Success message for printing to the console
	utils.Success("MX Records for " + endpoint + ".")
	
	// Prints all known MX Records to the console.
	for rows := range conn {
		fmt.Println("Host: " + conn[rows].Host + "\n")
	}
}

// Connects to a given domain, collects MX records, and writes them to the network output file.
func writeMXRecords() {
	// Connects to the given domain to retrieve MX records.
	conn, err := net.LookupMX(endpoint)
	// Crashes the program if unable to connect.
	networkCrashError(err, "MX records")
	
	// Opens a new writer to write records to output file
	write := bufio.NewWriter(outputFile.file)
	// Writes MX records title in output file
	if _, err := write.WriteString("\nMX Records: \n"); err != nil {
		utils.CrashCheck(err)
	}
	
	// Writes each MX records into output file.
	for rows := range conn {
		_, err := write.WriteString("Host: " + conn[rows].Host + "\n")
		utils.CrashCheck(err)
	}

	// Closes the writer when the function is complete
	write.Flush()
}

// Creates the network output file where all information is stored when swiss net gather is ran.
func initFileCreation() {
	if !utils.CheckFileExists(outputFile.name) {
		utils.MakeFile(outputFile.name, true)
		file, err := os.Create(outputFile.name)
		utils.CrashCheck(err)
		outputFile.file = file
	} else {
		// removes the old data
		os.Remove(outputFile.name)
		utils.MakeFile(outputFile.name, true)
		file, err := os.Create(outputFile.name)
		utils.CrashCheck(err)
		outputFile.file = file
	}
}

// Culminates all of the data in the above functions, then outputs all the data into a file for viewing.
func GatherData() {
	initFileCreation()
	writeAddresses()
	writeNameServer()
	writeCNameRecords()
	writeTXTRecords()
	writeMXRecords()
	utils.Success("Written data to " + outputFile.name + " for endpoint " + endpoint + ".")
}
