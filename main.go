package main

import (
	"bufio"
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	. "github.com/kkdai/oxford-face"
	"github.com/spf13/cobra"
)

//Remove file extension for renaming process
func removeExtension(full string) string {
	extension := filepath.Ext(full)
	return full[0 : len(full)-len(extension)]
}

//Draw a rectangle on src source when detect a face on it
func drawRectange(srcFile string, r image.Rectangle) error {
	//Check src file
	fileSrc, err := os.Open(srcFile)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer fileSrc.Close()
	imgSrc, _ := jpeg.Decode(fileSrc)

	fN := removeExtension(srcFile)
	newFileName := fmt.Sprintf("%s_update.%s", fN, filepath.Ext(srcFile))
	fileDst, err := os.Create(newFileName)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer fileDst.Close()

	jpg := image.NewRGBA(image.Rect(r.Min.X, r.Min.Y, r.Max.X, r.Max.Y))
	draw.Draw(jpg, imgSrc.Bounds().Add(image.Pt(10, 10)), imgSrc, imgSrc.Bounds().Min, draw.Src)
	jpeg.Encode(fileDst, jpg, nil)
	return nil
}

func toggleLogging(enable bool) {
	if enable {
		log.SetOutput(os.Stdout)
	} else {
		log.SetOutput(ioutil.Discard)
	}
}

func printConsole() {
	fmt.Println("Command:( C:Create S:Subscription P:Publish R:RemoveTopic V:Verbose G:Read Q:exit )")
	fmt.Printf(":>")
}

func main() {

	key := os.Getenv("MSFT_KEY")
	if key == "" {
		fmt.Println("Please export your key to environment first, `export MSFT_KEY=12234`")
		return
	}
	f := NewFace(key)

	//Detect
	ret, err := f.DetectFile(nil, "./1.jpg")
	fmt.Println("ret:", ret, " err=", err)

	var serverAddr string
	var verbose bool

	rootCmd := &cobra.Command{
		Use:   "oxford-face-client",
		Short: "Client to connect to Project Oxford Face API services",
		Run: func(ccmd *cobra.Command, args []string) {

			toggleLogging(verbose)

			fmt.Println("Connect to coapmq server:", serverAddr)
			client := NewClient(serverAddr)
			if client == nil {
				fmt.Println("Cannot connect to server, please check your setting.")
				return
			}
			quit := false
			scanner := bufio.NewScanner(os.Stdin)
			printConsole()
			for !quit {

				var topic, msg string

				if !scanner.Scan() {
					break
				}
				line := scanner.Text()
				parts := strings.Split(line, " ")
				cmd := parts[0]
				if len(parts) > 1 {
					topic = parts[1]
				}
				if len(parts) > 2 {
					msg = parts[2]
				}

				fmt.Println(topic, msg)
				//var err error
				switch cmd {
				case "D", "d": //CREATE TOPIC

				case "Q", "q":
					quit = true
				case "V", "v":
					verbose = !verbose
					toggleLogging(verbose)
					fmt.Println("Switch verbose to ", verbose)
				default:
					fmt.Println("Command not support.")
				}

				if quit != true {
					printConsole()
				}
			}
		},
	}
	rootCmd.Flags().StringVarP(&serverAddr, "key", "k", "", "Project Oxford key, please export your key to environment first, `export MSFT_KEY=12234`")
	rootCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Verbose")
	rootCmd.Execute()
}
