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

// File return object to store the name, os file type, and error of an output file.
type FileReturn struct {
	name string
	file *os.File
	err  string
}

// Creation of the output file object
var outputFile = FileReturn{
	name: "swiss_net_output",
	err:  "File already exists",
}

// Package variables to grab endpoint and port
var (
	endpoint = utils.CheckArguments(utils.Arguments, 3, 3)
	port     = utils.CheckArguments(utils.Arguments, 4, 4)
)

// Specific error message for networking related functions.
// Returns true if the lookup succeeded so callers can skip their section on failure.
func networkLookupOk(err error, data string) bool {
	if err != nil {
		utils.Error("Unable to gather " + data + " from " + endpoint + ".")
		utils.Reason(err.Error())
		return false
	}
	return true
}

// Connects to an IP Address or domain via TCP and returns an HTTP status code that states the result of the connection
func Connection() {
	// Connects to the domain or IP address
	conn, err := net.Dial("tcp", endpoint)
	if err != nil {
		utils.Error("Unable to connect to " + endpoint + ".")
		utils.Reason(err.Error())
		return
	}
	defer conn.Close()

	// Prints the status code neatly
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		utils.Error("Unable to read response from " + endpoint + ".")
		utils.Reason(err.Error())
		return
	}

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
	conn, err := net.LookupIP(endpoint)
	if !networkLookupOk(err, "IP addresses") {
		return
	}

	for rows := range conn {
		fmt.Println(conn[rows])
	}
}

// Connects to an address and writes collected addresses to a network output file.
func writeAddresses() {
	write := bufio.NewWriter(outputFile.file)
	if _, err := write.WriteString("\nIP Addresses: \n"); err != nil {
		utils.Error("Unable to write IP Addresses header.")
		utils.Reason(err.Error())
		return
	}

	conn, err := net.LookupIP(endpoint)
	if !networkLookupOk(err, "IP addresses") {
		write.WriteString("(lookup failed)\n")
		write.Flush()
		return
	}

	for rows := range conn {
		val := conn[rows].String()
		if _, err := write.WriteString(val + "\n"); err != nil {
			utils.Error("Unable to write IP address entry.")
			utils.Reason(err.Error())
			break
		}
	}

	write.Flush()
}

// Collects nameservers from a given address and prints it to the console.
func GetNameServer() {
	conn, err := net.LookupNS(endpoint)
	if !networkLookupOk(err, "Name servers") {
		return
	}

	utils.Success("\nNameservers for " + endpoint + ".")
	for rows := range conn {
		val := conn[rows].Host + "\n"
		fmt.Println(val)
	}
}

// Collects namesevers from a given address and writes them to a network output file.
func writeNameServer() {
	write := bufio.NewWriter(outputFile.file)
	if _, err := write.WriteString("\nName Servers: \n"); err != nil {
		utils.Error("Unable to write Name Servers header.")
		utils.Reason(err.Error())
		return
	}

	conn, err := net.LookupNS(endpoint)
	if !networkLookupOk(err, "Name servers") {
		write.WriteString("(lookup failed)\n")
		write.Flush()
		return
	}

	for rows := range conn {
		val := conn[rows].Host + "\n"
		if _, err := write.WriteString(val); err != nil {
			utils.Error("Unable to write nameserver entry.")
			utils.Reason(err.Error())
			break
		}
	}

	write.Flush()
}

// Connects to a given address and retrieves known CNAME records.
func GetCNameRecords() {
	conn, err := net.LookupCNAME(endpoint)
	if !networkLookupOk(err, "CNAME records") {
		return
	}

	utils.Success("CNAME Records for " + endpoint + ".")
	fmt.Println(conn)
}

// Connects to a given address, retrieves known CNAME records, and writes them to a network output file.
func writeCNameRecords() {
	write := bufio.NewWriter(outputFile.file)
	if _, err := write.WriteString("\nCNAME Records: \n"); err != nil {
		utils.Error("Unable to write CNAME Records header.")
		utils.Reason(err.Error())
		return
	}

	conn, err := net.LookupCNAME(endpoint)
	if !networkLookupOk(err, "CNAME records") {
		write.WriteString("(lookup failed)\n")
		write.Flush()
		return
	}

	if _, err := write.WriteString(conn); err != nil {
		utils.Error("Unable to write CNAME entry.")
		utils.Reason(err.Error())
	}

	write.Flush()
}

// Connects to an endpoint, retrieves TXT Records, and prints them to the console.
func GetTXTRecords() {
	conn, err := net.LookupTXT(endpoint)
	if !networkLookupOk(err, "TXT records") {
		return
	}

	utils.Success("TXT Records for " + endpoint + ".")
	for rows := range conn {
		fmt.Println(conn[rows])
	}
}

// Connects to a given endpoint, retrieves TXT records, and writes them to a network output file.
func writeTXTRecords() {
	write := bufio.NewWriter(outputFile.file)
	if _, err := write.WriteString("\n\nTXT Records: \n"); err != nil {
		utils.Error("Unable to write TXT Records header.")
		utils.Reason(err.Error())
		return
	}

	conn, err := net.LookupTXT(endpoint)
	if !networkLookupOk(err, "TXT records") {
		write.WriteString("(lookup failed)\n")
		write.Flush()
		return
	}

	for rows := range conn {
		if _, err := write.WriteString(conn[rows] + "\n"); err != nil {
			utils.Error("Unable to write TXT entry.")
			utils.Reason(err.Error())
			break
		}
	}

	write.Flush()
}

// Connects to a given endpoint, collects known MX records, and prints them to the console.
func GetMXRecords() {
	conn, err := net.LookupMX(endpoint)
	if !networkLookupOk(err, "MX records") {
		return
	}

	utils.Success("MX Records for " + endpoint + ".")
	for rows := range conn {
		fmt.Println("Host: " + conn[rows].Host + "\n")
	}
}

// Connects to a given domain, collects MX records, and writes them to the network output file.
func writeMXRecords() {
	write := bufio.NewWriter(outputFile.file)
	if _, err := write.WriteString("\nMX Records: \n"); err != nil {
		utils.Error("Unable to write MX Records header.")
		utils.Reason(err.Error())
		return
	}

	conn, err := net.LookupMX(endpoint)
	if !networkLookupOk(err, "MX records") {
		write.WriteString("(lookup failed)\n")
		write.Flush()
		return
	}

	for rows := range conn {
		if _, err := write.WriteString("Host: " + conn[rows].Host + "\n"); err != nil {
			utils.Error("Unable to write MX entry.")
			utils.Reason(err.Error())
			break
		}
	}

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
	defer outputFile.file.Close()
	writeAddresses()
	writeNameServer()
	writeCNameRecords()
	writeTXTRecords()
	writeMXRecords()
	utils.Success("Written data to " + outputFile.name + " for endpoint " + endpoint + ".")
}
