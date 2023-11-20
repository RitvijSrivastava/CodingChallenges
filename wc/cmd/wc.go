package cmd

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/spf13/cobra"
)

var (
	bBytes, bLines, bWords, bChars bool
)

type counts struct {
	nBytes, nLines, nWords, nChars int
}

func getCount(cmd *cobra.Command, args []string) {
	bBytes, _ = cmd.Flags().GetBool("bytes")
	bLines, _ = cmd.Flags().GetBool("lines")
	bWords, _ = cmd.Flags().GetBool("words")
	bChars, _ = cmd.Flags().GetBool("chars")

	if len(args) < 1 {
		c := count(os.Stdin, "")
		print(c, "")
	} else {
		
		total := counts{}
		
		for _, filename := range args {
			f, err := os.Open(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v: No such file found\n", filename)
				break
			}
			
			c := count(f, filename)

			f.Close()

			total.nBytes += c.nBytes
			total.nWords += c.nWords
			total.nLines += c.nLines
			total.nChars += c.nChars
			
			print(c, filename)
		}
		
		if len(args) > 1 {
			print(total, "total")
		}
	}
}

func count(r io.Reader, filename string) counts {
	scanner := bufio.NewReader(r)

	c := counts{}

	for {
		str, err := scanner.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Fprintf(os.Stderr, "error while reading file %v: %v", filename, err)
			return counts{}
		}


		c.nLines++;
		c.nBytes += len(str)
		c.nWords += len(bytes.Fields(str))
		c.nChars += utf8.RuneCount(str)
	}

	return c
}

func print(c counts, filename string) {
	values := []string{}
	
	if !(bBytes || bLines || bWords || bChars) {
		bBytes, bLines, bWords = true, true, true
	}

	if bLines {
		values = append(values, fmt.Sprintf("%8d", c.nLines))
	}

	if bWords {
		values = append(values,fmt.Sprintf("%8d", c.nWords))
	}

	if bBytes {
		values = append(values, fmt.Sprintf("%8d", c.nBytes))
	}

	if bChars {
		values = append(values, fmt.Sprintf("%8d", c.nChars))
	}

	if filename != "" {
		values = append(values, filename)
	}

	fmt.Println(strings.Join(values, " "))
}
