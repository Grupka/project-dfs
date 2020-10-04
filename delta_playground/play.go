package main

import (
	"github.com/monmohan/xferspdy"
	"os"
)

func main() {
	//Create fingerprint of a file
	//xferspdy.NewFingerprintFromReader creates a fingerprint using Reader
	fingerprint := xferspdy.NewFingerprint("./delta_playground/file_ver1.txt", 1024)

	//generate the diff after file was updated
	diff := xferspdy.NewDiff("./delta_playground/file_ver2.txt", *fingerprint)

	//diff is sufficient to recover/recreate the modified file, given the base/source and the diff.
	modifiedFile, _ := os.OpenFile("./delta_playground/file_v2_from_v1.txt", os.O_CREATE|os.O_WRONLY, 0777)

	//This writes the output=(initial file + difference) to modifiedFile (Writer).
	xferspdy.PatchFile(diff, "./delta_playground/file_ver1.txt", modifiedFile)

	//buf, _ := ioutil.ReadAll(modifiedFile)
	//fmt.Print(buf)
}
