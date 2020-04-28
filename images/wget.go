// +build !windows

package images

import "github.com/mhewedy/vermin/cmd"

func wget(url string, file string) (string, error) {
	return cmd.ExecuteP("wget", "-O", file, url)
}
