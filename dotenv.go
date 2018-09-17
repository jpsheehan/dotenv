package dotenv

import (
	"bufio"
	"os"
	"strings"
)

// Config sets environment variables according to the contents of filenames
// If filenames is not supplied it will read a file named '.env'
// If filenames is supplied then it will read all of the files specified
// It returns an error if one occurred or nil if everything worked
func Config(filenames ...string) error {

	// the filename to use if none are supplied
	const defaultFilename string = ".env"

	// make sure we only read the default file if no args are supplied.
	if len(filenames) == 0 {
		filenames = []string{defaultFilename}
	}

	// iterate through all the files, reading each one
	for i := 0; i < len(filenames); i++ {

		err := ConfigOne(filenames[i])
		if err != nil {
			return err
		}

	}

	return nil

}

// ConfigOne reads one configuration file and sets the environment variables
// accordingly.
// It returns an error if one occurred or nil if everything worked.
func ConfigOne(filename string) error {

	const (
		// the whitespace to trim from the keys and values
		cutset string = " \n\t\r"

		// the delimiter to split each line into a key-value pair
		delim string = ":"
	)

	// attempt to open the file
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	// create a scanner for the file
	scanner := bufio.NewScanner(file)

	// iterate over every line
	for scanner.Scan() {

		// split the strings on the delimiter
		parts := strings.SplitN(scanner.Text(), delim, 2)

		// make sure that we have exactly two parts, otherwise we
		// ignore the line
		if len(parts) == 2 {

			// trim the key and value
			key := strings.Trim(parts[0], cutset)
			val := strings.Trim(parts[1], cutset)

			// attempt to set the environment variable, returning
			// an error if something goes wrong
			err := os.Setenv(key, val)
			if err != nil {
				return err
			}

		}

	}

	// close the file
	file.Close()

	return nil
}
