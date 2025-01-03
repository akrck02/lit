package command

import "log"

// Show help in the standard output
func Help() {

	log.Println(" Lit help: ")
	log.Println("------------------------------------------------------")
	log.Println("-f | Configuration file to be loaded (Optional).")
	log.Println("-g | Generate a configuration file.")
	log.Println("-t | Trace.")
	log.Println("-h | Useful help.")
	log.Println("\n Lit creates a ./dist/ directory by default to store the generated sources.")

}
