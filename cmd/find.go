/*
Copyright © 2021 Pierre Galvez <dev@pierre-galvez.fr>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/hackerGatherer/scdscan/find"
	"github.com/spf13/cobra"
)

var urls []url.URL

var findCmd = &cobra.Command{
	Use:   "find URL",
	Short: "Find domains having repository exposed publically.",
	Long: `Find domains having repository exposed publically.
You can pass either a url in parameter or use the command with the pipe
	`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			scanner := bufio.NewScanner(os.Stdin)
			leastOne := false
			for scanner.Scan() {
				input := scanner.Text()
				if !isValidUrl(input) {
					continue
				}
				leastOne = true
			}
			if err := scanner.Err(); err != nil {
				return errors.New("Error on arguments")
			}
			if leastOne == false {
				return errors.New("Requires valid url")
			}
			return nil
		}
		if len(args) > 1 {
			return errors.New("Requires only one url argument")
		}
		if !isValidUrl(args[0]) {
			return fmt.Errorf("Invalid url specified: %s", args[0])
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		find.Execute(urls)
	},
}

func init() {
	rootCmd.AddCommand(findCmd)
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}

func isValidUrl(toTest string) bool {
	if !strings.HasPrefix(toTest, "http://") && !strings.HasPrefix(toTest, "https://") {
		toTest = "http://" + toTest
	}

	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	}

	u, err := url.Parse(toTest)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	urls = append(urls, *u)

	return true
}
